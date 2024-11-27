package resources

import (
	"strings"
	"testing"

	"github.com/no-code-api/api/internal/resources/domain/validations"
)

func TestPathsConflict(t *testing.T) {

	var tests = []struct {
		name         string
		endpointPath string
		path         string
		want         bool
	}{
		{" '/'vs'/:id' - Should pass", "/", "/:id", false},
		{" '/'vs'/teste' - Should pass", "/", "/teste", false},
		{" '/'vs'/:id/teste' - Should pass", "/", "/:id/teste", false},
		{" '/:id'vs'/:id/teste' - Should pass", "/:id", "/:id/teste", false},
		{" '/:name'vs'/:id' - Should not pass", "/:name", "/:id", true},
		{" '/:name'vs'/teste' - Should not pass", "/:name", "/teste", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t2 *testing.T) {
			endpointSegment := strings.Split(tt.endpointPath, "/")
			pathSegment := strings.Split(tt.path, "/")
			answer := validations.PathsConflict(endpointSegment, pathSegment)
			if answer != tt.want {
				t2.Errorf("recebeu %v, esperava %v", answer, tt.want)
			}
		})
	}
}
