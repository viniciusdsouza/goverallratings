package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateBookRequest struct {
	Title      string `json:"title"`
	AuthorName string `json:"authorName"`
	Rating     int64  `json:"rating"`
	Genre      string `json:"genre"`
}
type UpdateBookRequest struct {
	Title      string `json:"title"`
	AuthorName string `json:"authorName"`
	Rating     int64  `json:"rating"`
	Genre      string `json:"genre"`
}

func (r *CreateBookRequest) Validate() error {
	if r.Title == "" && r.AuthorName == "" && r.Rating <= 0 && r.Genre == "" {
		return fmt.Errorf("request body is empty or malformated")
	}
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.AuthorName == "" {
		return errParamIsRequired("authorName", "integer")
	}
	if r.Rating <= 0 {
		return errParamIsRequired("rating", "integer")
	}
	if r.Genre == "" {
		return errParamIsRequired("genre", "string")
	}

	return nil
}

func (r *UpdateBookRequest) Validate() error {
	if r.Title == "" && r.AuthorName == "" && r.Rating <= 0 && r.Genre == "" {
		return fmt.Errorf("request body is empty or malformated")
	}
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.AuthorName == "" {
		return errParamIsRequired("authorName", "integer")
	}
	if r.Rating <= 0 {
		return errParamIsRequired("rating", "integer")
	}
	if r.Genre == "" {
		return errParamIsRequired("genre", "string")
	}

	return nil
}
