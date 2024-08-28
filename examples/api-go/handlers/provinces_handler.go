package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuefii/NusantaraKita/examples/api-go/dtos"
	"github.com/yuefii/NusantaraKita/examples/api-go/models"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

func (handler *Handler) GetProvinces(ctx *gin.Context) {
	var province []models.Provinces
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

		result := handler.db.Find(&province)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		totalItems = int64(len(province))

		var provinceDTOs []dtos.ProvinceDTO
		for _, r := range province {
			provinceDTOs = append(provinceDTOs, dtos.ProvinceDTO{
				Code: r.Code,
				Name: r.Name,
			})
		}

		response := gin.H{
			"data": provinceDTOs,
		}
		ctx.JSON(http.StatusOK, response)
		return
	}

	result := handler.db.Model(&models.Provinces{}).Count(&totalItems)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	result = handler.db.Offset((page - 1) * perPage).Limit(perPage).Find(&province)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var provinceDTOs []dtos.ProvinceDTO
	for _, response := range province {
		provinceDTOs = append(provinceDTOs, dtos.ProvinceDTO{
			Code: response.Code,
			Name: response.Name,
		})
	}

	totalPages := int(totalItems) / perPage
	if int(totalItems)%perPage != 0 {
		totalPages++
	}

	response := dtos.ResponseDTO{
		Data: provinceDTOs,
		Pagination: dtos.PaginationDTO{
			CurrentPage: page,
			PerPage:     perPage,
			TotalItems:  int(totalItems),
			TotalPages:  totalPages,
		},
	}

	ctx.JSON(http.StatusOK, response)
}
