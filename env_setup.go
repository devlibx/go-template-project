package go_template_project

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed env/common.env
var commonEnv string

//go:embed env/dev.env
var devEnv string

//go:embed env/stage.env
var stageEnv string

//go:embed env/test.env
var testEnv string

//go:embed env/e2e_test.env
var e2eTestEnv string

type EnvSetupFunc func(envs map[string]string)

func DefaultEnvSetupFunc() EnvSetupFunc {
	return func(envs map[string]string) {
		for key, value := range envs {
			_ = os.Setenv(key, value)
			fmt.Printf("Setting %s to %s\n", key, value)
		}
	}
}

func buildFinalEnvAndCall(envs map[string]string, setupFunc EnvSetupFunc) {
	if setupFunc != nil {
		finalEnv := map[string]string{}
		for key, value := range envs {
			if val, ok := os.LookupEnv(key); ok {
				finalEnv[key] = val
			} else {
				finalEnv[key] = value
			}
		}
		setupFunc(finalEnv)
	}
}

func SetupCommonEnv(envs map[string]string, setupFunc EnvSetupFunc) {
	scanner := bufio.NewScanner(strings.NewReader(commonEnv))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			envs[parts[0]] = parts[1]
		}
	}

	buildFinalEnvAndCall(envs, setupFunc)
}

func SetupDevEnv(envs map[string]string, setupFunc EnvSetupFunc) {
	scanner := bufio.NewScanner(strings.NewReader(devEnv))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			envs[parts[0]] = parts[1]
		}
	}

	buildFinalEnvAndCall(envs, setupFunc)
}

func SetupTestEnv(envs map[string]string, setupFunc EnvSetupFunc) {
	scanner := bufio.NewScanner(strings.NewReader(testEnv))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			envs[parts[0]] = parts[1]
		}
	}

	buildFinalEnvAndCall(envs, setupFunc)
}

func SetupStageEnv(envs map[string]string, setupFunc EnvSetupFunc) {
	scanner := bufio.NewScanner(strings.NewReader(stageEnv))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			envs[parts[0]] = parts[1]
		}
	}

	buildFinalEnvAndCall(envs, setupFunc)
}

func SetupE2ETestEnv(envs map[string]string, setupFunc EnvSetupFunc) {
	SetupCommonEnv(envs, nil)
	SetupDevEnv(envs, nil)
	SetupStageEnv(envs, nil)
	SetupTestEnv(envs, nil)

	scanner := bufio.NewScanner(strings.NewReader(e2eTestEnv))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			envs[parts[0]] = parts[1]
		}
	}

	buildFinalEnvAndCall(envs, setupFunc)
}
