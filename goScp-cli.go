package main

import (
	"SecureCodePlatform-cli/Config"
	"SecureCodePlatform-cli/Login"
	"SecureCodePlatform-cli/Pull"
	"SecureCodePlatform-cli/Score"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"time"
)

func main() {
	app := &cli.App{
		Name: "SecureCodePlatform CLI",

		Version:  "v0.0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Hahanov Ivan",
				Email: "hahanov.i@explabs.ru",
			},
		},
		Usage:     "Work with SecureCodePlatform",
		UsageText: "command-line interface for easy and fast interaction with SecureCodePlatform",
	}
	app.Commands = []*cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Initialised structure for work",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "dir",
					Aliases:     []string{"d"},
					Value:       os.Getenv("HOME"),
					Usage:       "Root dir for program",
					DefaultText: os.Getenv("HOME"),
					EnvVars:     []string{"HOME"},
				},
				&cli.StringFlag{
					Name:        "config",
					Aliases:     []string{"c"},
					Value:       ".scp.yaml",
					Usage:       "Choose `FILE` to upload",
					DefaultText: ".scp.yaml",
				},
			},
			Action: func(context *cli.Context) error {
				rootDirectory := context.String("dir")
				configName := context.String("config")
				username, err := Login.EnterLoginAndPass()
				if err != nil {
					return err
				}
				token, err := Login.ValidateToken()
				if err != nil {
					return err
				}
				config := Config.YamlConfig{}
				config.Username = username
				config.RootDir = rootDirectory
				config.Token = token

				configPath := fmt.Sprintf("%s/%s", rootDirectory, configName)
				fmt.Println("Creating default conf file in ", configPath)

				if err = config.CreateConfig(); err != nil {
					return err
				}
				if err = config.UpdateConfig(); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "login",
			Aliases: []string{"l"},
			Usage:   "Login to your account",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "username",
					Aliases: []string{"u"},
					Usage:   "choose task `id` for upload",
				},
				&cli.StringFlag{
					Name:    "password",
					Aliases: []string{"p"},
					Usage:   "Choose `FILE` to upload",
				},
			},
			Action: func(context *cli.Context) error {
				username := context.String("username")
				password := context.String("password")
				if username != "" && password != "" {
					if _, err := Login.Auth(username, password); err != nil {
						return err
					}
				} else if username != "" && password == "" {
					if _, err := Login.EnterPass(username); err != nil {
						return err
					}
				} else if username == "" && password == "" {
					if _, err := Login.EnterLoginAndPass(); err != nil {
						return err
					}
				}
				return nil
			},
		},
		{
			Name:    "pull",
			Aliases: []string{"p"},
			Usage:   "complete a task on the list",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "task",
					Aliases: []string{"t"},
					Usage:   "choose task `id` for upload",
				},
			},
			Action: func(context *cli.Context) error {
				_, err := Login.ValidateToken()
				if err != nil {
					return err
				}
				_, err = Pull.CreateChallengeDir()
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "upload",
			Aliases: []string{"u"},
			Usage:   "options for task templates",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "task",
					Aliases: []string{"t"},
					Usage:   "choose task `id` for upload",
				},
				&cli.StringFlag{
					Name:    "file",
					Aliases: []string{"f"},
					Usage:   "Choose `FILE` to upload",
				},
			},
		},
		{
			Name:    "score",
			Aliases: []string{"s"},
			Usage:   "options for task templates",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "",
					Aliases: []string{"t"},
					Usage:   "choose task `id` for upload",
				},
				&cli.StringFlag{
					Name:    "file",
					Aliases: []string{"f"},
					Usage:   "Choose `FILE` to upload",
				},
			},
			Action: func(context *cli.Context) error {
				token, err := Login.ValidateToken()
				if err != nil {
					return err
				}
				if err = Score.PrintScoreboard(token); err != nil {
					return err
				}
				return nil
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
