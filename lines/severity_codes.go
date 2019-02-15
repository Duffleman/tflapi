package lines

import (
	"fmt"
)

type SeverityCode struct {
	Name        string `json:"modeName"`
	Level       int    `json:"severityLevel"`
	Description string `json:"description"`
}

// ListSeverityCodes gets a list of valid severity codes
func (l *LineWrapper) ListSeverityCodes() (scs []*SeverityCode, err error) {
	err = l.http.GetJSON(fmt.Sprintf("%s/Line/Meta/Severity", RootURL), &scs)

	return
}
