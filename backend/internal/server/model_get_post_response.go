/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetPostResponse struct {
	Post Post `json:"post"`
}

// AssertGetPostResponseRequired checks if the required fields are not zero-ed
func AssertGetPostResponseRequired(obj GetPostResponse) error {
	elements := map[string]interface{}{
		"post": obj.Post,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPostRequired(obj.Post); err != nil {
		return err
	}
	return nil
}

// AssertRecurseGetPostResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetPostResponse (e.g. [][]GetPostResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetPostResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetPostResponse, ok := obj.(GetPostResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetPostResponseRequired(aGetPostResponse)
	})
}
