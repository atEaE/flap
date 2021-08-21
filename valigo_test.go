package valigo_test

// stringPtr return string pointer.
func stringPtr(s string) *string {
	return &s
}

// stringSlicePtr return string pointer.
func stringSlicePtr(s []string) *[]string {
	return &s
}

// intPtr return integer pointer.
func intPtr(i int) *int {
	return &i
}

// int64Ptr return integer pointer.
func int64Ptr(i int64) *int64 {
	return &i
}

// int32Ptr return integer pointer.
func int32Ptr(i int32) *int32 {
	return &i
}

// float64Ptr return float64 pointer.
func float64Ptr(f float64) *float64 {
	return &f
}

// float32Ptr return float64 pointer.
func float32Ptr(f float32) *float32 {
	return &f
}
