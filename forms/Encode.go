package forms

type EncodeRequest struct {
	Url         string `json:"url" validate:"true"`
	IsTemporary int    `json:"isTemporary"`
}
