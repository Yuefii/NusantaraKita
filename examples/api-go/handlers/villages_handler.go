package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuefii/NusantaraKita/examples/api-go/dtos"
	"github.com/yuefii/NusantaraKita/examples/api-go/models"
)

func (handler *Handler) GetVillages(ctx *gin.Context) {
	var village []models.Villages
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

		result := handler.db.Find(&village)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		totalItems = int64(len(village))

		var villageDtos []dtos.ApiDTO
		for _, response := range village {
			villageDtos = append(villageDtos, dtos.ApiDTO{
				Code: response.Code,
				Name: response.Name,
			})
		}

		response := gin.H{
			"data": villageDtos,
		}

		ctx.JSON(http.StatusOK, response)
		return
	}

	result := handler.db.Model(&models.Villages{}).Count(&totalItems)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	result = handler.db.Offset((page - 1) * perPage).Limit(perPage).Find(&village)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var villageDTOs []dtos.ApiDTO
	for _, response := range village {
		villageDTOs = append(villageDTOs, dtos.ApiDTO{
			Code: response.Code,
			Name: response.Name,
		})
	}

	totalPages := int(totalItems) / perPage
	if int(totalItems)%perPage != 0 {
		totalPages++
	}

	response := dtos.ResponseDTO{
		Data: villageDTOs,
		Pagination: dtos.PaginationDTO{
			CurrentPage: page,
			PerPage:     perPage,
			TotalItems:  int(totalItems),
			TotalPages:  totalPages,
		},
	}
	ctx.JSON(http.StatusOK, response)
}
