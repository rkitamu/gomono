package lib

func TrimExclamation(s string) string {
	if len(s) > 0 && s[0] == '!' {
		return s[1:]
	}
	return s
}
