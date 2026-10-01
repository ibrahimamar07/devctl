package docker

import (
	"fmt"

	"github.com/ibrahimamar07/devctl/commandTemplate"
	"github.com/ibrahimamar07/devctl/internal/runner"
)

func Up() {
	upCommand := commandTemplate.GetCommandDocker()

	upCommand = append(upCommand, "up", "-d")

	err := runner.Run("docker", upCommand...)

	if err != nil {
		fmt.Println("eror", err)
	}
}
