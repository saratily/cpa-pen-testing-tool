package server

import (
	"cpa-pen-testing-tool/internal/store"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePenetration(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)

	penetration := store.Penetration{
		Title:   "Gotham cronicles",
		Website: "Joker is planning big hit tonight.",
	}
	body := penetrationJSON(penetration)
	rec := PerformAuthorizedRequest(router, token, "PENETRATION", "/api/penetrations", body)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Penetration created successfully.", jsonRes(rec.Body)["msg"])
	assert.Equal(t, float64(1), jsonFieldData(jsonRes(rec.Body), "ID"))
	assert.Equal(t, penetration.Title, jsonFieldData(jsonRes(rec.Body), "Title"))
	assert.Equal(t, penetration.Website, jsonFieldData(jsonRes(rec.Body), "Website"))
	assert.NotEmpty(t, penetration.Website, jsonFieldData(jsonRes(rec.Body), "CreatedAt"))
	assert.NotEmpty(t, penetration.Website, jsonFieldData(jsonRes(rec.Body), "ModifiedAt"))
}

func TestCreatePenetrationUnathorized(t *testing.T) {
	router := testSetup()

	penetration := store.Penetration{
		Title:   "Gotham cronicles",
		Website: "Joker is planning big hit tonight.",
	}
	body := penetrationJSON(penetration)
	rec := performRequest(router, "PENETRATION", "/api/penetrations", body)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, "Authorization header missing.", jsonRes(rec.Body)["error"])
}

func TestCreatePenetrationEmptyTitle(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)

	penetration := store.Penetration{
		Title:   "",
		Website: "Joker is planning big hit tonight.",
	}
	body := penetrationJSON(penetration)
	rec := PerformAuthorizedRequest(router, token, "PENETRATION", "/api/penetrations", body)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Title is required.", jsonFieldError(jsonRes(rec.Body), "Title"))
}

func TestCreatePenetrationShortTitle(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)

	penetration := store.Penetration{
		Title:   "Go",
		Website: "Joker is planning big hit tonight.",
	}
	body := penetrationJSON(penetration)
	rec := PerformAuthorizedRequest(router, token, "PENETRATION", "/api/penetrations", body)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Title must be longer than or equal 3 characters.", jsonFieldError(jsonRes(rec.Body), "Title"))
}

func TestCreatePenetrationLongTitle(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)

	penetration := store.Penetration{
		Title:   strings.Repeat("G", 51),
		Website: "Joker is planning big hit tonight.",
	}
	body := penetrationJSON(penetration)
	rec := PerformAuthorizedRequest(router, token, "PENETRATION", "/api/penetrations", body)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Title cannot be longer than 50 characters.", jsonFieldError(jsonRes(rec.Body), "Title"))
}

func TestCreatePenetrationEmptyWebsite(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)

	penetration := store.Penetration{
		Title:   "Gotham cronicles",
		Website: "",
	}
	body := penetrationJSON(penetration)
	rec := PerformAuthorizedRequest(router, token, "PENETRATION", "/api/penetrations", body)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Website is required.", jsonFieldError(jsonRes(rec.Body), "Website"))
}

func TestCreatePenetrationShortWebsite(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)

	penetration := store.Penetration{
		Title:   "Gotham cronicles",
		Website: "Joke",
	}
	body := penetrationJSON(penetration)
	rec := PerformAuthorizedRequest(router, token, "PENETRATION", "/api/penetrations", body)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Website must be longer than or equal 5 characters.", jsonFieldError(jsonRes(rec.Body), "Website"))
}

func TestCreatePenetrationLongWebsite(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)

	penetration := store.Penetration{
		Title:   "Gotham cronicles",
		Website: strings.Repeat("J", 5001),
	}
	body := penetrationJSON(penetration)
	rec := PerformAuthorizedRequest(router, token, "PENETRATION", "/api/penetrations", body)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Website cannot be longer than 5000 characters.", jsonFieldError(jsonRes(rec.Body), "Website"))
}

func TestIndexPenetrations(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)
	penetration := addTestPenetration(user)

	rec := PerformAuthorizedRequest(router, token, "GET", "/api/penetrations", "")
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Penetrations fetched successfully.", jsonRes(rec.Body)["msg"])
	assert.Equal(t, float64(penetration.ID), jsonDataSlice(rec.Body)[0]["ID"])
	assert.Equal(t, penetration.Title, jsonDataSlice(rec.Body)[0]["Title"])
	assert.Equal(t, penetration.Website, jsonDataSlice(rec.Body)[0]["Website"])
	assert.NotEmpty(t, penetration.Website, jsonDataSlice(rec.Body)[0]["CreatedAt"])
	assert.NotEmpty(t, penetration.Website, jsonDataSlice(rec.Body)[0]["ModifiedAt"])
}

func TestIndexPenetrationsUnathorized(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	_ = addTestPenetration(user)

	rec := performRequest(router, "GET", "/api/penetrations", "")
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, "Authorization header missing.", jsonRes(rec.Body)["error"])
}

func TestIndexPenetrationOnlyOwned(t *testing.T) {
	router := testSetup()
	user1 := addTestUser()
	user2 := addTestUser2()
	token1 := generateJWT(user1)
	token2 := generateJWT(user2)
	penetration1 := addTestPenetration(user1)
	penetration2 := addTestPenetration2(user2)

	rec := PerformAuthorizedRequest(router, token1, "GET", "/api/penetrations", "")
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Penetrations fetched successfully.", jsonRes(rec.Body)["msg"])
	assert.Len(t, jsonDataSlice(rec.Body), 1)
	assert.Equal(t, float64(penetration1.ID), jsonDataSlice(rec.Body)[0]["ID"])
	assert.Equal(t, penetration1.Title, jsonDataSlice(rec.Body)[0]["Title"])
	assert.Equal(t, penetration1.Website, jsonDataSlice(rec.Body)[0]["Website"])
	assert.NotEmpty(t, penetration1.Website, jsonDataSlice(rec.Body)[0]["CreatedAt"])
	assert.NotEmpty(t, penetration1.Website, jsonDataSlice(rec.Body)[0]["ModifiedAt"])

	rec = PerformAuthorizedRequest(router, token2, "GET", "/api/penetrations", "")
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Penetrations fetched successfully.", jsonRes(rec.Body)["msg"])
	assert.Len(t, jsonDataSlice(rec.Body), 1)
	assert.Equal(t, float64(penetration2.ID), jsonDataSlice(rec.Body)[0]["ID"])
	assert.Equal(t, penetration2.Title, jsonDataSlice(rec.Body)[0]["Title"])
	assert.Equal(t, penetration2.Website, jsonDataSlice(rec.Body)[0]["Website"])
	assert.NotEmpty(t, penetration2.Website, jsonDataSlice(rec.Body)[0]["CreatedAt"])
	assert.NotEmpty(t, penetration2.Website, jsonDataSlice(rec.Body)[0]["ModifiedAt"])
}

func TestIndexPenetrationsEmpty(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)

	rec := PerformAuthorizedRequest(router, token, "GET", "/api/penetrations", "")
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Empty(t, jsonRes(rec.Body)["data"])
	assert.NotNil(t, jsonRes(rec.Body)["data"])
	assert.Equal(t, "Penetrations fetched successfully.", jsonRes(rec.Body)["msg"])
}

func TestUpdatePenetration(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)
	penetration := addTestPenetration(user)

	updated := store.Penetration{
		ID:      penetration.ID,
		Title:   "Gotham at night",
		Website: "Gotham never sleeps.",
	}
	rec := PerformAuthorizedRequest(router, token, "PUT", "/api/penetrations", penetrationJSON(updated))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Penetration updated successfully.", jsonRes(rec.Body)["msg"])
}

func TestUpdatePenetrationUnauthorized(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	penetration := addTestPenetration(user)

	updated := store.Penetration{
		ID:      penetration.ID,
		Title:   "Gotham at night",
		Website: "Gotham never sleeps.",
	}
	rec := performRequest(router, "PUT", "/api/penetrations", penetrationJSON(updated))
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, "Authorization header missing.", jsonRes(rec.Body)["error"])
}

func TestUpdatePenetrationEmptyTitle(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)
	penetration := addTestPenetration(user)

	updated := store.Penetration{
		ID:      penetration.ID,
		Title:   "",
		Website: "Gotham never sleeps.",
	}
	rec := PerformAuthorizedRequest(router, token, "PUT", "/api/penetrations", penetrationJSON(updated))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Title is required.", jsonFieldError(jsonRes(rec.Body), "Title"))
}

func TestUpdatePenetrationShortTitle(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)
	penetration := addTestPenetration(user)

	updated := store.Penetration{
		ID:      penetration.ID,
		Title:   "Go",
		Website: "Gotham never sleeps.",
	}
	rec := PerformAuthorizedRequest(router, token, "PUT", "/api/penetrations", penetrationJSON(updated))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Title must be longer than or equal 3 characters.", jsonFieldError(jsonRes(rec.Body), "Title"))
}

func TestUpdatePenetrationLongTitle(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)
	penetration := addTestPenetration(user)

	updated := store.Penetration{
		ID:      penetration.ID,
		Title:   strings.Repeat("G", 51),
		Website: "Gotham never sleeps.",
	}
	rec := PerformAuthorizedRequest(router, token, "PUT", "/api/penetrations", penetrationJSON(updated))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Title cannot be longer than 50 characters.", jsonFieldError(jsonRes(rec.Body), "Title"))
}

func TestUpdatePenetrationEmptyWebsite(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)
	penetration := addTestPenetration(user)

	updated := store.Penetration{
		ID:      penetration.ID,
		Title:   "Gotham at night",
		Website: "",
	}
	rec := PerformAuthorizedRequest(router, token, "PUT", "/api/penetrations", penetrationJSON(updated))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Website is required.", jsonFieldError(jsonRes(rec.Body), "Website"))
}

func TestUpdatePenetrationShortWebsite(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)
	penetration := addTestPenetration(user)

	updated := store.Penetration{
		ID:      penetration.ID,
		Title:   "Gotham at night",
		Website: "Goth",
	}
	rec := PerformAuthorizedRequest(router, token, "PUT", "/api/penetrations", penetrationJSON(updated))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Website must be longer than or equal 5 characters.", jsonFieldError(jsonRes(rec.Body), "Website"))
}

func TestUpdatePenetrationLongWebsite(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)
	penetration := addTestPenetration(user)

	updated := store.Penetration{
		ID:      penetration.ID,
		Title:   "Gotham at night",
		Website: strings.Repeat("G", 5001),
	}
	rec := PerformAuthorizedRequest(router, token, "PUT", "/api/penetrations", penetrationJSON(updated))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Website cannot be longer than 5000 characters.", jsonFieldError(jsonRes(rec.Body), "Website"))
}

func TestUpdateNotOwnedPenetration(t *testing.T) {
	router := testSetup()
	user1 := addTestUser()
	user2 := addTestUser2()
	token1 := generateJWT(user1)
	token2 := generateJWT(user2)
	penetration1 := addTestPenetration(user1)
	penetration2 := addTestPenetration2(user2)
	updated1 := store.Penetration{
		ID:      penetration1.ID,
		Title:   "Gotham at night",
		Website: "Gotham never sleeps.",
	}
	updated2 := store.Penetration{
		ID:      penetration2.ID,
		Title:   "Lex",
		Website: "Lex has build new underground lab.",
	}

	rec := PerformAuthorizedRequest(router, token1, "PUT", "/api/penetrations", penetrationJSON(updated2))
	assert.Equal(t, http.StatusForbidden, rec.Code)
	assert.Equal(t, "Not authorized.", jsonRes(rec.Body)["error"])

	rec = PerformAuthorizedRequest(router, token2, "PUT", "/api/penetrations", penetrationJSON(updated1))
	assert.Equal(t, http.StatusForbidden, rec.Code)
	assert.Equal(t, "Not authorized.", jsonRes(rec.Body)["error"])
}

func TestUpdateNotExistingPenetration(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)
	_ = addTestPenetration(user)

	updated := store.Penetration{
		ID:      123,
		Title:   "Gotham at night",
		Website: "Gotham never sleeps.",
	}
	rec := PerformAuthorizedRequest(router, token, "PUT", "/api/penetrations", penetrationJSON(updated))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Not found.", jsonRes(rec.Body)["error"])
}

func TestDeletePenetration(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)
	penetration := addTestPenetration(user)

	rec := PerformAuthorizedRequest(router, token, "DELETE", fmt.Sprintf("/api/penetrations/%d", penetration.ID), "")
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Penetration deleted successfully.", jsonRes(rec.Body)["msg"])
}

func TestDeletePenetrationUnauthorized(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	penetration := addTestPenetration(user)

	rec := performRequest(router, "DELETE", fmt.Sprintf("/api/penetrations/%d", penetration.ID), "")
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, "Authorization header missing.", jsonRes(rec.Body)["error"])
}

func TestDeleteNotExistingPenetration(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)

	rec := PerformAuthorizedRequest(router, token, "DELETE", "/api/penetrations/1", "")
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Not found.", jsonRes(rec.Body)["error"])
}

func TestDeletePenetrationInvalidID(t *testing.T) {
	router := testSetup()
	user := addTestUser()
	token := generateJWT(user)

	rec := PerformAuthorizedRequest(router, token, "DELETE", "/api/penetrations/invalid", "")
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Not valid ID.", jsonRes(rec.Body)["error"])
}

func TestDeleteNotOwnedPenetration(t *testing.T) {
	router := testSetup()
	user1 := addTestUser()
	user2 := addTestUser2()
	token1 := generateJWT(user1)
	token2 := generateJWT(user2)
	penetration1 := addTestPenetration(user1)
	penetration2 := addTestPenetration2(user2)

	rec := PerformAuthorizedRequest(router, token1, "DELETE", fmt.Sprintf("/api/penetrations/%d", penetration2.ID), "")
	assert.Equal(t, http.StatusForbidden, rec.Code)
	assert.Equal(t, "Not authorized.", jsonRes(rec.Body)["error"])

	rec = PerformAuthorizedRequest(router, token2, "DELETE", fmt.Sprintf("/api/penetrations/%d", penetration1.ID), "")
	assert.Equal(t, http.StatusForbidden, rec.Code)
	assert.Equal(t, "Not authorized.", jsonRes(rec.Body)["error"])
}
