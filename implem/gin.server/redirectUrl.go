package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rH RouterHandler) listUrls(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	urls, err := rH.ucHandler.UrlList()
	if err != nil {
		log(err)
		c.Errors = append(c.Errors, &gin.Error{Err:errors.New("url not found"),Type:gin.ErrorTypePrivate})
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK,gin.H{"urls":urls})
}
