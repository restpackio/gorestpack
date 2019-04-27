package gorestpack

import (
	"os"
	"testing"
)

func Png(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x89 && buf[1] == 0x50 &&
		buf[2] == 0x4E && buf[3] == 0x47
}

func Test_Screenshot_InvalidToken(t *testing.T) {
	client := NewScreenshotClient("INVALID_TOKEN")
	_, err := client.Capture("https://google.com/")

	if err == nil {
		t.Errorf("Must return error")
	}

	if err.Error() != "The access token is invalid or you are not subscribed to any plan. Please visit the API console and choose your subscription plan." {
		t.Errorf("Must return error with invalid token warning")
	}
}

func Test_Screenshot_Capture(t *testing.T) {
	token := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(token)
	resp, err := client.Capture("https://google.com/")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if resp.Image == "" {
		t.Errorf("Must return screenshot image: get empty")
	}
}

func Test_Screenshot_Capture_404(t *testing.T) {
	token := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(token)
	_, err := client.Capture("https://google/")

	if err == nil {
		t.Errorf("Must return error")
	}

	if err.Error() != "net::ERR_NAME_NOT_RESOLVED at https://google/" {
		t.Errorf("Must return net::ERR_NAME_NOT_RESOLVED")
	}
}

func Test_Screenshot_CaptureToReader(t *testing.T) {
	token := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(token)
	resp, err := client.CaptureToReader("https://google.com")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	buffer := make([]byte, 4)
	_, err = resp.Read(buffer)

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if Png(buffer) != true {
		t.Errorf("Must return png file")
	}
}

func Test_Screenshot_CaptureToReader_404(t *testing.T) {
	token := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(token)
	_, err := client.CaptureToReader("https://google")

	if err == nil {
		t.Errorf("Must return error")
	}

	if err.Error() != "400 Bad Request" {
		t.Errorf("Must return 400 Bad Request")
	}
}

func Test_Screenshot_Image(t *testing.T) {
	token := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(token)
	_, err := client.CaptureToImage("https://google.com")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func Test_Screenshot_Image_404(t *testing.T) {
	token := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(token)
	_, err := client.CaptureToImage("https://google.coddm")

	if err == nil {
		t.Errorf("Must return error")
	}

	if err.Error() != "400 Bad Request" {
		t.Errorf("Must return 400 Bad Request")
	}
}

func Test_Screenshot_Capture_HTML(t *testing.T) {
	token := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(token)
	resp, err := client.CaptureHTML("<h1>Test</h1>")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if resp.URL != "" {
		t.Errorf("Url must be empty, get: %s", resp.URL)
	}
}

func Test_Screenshot_Capture_HTML_Image(t *testing.T) {
	token := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(token)
	_, err := client.CaptureHTMLToImage("<h1>Test</h1>")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func Test_Screenshot_Capture_HTML_Reader(t *testing.T) {
	token := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(token)
	resp, err := client.CaptureHTMLToReader("<h1>Test</h1>")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	buffer := make([]byte, 4)
	_, err = resp.Read(buffer)

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if Png(buffer) != true {
		t.Errorf("Must return png file")
	}
}
