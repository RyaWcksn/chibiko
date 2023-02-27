package entities

type SaveDatabase struct {
	Url string `json:"url"`
}

type GetDatabase struct {
	Id int `json:"id"`
}
