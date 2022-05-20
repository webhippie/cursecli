package forge

import (
	"net"
	"net/http"
	"time"

	"github.com/webhippie/cursecli/pkg/manifest"
)

// Option defines a single option function.
type Option func(o *Options)

// Options defines the available options for this package.
type Options struct {
	HTTPClient *http.Client
	Path       string
	APIKey     string
	Manifest   manifest.Manifest
}

// newOptions initializes the available default options.
func newOptions(opts ...Option) Options {
	opt := Options{
		HTTPClient: defaultHTTPClient(),
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// WithHTTPClient provides a function to set the httpclient option.
func WithHTTPClient(v *http.Client) Option {
	return func(o *Options) {
		o.HTTPClient = v
	}
}

// WithPath provides a function to set the destination option.
func WithPath(v string) Option {
	return func(o *Options) {
		o.Path = v
	}
}

// WithAPIKey provides a function to set the apikey option.
func WithAPIKey(v string) Option {
	return func(o *Options) {
		o.APIKey = v
	}
}

// WithManifest provides a function to set the manifest option.
func WithManifest(v manifest.Manifest) Option {
	return func(o *Options) {
		o.Manifest = v
	}
}

func defaultHTTPClient() *http.Client {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}

	return &http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           dialer.DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}
