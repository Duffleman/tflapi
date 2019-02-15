package lines

import (
	"fmt"
)

// ListServiceTypes gets a list of valid ServiceTypes to filter on
func (l *LineWrapper) ListServiceTypes() (codes []string, err error) {
	err = l.http.GetJSON(fmt.Sprintf("%s/Line/Meta/ServiceTypes", RootURL), &codes)

	return
}
