package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type ImageSetsNullClientV1 struct {
}

func NewImageSetsNullClientV1() *ImageSetsNullClientV1 {
	return &ImageSetsNullClientV1{}
}

func (c *ImageSetsNullClientV1) GetImageSets(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*ImageSetV1], error) {
	return *data.NewEmptyDataPage[*ImageSetV1](), nil
}

func (c *ImageSetsNullClientV1) GetImageSetById(ctx context.Context, correlationId string, imagesetId string) (*ImageSetV1, error) {
	return nil, nil
}

func (c *ImageSetsNullClientV1) CreateImageSet(ctx context.Context, correlationId string, imageset *ImageSetV1) (*ImageSetV1, error) {
	return imageset, nil
}

func (c *ImageSetsNullClientV1) UpdateImageSet(ctx context.Context, correlationId string, imageset *ImageSetV1) (*ImageSetV1, error) {
	return imageset, nil
}

func (c *ImageSetsNullClientV1) DeleteImageSetById(ctx context.Context, correlationId string, imagesetId string) (*ImageSetV1, error) {
	return nil, nil
}
