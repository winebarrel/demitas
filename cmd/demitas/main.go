package main

import (
	"log"

	"github.com/winebarrel/demitas"
)

func init() {
	log.SetFlags(0)
}

func main() {
	opts := parseArgs()

	containerDef, err := demitas.BuildContainerDefinition(opts.ContainerDefSrc, opts.ContainerDefOverrides)

	if err != nil {
		log.Fatal(err)
	}

	taskDef, err := demitas.BuildTaskDefinition(opts.TaskDefSrc, opts.TaskDefOverrides, containerDef)

	if err != nil {
		log.Fatal(err)
	}

	svrDef, err := demitas.BuildServiceDefinition(opts.ServiceDefSrc, opts.ServiceDefOverrides)

	if err != nil {
		log.Fatal(err)
	}

	ecsConf, err := demitas.BuildEcspressoConfig(opts.EcspressoConfigSrc, opts.EcspressoConfigOverrides)

	if err != nil {
		log.Fatal(err)
	}

	err = demitas.RunTask(&opts.RunOptions, ecsConf, svrDef, taskDef)

	if err != nil {
		log.Fatal(err)
	}
}
