package helper

func Pointer[TType any](in TType) *TType {
	return &in
}
