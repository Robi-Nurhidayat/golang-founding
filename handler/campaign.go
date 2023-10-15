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

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userId)
	if err != nil {
		response := helper.ApiResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if len(campaigns) == 0 {
		response := helper.ApiResponse("Error to get campaigns", http.StatusBadRequest, "error", []campaign.Campaign{})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// dari ini
	var newCampaigns []campaign.CompaignFormatter

	for _, value := range campaigns {
		campainFormatter := campaign.FormatterCampaign(value)
		newCampaigns = append(newCampaigns, campainFormatter)
	}

	// sampe atas ini, nanti kita repactor

	response := helper.ApiResponse("List Of campaigns", http.StatusOK, "success", newCampaigns)
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) GetCampaign(c *gin.Context) {

	var input campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.ApiResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignById, err := h.service.GetCampaign(input)

	if err != nil {
		response := helper.ApiResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Campaign detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignById))
	c.JSON(http.StatusOK, response)
}
