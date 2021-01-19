package main

import (
	"filestoreServer/handler"
	"fmt"
	"net/http"
)

const port = "8082"

func main()  {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)

	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)

	http.HandleFunc("/file/delete", handler.FileDeleteHandler)




	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("failed to start start err: %s", err.Error())
	}
}
