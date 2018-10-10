package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strastic"
)

func main() {
	var dir string
	var port string
	var isSPA bool

	flag.StringVar(&port, "port", "80", "port for strastic")
	flag.StringVar(&dir, "dir", "/var/www", "specify path to static dir")
	flag.BoolVar(&isSPA, "spa", false, "set true if your application is SPA")

	flag.Parse()

	if dir == "" {
		fmt.Println("Dir is empty. Do you specific the dir argument?")
		return
	}

	bEnv, _ := json.Marshal(strastic.GetStrasticEnv("ENV_"))

	fs := http.FileServer(strastic.FS{ServeFS: http.Dir(dir), IsSPA: isSPA})

	http.Handle("/", fs)
	http.HandleFunc("/config.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(bEnv)
	})

	fmt.Printf("Start strastic on %s \r\n", port)
	fmt.Printf("spa=%v dir=%s port=%s \r\n", isSPA, dir, port)
	fmt.Printf("Envs %s \r\n", bEnv)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Cannot start service. %s \r\n", err)
	}
}
