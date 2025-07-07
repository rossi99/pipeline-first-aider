package logic

// failOnNil takes a pointer to int, tries to set it,
// and returns the dereferenced value.
// If you pass in nil, this will panic.
func failOnNil(ptr *int) int {
	*ptr = 42   // panic if ptr == nil
	return *ptr // never reached when ptr is nil
}

func returnValue(num int32) int32 {
	return num
}
