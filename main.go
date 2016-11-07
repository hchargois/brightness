package main

import (
	"fmt"
	"os"
	"path"
)

func usage() {
	progname := path.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "Usage:\n"+
	"  %v value\n" +
	"\n"+
	"value: The brightness value to set as a percentage, absolute (e.g. \"90\") or relative (e.g. \"+10\")\n",
	progname)
}

func die(msg string, withUsage bool) {
	fmt.Fprintln(os.Stderr, msg)
	if withUsage {
		usage()
	}
	os.Exit(1)
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		die("Incorrect number of arguments", true)
	}

	switch args[0] {
	case "-h", "--help":
		usage()
		os.Exit(0)
	default:
		cfg := parseConfig()
		setBrightness(args[0], cfg)
	}
}
