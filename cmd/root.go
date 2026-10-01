package cmd

import (
	"fmt"

	"github.com/ibrahimamar07/devctl/cmd/artisan"
	"github.com/ibrahimamar07/devctl/cmd/docker"
)

func Execute(args []string) {
	if len(args) == 0 {
		fmt.Println("devctl - Development CLI-For Simplify docker compose -artisan command")
		fmt.Println()
		fmt.Println("dev --help")
		fmt.Println()
		fmt.Println()
		return
	}

	switch args[0] {

	//artisan
	case "clear":
		artisan.Clear()
	case "migrate":
		artisan.Migrate()
	case "seed":
		artisan.Seed()
	case "fresh":
		artisan.Refresh()

		//docker
	case "up":
		docker.Up()
	case "down":
		docker.Down()
	case "restart":
		docker.Restart()
	case "ps":
		docker.Ps()
	case "sh":
		docker.Sh()
	case "stop":
		docker.Stop()

	//help
	case "help":
		Help()

	default:
		fmt.Println("Unknown command:", args[0])
	}

	//docker compose exec app php artisan migrate:fresh --seed
}
