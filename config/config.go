package config

import (
	_ "embed"
	"github.com/TwiN/deepmerge"
	"os"
)

//go:embed app.yaml
var ApplicationConfigBytes []byte

//go:embed messaging.yaml
var MessagingConfig []byte

//go:embed http.yaml
var HttpConfigBytes []byte

func GetEnvExpandedMergedYamlApplicationConfig() (string, error) {
	if c, err := GetMergedYamlApplicationConfig(); err == nil {
		return os.ExpandEnv(c), nil
	} else {
		return "", err
	}
}

func GetMergedYamlApplicationConfig() (string, error) {
	if c, err := deepmerge.YAML(ApplicationConfigBytes, MessagingConfig); err != nil {
		return "", err
	} else {
		if c, err := deepmerge.YAML(c, HttpConfigBytes); err != nil {
			return "", err
		} else {
			return string(c), nil
		}
	}
}

func GetEnvExpandedMergedJsonApplicationConfig() (string, error) {
	if c, err := GetMergedJsonApplicationConfig(); err == nil {
		return os.ExpandEnv(c), nil
	} else {
		return "", err
	}
}

func GetMergedJsonApplicationConfig() (string, error) {
	return "", nil
}
