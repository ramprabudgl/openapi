package openapi

import (
	"strings"

	"github.com/ramprabudgl/openapi/models"
)

func SnssaiEqualFold(s, t models.Snssai) bool {
	if s.Sst == t.Sst && strings.EqualFold(s.Sd, t.Sd) {
		return true
	}
	return false
}

