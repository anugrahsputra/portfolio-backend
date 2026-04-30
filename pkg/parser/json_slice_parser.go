package parser

import "encoding/json"

func JsonSliceParser[T any](input any, profileID string) ([]T, error) {
	if input == nil {
		return nil, nil
	}

	data, err := json.Marshal(input)
	if err != nil || len(data) == 0 {
		return nil, err
	}

	var result []T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	// inject ProfileID if field exists
	for i := range result {
		if setter, ok := any(&result[i]).(interface {
			SetProfileID(string)
		}); ok {
			setter.SetProfileID(profileID)
		}
	}

	return result, nil
}
