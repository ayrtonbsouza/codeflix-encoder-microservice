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

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path/to/file.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "pending", video)

	require.Nil(t, err)

	jobRepo := repositories.JobRepositoryDb{Db: db}
	jobRepo.Insert(job)

	j, err := jobRepo.Find(job.ID)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, job.ID, j.ID)
	require.Equal(t, j.VideoID, video.ID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path/to/file.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "pending", video)

	require.Nil(t, err)

	jobRepo := repositories.JobRepositoryDb{Db: db}
	jobRepo.Insert(job)

	job.Status = "completed"

	jobRepo.Update(job)

	j, err := jobRepo.Find(job.ID)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.Status, job.Status)
}
