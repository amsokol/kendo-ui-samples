package kendoui

import (
	"net/url"
	"strings"
)

// filter[filters][0][field]=Title
// filter[filters][0][ignoreCase]=true
// filter[filters][0][operator]=startswith
// filter[filters][0][value]=the
// filter[logic]=and

type RequestParams map[string]interface{}

func ParseParams(vs url.Values) RequestParams {
	params := make(RequestParams)

	for key, value := range vs {
		p := params

		keys := strings.Split(key, "[")
		len := len(keys)
		key = keys[0]

		for i := 1; i < len; i++ {
			p1, ok := p[key]
			if !ok {
				p1 = make(RequestParams)
				p[key] = p1
			}
			p = (p1).(RequestParams)

			key = strings.TrimRight(keys[i], "]")
		}
		p[key] = value
	}

	return params
}
