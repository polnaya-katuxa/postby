/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ViewProfileUserResponse struct {
	User User `json:"user"`

	Subscribed bool `json:"subscribed"`

	Self bool `json:"self"`
}

// AssertViewProfileUserResponseRequired checks if the required fields are not zero-ed
func AssertViewProfileUserResponseRequired(obj ViewProfileUserResponse) error {
	elements := map[string]interface{}{
		"user":       obj.User,
		"subscribed": obj.Subscribed,
		"self":       obj.Self,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertUserRequired(obj.User); err != nil {
		return err
	}
	return nil
}

// AssertRecurseViewProfileUserResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ViewProfileUserResponse (e.g. [][]ViewProfileUserResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseViewProfileUserResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aViewProfileUserResponse, ok := obj.(ViewProfileUserResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertViewProfileUserResponseRequired(aViewProfileUserResponse)
	})
}
