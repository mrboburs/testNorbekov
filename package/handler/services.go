package handler

import (
	"github.com/mrboburs/Norbekov/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get services
// @Tags Service
// @Description get  service
// @ID get-services
// @Accept  json
// @Produce  json
// @Success 200 {object} model.allService
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /services/get [GET]
func (handler *Handler) GetAllService(ctx *gin.Context) {
	logrus := handler.logrus

	contacts, err := handler.services.Services.GetAllService(logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusInternalServerError, err.Error(), logrus)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": contacts,
	})
}

// @Summary Create ServicePost
// @Tags Service
// @Description create service_post
// @ID create-service_post
// @Accept  json
// @Produce  json
// @Param input body model.ServicePost true "home info"
// @Success 200 {object} ResponseSuccess
// @Failure 400,404 {object} errorResponse
// @Failure 409 {object} errorResponseData
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/service/create [post]
//@Security ApiKeyAuth
func (handler *Handler) CreateServicePost(ctx *gin.Context) {
	logrus := handler.logrus
	var input model.ServicePost
	err := ctx.BindJSON(&input)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	Id, err := handler.services.CreateServicePost(input, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":     Id,
		"status": "created",
	})
}

// @Summary Upload Service Image
// @Description Upload Service Image
// @ID uploadImgOfService
// @Tags   Service
// @Accept  json
// @Produce   json
// @Produce application/octet-stream
// @Produce image/png
// @Produce image/jpeg
// @Produce image/jpg
// @Param        id   path  int     true "Param ID"
// @Param file formData file true "file"
// @Accept multipart/form-data
// @Success      200   {object}      ResponseSuccess
// @Failure 400,404 {object} errorResponse
// @Failure 409 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router   /api/service/upload-img/{id} [PATCH]
//@Security ApiKeyAuth
func (handler *Handler) uploadServiceImage(ctx *gin.Context) {
	logrus := handler.logrus
	// homeId := ctx.Param("id")
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Errorf("syntax error")
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.Request.ParseMultipartForm(10 << 20)
	file, header, err := ctx.Request.FormFile("file")

	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	imageURL, err := handler.services.UploadImage(file, header, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	effectedRowsNum, err := handler.services.Services.UpdateServiceImage(id, imageURL, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	if effectedRowsNum == 0 {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, "User not found", logrus)
		return
	}
	ctx.JSON(http.StatusOK, ResponseSuccess{Message: "Uploaded", Data: imageURL})
}

// @Summary Update  Service By ID
// @Tags Service
// @Description Update service by id
// @ID update-service-id
// @Accept  json
// @Produce  json
// @Param        id   path  int     true "Param ID"
// @Param input body model.ServicePost true "home info"
// @Success 200 {object} ResponseSuccess
// @Failure 400,404 {object} errorResponse
// @Failure 409 {object} errorResponseData
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/service/update/{id} [PUT]
//@Security ApiKeyAuth
func (h *Handler) UpdateService(ctx *gin.Context) {
	logrus := h.logrus
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Errorf("syntax error")
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	var input model.ServicePost
	err = ctx.BindJSON(&input)

	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	effectedRowsNum, err := h.services.Services.UpdateService(id, input, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	if effectedRowsNum == 0 {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, "User not found", logrus)
		return
	}
	ctx.JSON(http.StatusOK, ResponseSuccess{Message: "Updated", Data: id})

}

// @Summary Get Service
// @Tags Service
// @Description get service post
// @ID get-service
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Success 200 {object}  model.ServiceFull
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /service/get [GET]
func (h *Handler) GetServiceById(ctx *gin.Context) {
	logrus := h.logrus
	id := ctx.Query("id")

	post, err := h.services.Services.GetServiceById(id, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusInternalServerError, err.Error(), logrus)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":   id,
		"post": post,
	})
}

// @Summary Get Delete service
// @Tags Service
// @Description delete service
// @ID delete-service
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Success 200 {object}  model.ServiceFull
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/service/delete [DELETE]
//@Security ApiKeyAuth
func (h *Handler) DeleteService(ctx *gin.Context) {

	logrus := h.logrus
	id := ctx.Query("id")

	h.services.Services.DeleteService(id, logrus)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":      id,
		"message": "deleted",
	})
}
