package errors

import (
	err "github.com/pkg/errors"
)

func NewWithPrefix(arg string, pref string) error {
	return err.Wrap(err.New(arg), pref+": ")
}

func WithPrefix(arg error, pref string) error {
	return err.Wrap(arg, pref+": ")
}

func New(message string) error {
	return err.New(message)
}

func Errorf(format string, args ...interface{}) error {
	return err.Errorf(format, args...)
}

func WithStack(arg error) error {
	return err.WithStack(arg)
}

func Wrap(arg error, message string) error {
	return err.Wrap(arg, message)
}

func Wrapf(arg error, format string, args ...interface{}) error {
	return err.Wrapf(arg, format, args...)
}

func WithMessage(arg error, message string) error {
	return err.WithMessage(arg, message)
}

func WithMessagef(arg error, format string, args ...interface{}) error {
	return err.WithMessagef(arg, format, args...)
}

func Cause(arg error) error {
	return err.Cause(arg)
}
