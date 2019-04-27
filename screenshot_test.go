package gorestpack

import (
	"os"
	"testing"
)

func Test_Screenshot_ValidToken(t *testing.T) {
	pdfToken := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(pdfToken)
	_, err := client.Capture("https://google.com/")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func Test_Screenshot_InvalidToken(t *testing.T) {
	client := NewScreenshotClient("INVALID_TOKEN")
	_, err := client.Capture("https://google.com/")

	if err == nil {
		t.Errorf("Must return error with invalid token")
	}
}

func Test_Screenshot_404(t *testing.T) {
	pdfToken := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(pdfToken)
	_, err := client.Capture("https://google/")

	if err == nil {
		t.Errorf("Must return error with net::ERR_NAME_NOT_RESOLVED: %s", err.Error())
	}
}

func Test_Screenshot_ValidToken_Render(t *testing.T) {
	pdfToken := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(pdfToken)
	_, err := client.CaptureToReader("https://google.com")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func Test_Screenshot_ValidToken_Render_404(t *testing.T) {
	pdfToken := os.Getenv("SS_TOKEN")

	client := NewScreenshotClient(pdfToken)
	_, err := client.CaptureToReader("https://google.coddm")

	if err == nil {
		t.Errorf("Must return error with net::ERR_NAME_NOT_RESOLVED: %s", err.Error())
	}
}
