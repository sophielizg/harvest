package utils

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
)

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
