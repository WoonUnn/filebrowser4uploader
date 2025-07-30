package http

import (
	"encoding/json"
	"net/http"

	"github.com/filebrowser/filebrowser/v2/rules"
	"github.com/filebrowser/filebrowser/v2/server"
	"github.com/filebrowser/filebrowser/v2/settings"
)

type settingsData struct {
	Signup                bool                  `json:"signup"`
	CreateUserDir         bool                  `json:"createUserDir"`
	MinimumPasswordLength uint                  `json:"minimumPasswordLength"`
	UserHomeBasePath      string                `json:"userHomeBasePath"`
	Defaults              settings.UserDefaults `json:"defaults"`
	Rules                 []rules.Rule          `json:"rules"`
	Branding              settings.Branding     `json:"branding"`
	Uploading 		   	  settings.Uploading    `json:"uploading"`
	Tus                   settings.Tus          `json:"tus"`
	Shell                 []string              `json:"shell"`
	Commands              map[string][]string   `json:"commands"`
}

type androidUploadingContent struct {
	EnableLogUpload bool `json:"enableLogUpload"`
	EnableImageUpload bool `json:"enableImageUpload"`
	EnableVideoUpload bool `json:"enableVideoUpload"`
}

var settingsGetHandler = withAdmin(func(w http.ResponseWriter, r *http.Request, d *data) (int, error) {
	data := &settingsData{
		Signup:                d.settings.Signup,
		CreateUserDir:         d.settings.CreateUserDir,
		MinimumPasswordLength: d.settings.MinimumPasswordLength,
		UserHomeBasePath:      d.settings.UserHomeBasePath,
		Defaults:              d.settings.Defaults,
		Rules:                 d.settings.Rules,
		Branding:              d.settings.Branding,
		Uploading: 		  	   d.settings.Uploading,
		Tus:                   d.settings.Tus,
		Shell:                 d.settings.Shell,
		Commands:              d.settings.Commands,
	}

	server.NowConfigContent = d.settings.Uploading.Content

	return renderJSON(w, r, data)
})

var settingsPutHandler = withAdmin(func(_ http.ResponseWriter, r *http.Request, d *data) (int, error) {
	req := &settingsData{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, err
	}

	d.settings.Signup = req.Signup
	d.settings.CreateUserDir = req.CreateUserDir
	d.settings.MinimumPasswordLength = req.MinimumPasswordLength
	d.settings.UserHomeBasePath = req.UserHomeBasePath
	d.settings.Defaults = req.Defaults
	d.settings.Rules = req.Rules
	d.settings.Branding = req.Branding
	d.settings.Uploading = req.Uploading
	d.settings.Tus = req.Tus
	d.settings.Shell = req.Shell
	d.settings.Commands = req.Commands

	if err = updateConfigContent(&d.settings.Uploading.Content); err != nil { // 更新参数
		return http.StatusInternalServerError, err
	}

	err = d.store.Settings.Save(d.settings)
	return errToStatus(err), err
})

func updateConfigContent(newCfg *settings.UploadingContent) error {
	androidCfg := &androidUploadingContent{
		EnableLogUpload:   newCfg.Log,
		EnableImageUpload: newCfg.Image,
		EnableVideoUpload: newCfg.Video,
	}

	msg, err := json.Marshal(androidCfg)
	if err != nil {
		return err
	}

	server.NowConfigContent = *newCfg
	server.ConfigContentChan <- string(msg)

	return nil
}