package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-content2/client-imagesets-go/version1"
	"github.com/stretchr/testify/assert"
)

type ImageSetsClientFixtureV1 struct {
	Client version1.IImageSetsClientV1

	IMAGESET1 *version1.ImageSetV1
	IMAGESET2 *version1.ImageSetV1
}

func NewImageSetsClientFixtureV1(client version1.IImageSetsClientV1) *ImageSetsClientFixtureV1 {
	return &ImageSetsClientFixtureV1{
		Client: client,
		IMAGESET1: &version1.ImageSetV1{
			Id:    "1",
			Title: "ImageSet 1",
			Pics: []*version1.AttachmentV1{
				version1.NewAttachmentV1("111", "", ""),
				version1.NewAttachmentV1("222", "", ""),
				version1.NewAttachmentV1("333", "", ""),
			},
		},
		IMAGESET2: &version1.ImageSetV1{
			Id:      "2",
			Tags:    []string{"TAG 1"},
			AllTags: []string{"tag1"},
			Title:   "ImageSet 2",
			Pics: []*version1.AttachmentV1{
				version1.NewAttachmentV1("444", "", ""),
				version1.NewAttachmentV1("555", "", ""),
				version1.NewAttachmentV1("777", "", ""),
			},
		},
	}
}

func (c *ImageSetsClientFixtureV1) TestCrudOperations(t *testing.T) {
	// Create one imageset
	imageset1, err := c.Client.CreateImageSet(context.Background(), "123", c.IMAGESET1)
	assert.Nil(t, err)

	assert.NotNil(t, imageset1)
	assert.Equal(t, imageset1.Title, c.IMAGESET1.Title)
	assert.Equal(t, imageset1.Pics, c.IMAGESET1.Pics)

	// Create another imageset
	imageset2, err := c.Client.CreateImageSet(context.Background(), "123", c.IMAGESET2)
	assert.Nil(t, err)

	assert.NotNil(t, imageset2)
	assert.Equal(t, imageset2.Title, c.IMAGESET2.Title)
	assert.Equal(t, imageset2.Pics, c.IMAGESET2.Pics)

	// Get all imagesets
	page, err := c.Client.GetImageSets(context.Background(), "123", nil, nil)
	assert.Nil(t, err)

	assert.NotNil(t, page)
	assert.Len(t, page.Data, 2)

	// Update the imageset
	imageset1.Title = "New Title 1"

	imageset, err := c.Client.UpdateImageSet(context.Background(), "123", imageset1)
	assert.Nil(t, err)

	assert.NotNil(t, imageset)
	assert.Equal(t, imageset.Title, "New Title 1")
	assert.Equal(t, imageset.Pics, c.IMAGESET1.Pics)

	imageset1 = imageset

	// Delete imageset
	_, err = c.Client.DeleteImageSetById(context.Background(), "123", imageset1.Id)
	assert.Nil(t, err)

	// Try to get delete imageset
	imageset, err = c.Client.GetImageSetById(context.Background(), "123", imageset1.Id)
	assert.Nil(t, err)
	assert.Nil(t, imageset)
}
