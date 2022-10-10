package domain_test

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/ayrtonbsouza/codeflix-encoder-microservice/domain"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestValidateIfVideoIDIsNotUUID(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "invalid-uuid"
	video.ResourceID = "test-resource-id"
	video.FilePath = "test/path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "test-resource-id"
	video.FilePath = "test/path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}
