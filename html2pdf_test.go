package gorestpack

import (
	"os"
	"testing"
)

func Pdf(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x25 && buf[1] == 0x50 &&
		buf[2] == 0x44 && buf[3] == 0x46
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

func Test_HTML2PDF_Capture(t *testing.T) {
	pdfToken := os.Getenv("PDF_TOKEN")

	client := NewHTMLToPDFClient(pdfToken)
	resp, err := client.Capture("https://google.com/")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if resp.Image == "" {
		t.Errorf("Must return pdf url")
	}

}

func Test_HTML2PDF_Capture_404(t *testing.T) {
	pdfToken := os.Getenv("PDF_TOKEN")

	client := NewHTMLToPDFClient(pdfToken)
	_, err := client.Capture("https://google/")

	if err == nil {
		t.Errorf("Must return error")
	}

	if err.Error() != "net::ERR_NAME_NOT_RESOLVED at https://google/" {
		t.Errorf("Must return net::ERR_NAME_NOT_RESOLVED")
	}
}

func Test_HTML2PDF_Capture_Reader(t *testing.T) {
	pdfToken := os.Getenv("PDF_TOKEN")

	client := NewHTMLToPDFClient(pdfToken)
	resp, err := client.CaptureToReader("https://google.com")

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	buffer := make([]byte, 4)
	_, err = resp.Read(buffer)

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if Pdf(buffer) != true {
		t.Errorf("Must return pdf file")
	}
}

func Test_HTML2PDF_Capture_Reader_404(t *testing.T) {
	pdfToken := os.Getenv("PDF_TOKEN")

	client := NewHTMLToPDFClient(pdfToken)
	_, err := client.CaptureToReader("https://google/")

	if err == nil {
		t.Errorf("Must return error")
	}

	if err.Error() != "400 Bad Request" {
		t.Errorf("Must return 400 Bad Request")
	}
}
