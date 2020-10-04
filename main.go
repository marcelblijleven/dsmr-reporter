package main

import (
	"flag"
	"github.com/marcelblijleven/dsmrreporter/dsmr"
	"github.com/marcelblijleven/dsmrreporter/reporters"
	"github.com/tarm/serial"
	"os"
)

var (
	device string
	baud   string
	parity string
)

func main() {
	reporter := reporters.NewStdOutReporter()

	err := parseFlags(os.Args[1:])

	if err != nil {
		reporter.Fatal(err)
	}

	br, err := dsmr.OpenPort(serial.OpenPort, device, baud, parity)

	if err != nil {
		reporter.Fatal(err)
	}


	dsmr.Read(br, reporter, "Europe/Amsterdam")
}

func parseFlags(args []string) error {
	flag.StringVar(&device, "device", "", "Device to listen on, e.g. /dev/ttyAMA0")
	flag.StringVar(&baud, "baud", "115200", "baud to use for the P1 port. Default is 115200")
	flag.StringVar(&parity, "parity", "", "parity to use for the P1 port. Choose from even, mark, none, odd or space")

	err := flag.CommandLine.Parse(args)

	if err != nil {
		return err
	}

	return nil
}