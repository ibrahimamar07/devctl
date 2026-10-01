package artisan

import (
	"fmt"

	"github.com/ibrahimamar07/devctl/commandTemplate"
	"github.com/ibrahimamar07/devctl/internal/runner"
)

func Migrate() {

	migrateCommand := commandTemplate.GetCommandArtisan()

	migrateCommand = append(migrateCommand, "migrate")

	err := runner.Run("docker", migrateCommand...)

	if err != nil {
		fmt.Println("error", err)
	}
}
