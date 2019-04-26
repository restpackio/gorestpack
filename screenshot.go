package gorestpack

import (
	"bytes"
	"image"
	"io"

	_ "image/jpeg"
	_ "image/png"

	"github.com/eknkc/request"
)

// Create a new Screenshot Client with supplied restpack.io access key
func NewScreenshotClient(accessToken string) ScreenshotClient {
	return &screenshotClient{
		client: &client{
			httpClient:  request.New(),
			accessToken: accessToken,
			basePath:    "https://restpack.io/api/screenshot/v3",
		},
	}
}

// Options supplied to the Restpack Screenshot API for conversion
type ScreenshotCaptureOptions struct {
	// Capturing mode.
	Mode string `json:"mode,omitempty"`
	// Preferred image output format. If you need a raw html string you can pass html as format
	Format string `json:"format,omitempty"`
	// Preferred viewport width in pixels.
	Width int `json:"width,omitempty"`
	// Preferred viewport height in pixels.
	Height int `json:"height,omitempty"`
	// In case you want a thumbnail image, provide a preferred width
	ThumbnailWidth int `json:"thumbnail_width,omitempty"`
	// Preferred thumbnail height, requires thumbnail_width to be set, unbounded if omitted.
	ThumbnailHeight int `json:"thumbnail_height,omitempty"`
	// Additional CSS string to be injected into the page before render.
	CSS string `json:"css,omitempty"`
	// Additional JS string to be injected into the page before render.
	JS string `json:"js,omitempty"`
	// Time in milliseconds to delay capture after page load
	Delay int `json:"delay,omitempty"`
	// Time in seconds for the resulting image to be cached for further requests.
	CacheTTL int `json:"cache_ttl,omitempty"`
	// Custom user-agent header string for the web request.
	UserAgent string `json:"user_agent,omitempty"`
	// Custom accept-language header string for the web request.
	AcceptLanguage string `json:"accept_language,omitempty"`
	// A CSS selector to be used with element rendering mode.
	ElementSelector string `json:"element_selector,omitempty"`
	// Generate retina sized screen capture (2x device pixel ratio).
	Retina bool `json:"retina,omitempty"`
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
	// Ensure that the captured document does not get cached / stored for further use
	Privacy bool `json:"privacy,omitempty"`
	// If specified, ensures that the resulting file is saved with the given name.
	Filename string `json:"filename,omitempty"`
	//Removes the ads on the page
	BlockAds bool `json:"block_ads,omitempty"`
	//Block / hide European Union cookie warnings before capture.
	BlockCookieWarnings bool `json:"block_cookie_warnings,omitempty"`
	// Do not render with default white background. You can use this option to generate transparent PNG images
	OmitBackground bool `json:"omit_background,omitempty"`
}

type screenshotCallOptions struct {
	ScreenshotCaptureOptions
	JSON bool   `json:"json,omitempty"`
	URL  string `json:"url,omitempty"`
	HTML string `json:"html,omitempty"`
}

// Capture result from screenshot API
type ScreenshotCaptureResult struct {
	Image        string `json:"image,omitempty"`
	Width        string `json:"width,omitempty"`
	Height       string `json:"height,omitempty"`
	RemoteStatus string `json:"remote_status,omitempty"`
	Cached       bool   `json:"cached,omitempty"`
	URL          string `json:"url,omitempty"`
}

// Restpack Screenshot API Client
type ScreenshotClient interface {
	// Capture a URL and return the information & cdn url
	Capture(url string, options ...ScreenshotCaptureOptions) (ScreenshotCaptureResult, error)
	// Capture a HTML snippet and return the information & cdn url
	CaptureHTML(url string, options ...ScreenshotCaptureOptions) (ScreenshotCaptureResult, error)

	// Capture a URL and return the image
	CaptureToImage(url string, options ...ScreenshotCaptureOptions) (image.Image, error)
	// Capture a HTML snippet and return the information & cdn url
	CaptureHTMLToImage(url string, options ...ScreenshotCaptureOptions) (image.Image, error)

	// Capture a URL and return a reader for resulting image
	CaptureToReader(url string, options ...ScreenshotCaptureOptions) (io.Reader, error)
	// Capture a HTML snippet and returna a reader for resulting image
	CaptureHTMLToReader(url string, options ...ScreenshotCaptureOptions) (io.Reader, error)
}

type screenshotClient struct {
	*client
}

func (me *screenshotClient) Capture(url string, options ...ScreenshotCaptureOptions) (ScreenshotCaptureResult, error) {
	opt := screenshotCallOptions{
		URL:  url,
		JSON: true,
	}

	if len(options) > 0 {
		opt.ScreenshotCaptureOptions = options[0]
	}

	var res ScreenshotCaptureResult
	_, _, err := me.do("POST", "/capture").JSON(opt).EndStruct(&res)
	return res, err
}

func (me *screenshotClient) CaptureHTML(html string, options ...ScreenshotCaptureOptions) (ScreenshotCaptureResult, error) {
	opt := screenshotCallOptions{
		HTML: html,
		JSON: true,
	}

	if len(options) > 0 {
		opt.ScreenshotCaptureOptions = options[0]
	}

	var res ScreenshotCaptureResult
	_, _, err := me.do("POST", "/capture").JSON(opt).EndStruct(&res)
	return res, err
}

func (me *screenshotClient) CaptureToImage(url string, options ...ScreenshotCaptureOptions) (image.Image, error) {
	opt := screenshotCallOptions{
		URL:  url,
		JSON: false,
	}

	if len(options) > 0 {
		opt.ScreenshotCaptureOptions = options[0]
	}

	_, body, err := me.do("POST", "/capture").JSON(opt).End()

	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(body))

	return img, err
}

func (me *screenshotClient) CaptureHTMLToImage(html string, options ...ScreenshotCaptureOptions) (image.Image, error) {
	opt := screenshotCallOptions{
		HTML: html,
		JSON: false,
	}

	if len(options) > 0 {
		opt.ScreenshotCaptureOptions = options[0]
	}

	_, body, err := me.do("POST", "/capture").JSON(opt).End()

	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(body))

	return img, err
}

func (me *screenshotClient) CaptureToReader(url string, options ...ScreenshotCaptureOptions) (io.Reader, error) {
	opt := screenshotCallOptions{
		URL:  url,
		JSON: false,
	}

	if len(options) > 0 {
		opt.ScreenshotCaptureOptions = options[0]
	}

	_, body, err := me.do("POST", "/capture").JSON(opt).End()

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(body), err
}

func (me *screenshotClient) CaptureHTMLToReader(html string, options ...ScreenshotCaptureOptions) (io.Reader, error) {
	opt := screenshotCallOptions{
		HTML: html,
		JSON: false,
	}

	if len(options) > 0 {
		opt.ScreenshotCaptureOptions = options[0]
	}

	_, body, err := me.do("POST", "/capture").JSON(opt).End()

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(body), err
}
