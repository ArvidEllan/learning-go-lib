package programming

import (
	"strings"

	"github.com/google/uuid"
)

func NewUuid(withoutHyphen bool) string {
	uuidWithHyphen := uuid.New()

	if withoutHyphen {

		return strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	}

	return uuidWithHyphen.String()

}

// NewUuid generates an UUID with the possibility
