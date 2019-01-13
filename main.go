package main

import (
	"BrowserUtils/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/cli"
)

type CliArgs struct {
	port int
	dir  string
}

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s", c.App.Version)
	}

	app := cli.NewApp()
	app.Name = "Browser Utils"
	app.Version = "v1.0"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port",
			Value: 36284,
			Usage: "Specify a port for the app to run",
		},
		cli.StringFlag{
			Name:  "dir",
			Usage: "Specify a default directory to download files",
		},
	}

	app.Action = func(c *cli.Context) error {
		var router = mux.NewRouter()
		var port = c.Int("port")

		router.HandleFunc("/about", routes.About(c)).Methods("GET")
		router.HandleFunc("/play/mpv", routes.RunWithMpv(c)).Methods("POST")
		router.HandleFunc("/youtubedl", routes.DownloadWithYoutubeDL(c)).Methods("POST")

		log.Printf("Listening on port %d", port)
		return http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
