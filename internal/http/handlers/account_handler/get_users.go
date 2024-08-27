package account_handler

import (
	"LoanTrackerApi/internal/http/handlers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMyProfile(c *gin.Context) {
	userId, err := handlers.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := h.usecase.GetMyProfile(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (h *Handler) AdminGetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	users, err := h.usecase.AdminGetAllUsers(c.Request.Context(), page, pageSize)

	baseUrl := c.Request.URL.Scheme + "://" + c.Request.Host + c.Request.URL.Path
	nextUrl := baseUrl + "?page=" + strconv.Itoa(page+1) + "&pageSize=" + strconv.Itoa(pageSize)
	prevUrl := baseUrl + "?page=" + strconv.Itoa(page) + "&pageSize=" + strconv.Itoa(pageSize)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users, "next": nextUrl, "prev": prevUrl})
}
