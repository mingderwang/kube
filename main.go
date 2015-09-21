package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "kube"
	app.Usage = "cli for sharing kubernetes resource files repos and management"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "Search global kubernetes resource files",
			Action: func(c *cli.Context) {
				println("searching")
			},
		},
		{
			Name:    "push",
			Aliases: []string{"p"},
			Usage:   "Push resource files to kube.hub with tag",
			Action: func(c *cli.Context) {
				println("pushing files")
			},
		},
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Get resource files from kube.hub by tag",
			Action: func(c *cli.Context) {
				println("downloading")
			},
		},
		{
			Name:    "like",
			Aliases: []string{"star"},
			Usage:   "Add a star for a specific resource tag",
			Action: func(c *cli.Context) {
				println("star +1")
			},
		},
	}

	app.Run(os.Args)
}
