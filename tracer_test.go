package traceredirects

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestTrace(t *testing.T) {
	tests := []struct {
		name        string
		handler     func(w http.ResponseWriter, r *http.Request)
		want        []string
		wantError   bool
		errContains string
	}{
		{
			name: "No redirect",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
			want: []string{"/1"},
		},
		{
			name: "Single redirect",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/1" {
					http.Redirect(w, r, "/2", http.StatusFound)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			},
			want: []string{"/1", "/2"},
		},
		{
			name: "Multiple redirects",
			handler: func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/1":
					http.Redirect(w, r, "/2", http.StatusFound)
				case "/2":
					http.Redirect(w, r, "/3", http.StatusFound)
				case "/3":
					w.WriteHeader(http.StatusOK)
				}
			},
			want: []string{"/1", "/2", "/3"},
		},
		{
			name: "Refresh header",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/1" {
					w.Header().Set("Refresh", "0; url=/2")
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			},
			want: []string{"/1", "/2"},
		},
		{
			name: "Loop detection",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/1" {
					http.Redirect(w, r, "/2", http.StatusFound)
				} else {
					http.Redirect(w, r, "/1", http.StatusFound)
				}
			},
			wantError:   true,
			errContains: "too many redirects",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(tt.handler))
			defer ts.Close()

			startURL := ts.URL + "/1"
			got, err := Trace(startURL)

			if tt.wantError {
				if err == nil {
					t.Errorf("Trace() error = nil, wantError %v", tt.wantError)
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("Trace() error = %v, want error containing %v", err, tt.errContains)
				}
				return
			}
			if err != nil {
				t.Errorf("Trace() error = %v, wantError %v", err, tt.wantError)
				return
			}

			var wantURLs []string
			for _, p := range tt.want {
				wantURLs = append(wantURLs, ts.URL+p)
			}

			if !reflect.DeepEqual(got, wantURLs) {
				t.Errorf("Trace() = %v, want %v", got, wantURLs)
			}
		})
	}
}
