package utils

import (
	"errors"
	"strconv"
)

func PointerFromJson(data *interface{}, jsonPath []string) (*interface{}, error) {
	curr := data
	for _, key := range jsonPath {
		if m, ok := (*curr).(map[string]interface{}); ok {
			val := m[key]
			curr = &val
			continue
		}

		if l, ok := (*curr).([]interface{}); ok {
			idx, err := strconv.Atoi(key)
			if err != nil {
				return nil, err
			}

			val := l[idx]
			curr = &val
			continue
		}

		return nil, errors.New("Could not cast data as valid json")
	}

	return curr, nil
}
