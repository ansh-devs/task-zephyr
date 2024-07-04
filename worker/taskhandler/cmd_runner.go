package taskhandler

import (
	"github.com/sirupsen/logrus"
	"os/exec"
)

func ShellCommandRunner(command string) error {
	cmd := exec.Command(command)
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	logrus.Info("Result for command :" + command + " is " + string(output))
	return nil
}
