package common

import (
	"bytes"
	"os"
	"os/exec"
)

var version string

// SetVersion overrides the detected version, typically during build time.
func SetVersion(v string) {
	version = v
}

// GetVersion returns the build version, falling back to git describe when unset.
func GetVersion() string {
	if version != "" {
		return version
	}

	gitVersion := getVersionFromGit()
	if gitVersion != "" {
		return gitVersion
	}

	return ""
}

func getVersionFromGit() string {
	if _, err := os.Stat(".git"); err != nil {
		return ""
	}

	tag, err := exec.Command("git", "describe", "--tags").Output()
	if err != nil {
		return ""
	}

	tag = bytes.TrimSpace(tag)

	if isRepoDirty() {
		tag = append(tag, '+')
	}

	return string(tag)
}

func isRepoDirty() bool {
	cmd := exec.Command("git", "status", "--porcelain")
	output, _ := cmd.Output()
	output = bytes.TrimSpace(output)
	return len(output) != 0
}
