package sreg

import "regexp"

func Quote(s string) string {
	return regexp.QuoteMeta(s)
}

func Validate(pattern string) error {
	_, err := regexp.Compile(pattern)
	return err
}

func IsMatch(pattern string, src []byte) bool {
	if r, err := regexp.Compile(pattern); err == nil {
		return r.Match(src)
	}
	return false
}

func IsMatchString(pattern string, src string) bool {
	return IsMatch(pattern, []byte(src))
}

func MatchString(pattern string, src string) ([]string, error) {
	if r, err := regexp.Compile(pattern); err == nil {
		return r.FindStringSubmatch(src), nil
	} else {
		return nil, err
	}
}

func MatchAll(pattern string, src []byte) ([][][]byte, error) {
	if r, err := regexp.Compile(pattern); err == nil {
		return r.FindAllSubmatch(src, -1), nil
	} else {
		return nil, err
	}
}

func Replace(pattern string, replace, src []byte) ([]byte, error) {
	if r, err := regexp.Compile(pattern); err == nil {
		return r.ReplaceAll(src, replace), nil
	} else {
		return nil, err
	}
}

func IsMail(str string) bool {
	return IsMatchString(RegEx_Mail, str)
}

func IsIP(str string) bool {
	return IsMatchString(RegEx_IP, str)
}
