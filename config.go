package main

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Monitors []Monitor
	Min      int
}

const CfgFilePath = "/etc/brightness.conf"

func parseConfig() *Config {
	cfgData, err := os.ReadFile(CfgFilePath)
	if err != nil {
		die("Error reading config file", false)
	}

	var c Config
	toml.Unmarshal(cfgData, &c)

	/*
		min := tree.GetDefault("min", 30)

		minInt, ok := min.(int64)

		if !ok {
			die("Error with min value in config file", false)
		}
		c.Min = int(minInt)

		monitors, ok := tree.Get("monitors").([]*toml.TomlTree)
		if !ok {
			die("Error with monitors value in config file", false)
		}

		for i, mon := range monitors {
			monMap := mon.ToMap()
			driver, ok := monMap["driver"].(string)
			if !ok {
				die(fmt.Sprintf("Error with driver value in config file for monitor #%v", i), false)
			}
			delete(monMap, "driver")

			gamma, ok := monMap["gamma"].(float64)
			if !ok {
				die(fmt.Sprintf("Error with gamma value in config file for monitor #%v", i), false)
			}
			delete(monMap, "gamma")

			scale, ok := monMap["scale"].(float64)
			if !ok {
				die(fmt.Sprintf("Error with scale value in config file for monitor #%v", i), false)
			}
			delete(monMap, "scale")
			if scale > 1 {
				die(fmt.Sprintf("Scale value in config file for monitor #%v cannot be > 1", i), false)
			}

			m := Monitor{
				Driver: driver,
				Gamma:  gamma,
				Scale:  scale,
				Opts:   monMap,
			}

			c.Monitors = append(c.Monitors, m)
		}
	*/

	return &c
}
