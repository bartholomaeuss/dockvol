/*
Copyright © 2025 Bartholomaeuss
*/
package backup

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func findContainer(volumeName string) (string, error) {

	dockerCmd := exec.Command("docker", "ps", "-aq", "--filter", "volume="+volumeName)

	stdout, err := dockerCmd.Output()
	if err != nil {
		return "", fmt.Errorf("docker ps failed: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(stdout)), "\n")
	countContainers := len(lines)
	if countContainers != 1 {
		return "", fmt.Errorf("architectural violation! Volume in question is not in exakt one container mounted")
	}
	if lines[0] == "" {
		return "", fmt.Errorf("volume is not mounted at all")
	}

	return lines[0], nil
}

func stopContainer(containerName string) error {

	dockerCmd := exec.Command("docker", "stop", containerName)

	stdout, err := dockerCmd.Output()
	if err != nil {
		return fmt.Errorf("idk: %w %v", err, stdout)
	}
	return nil
}

func backupContainer(volumeName string) error {

	dockerCmd := exec.Command(
		"docker", "run",
		"--rm",
		"-v", fmt.Sprintf("%s:%s:ro", volumeName, "/cdata"),
		"-v", fmt.Sprintf("%s:/backup", "backup"),
		"alpine",
		"tar", "czf",
		fmt.Sprintf("/backup/%s_%s.tar.gz", volumeName, time.Now().Format("20060102_150405")),
		"-C", "/cdata",
		".",
	)

	if err := dockerCmd.Run(); err != nil {
		return err
	}

	return nil
}

func startContainer(containerName string) error {
	dockerCmd := exec.Command("docker", "start", containerName)

	if err := dockerCmd.Run(); err != nil {
		return err
	}
	return nil
}

func Backup(volumeName string) error {
	containerName, err := findContainer(volumeName)
	if err != nil {
		return err
	}
	if err := stopContainer(containerName); err != nil {
		return err
	}
	if err := backupContainer(volumeName); err != nil {
		return err
	}
	if err := startContainer(containerName); err != nil {
		return err
	}
	return nil
}
