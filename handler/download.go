package handler

import (
	"filestoreServer/meta"
	"io/ioutil"
	"net/http"
	"os"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fileHash := r.Form["filehash"][0]
	fmeta := meta.GetFileMeta(fileHash)
	f, err := os.Open(fmeta.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("content-disposition", "attachment;filename=\"" + fmeta.FileName + "\"")
	w.Write(data)
}
