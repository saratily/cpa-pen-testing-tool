package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPenetration(t *testing.T) {
	testSetup()
	user, err := addTestUser()
	assert.NoError(t, err)

	penetration, err := addTestPenetration(user)
	assert.NoError(t, err)
	assert.Equal(t, 1, penetration.ID)
}

func TestFetchUserPenetrations(t *testing.T) {
	testSetup()
	user, err := addTestUser()
	assert.NoError(t, err)
	penetration, err := addTestPenetration(user)
	assert.NoError(t, err)

	err = FetchUserPenetrations(user)
	assert.NoError(t, err)
	assert.Equal(t, penetration, user.Penetrations[0])
}

func TestFetchUserPenetrationsEmpty(t *testing.T) {
	testSetup()
	user, err := addTestUser()
	assert.NoError(t, err)

	err = FetchUserPenetrations(user)
	assert.NoError(t, err)
	assert.Empty(t, user.Penetrations)
	assert.NotNil(t, user.Penetrations)
}

func TestFetchPenetration(t *testing.T) {
	testSetup()
	user, err := addTestUser()
	assert.NoError(t, err)
	penetration, err := addTestPenetration(user)
	assert.NoError(t, err)

	fetchedPenetration, err := FetchPenetration(penetration.ID)
	assert.NoError(t, err)
	assert.Equal(t, penetration.ID, fetchedPenetration.ID)
	assert.Equal(t, penetration.Title, fetchedPenetration.Title)
	assert.Equal(t, penetration.Website, fetchedPenetration.Website)
	assert.Equal(t, user.ID, fetchedPenetration.UserID)
}

func TestFetchNotExistingPenetration(t *testing.T) {
	testSetup()

	fetchedPenetration, err := FetchPenetration(1)
	assert.Error(t, err)
	assert.Nil(t, fetchedPenetration)
	assert.Equal(t, "Not found.", err.Error())
}

func TestUpdatePenetration(t *testing.T) {
	testSetup()
	user, err := addTestUser()
	assert.NoError(t, err)
	penetration, err := addTestPenetration(user)
	assert.NoError(t, err)

	penetration.Title = "New title"
	penetration.Website = "www.example.ocm"
	err = UpdatePenetration(penetration)
	assert.NoError(t, err)
}

func TestDeletePenetration(t *testing.T) {
	testSetup()
	user, err := addTestUser()
	assert.NoError(t, err)
	penetration, err := addTestPenetration(user)
	assert.NoError(t, err)

	err = DeletePenetration(penetration)
	assert.NoError(t, err)
}
