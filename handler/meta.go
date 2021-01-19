package handler

import (
	"encoding/json"
	"filestoreServer/meta"
	"fmt"
	"net/http"
	"os"
)

//获取原始文件
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fileHash := r.Form["filehash"][0]
	fmt.Printf("file hash %s", fileHash)
	fmeta := meta.GetFileMeta(fileHash)
	data, err := json.Marshal(fmeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fileHash := r.Form.Get("filehash")
	fileNewName := r.Form.Get("filename")
	opType := r.Form.Get("op")
	//0为删除
	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	curFileMeta := meta.GetFileMeta(fileHash)
	curFileMeta.FileName = fileNewName
	meta.UpdateFileMeta(curFileMeta)
	data, err := json.Marshal(curFileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func FileDeleteHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fileHash := r.Form.Get("filehash")

	fmata := meta.GetFileMeta(fileHash)
	os.Remove(fmata.Location)
	meta.DeleteFileMeta(fileHash)
	w.WriteHeader(http.StatusOK)
}