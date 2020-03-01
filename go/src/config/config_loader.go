package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const fieldSeparator = "="

//Load the configuration from the input external filego d
func GetConfigurationFrom(ConfigurationFile string) map[string]string {
	log.Printf("Getting the configuration from: %s \n", ConfigurationFile)
	configurations := make(map[string]string)
	configFile, err := os.Open(ConfigurationFile)
	reportError(err)
	scanner := bufio.NewScanner(configFile)
	for scanner.Scan() {
		line := scanner.Text()
		configurations = putInto(line, configurations)
	}

	defer configFile.Close()
	return configurations
}

func putInto(line string, configurationEntries map[string]string) map[string]string {
	if len(line) > 0 {
		contents := strings.Split(line, fieldSeparator)
		configurationEntries[contents[0]] = contents[1]
	}
	return configurationEntries
}

func reportError(reportedError error) {
	if reportedError != nil {
		log.Panic(reportedError)
	}
}
