package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/OpenFlag/OpenFlag/internal/app/openflag/grpc"

	"github.com/OpenFlag/OpenFlag/internal/app/openflag/engine"

	"github.com/OpenFlag/OpenFlag/internal/app/openflag/handler"
	"github.com/OpenFlag/OpenFlag/internal/app/openflag/model"

	"github.com/OpenFlag/OpenFlag/pkg/database"

	"github.com/OpenFlag/OpenFlag/internal/app/openflag/metric"
	"github.com/OpenFlag/OpenFlag/pkg/monitoring/prometheus"
	"github.com/carlescere/scheduler"

	"github.com/OpenFlag/OpenFlag/pkg/redis"

	"github.com/OpenFlag/OpenFlag/internal/app/openflag/router"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/OpenFlag/OpenFlag/internal/app/openflag/config"
	"github.com/spf13/cobra"
)

const (
	healthCheckInterval = 1
)

// nolint:funlen
func main(cfg config.Config) {
	e := router.New(cfg)

	dbCfg := cfg.Database

	dbMaster := database.WithRetry(database.Create, dbCfg.Driver, dbCfg.MasterConnStr, dbCfg.Options)
	dbSlave := database.WithRetry(database.Create, dbCfg.Driver, dbCfg.SlaveConnStr, dbCfg.Options)

	defer func() {
		if err := dbMaster.Close(); err != nil {
			logrus.Errorf("database master connection close error: %s", err.Error())
		}

		if err := dbSlave.Close(); err != nil {
			logrus.Errorf("database slave connection close error: %s", err.Error())
		}
	}()

	redisCfg := cfg.Redis

	redisMasterClient, redisMasterClose := redis.Create(redisCfg.MasterAddress, redisCfg.Options, true)
	redisSlaveClient, redisSlaveClose := redis.Create(redisCfg.SlaveAddress, redisCfg.Options, false)

	defer func() {
		if err := redisMasterClose(); err != nil {
			logrus.Errorf("redis master connection close error: %s", err.Error())
		}

		if err := redisSlaveClose(); err != nil {
			logrus.Errorf("redis slave connection close error: %s", err.Error())
		}
	}()

	_, err := scheduler.Every(healthCheckInterval).Seconds().Run(func() {
		metric.ReportDbStatus(dbMaster, "database_master")
		metric.ReportDbStatus(dbMaster, "database_slave")
		metric.ReportRedisStatus(redisMasterClient, "redis_master")
		metric.ReportRedisStatus(redisSlaveClient, "redis_slave")
	})
	if err != nil {
		logrus.Fatalf("failed to start metric scheduler: %s", err.Error())
	}

	e.GET("/healthz", func(c echo.Context) error { return c.NoContent(http.StatusNoContent) })

	flagRepo := model.SQLFlagRepo{Driver: dbCfg.Driver, MasterDB: dbMaster, SlaveDB: dbSlave}
	entityRepo := model.NewRedisEntityRepo(
		redisMasterClient, redisSlaveClient, cfg.Evaluation.EntityContextCacheExpiration,
	)

	evaluationLogger := engine.NewLogger(cfg.Logger.Evaluation)
	evaluationEngine := engine.New(evaluationLogger, flagRepo)

	if err := evaluationEngine.Fetch(); err != nil {
		logrus.Fatalf("failed to fetch flags: %s", err.Error())
	}

	if err := evaluationEngine.Start(cfg.Evaluation.UpdateFlagsCronPattern); err != nil {
		logrus.Fatalf("Failed to start evaluation engine: %s", err.Error())
	}

	flagHandler := handler.FlagHandler{FlagRepo: flagRepo}
	evaluationHandler := handler.EvaluationHandler{Engine: evaluationEngine, EntityRepo: entityRepo}

	v1 := e.Group("/api/v1")

	v1.POST("/flag", flagHandler.Create)
	v1.DELETE("/flag/:id", flagHandler.Delete)
	v1.PUT("/flag/:id", flagHandler.Update)
	v1.GET("/flag/:id", flagHandler.FindByID)
	v1.POST("/flag/tag", flagHandler.FindByTag)
	v1.POST("/flag/history", flagHandler.FindByFlag)
	v1.POST("/flags", flagHandler.FindFlags)

	v1.POST("/evaluation", evaluationHandler.Evaluate)

	e.Static("/", "browser/openflag-ui/build")

	grpcServer := grpc.New(evaluationEngine, entityRepo)

	go func() {
		if err := grpcServer.Start(cfg.Server.RPCAddress); err != nil {
			logrus.Fatalf("failed to start gRPC server: %s", err.Error())
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := e.Start(cfg.Server.Address); err != nil {
			logrus.Fatalf("failed to start openflag server: %s", err.Error())
		}
	}()

	go prometheus.StartServer(cfg.Monitoring.Prometheus)

	logrus.Info("start openflag server!")

	s := <-sig

	logrus.Infof("signal %s received", s)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.GracefulTimeout)
	defer cancel()

	e.Server.SetKeepAlivesEnabled(false)

	if err := e.Shutdown(ctx); err != nil {
		logrus.Errorf("failed to shutdown openflag server: %s", err.Error())
	}

	if err := grpcServer.Shutdown(ctx); err != nil {
		logrus.Errorf("failed to shutdown gRPC server: %s", err.Error())
	}
}

// Register registers server command for openflag binary.
func Register(root *cobra.Command, cfg config.Config) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run OpenFlag server component",
			Run: func(cmd *cobra.Command, args []string) {
				main(cfg)
			},
		},
	)
}
