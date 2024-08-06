package utils

func ErrorsToStrings(errors []error) []string {
	var errorMessages []string
	for _, err := range errors {
		errorMessages = append(errorMessages, err.Error())
	}
	return errorMessages
}
