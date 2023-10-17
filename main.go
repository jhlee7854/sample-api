package main

import (
	"fmt"
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
			fmt.Println(file.Name())
		}
		w.Write([]byte("hello"))
	})
	http.ListenAndServe(":8080", nil)
}
