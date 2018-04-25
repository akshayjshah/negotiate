package negotiate

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func fake(accept string) *http.Request {
	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	r.Header.Set("Accept", accept)
	return r
}

func Test(t *testing.T) {
	offers := []string{"text/plain", "text/html", "application/json"}
	tests := []struct {
		desc   string
		accept string
		expect string
	}{
		{"exact match", "application/json", "application/json"},
		{"prefer earlier", "text/*", "text/plain"},
		{"accept anything", "*/*", "text/plain"},
		{"empty accept", "", "text/plain"},
		{"honor weights", "text/html, text/plain;q=0.9, */*;q=0.8", "text/html"},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := ContentType(fake(tt.accept), offers)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.expect {
				t.Fatalf("failed: got %s, expected %s", got, tt.expect)
			}
		})
	}
}

func TestErrors(t *testing.T) {
	tests := []struct {
		desc      string
		accept    string
		offers    []string
		expect    string
		isNoMatch bool
	}{
		{"spaces in offer", "*/*", []string{"text/plain "}, "invalid offer: text/plain ", false},
		{"trailing slash", "*/*/", []string{"text/plain/"}, "invalid offer: text/plain/", false},
		{"no subtype", "*/*", []string{"text"}, "imprecise offer: text", false},
		{"imprecise offer", "*/*", []string{"text/*"}, "imprecise offer: text/*", false},
		{"no match", "text/*", []string{"application/json"}, "no matching offer", true},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			_, err := ContentType(fake(tt.accept), tt.offers)
			if err == nil {
				t.Fatalf("expected error %s, got nil", tt.expect)
			}
			if got := err.Error(); got != tt.expect {
				t.Fatalf("error mismatch: got %s, expected %s", got, tt.expect)
			}
			if tt.isNoMatch != IsNoMatch(err) {
				t.Fatalf("expected IsNoErr to be %v", tt.isNoMatch)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	r := fake("text/html, text/plain;q=0.9, text/*;q=0.8, */*;q=0.2")
	offers := []string{"text/html", "text/plain", "application/json"}
	for n := 0; n < b.N; n++ {
		ContentType(r, offers)
	}
}
