package models

type Tweet struct {
	Id int `json:"id"`

	CreatedAt string `json:"createdAt"`

	TwitterId int64 `json:"twitterId,string"`

	Text string `json:"text"`

	School string `json:"school"`
}
