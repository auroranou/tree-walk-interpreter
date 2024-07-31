package scan

func isAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		char == '_'
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func isAlphaNumeric(char rune) bool {
	return isAlpha(char) || isDigit(char)
}
