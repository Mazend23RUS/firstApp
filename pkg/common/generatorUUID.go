package common

import (
	"github.com/google/uuid"
)

type GeneratorUUID struct{}

func GetUUID() uuid.UUID {
    return uuid.New()
 
}
