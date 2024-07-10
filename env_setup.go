package go_template_project

import (
	"bufio"
	_ "embed"
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

func SetupCommonEnv() {
	scanner := bufio.NewScanner(strings.NewReader(commonEnv))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			_ = os.Setenv(parts[0], parts[1])
		}
	}
}

func SetupDevEnv() {
	scanner := bufio.NewScanner(strings.NewReader(devEnv))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			_ = os.Setenv(parts[0], parts[1])
		}
	}
}

func SetupTestEnv() {
	scanner := bufio.NewScanner(strings.NewReader(testEnv))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			_ = os.Setenv(parts[0], parts[1])
		}
	}
}

func SetupStageEnv() {
	scanner := bufio.NewScanner(strings.NewReader(stageEnv))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			_ = os.Setenv(parts[0], parts[1])
		}
	}
}

func SetupE2ETestEnv() {
	SetupCommonEnv()
	SetupDevEnv()
	SetupStageEnv()
	SetupTestEnv()

	scanner := bufio.NewScanner(strings.NewReader(e2eTestEnv))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			_ = os.Setenv(parts[0], parts[1])
		}
	}
}
