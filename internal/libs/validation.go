package libs

// 校验是否为空
func ValidateEmpty(value string) bool {
	return value == ""
}

// 校验是否为nil
func ValidateNil(value interface{}) bool {
	return value == nil
}
