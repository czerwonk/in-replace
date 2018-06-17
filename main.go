package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

const version = "0.1"

func main() {
	configFile := flag.String("config", "config.yml", "")
	showVersion := flag.Bool("v", false, "Shows the version information")
	flag.Parse()

	if *showVersion {
		showVersionInfo()
		os.Exit(0)
	}

	process(*configFile)
}

func showVersionInfo() {
	fmt.Println("in-replace")
	fmt.Println("Version:", version)
	fmt.Println("Author: Daniel Czerwonk")
}

func process(configFile string) {
	config, err := loadConfigFromFile(configFile)
	if err != nil {
		fmt.Println("failed loading config: ", err)
		os.Exit(1)
	}

	if !processFiles(config) {
		os.Exit(2)
	}
}

func processFiles(c *Config) bool {
	wg := sync.WaitGroup{}

	res := true
	for _, f := range c.Files {
		wg.Add(1)
		go func(f *File) {
			defer wg.Done()

			err := processFile(f)
			if err != nil {
				fmt.Printf("error (%s): %v", f.Path, err)
				res = false
			}
		}(f)
	}

	wg.Wait()
	return res
}
