// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTweet struct {
	CotentText string `json:"cotentText"`
	AuthoID    string `json:"authoId"`
}

type Tweet struct {
	ID              string `json:"id"`
	CotentText      string `json:"cotentText"`
	PublicationDate string `json:"publicationDate"`
	AuthoID         string `json:"authoId"`
}

type User struct {
	ID          string `json:"id"`
	ProfileName string `json:"profileName"`
}
