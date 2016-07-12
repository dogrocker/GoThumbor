package thumbor

import (
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha1"
)

type Thumbor struct {
  Secret string
  Server string
  imagePath string
  size string
	filters string
}

func NewThumbor(secret string, server string) *Thumbor {
	return &Thumbor{Secret: secret, Server: server}
}

func (t *Thumbor) SetImagePath(imgPath string) *Thumbor{
	t.imagePath = imgPath
	return t
}

func (t *Thumbor) Resize(size string) *Thumbor {
	if len([]rune(size)) > 0 {
		t.size = size + "/"
	}
	return t
}

func (t *Thumbor) Filters(filters string) *Thumbor {
	if len([]rune(filters)) > 0 {
		t.filters = "filters:" + filters + "/"
	}
	return t
}

func (t *Thumbor) BuildUrl() string {
	if t.imagePath == "" {
		return ""
	}
	return t.Server + "/" + t.getSecureKey() + "/" + t.getFullPath()
}

func (t *Thumbor) getSecureKey() string {
	if t.Secret == "" {
		return "unsafe"
	}
	return t.generateKey()
}

func (t *Thumbor) getFullPath() string {
	return t.size + t.filters + t.imagePath
}

func (t *Thumbor) generateKey() string {
	url := []byte(t.getFullPath())
	key := []byte(t.Secret)
	mac := hmac.New(sha1.New, key)
	mac.Write(url)
	expectedMAC := mac.Sum(nil)
	return base64.URLEncoding.EncodeToString(expectedMAC)
}
