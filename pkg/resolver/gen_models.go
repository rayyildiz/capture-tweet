// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package resolver

type ContactInput struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Message  string `json:"message"`
}

type Mutation struct {
}

type Query struct {
}

type SearchInput struct {
	Term string `json:"term"`
}
