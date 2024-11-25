package core

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUniqueId() string {
	uui := uuid.New()
	return strings.Replace(uui.String(), "-", "", -1)
}
