package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/victorbrugnolo/video-encoder/domain"

	uuid "github.com/satori/go.uuid"
)

func TestValidateIfVideoIsEmp(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNotAUuid(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "abc"
	video.ResourceID = "anyId"
	video.FilePath = "/path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceID = "anyId"
	video.FilePath = "/path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}
