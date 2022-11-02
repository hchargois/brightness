package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"sync"

	"github.com/hchargois/brightness/drivers"
)

type BrightnessVal struct {
	Relative bool
	Value    int
}

type Monitor struct {
	Driver     string
	Gamma      float64
	Scale      float64
	DriverOpts drivers.Options
}

func (m *Monitor) normalizeValue(val int) float64 {
	return math.Pow(float64(val)/100, 1/m.Gamma) * m.Scale
}

func (m *Monitor) SetBrightness(b int) {
	drv, err := drivers.New(m.Driver, &m.DriverOpts)
	if err != nil {
		log.Printf(`Error with driver %v: %v`, m.Driver, err)
		return
	}
	norm := m.normalizeValue(b)
	drv.SetBrightness(norm)
}

func parseValue(v string) BrightnessVal {
	var val BrightnessVal
	var err error
	val.Relative = v[0] == '+' || v[0] == '-'
	val.Value, err = strconv.Atoi(v)
	if err != nil {
		die(fmt.Sprintf("Invalid brightness value \"%v\"", v), false)
	}
	return val
}

func setBrightness(valStr string, cfg *Config) {
	var absVal int
	val := parseValue(valStr)
	if val.Relative {
		absVal = LoadBrightness() + val.Value
	} else {
		absVal = val.Value
	}

	if absVal > 100 {
		absVal = 100
	}

	if absVal < cfg.Min {
		log.Printf("Asking for less than minimum brightness (%v < %v), setting %v", absVal, cfg.Min, cfg.Min)
		absVal = cfg.Min
	}

	var wg sync.WaitGroup
	for _, mon := range cfg.Monitors {
		wg.Add(1)
		go func(mon Monitor) {
			mon.SetBrightness(absVal)
			defer wg.Done()
		}(mon)
	}
	wg.Wait()

	SaveBrightness(absVal)
}
