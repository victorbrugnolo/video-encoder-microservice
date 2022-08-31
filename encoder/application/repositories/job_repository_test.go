package repositories_test

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/victorbrugnolo/video-encoder/application/repositories"
	"github.com/victorbrugnolo/video-encoder/domain"
	"github.com/victorbrugnolo/video-encoder/framework/database"
)

func TestJobRepositoryInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	found, err := repoJob.Find(job.ID)

	require.NotEmpty(t, found.ID)
	require.Nil(t, err)
	require.Equal(t, found.ID, job.ID)
	require.Equal(t, found.VideoID, video.ID)
}

func TestJobRepositoryUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	job.Status = "Complete"

	repoJob.Update(job)

	found, err := repoJob.Find(job.ID)

	require.NotEmpty(t, found.ID)
	require.Nil(t, err)
	require.Equal(t, found.ID, job.ID)
	require.Equal(t, found.Status, job.Status)

}
