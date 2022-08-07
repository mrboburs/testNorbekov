package handler

import (
	// "fmt"
	"fmt"
	// "net/http"
	// "norbekov/docs"
	// "github.com/gin-contrib/cors"
	"github.com/mrboburs/Norbekov/docs"
	"github.com/mrboburs/Norbekov/package/service"
	"github.com/mrboburs/Norbekov/util/logrus"

	"github.com/mrboburs/Norbekov/configs"
	_ "github.com/mrboburs/Norbekov/docs"

	// "github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
	logrus   *logrus.Logger
	config   *configs.Configs
}

func NewHandler(services *service.Service, logrus *logrus.Logger, config *configs.Configs) *Handler {
	return &Handler{services: services, logrus: logrus, config: config}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	config := handler.config
	fmt.Println(config)

	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.New()

	// docs.SwaggerInfo_swagger.Version = config.Version
	// docs.SwaggerInfo_swagger.Host = config.ServiceHost + config.HTTPPort
	// docs.SwaggerInfo_swagger.Schemes = []string{ "https"}
	router.Use(CORSMiddleware())
	path := config.PhotoPath
	router.Static("/public", path)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	admin := router.Group("/admin")
	{
		admin.POST("/create", handler.CreateAdmin)
		admin.DELETE("/delete", handler.DeleteAdmin)
		admin.POST("/login", handler.LoginAdmin)
	}
	api := router.Group("/api", handler.userIdentity)
	home := api.Group("/home")
	homes := router.Group("/homes")
	{
		home.POST("/create", handler.CreateHomePost)
		home.PATCH("/upload-img/:id", handler.uploadHomeImage)
		home.PUT("/update/:id", handler.updateHome)
		router.Group("/home").GET("/get", handler.GetHomeById)
		home.DELETE("/delete", handler.DeleteHome)
		homes.GET("/get", handler.GetAllHome)

	}
	news := api.Group("/news")
	newS := router.Group("/newsS")
	{
		news.POST("/create", handler.CreateNewsPost)
		news.PATCH("/upload-img/:id", handler.uploadNewsImage)
		news.PUT("/update/:id", handler.updateNews)
		router.Group("/news").GET("/get", handler.GetNewsById)
		news.DELETE("/delete", handler.DeleteNews)
		newS.GET("/get", handler.GetAllNews)
	}
	service := api.Group("/service")
	services := router.Group("/services")
	{
		service.POST("/create", handler.CreateServicePost)
		service.PATCH("/upload-img/:id", handler.uploadServiceImage)
		service.PUT("/update/:id", handler.UpdateService)
		router.Group("/service").GET("/get", handler.GetServiceById)
		service.DELETE("/delete", handler.DeleteService)
		services.GET("/get", handler.GetAllService)
	}
	table := api.Group("/table")
	tables := router.Group("/tables")
	{
		table.POST("/create", handler.CreateTablePost)
		table.POST("/course/create", handler.CreateCoursePost)
		table.PATCH("/upload-img/:id", handler.uploadTableImage)
		table.PUT("/update/:id", handler.UpdateTable)
		router.Group("/table").GET("/get", handler.GetTableById)
		table.DELETE("/delete", handler.DeleteTable)
		table.DELETE("/course/delete", handler.DeleteCourse)
		tables.GET("/get", handler.GetAllTable)
		tables.GET("/course/get", handler.GetAllCourse)
	}

	contact := router.Group("/contact")

	{
		contact.POST("/create", handler.CreateContactPost)
		contact.GET("/get", handler.GetAllContact)
	}

	return router
}
