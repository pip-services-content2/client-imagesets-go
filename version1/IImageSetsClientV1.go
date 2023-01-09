package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IImageSetsClientV1 interface {
	GetImageSets(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*ImageSetV1], error)

	GetImageSetById(ctx context.Context, correlationId string, imagesetId string) (*ImageSetV1, error)

	CreateImageSet(ctx context.Context, correlationId string, imageset *ImageSetV1) (*ImageSetV1, error)

	UpdateImageSet(ctx context.Context, correlationId string, imageset *ImageSetV1) (*ImageSetV1, error)

	DeleteImageSetById(ctx context.Context, correlationId string, imagesetId string) (*ImageSetV1, error)
}
