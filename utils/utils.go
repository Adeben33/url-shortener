package utils

func Base62Converter(num uint64) string {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyz"
	charsetLen := len(charset)
	if num == 0 {
		return "0"
	}
	var result []byte
	for num > 0 {
		rem := num % uint64(charsetLen)
		result = append(result, charset[rem])
		num /= uint64(charsetLen)
	}
	return string(result)
}
