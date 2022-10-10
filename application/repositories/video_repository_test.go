package repositories_test

import (
	"testing"
	"time"

	"github.com/ayrtonbsouza/codeflix-encoder-microservice/application/repositories"
	"github.com/ayrtonbsouza/codeflix-encoder-microservice/domain"
	"github.com/ayrtonbsouza/codeflix-encoder-microservice/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path/to/file.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, video.ID, v.ID)
}
