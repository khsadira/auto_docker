package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

type JsonPost struct {
	PushData struct {
		Tag string `json:"tag"`
	} `json:"push_data"`
	Repository struct {
		Name string `json:"name"`
		RepoName string `json:"repo_name"`
		Status string `json:"status"`
		Namespace string `json:"namespace"`
	} `json:"repository""`
}

func dockerWebhooks(w http.ResponseWriter, r *http.Request) {
	var scan JsonPost
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("%s\n", err)
		}
		json.Unmarshal(body ,&scan)
		if scan.Repository.Status == "Active" && scan.Repository.Namespace == "khsadira" {
			req := fmt.Sprintf("%s:%s", scan.Repository.RepoName, scan.PushData.Tag)
		//	println(req)
			out, _ := exec.Command("./script.py", scan.Repository.Name, req).Output()
			fmt.Printf("%s\n", out)

		}
	} else if r.Method == "GET" {
		println("GET")
	}
}

func main() {
	http.HandleFunc("/", dockerWebhooks)
	log.Fatal(http.ListenAndServe(":9191", nil))
}
