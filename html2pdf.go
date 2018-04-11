package gorestpack

import (
	"bytes"
	"io"

	"github.com/eknkc/request"
)

// Create a new Screenshot Client with supplied restpack.io access key
func NewHTMLToPDFClient(accessToken string) HTMLToPDFClient {
	return &htmlToPDFClient{
		client: &client{
			httpClient:  request.New(),
			accessToken: accessToken,
			basePath:    "https://restpack.io/api/html2pdf/v3",
		},
	}
}

// Options supplied to the Restpack Screenshot API for conversion
type HTMLToPDFCaptureOptions struct {
	// Force rendering a new pdf disregarding the cache status.
	Fresh bool `json:"fresh,omitempty"`
	// Custom page size for created document
	PDFPage string `json:"pdf_page,omitempty"`
	// CSS style margin sizes.
	PDFMargins string `json:"pdf_margins,omitempty"`
	// Page Orientation
	PDFOrientation string `json:"pdf_orientation,omitempty"`
	// Additional CSS string to be injected into the page before render.
	CSS string `json:"css,omitempty"`
	// Additional JS string to be injected into the page before render.
	JS string `json:"js,omitempty"`
	// Time in milliseconds to delay capture after page load
	Delay int `json:"delay,omitempty"`
	// Time in milliseconds for the resulting image to be cached for further requests.
	TTL int `json:"ttl,omitempty"`
	// Custom user-agent header string for the web request.
	UserAgent string `json:"user_agent,omitempty"`
	// Custom accept-language header string for the web request.
	AcceptLanguage string `json:"accept_language,omitempty"`
	// Additional headers seperated with newline
	Headers string `json:"headers,omitempty"`
	// Force CSS media emulation for print or screen.
	EmulateMedia string `json:"emulate_media,omitempty"`
	// By default, any response from remote server outside http 200-299 status codes generates an error. If you wish to capture error pages, pass true.
	AllowFailed bool `json:"allow_failed,omitempty"`
	// Wait until window load event fires or network becomes idle before capturing the page.
	Wait string `json:"wait,omitempty"`
	// Wait until a DOM element matching the provided css selector becomes present on the page.
	Shutter string `json:"shutter,omitempty"`
}

type htmlToPDFCallOptions struct {
	HTMLToPDFCaptureOptions
	JSON bool   `json:"json,omitempty"`
	URL  string `json:"url,omitempty"`
	HTML string `json:"html,omitempty"`
}

// Capture result from screenshot API
type HTMLToPDFCaptureResult struct {
	Image        string `json:"image,omitempty"`
	Width        string `json:"width,omitempty"`
	Height       string `json:"height,omitempty"`
	RemoteStatus string `json:"remote_status,omitempty"`
	Cached       bool   `json:"cached,omitempty"`
	URL          string `json:"url,omitempty"`
}

// Restpack Screenshot API Client
type HTMLToPDFClient interface {
	// Capture a URL and return the information & cdn url
	Capture(url string, options ...HTMLToPDFCaptureOptions) (HTMLToPDFCaptureResult, error)
	// Capture a HTML snippet and return the information & cdn url
	CaptureHTML(url string, options ...HTMLToPDFCaptureOptions) (HTMLToPDFCaptureResult, error)

	// Capture a URL and return a reader for resulting pdf
	CaptureToReader(url string, options ...HTMLToPDFCaptureOptions) (io.Reader, error)
	// Capture a HTML snippet and returna a reader for resulting pdf
	CaptureHTMLToReader(url string, options ...HTMLToPDFCaptureOptions) (io.Reader, error)
}

type htmlToPDFClient struct {
	*client
}

func (me *htmlToPDFClient) Capture(url string, options ...HTMLToPDFCaptureOptions) (HTMLToPDFCaptureResult, error) {
	opt := htmlToPDFCallOptions{
		URL:  url,
		JSON: true,
	}

	if len(options) > 0 {
		opt.HTMLToPDFCaptureOptions = options[0]
	}

	var res HTMLToPDFCaptureResult
	_, _, err := me.do("POST", "/convert").JSON(opt).EndStruct(&res)
	return res, err
}

func (me *htmlToPDFClient) CaptureHTML(html string, options ...HTMLToPDFCaptureOptions) (HTMLToPDFCaptureResult, error) {
	opt := htmlToPDFCallOptions{
		HTML: html,
		JSON: true,
	}

	if len(options) > 0 {
		opt.HTMLToPDFCaptureOptions = options[0]
	}

	var res HTMLToPDFCaptureResult
	_, _, err := me.do("POST", "/convert").JSON(opt).EndStruct(&res)
	return res, err
}

func (me *htmlToPDFClient) CaptureToReader(url string, options ...HTMLToPDFCaptureOptions) (io.Reader, error) {
	opt := htmlToPDFCallOptions{
		URL:  url,
		JSON: false,
	}

	if len(options) > 0 {
		opt.HTMLToPDFCaptureOptions = options[0]
	}

	_, body, err := me.do("POST", "/convert").JSON(opt).End()

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(body), err
}

func (me *htmlToPDFClient) CaptureHTMLToReader(html string, options ...HTMLToPDFCaptureOptions) (io.Reader, error) {
	opt := htmlToPDFCallOptions{
		HTML: html,
		JSON: false,
	}

	if len(options) > 0 {
		opt.HTMLToPDFCaptureOptions = options[0]
	}

	_, body, err := me.do("POST", "/convert").JSON(opt).End()

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(body), err
}
