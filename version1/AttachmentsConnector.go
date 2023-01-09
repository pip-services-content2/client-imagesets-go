package version1

import (
	"context"

	aclients "github.com/pip-services-content2/client-attachments-go/version1"
)

type AttachmentsConnector struct {
	attachmentsClient aclients.IAttachmentsClientV1
}

func NewAttachmentsConnector(client aclients.IAttachmentsClientV1) *AttachmentsConnector {
	return &AttachmentsConnector{
		attachmentsClient: client,
	}
}

func (c *AttachmentsConnector) extractAttachmentIds(imageset *ImageSetV1) []string {
	ids := make([]string, 0)

	if imageset == nil {
		return ids
	}

	for _, pic := range imageset.Pics {
		if pic.Id != "" {
			ids = append(ids, pic.Id)
		}
	}

	return ids
}

func (c *AttachmentsConnector) AddAttachments(ctx context.Context, correlationId string, imageset *ImageSetV1) error {
	if c.attachmentsClient == nil || imageset == nil {
		return nil
	}

	ids := c.extractAttachmentIds(imageset)
	reference := aclients.NewReferenceV1(imageset.Id, "imageset", "")
	_, err := c.attachmentsClient.AddAttachments(ctx, correlationId, reference, ids)
	return err
}

func (c *AttachmentsConnector) UpdateAttachments(ctx context.Context, correlationId string, oldImageset *ImageSetV1, newImageset *ImageSetV1) error {
	if c.attachmentsClient == nil || oldImageset == nil || newImageset == nil {
		return nil
	}

	oldIds := c.extractAttachmentIds(oldImageset)
	newIds := c.extractAttachmentIds(newImageset)
	reference := aclients.NewReferenceV1(newImageset.Id, "imageset", "")
	_, err := c.attachmentsClient.UpdateAttachments(ctx, correlationId, reference, oldIds, newIds)
	return err
}

func (c *AttachmentsConnector) RemoveAttachments(ctx context.Context, correlationId string, imageset *ImageSetV1) error {
	if c.attachmentsClient == nil || imageset == nil {
		return nil
	}

	ids := c.extractAttachmentIds(imageset)
	reference := aclients.NewReferenceV1(imageset.Id, "imageset", "")
	_, err := c.attachmentsClient.RemoveAttachments(ctx, correlationId, reference, ids)
	return err
}
