package handler

import (
	"net/http"
	// "norbekov/model"
	// "norbekov/util/logrus"

	"github.com/gin-gonic/gin"
	"github.com/mrboburs/Norbekov/model"
)

// @Summary Create Admin
// @Tags Admin
// @Description create admin_post
// @ID create-admin_post
// @Accept  json
// @Produce  json
// @Param input body model.Admin true "home info"
// @Success 200 {object} ResponseSuccess
// @Failure 400,404 {object} errorResponse
// @Failure 409 {object} errorResponseData
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /admin/create [post]
func (handler *Handler) CreateAdmin(ctx *gin.Context) {
	logrus := handler.logrus
	var input model.Admin
	err := ctx.BindJSON(&input)

	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	Id, err := handler.services.Admin.CreateAdmin(input, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, ResponseSuccess{Data: Id, Message: "DONE"})
}

// @Summary Get Delete Admin
// @Tags Admin
// @Description delete admin
// @ID delete-admin
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Success 200 {object} ResponseSuccess
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /admin/delete [DELETE]
func (h *Handler) DeleteAdmin(ctx *gin.Context) {

	logrus := h.logrus
	id := ctx.Query("id")

	h.services.Admin.DeleteAdmin(id, logrus)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":      id,
		"message": "deleted",
	})
}

// @Summary Login Admin
// @Tags Admin
// @Description login admin
// @ID login-admin
// @Accept  json
// @Produce  json
// @Param input body model.Admin true "admin info"
// @Success 200 {object} ResponseSuccess
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /admin/login [post]
func (h *Handler) LoginAdmin(ctx *gin.Context) {
	logrus := h.logrus
	var input model.Admin

	if err := ctx.BindJSON(&input); err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	token, err := h.services.Admin.GenerateToken(input.UserName, input.Passord, logrus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"server_reply": "wrong passord or username",
		})
		return
	}
	_, err = h.services.Admin.GetAdmin(input.UserName, input.Passord, logrus)
	if err != nil {
		logrus.Fatal(err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"server_reply": "wrong passord or username",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"loginStatus": "successfully logged in ",
		"data":        input,
		"toke":        token,
	})

}
