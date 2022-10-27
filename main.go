package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/vaibhavchalse99/app"
	"github.com/vaibhavchalse99/config"
	"github.com/vaibhavchalse99/db"
	"github.com/vaibhavchalse99/server"
)

func main() {
	config.Load()
	db.LoadRoles()
	app.Init()
	defer app.Close()

	cliApp := cli.NewApp()
	cliApp.Name = "Golang app"
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start server",
			Action: func(c *cli.Context) error {
				server.StartAPIServer()
				return nil
			},
		},
		{
			Name:  "create-migration",
			Usage: "create migration file",
			Action: func(c *cli.Context) error {
				return db.CreateMigration(c.Args().Get(0))
			},
		},
		{
			Name:  "migrate",
			Usage: "run db migrations",
			Action: func(c *cli.Context) error {
				return db.RunMigration()
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback migrations",
			Action: func(c *cli.Context) error {
				return db.RollbackMigrations(c.Args().Get(0))
			},
		},
	}
	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}
