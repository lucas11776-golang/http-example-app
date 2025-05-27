package slices

import "encoding/json"

// Comment
func Map[T any, R any](items []T, callback func(item T) R) []R {
	mapped := []R{}

	for _, item := range items {
		mapped = append(mapped, callback(item))
	}

	return mapped
}

// Comment
func StructToMap(obj interface{}) (map[string]interface{}, error) {
	var out map[string]interface{}

	data, err := json.Marshal(obj)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &out)

	if err != nil {
		return nil, err
	}

	return out, nil
}
