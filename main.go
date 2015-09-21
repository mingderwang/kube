package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "kube"
	app.Usage = "cli for sharing kubernetes resource files repos and management"
	app.Version = "1.0.0"
	// global level flags
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Show more output",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "Search global kubernetes resource with keyword.",
			Action: func(c *cli.Context) {
				keyword := ""
				if len(c.Args()) > 1 {
					println("NOTICE: Only support one keyword in v1")
				}
				if len(c.Args()) > 0 {
					keyword = c.Args()[0]
					searchTag(keyword)
				} else {
					println("you need at least one keyword for search. for example: kube search redis")

				}
			},
		},
		{
			Name:    "push",
			Aliases: []string{"p"},
			Usage:   "Push resource files to kube.hub with tag",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "tag, t",
					Usage: "Tag your kubernetes resource files repo a name.",
				},
				cli.StringFlag{
					Name:  "description, d",
					Usage: "Write a description for your kubernetes resource repo.",
				},
			},
			Action: func(c *cli.Context) {
				filename := ""
				tag := ""
				description := ""
				if len(c.Args()) > 1 {
					println("NOTICE: Only pushing one resource file in v1")
				}
				if len(c.Args()) > 0 {
					if c.String("tag") == "" {
						println("You need to specify a tag")
						return
					}
					filename = c.Args()[0]
					tag = c.String("tag")
					description = c.String("description")
					println("pushing file: ", filename)
					println("with tag: ", tag)
					println("with description: ", description)
					pushFile(filename, tag, description)
				} else {
					println("You need to push with files, for example: kube push my-redis-rc.json -t ming.redis")

				}
			},
		},
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Get resource files from kube.hub by tag",
			Action: func(c *cli.Context) {
				tag := ""
				if len(c.Args()) > 1 {
					println("NOTICE: Only get one repo at a time")
				}
				if len(c.Args()) > 0 {
					tag = c.Args().First()
					println("get tag: ", tag)
					println("NOTICE: You will get the resource files, you can use kubectl command after that.")
					getFile(tag)
				} else {
					println("NOTICE: You need to specify a tag for download, for example: kube get ming.redis")
					println("NOTICE: You can use subcommand search to find one, for example: kube search redis.")
				}
			},
		},
		{
			Name:    "like",
			Aliases: []string{"star"},
			Usage:   "Add a star for a specific resource tag",
			Action: func(c *cli.Context) {
				tag := ""
				if len(c.Args()) > 1 {
					println("NOTICE: Only like one repo at a time")
				}
				if len(c.Args()) > 0 {
					tag = c.Args().First()
					println("I like tag: ", tag)
					likeIt(tag)
				} else {
					println("NOTICE: You need to specify a tag to give it a start")
				}
			},
		},
	}

	app.Run(os.Args)
}
