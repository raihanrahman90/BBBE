package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationRequest struct {
	Page   int
	Limit  int
	Offset int
}

type PaginationResponse struct {
	Page      int         `json:"page"`
	DataCount int         `json:"dataCount"`
	PageCount int         `json:"pageCount"`
	Data      interface{} `json:"data"`
}

func GeneratePaginationRequest(c *gin.Context) PaginationRequest {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	return PaginationRequest{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
}

func GeneratePaginationResponse(data interface{}, dataCount int64, pagination PaginationRequest) PaginationResponse {
	pageCount := int(dataCount) / pagination.Limit
	if (int(dataCount) % pagination.Limit) > 0 {
		pageCount++
	}
	return PaginationResponse{
		Page:      pagination.Page,
		DataCount: int(dataCount),
		Data:      data,
		PageCount: pageCount,
	}
}
