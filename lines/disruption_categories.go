package lines

import (
	"fmt"
)

// ListDisruptionCategories gets a list of valid disruption categories
func (l *LineWrapper) ListDisruptionCategories() (codes []string, err error) {
	err = l.http.GetJSON(fmt.Sprintf("%s/Line/Meta/DisruptionCategories", RootURL), &codes)

	return
}
