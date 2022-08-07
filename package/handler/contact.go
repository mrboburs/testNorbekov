package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrboburs/Norbekov/model"
	// "norbekov/model"
)

// @Summary Create Contact
// @Tags Contact
// @Description create contact_post
// @ID create-contact_post
// @Accept  json
// @Produce  json
// @Param input body model.Contact true "home info"
// @Success 200 {object} ResponseSuccess
// @Failure 400,404 {object} errorResponse
// @Failure 409 {object} errorResponseData
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /contact/create [post]
func (handler *Handler) CreateContactPost(ctx *gin.Context) {
	logrus := handler.logrus
	var input model.Contact
	err := ctx.BindJSON(&input)

	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	Id, err := handler.services.Contact.CreateContactPost(input, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, ResponseSuccess{Data: Id, Message: "DONE"})
}

// @Summary Get contacts
// @Tags Contact
// @Description get  contacts
// @ID get-contacts
// @Accept  json
// @Produce  json
// @Success 200 {object} model.allContacts
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /contact/get [GET]
func (handler *Handler) GetAllContact(ctx *gin.Context) {
	logrus := handler.logrus

	contacts, err := handler.services.Contact.GetAllContact(logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusInternalServerError, err.Error(), logrus)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": contacts,
	})
}
