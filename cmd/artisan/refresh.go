package artisan

import (
	"fmt"

	"github.com/ibrahimamar07/devctl/commandTemplate"
	"github.com/ibrahimamar07/devctl/internal/runner"
)

func Refresh() {

	refreshcommand := commandTemplate.GetCommandArtisan()

	refreshcommand = append(refreshcommand, "migrate:fresh --seed")

	err := runner.Run("docker", refreshcommand...)

	if err != nil {
		fmt.Println("error", err)
	}
}
