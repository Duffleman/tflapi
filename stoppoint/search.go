package stoppoint

import (
	"fmt"
	"net/url"
)

type SearchParams struct {
	Modes                               []string
	FaresOnly                           *bool
	MaxResults                          *int
	Lines                               []string
	IncludeHubs                         *bool
	TFLOperatedNationalRailStationsOnly *bool
}

type SearchResult struct {
	Query    string   `json:"query"`
	From     *int     `json:"from,omitempty"`
	Page     *int     `json:"page,omitempty"`
	PageSize *int     `json:"pageSize,omitempty"`
	Provider *string  `json:"provider,omitempty"`
	Total    int      `json:"total"`
	Matches  []*Match `json:"matches"`
	MaxScore *int     `json:"maxScore,omitempty"`
}

type Match struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	ICSID string   `json:"icsId"`
	Modes []string `json:"modes"`
	Zone  *string  `json:"zone,omitempty"`
	Lat   float64  `json:"lat"`
	Lng   float64  `json:"lon"`
}

func (s *StopPointWrapper) Search(query string, params *SearchParams) (res *SearchResult, err error) {
	q := url.Values{
		"query": []string{query},
	}

	if params != nil {
		if params.Modes != nil && len(params.Modes) > 0 {
			for _, mode := range params.Modes {
				q.Add("modes[]", mode)
			}
		}

		if params.FaresOnly != nil {
			q.Add("faresOnly", fmt.Sprintf("%t", *params.FaresOnly))
		}

		if params.MaxResults != nil {
			q.Add("maxResults", fmt.Sprintf("%d", *params.MaxResults))
		}

		if params.Lines != nil && len(params.Lines) > 0 {
			for _, line := range params.Lines {
				q.Add("lines[]", line)
			}
		}

		if params.IncludeHubs != nil {
			q.Add("includeHubs", fmt.Sprintf("%t", *params.IncludeHubs))
		}
	}

	url := fmt.Sprintf("%s/StopPoint/Search?%s", RootURL, q.Encode())

	fmt.Println(url)

	err = s.http.GetJSON(url, &res)

	return
}
