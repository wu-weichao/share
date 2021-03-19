package funcs

import "time"

func UnixToDateTime(u interface{}) string {
	var t int64
	switch v := u.(type) {
	case int:
		t = int64(v)
	case int32:
		t = int64(v)
	case int64:
		t = v
	default:
		return ""
	}
	if t == 0 {
		return ""
	}
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

func UnixToDate(u interface{}) string {
	var t int64
	switch v := u.(type) {
	case int:
		t = int64(v)
	case int32:
		t = int64(v)
	case int64:
		t = v
	default:
		return ""
	}
	if t == 0 {
		return ""
	}
	return time.Unix(t, 0).Format("2006-01-02")
}

func UnixToFormat(u interface{}, format string) string {
	var t int64
	switch v := u.(type) {
	case int:
		t = int64(v)
	case int32:
		t = int64(v)
	case int64:
		t = v
	default:
		return ""
	}
	if t == 0 {
		return ""
	}
	return time.Unix(t, 0).Format(format)
}
