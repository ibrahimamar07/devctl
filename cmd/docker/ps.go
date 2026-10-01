package docker

import (
	"fmt"

	"github.com/ibrahimamar07/devctl/commandTemplate"
	"github.com/ibrahimamar07/devctl/internal/runner"
)

func Ps() {

	psCommand := commandTemplate.GetCommandDocker()

	psCommand = append(psCommand, "ps")

	err := runner.Run("docker", psCommand...)

	if err != nil {
		fmt.Println("error", err)
	}
}
