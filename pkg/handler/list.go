package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) creatList(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSONP(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}
