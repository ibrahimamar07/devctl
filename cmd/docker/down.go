package docker

import (
	"fmt"

	"github.com/ibrahimamar07/devctl/commandTemplate"
	"github.com/ibrahimamar07/devctl/internal/runner"
)

func Down() {

	downCommand := commandTemplate.GetCommandDocker()

	downCommand = append(downCommand, "down")

	err := runner.Run("docker", downCommand...)

	if err != nil {
		fmt.Println("error", err)
	}

}
