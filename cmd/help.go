package cmd

import "fmt"

func Help() {
	fmt.Println("devctl - Development CLI")
	fmt.Println("Simplify Docker Compose & Laravel Artisan commands")
	fmt.Println()

	fmt.Println("Usage:")
	fmt.Println("  dev <command>")
	fmt.Println()

	fmt.Println("Docker commands:")
	fmt.Println("  up        Start containers")
	fmt.Println("  down      Stop and remove containers")
	fmt.Println("  restart   Restart containers")
	fmt.Println("  ps        Show running containers")
	fmt.Println("  logs      Show container logs")
	fmt.Println("  shell     Open shell inside app container")
	fmt.Println()

	fmt.Println("Laravel commands:")
	fmt.Println("  migrate   Run database migrations")
	fmt.Println("  seed      Run database seeder")
	fmt.Println("  test      Run Laravel tests")
	fmt.Println("  clear     Clear Laravel cache")
	fmt.Println("  fresh     Refresh database and run seeder")
	fmt.Println("  tinker    Open Laravel Tinker")
	fmt.Println()

	fmt.Println("Other commands:")
	// fmt.Println("  artisan   Run any Laravel Artisan command")
	// fmt.Println("  docker    Run any Docker command")
	fmt.Println("  help      Show this help message")
	fmt.Println()

	fmt.Println("Examples:")
	fmt.Println("  dev up")
	fmt.Println("  dev logs")
	// fmt.Println("  devctl logs app")
	fmt.Println("  dev migrate")
	// fmt.Println("  devctl artisan route:list")
	// fmt.Println("  devctl docker compose ps")
}
