package server

import "github.com/filebrowser/filebrowser/v2/settings"

const (
	ADDR = "0.0.0.0:32111"
	BASEURL = "statistics"
)

type androidUploadingContent struct {
	EnableLogUpload bool `json:"enableLogUpload"`
	EnableImageUpload bool `json:"enableImageUpload"`
	EnableVideoUpload bool `json:"enableVideoUpload"`
}

var ConfigContentChan = make(chan string)
var NowConfigContent settings.UploadingContent = settings.UploadingContent {
	Log:  false,
	Image: false,
	Video: false,
}