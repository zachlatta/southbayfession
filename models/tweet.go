package models

type Tweet struct {
	Id int `json:"id"`

	CreatedAt string `json:"createdAt"`

	TwitterId string `json:"twitterId"`

	Text string `json:"text"`

	School string `json:"school"`
}
