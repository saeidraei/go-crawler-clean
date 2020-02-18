package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saeidraei/go-crawler-clean/uc"
)

type RouterHandler struct {
	ucHandler uc.Handler
	Logger    uc.Logger
}

func NewRouter(i uc.Handler) RouterHandler {
	return RouterHandler{
		ucHandler: i,
	}
}

func NewRouterWithLogger(i uc.Handler, logger uc.Logger) RouterHandler {
	return RouterHandler{
		ucHandler: i,
		Logger:    logger,
	}
}

func (rH RouterHandler) SetRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(rH.errorCatcher())
	rH.urlRoutes(api)

}

func (rH RouterHandler) urlRoutes(api *gin.RouterGroup) {
	profiles := api.Group("url")
	profiles.GET("/list", rH.listUrls)
	profiles.POST("", rH.addUrl) //
}

func (rH RouterHandler) errorCatcher() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() > 399 {
			rb := c.Errors.JSON()
			//if there was no error message show this message
			if c.Errors == nil {
				rb = gin.H{"error": "something went wrong"}
			}
			c.JSON(c.Writer.Status(), rb)
		}
	}
}

// log is used to "partially apply" the title to the rH.logger.Log function
// so we can see in the logs from which route the log comes from
func (rH RouterHandler) log(title string) func(...interface{}) {
	return func(logs ...interface{}) {
		rH.Logger.Log(title, logs)
	}
}

func (RouterHandler) MethodAndPath(c *gin.Context) string {
	return fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path)
}
