/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UserInfoRequest struct {
	Token string `json:"token"`
}

// AssertUserInfoRequestRequired checks if the required fields are not zero-ed
func AssertUserInfoRequestRequired(obj UserInfoRequest) error {
	elements := map[string]interface{}{
		"token": obj.Token,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseUserInfoRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UserInfoRequest (e.g. [][]UserInfoRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserInfoRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUserInfoRequest, ok := obj.(UserInfoRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUserInfoRequestRequired(aUserInfoRequest)
	})
}
