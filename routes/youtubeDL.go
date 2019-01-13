package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os/exec"

	"github.com/urfave/cli"
)

type youtubedlPostData struct {
	URL        string
	IsMP3      bool
	WorkingDir string
	Output     string
}

func DownloadWithYoutubeDL(c *cli.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := httputil.DumpRequest(r, true)

		fmt.Println(string(bytes))

		decoder := json.NewDecoder(r.Body)
		var body youtubedlPostData
		err := decoder.Decode(&body)
		if err != nil {
			panic(err)
		}

		args := []string{}

		if body.IsMP3 == true {
			args = append(args, "--extract-audio", "--audio-format", "mp3")
		}

		if len(body.Output) > 0 {
			args = append(args, "--output", fmt.Sprintf("\"%s\"", body.Output))
		}

		args = append(args, body.URL)

		cmd := exec.Command("youtube-dl", args...)

		if len(body.WorkingDir) > 0 {
			cmd.Dir = body.WorkingDir
		} else if len(c.String("dir")) > 0 {
			cmd.Dir = c.String("dir")
		}

		out, err := cmd.CombinedOutput()

		if err != nil {
			io.WriteString(w, fmt.Sprintf("Error:%s\nOutput:%s", err.Error(), string(out)))
			return
		}

		io.WriteString(w, fmt.Sprintf("Done! Output:%s", string(out)))
	}
}
