package endpoint

// Endpoint defines the input and output shapes for a TRPC endpoint.
// Implementations should return zero values for InputShape and OutputShape.
type Endpoint[TInput any, TOutput any] interface {
	// Name returns the remote endpoint name.
	Name() string
	// InputShape returns the zero value of the request DTO.
	InputShape() TInput
	// OutputShape returns the zero value of the response DTO.
	OutputShape() TOutput
	// Mutates reports whether the endpoint must be sent as a mutating request.
	Mutates() bool
}
