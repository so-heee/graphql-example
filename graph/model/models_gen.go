// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type SearchResultItemConnection struct {
	After  *string    `json:"after"`
	Before *string    `json:"before"`
	First  *int       `json:"first"`
	Last   *int       `json:"last"`
	Query  string     `json:"query"`
	Type   SearchType `json:"type"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SearchType string

const (
	SearchTypeUser SearchType = "USER"
)

var AllSearchType = []SearchType{
	SearchTypeUser,
}

func (e SearchType) IsValid() bool {
	switch e {
	case SearchTypeUser:
		return true
	}
	return false
}

func (e SearchType) String() string {
	return string(e)
}

func (e *SearchType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SearchType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SearchType", str)
	}
	return nil
}

func (e SearchType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
