package fieldslog

import (
	"fmt"
	"reflect"
)

func MergeMaps(maps ...map[string]any) map[string]any {
	result := make(map[string]any)

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

func MergeTypedMaps(maps ...map[string]any) (map[string]any, error) {
	result := make(map[string]any)

	for _, m := range maps {
		for k, v := range m {
			if existingVal, exists := result[k]; exists {
				if reflect.TypeOf(existingVal) != reflect.TypeOf(v) {
					return nil, fmt.Errorf("type mismatch for key '%s'", k)
				}
			}
			result[k] = v
		}
	}

	return result, nil
}
