package models

type URL struct {
	ShortCode      string
	OriginalURL    string
	CreatedBy      string
	CreationTime   string
	ExpirationTime *string
}
