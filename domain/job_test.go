package domain_test

import (
	"testing"
	"time"

	"github.com/ayrtonbsouza/codeflix-encoder-microservice/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path/to/file.mp4"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path/to/output", "converted", video)

	require.Nil(t, err)
	require.NotNil(t, job)
}
