package groupi

func PopDoublon(tab []string) []string {
	cleaned := []string{}

	for _, value := range tab {

		if !stringInSlice(value, cleaned) {
			cleaned = append(cleaned, value)
		}
	}
	return cleaned
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
