package handler

import "github.com/gin-gonic/gin"

func (h *handler) CreateStudent(c *gin.Context) {
	h.storage.Student().CreateStudent()
}