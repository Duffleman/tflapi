package lines

import (
	"fmt"
)

// ListLinesByMode lists the TFL lines for the specified mode of transport
func (l *LineWrapper) ListLinesByMode(mode string) (lines []*Line, err error) {
	url := fmt.Sprintf("%s/Line/Mode/%s", RootURL, mode)

	err = l.http.GetJSON(url, &lines)

	return
}
