package artisan

import (
	"fmt"

	"github.com/ibrahimamar07/devctl/commandTemplate"
	"github.com/ibrahimamar07/devctl/internal/runner"
)

func Clear() {

	clearCommand := commandTemplate.GetCommandArtisan()

	clearCommand = append(clearCommand, "optimize:clear")

	err := runner.Run("docker", clearCommand...)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
