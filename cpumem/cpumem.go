package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/projecteru2/core-plugins/cpumem/cmd"
	"github.com/projecteru2/core/version"
)

var debug bool

func setupLog(l string) error {
	logrus.SetOutput(os.Stderr)
	level, err := logrus.ParseLevel(l)
	if err != nil {
		return err
	}
	logrus.SetLevel(level)

	formatter := &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	}
	logrus.SetFormatter(formatter)
	return nil
}

func main() {
	app := &cli.App{
		Name:    "CPU-Mem",
		Usage:   "The resource plugin to manage cpu and memory",
		Version: version.VERSION,
		Commands: cmd.Commands,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Usage:       "enable debug",
				Aliases:     []string{"d"},
				Value:       false,
				Destination: &debug,
			},
			&cli.StringFlag{
				Name: "config",
				Usage: "config file path",
				Value: "./cpumem.yaml",
				EnvVars: []string{"CPUMEM_CONFIG"},
			},
		},
	}

	var loglevel string
	if debug {
		loglevel = "DEBUG"
	} else {
		loglevel = "INFO"
	}

	if err := setupLog(loglevel); err != nil {
		fmt.Printf("Error setup log: %v\n", err)
		os.Exit(-1)
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Errorf("Error running cpumem: %v\n", err)
		os.Exit(-1)
	}
}
