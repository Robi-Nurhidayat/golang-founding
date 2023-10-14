package handler

import (
	"bwa/golang/campaign"
	"bwa/golang/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type campaignHandler struct {
	service campaign.CampaignService
}

func NewCampaignHandler(service campaign.CampaignService) *campaignHandler {
	return &campaignHandler{service: service}
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userId)
	if err != nil {
		response := helper.ApiResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("List Of campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)

}
