package docker

import (
	"fmt"

	"github.com/ibrahimamar07/devctl/commandTemplate"
	"github.com/ibrahimamar07/devctl/internal/runner"
)

func Stop() {

	stopCommand := commandTemplate.GetCommandDocker()

	stopCommand = append(stopCommand, "stop")

	err := runner.Run("docker", stopCommand...)

	if err != nil {
		fmt.Println("error", err)
	}
}
