package drivers

import (
	"os/exec"
	"fmt"
	"log"
	"strings"
)

type Xrandr struct {
	output string
}

func NewXrandr(opts *Options) (Driver, error) {
	xr := Xrandr{}
	output := (*opts)["output"]
	outputStr, ok := output.(string)
	if !ok {
		return nil, fmt.Errorf(`Option "output" for XRandR driver should be a string`)
	}
	xr.output = outputStr
	return xr, nil
}

func (xr Xrandr) SetBrightness(val float64) {
	path := "xrandr"
	args := []string{"--output", xr.output, "--brightness", fmt.Sprintf("%.3f", val)}
	log.Printf("Calling %v %v", path, strings.Join(args, " "))
	cmd := exec.Command(path, args...)
	cmd.Run()
}
