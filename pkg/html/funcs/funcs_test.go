package funcs

import (
	"testing"
	"time"
)

func TestUnixToDate(t *testing.T) {
	now := time.Now().Unix()
	t.Log(UnixToDate(now))
	t.Log(UnixToDateTime(now))
	t.Log(UnixToFormat(now, "2006-01-02 15:04"))
}
