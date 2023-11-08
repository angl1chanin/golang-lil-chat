package handlers

import (
	"chat/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type handlers struct {
	message usecase.MessageUseCase
}

func Init(message usecase.MessageUseCase) *handlers {
	return &handlers{message: message}
}

func (h *handlers) Get(ctx *gin.Context) {
	// get body from request
	body, err := ctx.GetRawData()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get messages
	items, err := h.message.Get(body)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// check if messages are empty
	if len(items) == 0 {
		ctx.IndentedJSON(http.StatusAccepted, gin.H{"status": "empty"})
		return
	}

	// return messages
	ctx.IndentedJSON(http.StatusOK, items)
}

func (h *handlers) Create(ctx *gin.Context) {
	// get body from request
	body, err := ctx.GetRawData()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	// create message
	if err = h.message.Create(body); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	// return success
	ctx.IndentedJSON(http.StatusOK, map[string]string{"status": "success"})
}

func (h *handlers) Delete(ctx *gin.Context) {
	// get id from url
	id := ctx.Param("id")

	// convert id to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	// delete message
	if err = h.message.Delete(intID); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	// return success
	ctx.IndentedJSON(http.StatusOK, map[string]string{"status": "success"})
}
