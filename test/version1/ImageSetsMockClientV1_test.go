package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-content2/client-imagesets-go/version1"
)

type ImageSetsMockClientV1 struct {
	client  *version1.ImageSetsMockClientV1
	fixture *ImageSetsClientFixtureV1
}

func newImageSetsMockClientV1() *ImageSetsMockClientV1 {
	return &ImageSetsMockClientV1{}
}

func (c *ImageSetsMockClientV1) setup(t *testing.T) *ImageSetsClientFixtureV1 {
	c.client = version1.NewImageSetsMockClientV1()
	c.fixture = NewImageSetsClientFixtureV1(c.client)
	return c.fixture
}

func (c *ImageSetsMockClientV1) teardown(t *testing.T) {
	imagesets, _ := c.client.GetImageSets(context.Background(), "123", nil, nil)
	for _, imageset := range imagesets.Data {
		c.client.DeleteImageSetById(context.Background(), "123", imageset.Id)
	}
}

func TestMockCrudOperations(t *testing.T) {
	c := newImageSetsMockClientV1()

	fixture := c.setup(t)
	fixture.TestCrudOperations(t)
	c.teardown(t)
}
