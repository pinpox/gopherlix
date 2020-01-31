package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

func main() {

	var flagConfig = flag.String("config", "config.ini", "Configuration file to use")
	// Set up commandline flags
	flag.Parse()

	cfg, err := ini.Load(*flagConfig)
	if err != nil {
		log.Fatalf("Fail to read configuration file: %v", err)
		os.Exit(1)
	}

	log.Info("Data Path: ", cfg.Section("paths").Key("content").String())
	log.Info("Templates Path: ", cfg.Section("paths").Key("templates").String())

	log.Info("Server Port: ", cfg.Section("server").Key("port").String())
	log.Info("Server Domain: ", cfg.Section("server").Key("domain").String())
	log.Info("Server Host: ", cfg.Section("server").Key("host").String())

	sv := NewGopherServer(
		cfg.Section("server").Key("port").String(),
		cfg.Section("server").Key("domain").String(),
		cfg.Section("server").Key("host").String(),
		cfg.Section("paths").Key("content").String(),
		cfg.Section("paths").Key("templates").String(),
	)
	sv.Run()
}
