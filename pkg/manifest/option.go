package manifest

// Option defines a single option function.
type Option func(o *Options)

// Options defines the available options for this package.
type Options struct {
	Path string
}

// newOptions initializes the available default options.
func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// WithPath provides a function to set the path option.
func WithPath(v string) Option {
	return func(o *Options) {
		o.Path = v
	}
}
