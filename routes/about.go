package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/urfave/cli"
)

type about struct {
	Version string
}

func About(c *cli.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var data = &about{
			Version: c.App.Version,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			io.WriteString(w, "Error")
		}

		io.WriteString(w, string(jsonData))
	}

}
