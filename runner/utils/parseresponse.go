package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"golang.org/x/sync/errgroup"
)

func pointerFromJson(data *interface{}, jsonPath []string) (*interface{}, error) {
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

func valueFromJsonWithWildcard(ctx context.Context, data interface{}, jsonPath []string) (interface{}, error) {
	g, ctx := errgroup.WithContext(ctx)

	recurse := func(res []interface{}, resIdx int, val interface{}, keyIdx int) {
		g.Go(func() error {
			var err error
			res[resIdx], err = valueFromJsonWithWildcard(ctx, val, jsonPath[keyIdx+1:])
			return err
		})
	}

	curr := data
	for keyIdx, key := range jsonPath {
		if m, ok := curr.(map[string]interface{}); ok {
			if key == "*" {
				res := make([]interface{}, len(m))
				resIdx := 0

				for _, val := range m {
					recurse(res, resIdx, val, keyIdx)
					resIdx += 1
				}

				return res, g.Wait()
			}

			curr = m[key]
		} else if l, ok := curr.([]interface{}); ok {
			if key == "*" {
				res := make([]interface{}, len(l))

				for resIdx, val := range l {
					recurse(res, resIdx, val, keyIdx)
				}

				return res, g.Wait()
			}

			idx, err := strconv.Atoi(key)
			if err != nil {
				return nil, err
			}

			curr = l[idx]
		} else {
			return nil, errors.New("Could not cast data as valid json")
		}

	}

	return curr, nil
}

func GetFromJson(data []byte, jsonPath []string) (string, error) {
	var j interface{}
	err := json.Unmarshal(data, &j)
	if err != nil {
		return "", err
	}

	val, err := valueFromJsonWithWildcard(context.Background(), &j, jsonPath)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", *val), nil
}

func incrementNum(val interface{}) (int, error) {
	if num, ok := val.(float64); ok {
		return int(num) + 1, nil
	} else if num, ok := val.(int); ok {
		return num + 1, nil
	} else if numStr, ok := val.(string); ok {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, err
		}

		return num + 1, nil
	}
	return 0, errors.New("Could not cast value as valid type")
}

func IncrementBody(data []byte, jsonPath []string) ([]byte, error) {
	var j interface{}
	err := json.Unmarshal(data, &j)
	if err != nil {
		return nil, err
	}

	parentVal, err := pointerFromJson(&j, jsonPath[:len(jsonPath)-1])
	if err != nil {
		return nil, err
	}

	key := jsonPath[len(jsonPath)-1]
	if m, ok := (*parentVal).(map[string]interface{}); ok {
		m[key], err = incrementNum(m[key])
	} else if l, ok := (*parentVal).([]interface{}); ok {
		idx, err := strconv.Atoi(key)
		if err != nil {
			return nil, err
		}
		l[idx], err = incrementNum(m[key])
	}

	if err != nil {
		return nil, err
	}

	return json.Marshal(j)
}

func IncrementUrl(url string, urlRegex string) (string, error) {
	re := regexp.MustCompile(urlRegex)
	matchIndices := re.FindStringSubmatchIndex(url)
	if len(matchIndices) == 4 {
		start, end := matchIndices[2], matchIndices[3]
		newNum, err := incrementNum(url[start:end])

		if err != nil {
			return "", err
		}

		return url[:start] + strconv.Itoa(newNum) + url[end:], nil
	}

	return "", errors.New("No regex matches found")
}
