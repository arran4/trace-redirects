package traceredirects

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
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
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 30 * time.Second,
	}
	for len(next) > 0 {
		if len(r) > 20 {
			return nil, fmt.Errorf("too many redirects")
		}
		r = append(r, next)
		resp, err := client.Get(next)
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
				s := strings.Split(resp.Header.Get(header.Header), ";")
				if len(s) > 1 {
					val := strings.TrimSpace(s[1])
					if strings.HasPrefix(strings.ToLower(val), "url=") {
						next = val[4:]
					}
				}
			default:
				next = resp.Header.Get(header.Header)
			}
			if len(next) > 0 {
				nextURL, err := url.Parse(next)
				if err == nil {
					next = resp.Request.URL.ResolveReference(nextURL).String()
				}
				break
			}
		}
		resp.Body.Close()
	}
	return r, nil
}
