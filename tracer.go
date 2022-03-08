package traceredirects

import (
	"net/http"
	"strings"
)

type RedirectHeaderType int

const (
	Plain RedirectHeaderType = iota
	DelayURL
)

type RedirectHeader struct {
	Header string
	Type   RedirectHeaderType
}

func Trace(u string) ([]string, error) {
	next := u
	var r []string
	for len(next) > 0 {
		r = append(r, next)
		r, err := http.Get(next)
		if err != nil {
			return nil, err
		}
		for _, header := range []*RedirectHeader{
			{Header: "Location"},
			{Header: "Redirect"},
			{Header: "Refresh", Type: DelayURL},
		} {
			switch header.Type {
			case DelayURL:
				s := strings.Split(r.Header.Get(header.Header), ";")
				if len(s) > 1 {
					next = strings.TrimPrefix(strings.TrimSpace(s[1]), "URL=")
				}
			default:
				next = r.Header.Get(header.Header)
			}
			if len(next) > 0 {
				break
			}
		}
	}
	return r, nil
}
