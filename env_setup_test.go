package go_template_project

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildFinalEnvAndCall(t *testing.T) {
	// Setup
	keyOverride := "TEST_OVERRIDE_VAR"
	keyNew := "TEST_NEW_VAR"
	
	valSystem := "system_value"
	valFile := "file_value"
	valNew := "new_value"

	os.Setenv(keyOverride, valSystem)
	defer os.Unsetenv(keyOverride)
	// Ensure keyNew is not set
	os.Unsetenv(keyNew)

	// Input map simulating values read from files
	envs := map[string]string{
		keyOverride: valFile,
		keyNew:      valNew,
	}

	// Capture result
	var finalEnvs map[string]string
	captureFunc := func(e map[string]string) {
		finalEnvs = e
	}

	// Execute
	buildFinalEnvAndCall(envs, captureFunc)

	// Verify
	// Expectation:
	// 1. keyOverride should be valSystem (System takes precedence)
	// 2. keyNew should be valNew (File value used if not in System)
	
	// Check current implementation behavior (likely to fail based on analysis)
	// We want the code to pass these assertions eventually.
	assert.Equal(t, valSystem, finalEnvs[keyOverride], "Expected system value to take precedence")
	assert.Equal(t, valNew, finalEnvs[keyNew], "Expected file value to be used when system env is missing")
}
