package handlers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"math/rand"
	"rest_service/api/http"
	"rest_service/models"
	"time"
)

// @Router /v1/phone/{ID} [GET]
// @Summary returns phone from server
// @Description this returns phone if exists
// @Accept json
// @Produce json
// @Param ID path string true "ID"
// @Success 200 {object} http.Response{data=string} "Response data"
// @Failure 500 {object} http.Response{}
func (h *Handler) GetPhone(c *gin.Context) {
	id := c.Param("ID")

	resp, notExists, err := h.storage.Phone().Get(context.Background(), &models.Phone{
		ID: id,
	})

	if notExists {
		h.handleResponse(c, http.NotFound, err)
		return
	}

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Int()

	if random%3 == 0 {
		h.handleResponse(c, http.InternalServerError, errors.New("Random generated error."))
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Router /v1/phone [GET]
// @Summary returns all phone from server
// @Description this returns phone message is exists
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=string} "Response data"
// @Failure 500 {object} http.Response{}
func (h *Handler) GetAllPhones(c *gin.Context) {
	id := c.Param("ID")

	resp, err := h.storage.Phone().GetAll(context.Background(), &models.Phone{
		ID: id,
	})

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err)
		return
	}

	h.handleResponse(c, http.OK, resp)
}
