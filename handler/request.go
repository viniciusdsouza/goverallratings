package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create Movie
type CreateMovieRequest struct {
	Title    string `json:"title"`
	Duration int64  `json:"duration"`
	Rating   int64  `json:"rating"`
	Genre    string `json:"genre"`
}

func (r *CreateMovieRequest) Validate() error {
	if r.Title == "" && r.Duration <= 0 && r.Rating <= 0 && r.Genre == ""{
		return fmt.Errorf("request body is empty or malformated")
	}
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.Duration <= 0 {
		return errParamIsRequired("duration", "integer")
	}
	if r.Rating <= 0 {
		return errParamIsRequired("rating", "integer")
	}
	if r.Genre == "" {
		return errParamIsRequired("genre", "string")
	}

	return nil
}
