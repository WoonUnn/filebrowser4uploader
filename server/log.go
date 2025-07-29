package server

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

func logUploadHandler(w http.ResponseWriter, r *http.Request) {
	/* 处理日志上传 */
	fileName := r.Header.Get("File-Name")
	fileOffset := r.Header.Get("File-Offset")

	offset, err := strconv.ParseInt(fileOffset, 10, 64)
	if err != nil {
		http.Error(w, "Invalid File-Offset header", http.StatusBadRequest)
		return
	}

	fileDate := r.FormValue("date")
	imei := r.FormValue("imei")
	uploadDir := fmt.Sprintf("./%s/%s/logs/%s", BASEURL, imei, fileDate)
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		http.Error(w, "Create the directory error", http.StatusInternalServerError)
		return
	}

	filePath := uploadDir + "/" + fileName
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		http.Error(w, "Open file error", http.StatusInternalServerError)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			http.Error(w, "File close failed", http.StatusInternalServerError)
			return
		}
	}(file)

	if _, err := file.Seek(offset, io.SeekStart); err != nil {
		http.Error(w, "File seek failed", http.StatusInternalServerError)
		return
	}

	filePart, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer func(filePart multipart.File) {
		err := filePart.Close()
		if err != nil {
			http.Error(w, "Close the file error", http.StatusInternalServerError)
			return
		}
	}(filePart)

	if _, err := io.Copy(file, filePart); err != nil {
		http.Error(w, "Save the file error", http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "Log part uploaded successfully, Offset: %s (File name: %s)\n", fileOffset, fileName)
	if err != nil {
		http.Error(w, "Communication between server and client error", http.StatusInternalServerError)
		return
	}
}