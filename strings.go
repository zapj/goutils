package goutils

func IsUpperASCII(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func IsLowerASCII(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func ToUpperASCII(c byte) byte {
	if IsLowerASCII(c) {
		return c - ('a' - 'A')
	}
	return c
}

func ToLowerASCII(c byte) byte {
	if IsUpperASCII(c) {
		return c + 'a' - 'A'
	}
	return c
}
