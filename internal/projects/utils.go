package projects

import (
	"strings"

	"github.com/google/uuid"
)

func generateUniqueId() string {
	uui := uuid.New()
	return strings.Replace(uui.String(), "-", "", -1)
}
