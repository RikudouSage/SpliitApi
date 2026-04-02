package helper

// Pointer returns a pointer to in.
func Pointer[TType any](in TType) *TType {
	return &in
}
