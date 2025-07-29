package server

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func frameHandler(w http.ResponseWriter, r *http.Request) {
	/* 处理视频的帧上传 */
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			http.Error(w, "Close the file error", http.StatusInternalServerError)
			return
		}
	}(file)

	fileDate := r.FormValue("date")
	imei := r.FormValue("imei")
	videoName := r.FormValue("videoName")
	uploadDir := fmt.Sprintf("./%s/%s/videos/%s/%s", BASEURL, imei, fileDate, videoName)
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		http.Error(w, "Create the directory error", http.StatusInternalServerError)
		return
	}

	dst, err := os.Create(filepath.Join(uploadDir, handler.Filename))
	if err != nil {
		http.Error(w, "Create the file error", http.StatusInternalServerError)
		return
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			http.Error(w, "Close the file error", http.StatusInternalServerError)
			return
		}
	}(dst)

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Save the file error", http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
	if err != nil {
		http.Error(w, "Communication between server and client error", http.StatusInternalServerError)
		return
	}
}