/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CommentRequest struct {
	Content string `json:"content"`
}

// AssertCommentRequestRequired checks if the required fields are not zero-ed
func AssertCommentRequestRequired(obj CommentRequest) error {
	elements := map[string]interface{}{
		"content": obj.Content,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseCommentRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CommentRequest (e.g. [][]CommentRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCommentRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCommentRequest, ok := obj.(CommentRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCommentRequestRequired(aCommentRequest)
	})
}