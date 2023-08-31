package server

import (
	"cpa-pen-testing-tool/internal/store"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/niemeyer/pretty"
)

func createPenetration(ctx *gin.Context) {
	pretty.Print("inside create")
	penetration := ctx.MustGet(gin.BindKey).(*store.Penetration)
	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := store.AddPenetration(user, penetration); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Penetration created successfully.",
		"data": penetration,
	})
}

func indexPenetrations(ctx *gin.Context) {
	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := store.FetchUserPenetrations(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Penetrations fetched successfully.",
		"data": user.Penetrations,
	})
}

func updatePenetration(ctx *gin.Context) {
	jsonPenetration := ctx.MustGet(gin.BindKey).(*store.Penetration)
	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
		return
	}
	dbPenetration, err := store.FetchPenetration(jsonPenetration.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.ID != dbPenetration.UserID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not authorized."})
		return
	}
	jsonPenetration.ModifiedAt = time.Now()
	if err := store.UpdatePenetration(jsonPenetration); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Penetration updated successfully.",
		"data": jsonPenetration,
	})
}

func deletePenetration(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
		return
	}
	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
		return
	}
	penetration, err := store.FetchPenetration(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.ID != penetration.UserID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not authorized."})
		return
	}
	if err := store.DeletePenetration(penetration); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Penetration deleted successfully."})
}
