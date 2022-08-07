package handler

import (
	"net/http"
	"strconv"

	// "strconv"

	"github.com/gin-gonic/gin"

	"github.com/mrboburs/Norbekov/model"
	// "norbekov/util/logrus"
)

// @Summary Get newS
// @Tags News
// @Description get  news
// @ID get-newsS
// @Accept  json
// @Produce  json
// @Success 200 {object} model.allNews
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /newsS/get [GET]
func (handler *Handler) GetAllNews(ctx *gin.Context) {
	logrus := handler.logrus

	contacts, err := handler.services.News.GetAllNews(logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusInternalServerError, err.Error(), logrus)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": contacts,
	})
}

// @Summary Create NewsPost
// @Tags News
// @Description create news_post
// @ID create-news_post
// @Accept  json
// @Produce  json
// @Param input body model.NewsPost true "home info"
// @Success 200 {object} ResponseSuccess
// @Failure 400,404 {object} errorResponse
// @Failure 409 {object} errorResponseData
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/news/create [post]
//@Security ApiKeyAuth
func (handler *Handler) CreateNewsPost(ctx *gin.Context) {
	logrus := handler.logrus
	var input model.NewsPost
	err := ctx.BindJSON(&input)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	homeId, err := handler.services.CreateNewsPost(input, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, ResponseSuccess{Data: homeId, Message: "DONE"})
}

// @Summary Upload News Image
// @Description Upload News Image
// @ID uploadImg
// @Tags   News
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
// @Router   /api/news/upload-img/{id} [PATCH]
//@Security ApiKeyAuth
func (handler *Handler) uploadNewsImage(ctx *gin.Context) {
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

	effectedRowsNum, err := handler.services.News.UpdateNewsImage(id, imageURL, logrus)
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

// @Summary Update  News By ID
// @Tags News
// @Description Update news by id
// @ID update-news-id
// @Accept  json
// @Produce  json
// @Param        id   path  int     true "Param ID"
// @Param input body model.NewsPost true "home info"
// @Success 200 {object} ResponseSuccess
// @Failure 400,404 {object} errorResponse
// @Failure 409 {object} errorResponseData
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/news/update/{id} [PUT]
//@Security ApiKeyAuth
func (h *Handler) updateNews(ctx *gin.Context) {
	logrus := h.logrus
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Errorf("syntax error")
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	var input model.NewsPost
	err = ctx.BindJSON(&input)

	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	effectedRowsNum, err := h.services.News.UpdateNews(id, input, logrus)
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

// @Summary Get news
// @Tags News
// @Description get news post
// @ID get-news
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Success 200 {object}  model.NewsFull
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /news/get [GET]
func (h *Handler) GetNewsById(ctx *gin.Context) {
	logrus := h.logrus
	id := ctx.Query("id")

	post, err := h.services.News.GetNewsById(id, logrus)
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusInternalServerError, err.Error(), logrus)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":   id,
		"post": post,
	})
}

// @Summary Get Delete news
// @Tags News
// @Description delete news
// @ID delete-news
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Success 200 {object} model.NewsFull
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/news/delete [DELETE]
//@Security ApiKeyAuth
func (h *Handler) DeleteNews(ctx *gin.Context) {

	logrus := h.logrus
	id := ctx.Query("id")

	h.services.News.DeleteNews(id, logrus)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":      id,
		"message": "deleted",
	})
}
