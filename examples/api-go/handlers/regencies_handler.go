package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuefii/NusantaraKita/examples/api-go/dtos"
	"github.com/yuefii/NusantaraKita/examples/api-go/models"
)

func (handler *Handler) GetRegencies(ctx *gin.Context) {
	var regency []models.Regencies
	var totalItems int64

	showAll := ctx.Query("show_all") == "true"
	pageStr := ctx.Query("page")
	perPageStr := ctx.Query("per_page")

	page := 1
	perPage := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if perPageStr != "" {
		if pp, err := strconv.Atoi(perPageStr); err == nil && pp > 0 {
			perPage = pp
		}
	}

	if showAll {

		result := handler.db.Find(&regency)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		totalItems = int64(len(regency))

		var regenciesDTOs []dtos.ApiDTO
		for _, response := range regency {
			regenciesDTOs = append(regenciesDTOs, dtos.ApiDTO{
				Code: response.Code,
				Name: response.Name,
			})
		}

		response := gin.H{
			"data": regenciesDTOs,
		}

		ctx.JSON(http.StatusOK, response)
		return
	}

	result := handler.db.Model(&models.Regencies{}).Count(&totalItems)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	result = handler.db.Offset((page - 1) * perPage).Limit(perPage).Find(&regency)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var regenciesDTOs []dtos.ApiDTO
	for _, response := range regency {
		regenciesDTOs = append(regenciesDTOs, dtos.ApiDTO{
			Code: response.Code,
			Name: response.Name,
		})
	}

	totalPages := int(totalItems) / perPage
	if int(totalItems)%perPage != 0 {
		totalPages++
	}

	response := dtos.ResponseDTO{
		Data: regenciesDTOs,
		Pagination: dtos.PaginationDTO{
			CurrentPage: page,
			PerPage:     perPage,
			TotalItems:  int(totalItems),
			TotalPages:  totalPages,
		},
	}
	ctx.JSON(http.StatusOK, response)
}
