package sreg

const (
	RegEx_Mail = `^[a-zA-Z0-9][a-zA-Z0-9_.-]*[a-zA-Z0-9]@[a-zA-Z0-9][a-zA-Z0-9.-]*[a-zA-Z0-9]\.[a-zA-Z]{2,}$`
	RegEx_IP   = `((?:(?:25[0-5]|2[0-4]\\d|[01]?\\d?\\d)\\.){3}(?:25[0-5]|2[0-4]\\d|[01]?\\d?\\d))`
	RegEx_ZH   = `[\u4e00-\u9fa5]+`
)
