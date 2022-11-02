package drivers

import (
	"fmt"
)

type Driver interface {
	SetBrightness(float64)
}

type Options map[string]interface{}

func New(name string, opts *Options) (Driver, error) {
	var err error
	var d Driver
	switch name {
	case "xrandr":
		d, err = NewXrandr(opts)
	case "acpi":
		d, err = NewAcpi(opts)
	default:
		return nil, fmt.Errorf(`no driver for type "%v"`, name)
	}
	if err != nil {
		return nil, err
	}
	return d, nil
}
