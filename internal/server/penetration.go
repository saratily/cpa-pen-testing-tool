package server

import (
	"bytes"
	"cpa-pen-testing-tool/internal/store"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func createPenetration(ctx *gin.Context) {
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

	data := map[string]interface{}{
		"domain":     penetration.Website,
		"URL":        "http://www." + penetration.Website,
		"ip_address": "",
		"Roles":      []string{"dbteam", "uiteam", "tester"},
	}

	default_tools, _ := store.FetchDefaultTool()

	fmt.Println(penetration)
	for i, default_tool := range default_tools {

		options := template.Must(template.New("options").Parse(default_tool.Options))
		optionsBuf := &bytes.Buffer{}
		if err := options.Execute(optionsBuf, data); err != nil {
			panic(err)
		}

		command := template.Must(template.New("command").Parse(default_tool.Format))
		commandBuf := &bytes.Buffer{}
		if err := command.Execute(commandBuf, data); err != nil {
			panic(err)
		}

		tool := &store.Tool{
			Unique_ID:     uuid.NewV4(),
			Type:          default_tools[i].Type,
			Category:      default_tools[i].Category,
			Options:       fmt.Sprintf("%s", optionsBuf.String()),
			Command:       fmt.Sprintf("%s", commandBuf.String()),
			Output:        "",
			CanChange:     default_tools[i].CanChange,
			Selected:      default_tools[i].Selected,
			PenetrationID: penetration.ID,
			CreatedAt:     time.Now(),
			ModifiedAt:    time.Now(),
		}
		store.AddTool(penetration, tool)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Penetration created successfully.",
		"data": penetration,
	})
}

func getPenetration(ctx *gin.Context) {
	paramID := ctx.Param("id")
	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
		return
	}
	penetration, err := store.FetchPenetration(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.ID != penetration.UserID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not authorized."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Penetrations fetched successfully.",
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
	// jsonPenetration := ctx.MustGet(gin.BindKey).(*store.Penetration)
	// user, err := currentUser(ctx)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
	// 	return
	// }
	// dbPenetration, err := store.FetchPenetration(jsonPenetration.ID)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// if user.ID != dbPenetration.UserID {
	// 	ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not authorized."})
	// 	return
	// }
	// jsonPenetration.ModifiedAt = time.Now()
	// if err := store.UpdatePenetration(jsonPenetration); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"msg":  "Penetration updated successfully.",
	// 	"data": jsonPenetration,
	// })
}

func deletePenetration(ctx *gin.Context) {
	paramID := ctx.Param("id")
	// id, err := strconv.Atoi(paramID)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
	// 	return
	// }
	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
		return
	}
	penetration, err := store.FetchPenetration(paramID)
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
