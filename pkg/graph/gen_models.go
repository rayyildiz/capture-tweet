// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

type ContactInput struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Message  string `json:"message"`
}

type SearchInput struct {
	Term string `json:"term"`
}
