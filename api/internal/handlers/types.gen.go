// Package handlers provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package handlers

// Problem defines model for Problem.
type Problem struct {
	Detail *string `json:"detail,omitempty"`
	Title  *string `json:"title,omitempty"`
	Type   *string `json:"type,omitempty"`
}

// User defines model for User.
type User struct {
	Email *string `json:"email,omitempty"`
	Id    *string `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
}

// PostUserJSONRequestBody defines body for PostUser for application/json ContentType.
type PostUserJSONRequestBody = User
