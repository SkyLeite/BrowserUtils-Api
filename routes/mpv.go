package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os/exec"
)

type mpvPostData struct {
	URL      string
	IsOnTop  bool
	Geometry string
}

func RunWithMpv(w http.ResponseWriter, r *http.Request) {
	bytes, _ := httputil.DumpRequest(r, true)

	fmt.Println(string(bytes))

	decoder := json.NewDecoder(r.Body)
	var body mpvPostData
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}

	args := []string{body.URL}

	if body.IsOnTop == true {
		args = append(args, "--ontop")
	}

	if len(body.Geometry) > 0 {
		args = append(args, fmt.Sprintf("--geometry=%s", body.Geometry))
	}

	exec.Command("mpv", args...).Start()

	io.WriteString(w, "Done!")
}
