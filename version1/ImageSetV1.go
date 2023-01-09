package version1

import "time"

type ImageSetV1 struct {
	// Identification
	Id string `json:"id"`

	// Automatically set fields
	CreateTime time.Time `json:"create_time"`

	// Content
	Title string          `json:"title"`
	Pics  []*AttachmentV1 `json:"pics"`

	// Search
	Tags    []string `json:"tags"`
	AllTags []string `json:"all_tags"`
}

func NewImageSetV1(id, title string, picIds []string) *ImageSetV1 {
	return &ImageSetV1{
		Id:         id,
		Title:      title,
		Pics:       make([]*AttachmentV1, 0),
		CreateTime: time.Now(),
	}
}
