package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		targetDir := "/eks-share/mnt/s3"
		files, err := os.ReadDir(targetDir)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("something is wrong!!!")
			return
		}
		for _, file := range files {
			log.Default().Println(file.Name())
			w.Write([]byte(file.Name()))
		}
		w.Write([]byte("===="))
	})
	http.ListenAndServe(":8080", nil)
}
