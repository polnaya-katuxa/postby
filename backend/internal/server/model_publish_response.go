/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PublishResponse struct {
	Post Post `json:"post"`

	Published bool `json:"published"`
}

// AssertPublishResponseRequired checks if the required fields are not zero-ed
func AssertPublishResponseRequired(obj PublishResponse) error {
	elements := map[string]interface{}{
		"post":      obj.Post,
		"published": obj.Published,
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

// AssertRecursePublishResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PublishResponse (e.g. [][]PublishResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePublishResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPublishResponse, ok := obj.(PublishResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPublishResponseRequired(aPublishResponse)
	})
}
