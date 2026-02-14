package traffic

import (
	"strings"

	domain "github.com/diteria/project_llot/backend/internal/domain/traffic"
)

var botSignatures = []string{
	"bot",
	"spider",
	"crawler",
	"slurp",
	"curl/",
	"wget/",
	"python-requests",
	"httpclient",
	"headless",
}

func classifyUserAgent(userAgent string) (domain.Classification, float64) {
	ua := strings.ToLower(strings.TrimSpace(userAgent))
	if ua == "" {
		return domain.ClassificationUnknown, 0.5
	}
	for _, signature := range botSignatures {
		if strings.Contains(ua, signature) {
			return domain.ClassificationBot, 0.9
		}
	}
	return domain.ClassificationHuman, 0.1
}
