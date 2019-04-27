package gorestpack

import (
	"os"
	"testing"
)

func Test_Screenshot_ValidToken(t *testing.T) {
	token := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(token)
	_, err := client.Capture("https://google.com/")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
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

func Test_Screenshot_404(t *testing.T) {
	token := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(token)
	_, err := client.Capture("https://google/")

	if err == nil {
		t.Errorf("Must return error with net::ERR_NAME_NOT_RESOLVED: %s", err.Error())
	}
}

func Test_Screenshot_ValidToken_Render(t *testing.T) {
	token := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(token)
	_, err := client.CaptureToReader("https://google.com")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func Test_Screenshot_ValidToken_Render_404(t *testing.T) {
	token := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(token)
	_, err := client.CaptureToReader("https://google.coddm")

	if err == nil {
		t.Errorf("Must return error with net::ERR_NAME_NOT_RESOLVED: %s", err.Error())
	}
}

func Test_Screenshot_ValidToken_Image(t *testing.T) {
	token := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(token)
	_, err := client.CaptureToImage("https://google.com")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func Test_Screenshot_ValidToken_Image_404(t *testing.T) {
	token := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(token)
	_, err := client.CaptureToImage("https://google.coddm")

	if err == nil {
		t.Errorf("Must return error with net::ERR_NAME_NOT_RESOLVED: %s", err.Error())
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
	_, err := client.CaptureHTMLToReader("<h1>Test</h1>")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}
