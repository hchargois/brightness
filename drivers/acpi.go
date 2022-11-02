package drivers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Acpi struct {
	path string
}

func NewAcpi(opts *Options) (Driver, error) {
	a := Acpi{}
	path := (*opts)["path"]
	pathStr, ok := path.(string)
	if !ok {
		return nil, fmt.Errorf(`option "path" for ACPI driver should be a string`)
	}
	if !strings.HasPrefix(pathStr, "/sys/class/backlight/") {
		return nil, fmt.Errorf(`option "path" for ACPI driver should start with "/sys/class/backlight/"`)
	}
	a.path = pathStr
	return a, nil
}

func retrieveMaxBrightness(path string) (int, error) {
	fp := filepath.Join(path, "max_brightness")
	contents, err := os.ReadFile(fp)
	if err != nil {
		log.Printf(`Error while reading max brightness from %v: %v`, fp, err)
		return 0, err
	}
	value, err := strconv.Atoi(strings.TrimSpace(string(contents)))
	if err != nil {
		log.Printf(`Error while reading max brightness from %v: %v`, fp, err)
		return 0, err
	}
	return value, nil
}

func writeBrightness(path string, b int) error {
	fp := filepath.Join(path, "brightness")
	f, err := os.OpenFile(fp, os.O_WRONLY|os.O_TRUNC, 0)
	if err != nil {
		log.Printf(`Error while writing brightness to %v: %v`, fp, err)
		return err
	}
	defer f.Close()

	_, err = f.WriteString(strconv.Itoa(b))
	if err != nil {
		log.Printf(`Error while writing brightness to %v: %v`, fp, err)
		return err
	}
	return nil
}

func (a Acpi) SetBrightness(val float64) {
	max, err := retrieveMaxBrightness(a.path)
	if err != nil {
		return
	}

	valToSet := int(val * float64(max))
	log.Printf(`Setting %v to %v (max. %v)`, a.path, valToSet, max)

	err = writeBrightness(a.path, valToSet)
	if err != nil {
		return
	}
}
