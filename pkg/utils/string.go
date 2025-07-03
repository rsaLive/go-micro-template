package utils

import (
	"fmt"
	"strings"
)

func SliceToString[T ~int32](data []T) string {
	if len(data) == 0 {
		return ""
	}

	ids := make([]string, len(data))
	for i, e := range data {
		ids[i] = fmt.Sprintf("%d", int32(e))
	}
	return strings.Join(ids, ",")
}
