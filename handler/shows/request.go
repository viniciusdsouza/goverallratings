package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}
type CreateShowRequest struct {
	Title    string `json:"title"`
	Episodes int64  `json:"episodes"`
	Seasons  int64  `json:"seasons"`
	Rating   int64  `json:"rating"`
	Genre    string `json:"genre"`
}
type UpdateShowRequest struct {
	Title    string `json:"title"`
	Episodes int64  `json:"episodes"`
	Seasons  int64  `json:"seasons"`
	Rating   int64  `json:"rating"`
	Genre    string `json:"genre"`
}

func (r *CreateShowRequest) Validate() error {
	if r.Title == "" && r.Episodes <= 0 && r.Seasons <= 0 && r.Rating <= 0 && r.Genre == "" {
		return fmt.Errorf("request body is empty or malformated")
	}
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.Episodes <= 0 {
		return errParamIsRequired("episodes", "integer")
	}
	if r.Seasons <= 0 {
		return errParamIsRequired("seasons", "integer")
	}
	if r.Rating <= 0 {
		return errParamIsRequired("rating", "integer")
	}
	if r.Genre == "" {
		return errParamIsRequired("genre", "string")
	}

	return nil
}

func (r *UpdateShowRequest) Validate() error {
	if r.Title == "" && r.Episodes <= 0 && r.Seasons <= 0 && r.Rating <= 0 && r.Genre == "" {
		return fmt.Errorf("request body is empty or malformated")
	}
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.Episodes <= 0 {
		return errParamIsRequired("episodes", "integer")
	}
	if r.Seasons <= 0 {
		return errParamIsRequired("seasons", "integer")
	}
	if r.Rating <= 0 {
		return errParamIsRequired("rating", "integer")
	}
	if r.Genre == "" {
		return errParamIsRequired("genre", "string")
	}

	return nil
}
