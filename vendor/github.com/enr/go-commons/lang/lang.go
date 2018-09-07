package lang

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
)

// SliceContainsString returns true if the given string is present in a strings slice.
// Case sensitive.
func SliceContainsString(list []string, a string) bool {
	sort.Strings(list)
	i := sort.SearchStrings(list, a)
	return (i < len(list) && list[i] == a)
}

// ExtractJSONFieldValue returns the value for a given field of a JSON structure.
func ExtractJSONFieldValue(data []byte, key string) (interface{}, error) {
	var b map[string]interface{}
	err := json.Unmarshal(data, &b)
	if err != nil {
		return nil, err
	}
	if value, keyExists := b[key]; keyExists {
		return value, nil
	}
	return nil, errors.New(`field "%s" not found`)
}

// JSONArrayToStringSlice coerces a JSON array to slice of strings.
// Copyright Am Laher
// https://github.com/laher/goxc/blob/master/typeutils/mapstringinterfaceutils.go
func JSONArrayToStringSlice(v interface{}, k string) ([]string, error) {
	ret := []string{}
	switch typedV := v.(type) {
	case []interface{}:
		for _, i := range typedV {
			ret = append(ret, i.(string))
		}
		return ret, nil
	}
	return ret, fmt.Errorf("%s should be a `json array`, got a %T", k, v)
}
