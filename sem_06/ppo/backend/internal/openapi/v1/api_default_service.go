/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"net/http"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct{}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// ChangePostPerms -
func (s *DefaultApiService) ChangePostPerms(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update ChangePostPerms with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ChangePermsResponse{}) or use other options such as http.Ok ...
	// return Response(200, ChangePermsResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ChangePostPerms method not implemented")
}

// Comment -
func (s *DefaultApiService) Comment(ctx context.Context, id string, commentRequest CommentRequest) (ImplResponse, error) {
	// TODO - update Comment with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, CommentResponse{}) or use other options such as http.Ok ...
	// return Response(201, CommentResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("Comment method not implemented")
}

// DeletePost -
func (s *DefaultApiService) DeletePost(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update DeletePost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, DeletePostResponse{}) or use other options such as http.Ok ...
	// return Response(200, DeletePostResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeletePost method not implemented")
}

// DeleteUser -
func (s *DefaultApiService) DeleteUser(ctx context.Context, login string) (ImplResponse, error) {
	// TODO - update DeleteUser with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, DeleteUserResponse{}) or use other options such as http.Ok ...
	// return Response(200, DeleteUserResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteUser method not implemented")
}

// GetPost -
func (s *DefaultApiService) GetPost(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update GetPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetPostResponse{}) or use other options such as http.Ok ...
	// return Response(200, GetPostResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPost method not implemented")
}

// Login -
func (s *DefaultApiService) Login(ctx context.Context, loginRequest LoginRequest) (ImplResponse, error) {
	// TODO - update Login with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, AuthResponse{}) or use other options such as http.Ok ...
	// return Response(200, AuthResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("Login method not implemented")
}

// Publish -
func (s *DefaultApiService) Publish(ctx context.Context, publishRequest PublishRequest) (ImplResponse, error) {
	// TODO - update Publish with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, PublishResponse{}) or use other options such as http.Ok ...
	// return Response(201, PublishResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("Publish method not implemented")
}

// React -
func (s *DefaultApiService) React(ctx context.Context, id string, reactRequest ReactRequest) (ImplResponse, error) {
	// TODO - update React with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ReactResponse{}) or use other options such as http.Ok ...
	// return Response(200, ReactResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("React method not implemented")
}

// Register -
func (s *DefaultApiService) Register(ctx context.Context, registerRequest RegisterRequest) (ImplResponse, error) {
	// TODO - update Register with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, AuthResponse{}) or use other options such as http.Ok ...
	// return Response(201, AuthResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("Register method not implemented")
}

// Subscribe -
func (s *DefaultApiService) Subscribe(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update Subscribe with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, SubscribeResponse{}) or use other options such as http.Ok ...
	// return Response(200, SubscribeResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("Subscribe method not implemented")
}

// Uncomment -
func (s *DefaultApiService) Uncomment(ctx context.Context, id string, uncommentRequest UncommentRequest) (ImplResponse, error) {
	// TODO - update Uncomment with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, UncommentResponse{}) or use other options such as http.Ok ...
	// return Response(200, UncommentResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("Uncomment method not implemented")
}

// UserInfo -
func (s *DefaultApiService) UserInfo(ctx context.Context) (ImplResponse, error) {
	// TODO - update UserInfo with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, UserInfoResponse{}) or use other options such as http.Ok ...
	// return Response(200, UserInfoResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UserInfo method not implemented")
}

// ViewComments -
func (s *DefaultApiService) ViewComments(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update ViewComments with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ViewCommentsResponse{}) or use other options such as http.Ok ...
	// return Response(200, ViewCommentsResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ViewComments method not implemented")
}

// ViewFeed -
func (s *DefaultApiService) ViewFeed(ctx context.Context) (ImplResponse, error) {
	// TODO - update ViewFeed with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ViewFeedResponse{}) or use other options such as http.Ok ...
	// return Response(200, ViewFeedResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ViewFeed method not implemented")
}

// ViewProfilePosts -
func (s *DefaultApiService) ViewProfilePosts(ctx context.Context, login string) (ImplResponse, error) {
	// TODO - update ViewProfilePosts with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ViewProfilePostsResponse{}) or use other options such as http.Ok ...
	// return Response(200, ViewProfilePostsResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ViewProfilePosts method not implemented")
}

// ViewProfileUser -
func (s *DefaultApiService) ViewProfileUser(ctx context.Context, login string) (ImplResponse, error) {
	// TODO - update ViewProfileUser with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ViewProfileUserResponse{}) or use other options such as http.Ok ...
	// return Response(200, ViewProfileUserResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ViewProfileUser method not implemented")
}

// ViewUsers -
func (s *DefaultApiService) ViewUsers(ctx context.Context) (ImplResponse, error) {
	// TODO - update ViewUsers with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ViewUsersResponse{}) or use other options such as http.Ok ...
	// return Response(200, ViewUsersResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ViewUsers method not implemented")
}