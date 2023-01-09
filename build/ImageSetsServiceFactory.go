package build

import (
	clients1 "github.com/pip-services-content2/client-imagesets-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type ImageSetsServiceFactory struct {
	*cbuild.Factory
}

func NewImageSetsServiceFactory() *ImageSetsServiceFactory {
	c := &ImageSetsServiceFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-imagesets", "client", "null", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-imagesets", "client", "mock", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-imagesets", "client", "commandable-http", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewImageSetsNullClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewImageSetsMockClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewImageSetsCommandableHttpClientV1)

	return c
}
