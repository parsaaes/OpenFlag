package config

//nolint:lll
const defaultConfig = `
logger:
  access:
    enabled: true
    path: "./logs/access.log"
    format: "${remote_ip} - - [${time_rfc3339}] \"${method} ${uri} HTTP/1.1\" ${status} \
    ${bytes_out} ${bytes_in} ${latency} \"${referer}\" \"${user_agent}\"\n"
    max-size: 1024
    max-backups: 7
    max-age: 7
  app:
    level: debug
    path: "./logs/app.log"
    max-size: 1024
    max-backups: 7
    max-age: 7
    stdout: true

server:
  address: :7677
  read-timeout: 20s
  write-timeout: 20s
  graceful-timeout: 5s

postgres:
  host: 127.0.0.1
  port: 5432
  user: openflag
  pass: secret
  dbname: openflag
  connect-timeout: 30s
  connection-lifetime: 30m
  max-open-connections: 10
  max-idle-connections: 5

redis:
  master:
    address: 127.0.0.1:6379
    pool-size: 0
    min-idle-conns: 20
    dial-timeout: 5s
    read-timeout: 3s
    write-timeout: 3s
    pool-timeout: 4s
    idle-timeout: 5m
    max-retries: 5
    min-retry-backoff: 1s
    max-retry-backoff: 3s
  slave:
    address: 127.0.0.1:6379
    pool-size: 0
    min-idle-conns: 20
    dial-timeout: 5s
    read-timeout: 3s
    write-timeout: 3s
    pool-timeout: 4s
    idle-timeout: 5m
    max-retries: 5
    min-retry-backoff: 1s
    max-retry-backoff: 3s

monitoring:
  prometheus:
    enabled: true
    address: ":9001"
`