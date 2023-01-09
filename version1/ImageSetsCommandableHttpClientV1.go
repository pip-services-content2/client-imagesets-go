package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type ImageSetsCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewImageSetsCommandableHttpClientV1() *ImageSetsCommandableHttpClientV1 {
	return &ImageSetsCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/imagesets"),
	}
}

func (c *ImageSetsCommandableHttpClientV1) GetImageSets(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*ImageSetV1], error) {
	params := data.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(ctx, "get_imagesets", correlationId, params)
	if err != nil {
		return *data.NewEmptyDataPage[*ImageSetV1](), err
	}

	return clients.HandleHttpResponse[data.DataPage[*ImageSetV1]](res, correlationId)
}

func (c *ImageSetsCommandableHttpClientV1) GetImageSetById(ctx context.Context, correlationId string, imagesetId string) (*ImageSetV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"imageset_id", imagesetId,
	)

	res, err := c.CallCommand(ctx, "get_imageset_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ImageSetV1](res, correlationId)
}

func (c *ImageSetsCommandableHttpClientV1) CreateImageSet(ctx context.Context, correlationId string, imageset *ImageSetV1) (*ImageSetV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"imageset", imageset,
	)

	res, err := c.CallCommand(ctx, "create_imageset", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ImageSetV1](res, correlationId)
}

func (c *ImageSetsCommandableHttpClientV1) UpdateImageSet(ctx context.Context, correlationId string, imageset *ImageSetV1) (*ImageSetV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"imageset", imageset,
	)

	res, err := c.CallCommand(ctx, "update_imageset", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ImageSetV1](res, correlationId)
}

func (c *ImageSetsCommandableHttpClientV1) DeleteImageSetById(ctx context.Context, correlationId string, imagesetId string) (*ImageSetV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"imageset_id", imagesetId,
	)

	res, err := c.CallCommand(ctx, "delete_imageset_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ImageSetV1](res, correlationId)
}
