package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type config struct {
	port string
	dir  string

	gitPath    string
	goPath     string
	screenPath string
	screenName string
}

func main() {
	conf := config{}
	flag.StringVar(&conf.port, "port", ":9178", "port to listen on < 65535, default: :9178")
	flag.StringVar(&conf.dir, "dir", "./", "directory to update")
	flag.StringVar(&conf.gitPath, "git", "/usr/bin/git", "path/to/git")
	flag.StringVar(&conf.goPath, "go", "/usr/bin/go", "path/to/go")
	flag.StringVar(&conf.screenPath, "screen", "/usr/bin/screen", "path/to/screen")
	flag.StringVar(&conf.screenName, "name", "SESS", "name for screen")
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Fatal("usage: autodeploy deploy/bin pass through options")
	}

	fmt.Println("listening on ", conf.port)
	http.HandleFunc("/", conf.handler)
	http.ListenAndServe(conf.port, nil)
}

type Payload struct {
	Commits []CommitPayload
}
type CommitPayload struct {
	Message string
}

func (c config) handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	payload := Payload{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Println("decoding payload failed: ", err)
		return
	}

	// update
	gl := exec.Command(c.gitPath, "pull")
	gl.Dir = c.dir
	if err := gl.Run(); err != nil {
		log.Println("pull failed: ", err)
		return
	}

	// build
	gb := exec.Command(c.goPath, "build")
	gb.Dir = c.dir
	if err := gb.Run(); err != nil {
		log.Println("build failed: ", err)
		return
	}

	// stop
	sk := exec.Command(c.screenPath, "-XS", c.screenName, "quit")
	if err := sk.Run(); err != nil {
		log.Println("screen kill failed: ", err)
	}

	// restart
	opts := []string{"-dms", c.screenName}
	ss := exec.Command(c.screenPath, append(opts, flag.Args()...)...)
	ss.Dir = c.dir
	if err := ss.Run(); ss != nil {
		log.Println("restart failed: ", err)
		return
	}

	log.Println("updated")
}
