/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ViewUsersResponse struct {
	Users []User `json:"users"`
}

// AssertViewUsersResponseRequired checks if the required fields are not zero-ed
func AssertViewUsersResponseRequired(obj ViewUsersResponse) error {
	elements := map[string]interface{}{
		"users": obj.Users,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Users {
		if err := AssertUserRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseViewUsersResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ViewUsersResponse (e.g. [][]ViewUsersResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseViewUsersResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aViewUsersResponse, ok := obj.(ViewUsersResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertViewUsersResponseRequired(aViewUsersResponse)
	})
}
