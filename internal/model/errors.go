package model

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidID       = errors.New("batcycle: invalid identifier")
	ErrNotFound        = errors.New("batcycle: entity not found")
	ErrConflict        = errors.New("batcycle: state conflict")
	ErrInterlock       = errors.New("batcycle: interlock denied")
	ErrMoistureHold    = errors.New("batcycle: moisture hold active")
	ErrAirflowSetpoint = errors.New("batcycle: airflow setpoint violation")
	ErrFanFault        = errors.New("batcycle: fan fault")
	ErrScheduleEmpty   = errors.New("batcycle: schedule empty")
	ErrSoakNotDue      = errors.New("batcycle: soak window not yet due")
	ErrGradient        = errors.New("batcycle: moisture gradient violation")
	ErrMoistureDrift   = errors.New("batcycle: moisture drift exceeded")
	ErrHeatOvertemp    = errors.New("batcycle: heat overtemperature")
	ErrGradientHold    = errors.New("batcycle: gradient hold not satisfied")
	ErrContextCanceled = errors.New("batcycle: operation canceled")
)

type DomainError struct {
	Op   string
	Code string
	Err  error
}

func (e *DomainError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Err != nil {
		return fmt.Sprintf("batcycle %s [%s]: %v", e.Op, e.Code, e.Err)
	}
	return fmt.Sprintf("batcycle %s [%s]", e.Op, e.Code)
}

func (e *DomainError) Unwrap() error { return e.Err }

func Wrap(op, code string, err error) error {
	if err == nil {
		return nil
	}
	return &DomainError{Op: op, Code: code, Err: err}
}

func Is(err, target error) bool   { return errors.Is(err, target) }
func As(err error, target any) bool { return errors.As(err, target) }
