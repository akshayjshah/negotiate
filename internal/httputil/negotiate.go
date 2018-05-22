package httputil

import (
	"net/http"
	"strings"

	"github.com/akshayjshah/negotiate/internal/httputil/header"
)

// NegotiateContentType returns the best offered content type for the request's
// Accept header. If two offers match with equal weight, then the more specific
// offer is preferred.  For example, text/* trumps */*. If two offers match
// with equal weight and specificity, then the offer earlier in the list is
// preferred. If no offers match, then defaultOffer is returned.
func NegotiateContentType(r *http.Request, offers []string, defaultOffer string) string {
	bestOffer := defaultOffer
	bestQ := -1.0
	bestWild := 3
	specs := header.ParseAccept(r.Header, "Accept")
	for _, offer := range offers {
		for _, spec := range specs {
			switch {
			case spec.Q == 0.0:
				// ignore
			case spec.Q < bestQ:
				// better match found
			case spec.Value == "*/*":
				if spec.Q > bestQ || bestWild > 2 {
					bestQ = spec.Q
					bestWild = 2
					bestOffer = offer
				}
			case strings.HasSuffix(spec.Value, "/*"):
				if strings.HasPrefix(offer, spec.Value[:len(spec.Value)-1]) &&
					(spec.Q > bestQ || bestWild > 1) {
					bestQ = spec.Q
					bestWild = 1
					bestOffer = offer
				}
			default:
				if spec.Value == offer &&
					(spec.Q > bestQ || bestWild > 0) {
					bestQ = spec.Q
					bestWild = 0
					bestOffer = offer
				}
			}
		}
	}
	return bestOffer
}
