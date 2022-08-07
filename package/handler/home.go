package handler

import (
	"net/http"
	"strconv"

	// "strconv"

	"github.com/gin-gonic/gin"

	"github.com/mrboburs/Norbekov/model"
	// "norbekov/util/logrus"
	// "norbekov/util/logrus"
)

// @Summary Get homes
// @Tags Home
// @Description get  homes
// @ID get-homes
// @Accept  json
// @Produce  json
// @Success 200 {object} model.allHome
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /homes/get [GET]
func (handler *Handler) GetAllHome(ctx *gin.Context) {
	logrus := handler.logrus

	contacts, err := handler.services.Home.GetAllHome(logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusInternalServerError, err.Error(), logrus)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": contacts,
	})
}

// @Summary Create HomePost
// @Tags Home
// @Description create home_post
// @ID create-home_post
// @Accept  json
// @Produce  json
// @Param input body model.HomePost true "home info"
// @Success 200 {object} ResponseSuccess
// @Failure 400,404 {object} errorResponse
// @Failure 409 {object} errorResponseData
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/home/create [post]
//@Security ApiKeyAuth
func (handler *Handler) CreateHomePost(ctx *gin.Context) {
	logrus := handler.logrus
	var input model.HomePost

	err := ctx.BindJSON(&input)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	homeId, err := handler.services.CreateHomePost(input, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "post created",
		"id":     homeId,
	})
}

// @Summary Upload Home Image
// @Description Upload Home Image
// @ID upload-image
// @Tags   Home
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
// @Router   /api/home/upload-img/{id} [PATCH]
//@Security ApiKeyAuth
func (handler *Handler) uploadHomeImage(ctx *gin.Context) {
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

	imageURL, err := handler.services.Home.UploadImage(file, header, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	effectedRowsNum, err := handler.services.Home.UpdateHomeImage(id, imageURL, logrus)
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

// @Summary Update  Home By ID
// @Tags Home
// @Description Update home by id
// @ID update-home-id
// @Accept  json
// @Produce  json
// @Param        id   path  int     true "Param ID"
// @Param input body model.HomePost true "home info"
// @Success 200 {object} ResponseSuccess
// @Failure 400,404 {object} errorResponse
// @Failure 409 {object} errorResponseData
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/home/update/{id} [PUT]
//@Security ApiKeyAuth
func (h *Handler) updateHome(ctx *gin.Context) {
	logrus := h.logrus
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Errorf("syntax error")
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	var input model.HomePost
	err = ctx.BindJSON(&input)

	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	effectedRowsNum, err := h.services.Home.UpdateHome(id, input, logrus)
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

// @Summary Get home
// @Tags Home
// @Description get home post
// @ID get-home
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Success 200 {object} model.HomeFull
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /home/get [GET]
func (h *Handler) GetHomeById(ctx *gin.Context) {
	logrus := h.logrus
	id := ctx.Query("id")

	post, err := h.services.Home.GetHomeById(id, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusInternalServerError, err.Error(), logrus)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":   id,
		"post": post,
	})
}

// @Summary Get Delete Post
// @Tags Home
// @Description delete post
// @ID delete-posts
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Success 200 {object} model.HomeFull
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/home/delete [DELETE]
//@Security ApiKeyAuth
func (h *Handler) DeleteHome(ctx *gin.Context) {

	logrus := h.logrus
	id := ctx.Query("id")

	h.services.Home.DeleteHome(id, logrus)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":      id,
		"message": "deleted",
	})
}
