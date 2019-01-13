package routes

import (
	"encoding/json"
	"io"
	"net/http"
)

type about struct {
	Version string
}

func About(w http.ResponseWriter, r *http.Request) {
	var data = &about{
		Version: "0.1",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		io.WriteString(w, "Error")
	}

	io.WriteString(w, string(jsonData))
}
