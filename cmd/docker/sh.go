package docker

import (
	"fmt"

	"github.com/ibrahimamar07/devctl/commandTemplate"
	"github.com/ibrahimamar07/devctl/internal/runner"
)

func Sh() {

	shCommand := commandTemplate.GetCommandDocker()

	shCommand = append(shCommand, "exec", "app", "sh")

	err := runner.Run("docker", shCommand...)

	if err != nil {
		fmt.Println("error", err)
	}
}
