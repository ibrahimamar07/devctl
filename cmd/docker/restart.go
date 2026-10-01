package docker

import (
	"fmt"

	"github.com/ibrahimamar07/devctl/commandTemplate"
	"github.com/ibrahimamar07/devctl/internal/runner"
)

func Restart() {
	restartCommand := commandTemplate.GetCommandDocker()

	restartCommand = append(restartCommand, "restart")

	err := runner.Run("docker", restartCommand...)

	if err != nil {
		fmt.Println("error", err)
	}

}
