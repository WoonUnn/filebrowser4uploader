package settings

type UploadingContent struct {
	Log   bool `json:"log"`
	Image bool `json:"image"`
	Video bool `json:"video"`
}

type Uploading struct {
	Method   string `json:"method"` // "periodic" or "onetime"
	Content  UploadingContent `json:"content"`
}