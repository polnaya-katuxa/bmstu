/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UserInfoResponse struct {
	User User `json:"user"`
}

// AssertUserInfoResponseRequired checks if the required fields are not zero-ed
func AssertUserInfoResponseRequired(obj UserInfoResponse) error {
	elements := map[string]interface{}{
		"user": obj.User,
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

// AssertRecurseUserInfoResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UserInfoResponse (e.g. [][]UserInfoResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserInfoResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUserInfoResponse, ok := obj.(UserInfoResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUserInfoResponseRequired(aUserInfoResponse)
	})
}
