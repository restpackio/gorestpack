package gorestpack

import (
	"os"
	"testing"
)

func Test_HTML2PDF_ValidToken(t *testing.T) {
	pdfToken := os.Getenv("PDF_TOKEN")

	client := NewHTMLToPDFClient(pdfToken)
	_, err := client.Capture("https://google.com/")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func Test_HTML2PDF_InvalidToken(t *testing.T) {
	client := NewHTMLToPDFClient("INVALID_TOKEN")
	_, err := client.Capture("https://google.com/")

	if err == nil {
		t.Errorf("Must return error")
	}

	if err.Error() != "The access token is invalid or you are not subscribed to any plan. Please visit the API console and choose your subscription plan." {
		t.Errorf("Must return error with invalid token warning")
	}
}

func Test_HTML2PDF_404(t *testing.T) {
	pdfToken := os.Getenv("PDF_TOKEN")

	client := NewHTMLToPDFClient(pdfToken)
	_, err := client.Capture("https://google/")

	if err == nil {
		t.Errorf("Must return error with net::ERR_NAME_NOT_RESOLVED: %s", err.Error())
	}
}

func Test_HTML2PDF_ValidToken_Render(t *testing.T) {
	pdfToken := os.Getenv("PDF_TOKEN")

	client := NewHTMLToPDFClient(pdfToken)
	_, err := client.CaptureToReader("https://google.com")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func Test_HTML2PDF_ValidToken_Render_404(t *testing.T) {
	pdfToken := os.Getenv("PDF_TOKEN")

	client := NewHTMLToPDFClient(pdfToken)
	_, err := client.CaptureToReader("https://google.coddm")

	if err == nil {
		t.Errorf("Must return error with net::ERR_NAME_NOT_RESOLVED: %s", err.Error())
	}
}
