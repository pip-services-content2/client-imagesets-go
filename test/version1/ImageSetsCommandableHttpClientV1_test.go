package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-imagesets-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type ImageSetsCommandableHttpClientV1 struct {
	client  *version1.ImageSetsCommandableHttpClientV1
	fixture *ImageSetsClientFixtureV1
}

func newImageSetsCommandableHttpClientV1() *ImageSetsCommandableHttpClientV1 {
	return &ImageSetsCommandableHttpClientV1{}
}

func (c *ImageSetsCommandableHttpClientV1) setup(t *testing.T) *ImageSetsClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewImageSetsCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewImageSetsClientFixtureV1(c.client)

	return c.fixture
}

func (c *ImageSetsCommandableHttpClientV1) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newImageSetsCommandableHttpClientV1()

	fixture := c.setup(t)
	fixture.TestCrudOperations(t)
	c.teardown(t)
}
