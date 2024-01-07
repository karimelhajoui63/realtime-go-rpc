package colorrepo_test

import (
	"testing"

	"rpc-server/internal/adapter/outbound/colorrepo"
	"rpc-server/internal/core/domain/enum"

	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryColorRepository(t *testing.T) {
	repo, err := colorrepo.NewInMemoryColorRepository()

	// Test repository creation
	assert.NoError(t, err)
	assert.NotNil(t, repo)
}

func TestInMemoryColorRepository_Get(t *testing.T) {
	repo, _ := colorrepo.NewInMemoryColorRepository()

	// Test initial state
	color, err := repo.Get()
	assert.NoError(t, err)
	assert.Equal(t, enum.Unspecified, color)

	// Test after an update
	expectedColor := enum.Red
	repo.Update(expectedColor)
	color, err = repo.Get()
	assert.NoError(t, err)
	assert.Equal(t, expectedColor, color)
}

func TestInMemoryColorRepository_Update(t *testing.T) {
	repo, _ := colorrepo.NewInMemoryColorRepository()

	// Test updating color
	expectedColor := enum.Blue
	repo.Update(expectedColor)
	color, err := repo.Get()
	assert.NoError(t, err)
	assert.Equal(t, expectedColor, color)
}
