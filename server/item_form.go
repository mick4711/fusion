package server

import "time"

type ItemFeed struct {
	ID   uint    `json:"id"`
	Name *string `json:"name"`
}
type ItemForm struct {
	ID      uint       `json:"id"`
	Title   *string    `json:"title"`
	Link    *string    `json:"link"`
	GUID    *string    `json:"guid"`
	Content *string    `json:"content"`
	PubDate *time.Time `json:"pub_date"`
	Unread  *bool      `json:"unread"`
	Feed    ItemFeed   `json:"feed"`
}

type ReqItemList struct {
	Paginate
	Keyword *string `query:"keyword"`
	FeedID  *uint   `query:"feed_id"`
	Unread  *bool   `query:"unread"`
}

type RespItemList struct {
	Total *int        `json:"total"`
	Items []*ItemForm `json:"items"`
}

type ReqItemGet struct {
	ID uint `param:"id" validate:"required"`
}

type RespItemGet ItemForm

type ReqItemUpdate struct {
	ID     uint  `param:"id" validate:"required"`
	Unread *bool `json:"unread"`
}

type ReqItemDelete struct {
	ID uint `param:"id" validate:"required"`
}
