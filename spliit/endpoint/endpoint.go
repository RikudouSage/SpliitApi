package endpoint

// Endpoint defines the input and output shapes for a TRPC endpoint.
// Implementations should return zero values for InputShape and OutputShape.
type Endpoint[TInput any, TOutput any] interface {
	Name() string
	InputShape() TInput
	OutputShape() TOutput
	Mutates() bool
}
