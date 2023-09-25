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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service      DefaultApiServicer
	errorHandler ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultApiOption func(*DefaultApiController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultApiErrorHandler(h ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer, opts ...DefaultApiOption) Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{
		{
			"ChangePostPerms",
			strings.ToUpper("Put"),
			"/posts/{id}/perms",
			c.ChangePostPerms,
		},
		{
			"Comment",
			strings.ToUpper("Post"),
			"/posts/{id}/comments",
			c.Comment,
		},
		{
			"DeletePost",
			strings.ToUpper("Delete"),
			"/posts/{id}",
			c.DeletePost,
		},
		{
			"DeleteUser",
			strings.ToUpper("Delete"),
			"/users/{login}",
			c.DeleteUser,
		},
		{
			"GetPost",
			strings.ToUpper("Get"),
			"/posts/{id}",
			c.GetPost,
		},
		{
			"Login",
			strings.ToUpper("Post"),
			"/login",
			c.Login,
		},
		{
			"Publish",
			strings.ToUpper("Post"),
			"/posts",
			c.Publish,
		},
		{
			"React",
			strings.ToUpper("Post"),
			"/posts/{id}/reactions",
			c.React,
		},
		{
			"Register",
			strings.ToUpper("Post"),
			"/register",
			c.Register,
		},
		{
			"Subscribe",
			strings.ToUpper("Post"),
			"/users/{id}/subscribtions",
			c.Subscribe,
		},
		{
			"Uncomment",
			strings.ToUpper("Delete"),
			"/posts/{id}/comments",
			c.Uncomment,
		},
		{
			"UserInfo",
			strings.ToUpper("Get"),
			"/users/me",
			c.UserInfo,
		},
		{
			"ViewComments",
			strings.ToUpper("Get"),
			"/posts/{id}/comments",
			c.ViewComments,
		},
		{
			"ViewFeed",
			strings.ToUpper("Get"),
			"/feed",
			c.ViewFeed,
		},
		{
			"ViewProfilePosts",
			strings.ToUpper("Get"),
			"/users/{login}/posts",
			c.ViewProfilePosts,
		},
		{
			"ViewProfileUser",
			strings.ToUpper("Get"),
			"/users/{login}",
			c.ViewProfileUser,
		},
		{
			"ViewUsers",
			strings.ToUpper("Get"),
			"/users",
			c.ViewUsers,
		},
	}
}

// ChangePostPerms -
func (c *DefaultApiController) ChangePostPerms(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	result, err := c.service.ChangePostPerms(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// Comment -
func (c *DefaultApiController) Comment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	commentRequestParam := CommentRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&commentRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCommentRequestRequired(commentRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.Comment(r.Context(), idParam, commentRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeletePost -
func (c *DefaultApiController) DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	result, err := c.service.DeletePost(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteUser -
func (c *DefaultApiController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	loginParam := params["login"]
	result, err := c.service.DeleteUser(r.Context(), loginParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetPost -
func (c *DefaultApiController) GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	result, err := c.service.GetPost(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// Login -
func (c *DefaultApiController) Login(w http.ResponseWriter, r *http.Request) {
	loginRequestParam := LoginRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&loginRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertLoginRequestRequired(loginRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.Login(r.Context(), loginRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// Publish -
func (c *DefaultApiController) Publish(w http.ResponseWriter, r *http.Request) {
	publishRequestParam := PublishRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&publishRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPublishRequestRequired(publishRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.Publish(r.Context(), publishRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// React -
func (c *DefaultApiController) React(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	reactRequestParam := ReactRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&reactRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertReactRequestRequired(reactRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.React(r.Context(), idParam, reactRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// Register -
func (c *DefaultApiController) Register(w http.ResponseWriter, r *http.Request) {
	registerRequestParam := RegisterRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&registerRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertRegisterRequestRequired(registerRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.Register(r.Context(), registerRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// Subscribe -
func (c *DefaultApiController) Subscribe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	result, err := c.service.Subscribe(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// Uncomment -
func (c *DefaultApiController) Uncomment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	uncommentRequestParam := UncommentRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&uncommentRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUncommentRequestRequired(uncommentRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.Uncomment(r.Context(), idParam, uncommentRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UserInfo -
func (c *DefaultApiController) UserInfo(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.UserInfo(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ViewComments -
func (c *DefaultApiController) ViewComments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	result, err := c.service.ViewComments(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ViewFeed -
func (c *DefaultApiController) ViewFeed(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ViewFeed(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ViewProfilePosts -
func (c *DefaultApiController) ViewProfilePosts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	loginParam := params["login"]
	result, err := c.service.ViewProfilePosts(r.Context(), loginParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ViewProfileUser -
func (c *DefaultApiController) ViewProfileUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	loginParam := params["login"]
	result, err := c.service.ViewProfileUser(r.Context(), loginParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ViewUsers -
func (c *DefaultApiController) ViewUsers(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ViewUsers(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
