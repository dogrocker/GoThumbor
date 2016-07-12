package thumbor

import "testing"

const (
  server    = "www.thumbor-server.com"
  imagePath  = "http://localhost/some/path/to/image.jpg"
  secret    = "test"
  resize    = "100x50"
)

func TestGetImageUrlWithNoSecretKey (t *testing.T) {
  t.Log("Try to get image url with out secret key")
  noSecretKeyThumbor := Thumbor{Server: server}
  expected := "www.thumbor-server.com/unsafe/http://localhost/some/path/to/image.jpg"
  result := noSecretKeyThumbor.SetImagePath(imagePath).BuildUrl()
  if result != expected {
    t.Errorf("Got an unxpected result url: %s != expected url %s", result, expected)
  }
}

func TestGetImageUrlWithNoImagePath (t *testing.T) {
  t.Log("Try to get image url with out image path")
  noSecretKeyThumbor := Thumbor{Server: server}
  expected := ""
  result := noSecretKeyThumbor.SetImagePath("").BuildUrl()
  if result != expected {
    t.Errorf("Got an unxpected result url: %s != expected url %s", result, expected)
  }
}

func TestGetImageUrlWithOutAnyOption (t *testing.T) {
  t.Log("Try to get image url with out any option")
  noSecretKeyThumbor := Thumbor{}
  expected := "/unsafe/http://localhost/some/path/to/image.jpg"
  result := noSecretKeyThumbor.SetImagePath(imagePath).BuildUrl()
  if result != expected {
    t.Errorf("Got an unxpected result url: %s != expected url %s", result, expected)
  }
}

var thumbor = Thumbor{Server: server, Secret: secret}

func TestGetImageUrlWithKey (t *testing.T) {
  t.Log("Try to get image url with secret key")
  expected := "www.thumbor-server.com/1fSvqEDFBolBrQRPM3lreU05GEQ=/http://localhost/some/path/to/image.jpg"
  result := thumbor.SetImagePath(imagePath).BuildUrl()
  if result != expected {
    t.Errorf("Got an unxpected result url: %s != expected url %s", result, expected)
  }
}

func TestGetImageUrlWithKeyAndWidthAndHeight (t *testing.T) {
  t.Log("Try to get image url with secret key and weight and height")
  expected := "www.thumbor-server.com/SQ2o2zwO7zR4s7-qBbxcrjM4Spk=/100x50/http://localhost/some/path/to/image.jpg"
  result := thumbor.SetImagePath(imagePath).Resize(resize).BuildUrl()
  if result != expected {
    t.Errorf("Got an unxpected result url: %s != expected url %s", result, expected)
  }
}

func TestGetImageUrlWithKeyAndWidthAndHeightAndFilters (t *testing.T) {
  t.Log("Try to get image url with secret key and weight height and filters")
  filters := "blur(50)"
  expected := "www.thumbor-server.com/CiDgnPFBBnAnduwlzNtEZ_xIIQc=/100x50/filters:blur(50)/http://localhost/some/path/to/image.jpg"
  result := thumbor.SetImagePath(imagePath).Resize(resize).Filters(filters).BuildUrl()
  if result != expected {
    t.Errorf("Got an unxpected result url: %s != expected url %s", result, expected)
  }
}

func TestGetImageUrlWithKeyAndWidthAndHeightAndMoreFilters (t *testing.T) {
  t.Log("Try to get image url with secret key and weight height and more filters")
  filters := "blur(50):rotate(90):saturation(0.20)"
  expected := "www.thumbor-server.com/9cm0XbdNUM6FFfK69VhOJ61U1To=/100x50/filters:blur(50):rotate(90):saturation(0.20)/http://localhost/some/path/to/image.jpg"
  result := thumbor.SetImagePath(imagePath).Resize(resize).Filters(filters).BuildUrl()
  if result != expected {
    t.Errorf("Got an unxpected result url: %s != expected url %s", result, expected)
  }
}
