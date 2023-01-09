package version1

import (
	"context"
	"strings"
	"time"

	aclients "github.com/pip-services-content2/client-attachments-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

type ImageSetsMockClientV1 struct {
	imagesets []*ImageSetV1

	attachmentsConnector *AttachmentsConnector
	attachmentsClient    *aclients.AttachmentsMockClientV1
}

func NewImageSetsMockClientV1() *ImageSetsMockClientV1 {
	c := &ImageSetsMockClientV1{
		imagesets:         make([]*ImageSetV1, 0),
		attachmentsClient: aclients.NewAttachmentsMockClientV1(),
	}

	c.attachmentsConnector = NewAttachmentsConnector(c.attachmentsClient)

	return c
}

func (c *ImageSetsMockClientV1) contains(array1 []string, array2 []string) bool {
	if array1 == nil || array2 == nil {
		return false
	}

	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[i] {
				return true
			}
		}
	}

	return false
}

func (c *ImageSetsMockClientV1) matchString(value string, search string) bool {
	if value == "" && search == "" {
		return true
	}
	if value == "" || search == "" {
		return false
	}

	return strings.Index(strings.ToLower(value), search) >= 0
}

func (c *ImageSetsMockClientV1) matchSearch(item *ImageSetV1, search string) bool {
	search = strings.ToLower(search)
	if c.matchString(item.Title, search) {
		return true
	}

	return false
}

func (c *ImageSetsMockClientV1) composeFilter(filter *data.FilterParams) func(item *ImageSetV1) bool {
	if filter == nil {
		filter = data.NewEmptyFilterParams()
	}

	search, searchOk := filter.GetAsNullableString("search")
	id, idOk := filter.GetAsNullableString("id")
	tagsString := filter.GetAsString("tags")
	tags := make([]string, 0)
	if tagsString != "" {
		tags = data.TagsProcessor.CompressTags([]string{tagsString})
	}

	return func(item *ImageSetV1) bool {
		if idOk && id != item.Id {
			return false
		}
		if len(tags) > 0 && !c.contains(item.AllTags, tags) {
			return false
		}
		if searchOk && !c.matchSearch(item, search) {
			return false
		}
		return true
	}
}

func (c *ImageSetsMockClientV1) GetImageSets(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*ImageSetV1], error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*ImageSetV1, 0)
	for _, v := range c.imagesets {
		item := *v
		if filterFunc(&item) {
			items = append(items, &item)
		}
	}
	return *data.NewDataPage(items, len(c.imagesets)), nil
}

func (c *ImageSetsMockClientV1) GetImageSetById(ctx context.Context, correlationId string, imagesetId string) (result *ImageSetV1, err error) {
	for _, v := range c.imagesets {
		if v.Id == imagesetId {
			buf := *v
			result = &buf
			break
		}
	}
	return result, err
}

func (c *ImageSetsMockClientV1) CreateImageSet(ctx context.Context, correlationId string, imageset *ImageSetV1) (*ImageSetV1, error) {
	imageset.CreateTime = time.Now()
	imageset.AllTags = data.TagsProcessor.ExtractHashTags("#title")

	buf := *imageset
	c.imagesets = append(c.imagesets, &buf)

	err := c.attachmentsConnector.AddAttachments(ctx, correlationId, imageset)
	return imageset, err
}

func (c *ImageSetsMockClientV1) UpdateImageSet(ctx context.Context, correlationId string, imageset *ImageSetV1) (*ImageSetV1, error) {
	if imageset == nil {
		return nil, nil
	}

	var oldImageSet *ImageSetV1
	var newImageSet ImageSetV1

	imageset.AllTags = data.TagsProcessor.ExtractHashTags("#title")

	var index = -1
	for i, v := range c.imagesets {
		if v.Id == imageset.Id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, errors.NewNotFoundError(
			correlationId,
			"IMAGESET_NOT_FOUND",
			"ImageSet "+imageset.Id+" wass not found",
		).WithDetails("imageset_id", imageset.Id)
	}

	c.imagesets[index] = oldImageSet
	newImageSet = *imageset

	c.imagesets[index] = &newImageSet

	err := c.attachmentsConnector.UpdateAttachments(ctx, correlationId, oldImageSet, &newImageSet)
	return imageset, err

}

func (c *ImageSetsMockClientV1) DeleteImageSetById(ctx context.Context, correlationId string, imagesetId string) (*ImageSetV1, error) {
	var index = -1
	for i, v := range c.imagesets {
		if v.Id == imagesetId {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}
	var oldImageset = c.imagesets[index]
	if index < len(c.imagesets) {
		c.imagesets = append(c.imagesets[:index], c.imagesets[index+1:]...)
	} else {
		c.imagesets = c.imagesets[:index]
	}

	err := c.attachmentsConnector.RemoveAttachments(ctx, correlationId, oldImageset)
	return oldImageset, err
}
