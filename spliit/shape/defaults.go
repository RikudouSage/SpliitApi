package shape

// RequestDefaults is implemented by request shapes that can fill in default values.
type RequestDefaults interface {
	// ApplyDefaults mutates the request to fill in implicit API defaults.
	ApplyDefaults()
}
