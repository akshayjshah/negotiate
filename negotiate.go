// Package negotiate provides a function for HTTP content type negotiation.
//
// Once https://github.com/golang/go/issues/19307 is resolved, this package
// will no longer be maintained.
package negotiate

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang/gddo/httputil"
)

const fakeType = "not/real"

var errNoMatch = errors.New("no matching offer")

// ContentType returns the best offered content type for the request's Accept
// header. Offers must include both a MIME type and subtype, e.g. text/plain
// or image/png. If two offers match with equal weight and specificity, then
// the offer earlier in the list is preferred. If no offers match, an error is
// returned.
func ContentType(r *http.Request, offers []string) (string, error) {
	for _, o := range offers {
		if err := checkOffer(o); err != nil {
			return "", err
		}
	}
	if t := httputil.NegotiateContentType(r, offers, fakeType); t != fakeType {
		return t, nil
	}
	return "", errNoMatch
}

func checkOffer(offer string) error {
	if strings.Contains(offer, " ") {
		return fmt.Errorf("invalid offer: %s", offer)
	}
	slashes := strings.Count(offer, "/")
	stars := strings.Count(offer, "*")
	if slashes > 1 {
		return fmt.Errorf("invalid offer: %s", offer)
	}
	if slashes < 1 || stars > 0 {
		return fmt.Errorf("imprecise offer: %s", offer)
	}
	return nil
}
