package handlers

import (
	" hery-ciaputra/final-project-backend/dto"
	" hery-ciaputra/final-project-backend/httperror"
	" hery-ciaputra/final-project-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) Profile(c *gin.Context) {
	user, _ := c.Get("user")
	ur := user.(models.User)

	user, err := h.authService.Profile(ur.ID)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	user, _ := c.Get("user")
	ur := user.(models.User)
	payload, _ := c.Get("payload")
	profileReq := payload.(*dto.ProfileReq)

	user, err := h.authService.UpdateProfile(ur.ID, profileReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) TopUp(c *gin.Context) {
	user, _ := c.Get("user")
	ur := user.(models.User)
	value, _ := c.Get("payload")
	topUpReq := value.(*dto.TopUpReq)

	result, err := h.authService.TopUp(ur.ID, topUpReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) CreateShipping(c *gin.Context) {
	payload, _ := c.Get("payload")
	shippingReq := payload.(*dto.ShippingReq)

	result, err := h.authService.CreateShipping(shippingReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) ShippingList(c *gin.Context) {
	user, _ := c.Get("user")
	ur := user.(models.User)

	shippingList, err := h.authService.ShippingList(ur.ID)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, shippingList)
}

func (h *Handler) Address(c *gin.Context) {
	user, _ := c.Get("user")
	ur := user.(models.User)

	addresses, err := h.authService.Address(ur.ID)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, addresses)
}

func (h *Handler) CreateAddress(c *gin.Context) {
	user, _ := c.Get("user")
	ur := user.(models.User)
	payload, _ := c.Get("payload")
	addressReq := payload.(*dto.AddressReq)

	result, err := h.authService.CreateAddress(ur.ID, addressReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) EditAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	payload, _ := c.Get("payload")
	addressReq := payload.(*dto.AddressReq)

	result, err := h.authService.EditAddress(id, addressReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) Payment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("shipping-id"))
	payload, _ := c.Get("payload")
	paymentReq := payload.(*dto.PaymentReq)

	result, err := h.authService.Payment(id, paymentReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) AddPromo(c *gin.Context) {
	payload, _ := c.Get("payload")
	promoReq := payload.(*dto.PromoReq)

	result, err := h.authService.AddPromo(promoReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) PromoList(c *gin.Context) {
	promos, err := h.authService.PromoList()
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, promos)
}

func (h *Handler) EditPromo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	payload, _ := c.Get("payload")
	promoReq := payload.(*dto.PromoReq)

	promo, err := h.authService.EditPromo(id, promoReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, promo)
}

func (h *Handler) SizeList(c *gin.Context) {
	sizes, err := h.authService.SizeList()
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, sizes)
}

func (h *Handler) CategoryList(c *gin.Context) {
	categories, err := h.authService.CategoryList()
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *Handler) AddOnList(c *gin.Context) {
	addOns, err := h.authService.AddOnList()
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, addOns)
}
