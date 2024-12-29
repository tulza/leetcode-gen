package utils

func GetKey(args []string) string {
	// check if args is empty
	if len(args) == 0 {
		return ""
	}

	// check if the first arg is a flag
	char := string([]rune(args[0])[0])
	if char == "-" {
		return ""
	}

	return args[0]
}