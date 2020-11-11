package constraint

import (
	"strconv"

	"github.com/OpenFlag/OpenFlag/internal/app/openflag/model"

	"github.com/sirupsen/logrus"
)

// LessThanConstraint represents Openflag less than constraint.
type LessThanConstraint struct {
	Value    float64 `json:"value"`
	Property string  `json:"property,omitempty"`
}

// Name is an implementation for the Constraint interface.
func (l LessThanConstraint) Name() string {
	return LessThanConstraintName
}

// Validate is an implementation for the Constraint interface.
func (l LessThanConstraint) Validate() error {
	return nil
}

// Initialize is an implementation for the Constraint interface.
func (l *LessThanConstraint) Initialize() error {
	return nil
}

// Evaluate is an implementation for the Constraint interface.
func (l LessThanConstraint) Evaluate(e model.Entity) bool {
	property, ok := GetProperty(l.Property, e)
	if !ok {
		return false
	}

	value, err := strconv.ParseFloat(property, 64)
	if err != nil {
		logrus.Errorf(
			"invalid property for less than constraint => property: %s, value: %s, err: %s",
			l.Property, property, err.Error(),
		)

		return false
	}

	return value < l.Value
}
