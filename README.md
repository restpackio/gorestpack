

# gorestpack
`import "github.com/restpackio/gorestpack"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
This package provides easy access to <a href="http://restpack.io">http://restpack.io</a> API services from Go applications




## <a name="pkg-index">Index</a>
* [type HTMLToPDFCaptureOptions](#HTMLToPDFCaptureOptions)
* [type HTMLToPDFCaptureResult](#HTMLToPDFCaptureResult)
* [type HTMLToPDFClient](#HTMLToPDFClient)
  * [func NewHTMLToPDFClient(accessToken string) HTMLToPDFClient](#NewHTMLToPDFClient)
* [type ScreenshotCaptureOptions](#ScreenshotCaptureOptions)
* [type ScreenshotCaptureResult](#ScreenshotCaptureResult)
* [type ScreenshotClient](#ScreenshotClient)
  * [func NewScreenshotClient(accessToken string) ScreenshotClient](#NewScreenshotClient)


#### <a name="pkg-files">Package files</a>
[client.go](/src/github.com/restpackio/gorestpack/client.go) [html2pdf.go](/src/github.com/restpackio/gorestpack/html2pdf.go) [screenshot.go](/src/github.com/restpackio/gorestpack/screenshot.go) 






## <a name="HTMLToPDFCaptureOptions">type</a> [HTMLToPDFCaptureOptions](/src/target/html2pdf.go?s=446:2124#L22)
``` go
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
```
Options supplied to the Restpack Screenshot API for conversion










## <a name="HTMLToPDFCaptureResult">type</a> [HTMLToPDFCaptureResult](/src/target/html2pdf.go?s=2337:2659#L63)
``` go
type HTMLToPDFCaptureResult struct {
    Image        string `json:"image,omitempty"`
    Width        string `json:"width,omitempty"`
    Height       string `json:"height,omitempty"`
    RemoteStatus string `json:"remote_status,omitempty"`
    Cached       bool   `json:"cached,omitempty"`
    URL          string `json:"url,omitempty"`
}
```
Capture result from screenshot API










## <a name="HTMLToPDFClient">type</a> [HTMLToPDFClient](/src/target/html2pdf.go?s=2695:3325#L73)
``` go
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
```
Restpack Screenshot API Client







### <a name="NewHTMLToPDFClient">func</a> [NewHTMLToPDFClient](/src/target/html2pdf.go?s=147:206#L11)
``` go
func NewHTMLToPDFClient(accessToken string) HTMLToPDFClient
```
Create a new Screenshot Client with supplied restpack.io access key





## <a name="ScreenshotCaptureOptions">type</a> [ScreenshotCaptureOptions](/src/target/screenshot.go?s=492:2773#L26)
``` go
type ScreenshotCaptureOptions struct {
    // Force rendering a new screenshot disregarding the cache status.
    Fresh bool `json:"fresh,omitempty"`
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
    // Time in milliseconds for the resulting image to be cached for further requests.
    TTL int `json:"ttl,omitempty"`
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
}
```
Options supplied to the Restpack Screenshot API for conversion










## <a name="ScreenshotCaptureResult">type</a> [ScreenshotCaptureResult](/src/target/screenshot.go?s=2988:3311#L77)
``` go
type ScreenshotCaptureResult struct {
    Image        string `json:"image,omitempty"`
    Width        string `json:"width,omitempty"`
    Height       string `json:"height,omitempty"`
    RemoteStatus string `json:"remote_status,omitempty"`
    Cached       bool   `json:"cached,omitempty"`
    URL          string `json:"url,omitempty"`
}
```
Capture result from screenshot API










## <a name="ScreenshotClient">type</a> [ScreenshotClient](/src/target/screenshot.go?s=3347:4268#L87)
``` go
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
```
Restpack Screenshot API Client







### <a name="NewScreenshotClient">func</a> [NewScreenshotClient](/src/target/screenshot.go?s=188:249#L15)
``` go
func NewScreenshotClient(accessToken string) ScreenshotClient
```
Create a new Screenshot Client with supplied restpack.io access key









- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
