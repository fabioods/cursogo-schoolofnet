package hello

import (
	"fmt"

	"github.com/google/uuid"
)

func ShowUUID() {
	id := uuid.New()
	fmt.Printf("UUID %v \n\n", id.String())
}
