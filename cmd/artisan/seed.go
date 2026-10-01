package artisan

import (
	"fmt"

	"github.com/ibrahimamar07/devctl/commandTemplate"
	"github.com/ibrahimamar07/devctl/internal/runner"
)

func Seed() {

	seedCommand := commandTemplate.GetCommandArtisan()

	seedCommand = append(seedCommand, "db:seed")

	err := runner.Run("docker", seedCommand...)

	if err != nil {
		fmt.Println("error", err)
	}
}
