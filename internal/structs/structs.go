package structs

import "fmt"

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

func NewInvalidType(v any, err error) *ErrorResponse {
	return &ErrorResponse{
		Error: fmt.Sprintf("Был получен неправильный тип: %v.\n%v", v, err),
	}
}

type UniqueBanner struct {
	Title     string `json:"title,omitempty"`
	Text      string `json:"text,omitempty"`
	Url       string `json:"url,omitempty"`
	IsActive  bool   `json:"is_active,omitempty"`
	TagId     int    `json:"tag_id,omitempty"`
	FeatureId int    `json:"feature_id,omitempty"`
}

type Banner struct {
	Title    string `json:"title,omitempty"`
	Text     string `json:"text,omitempty"`
	Url      string `json:"url,omitempty"`
	IsActive bool   `json:"is_active,omitempty"`
}

func (b *UniqueBanner) ToBanner() *Banner {
	return &Banner{
		Title:    b.Title,
		Text:     b.Text,
		Url:      b.Url,
		IsActive: b.IsActive,
	}
}
