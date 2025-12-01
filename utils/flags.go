package utils

import (
	"flag"
	"fmt"
)

type Flags struct {
	Part      int
	InputPath string
}

func ParseFlags() (Flags, error) {
	var flags Flags

	flag.IntVar(&flags.Part, "part", 1, "passes which part of the day to run")
	flag.StringVar(&flags.InputPath, "input_path", "", "path to the input file")

	flag.Parse()

	if flags.InputPath == "" {
		return Flags{}, fmt.Errorf("input_path is required")
	}
	if flags.Part < 1 || flags.Part > 2 {
		return Flags{}, fmt.Errorf("part must be either 1 or 2")
	}

	return flags, nil
}
