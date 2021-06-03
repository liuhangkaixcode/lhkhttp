package common

import (
	"fmt"
	"testing"
)

func TestGetUUID(t *testing.T) {
	fmt.Println(GetUUID())
	fmt.Println(GetSnowflakeId())
}