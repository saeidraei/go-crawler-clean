package server

import (
	"github.com/saeidraei/go-crawler-clean/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UrlReq struct {
	Url struct {
		Address *string `json:"address"`
	} `json:"url,required"`
}

func urlFromRequest(req *UrlReq) domain.Url {
	return domain.Url{
		Address: *req.Url.Address,
	}
}

func (rH RouterHandler) addUrl(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	req := &UrlReq{}
	if err := c.BindJSON(req); err != nil {
		log(err)
		c.Errors = append(c.Errors, &gin.Error{Err:err,Type:gin.ErrorTypePublic})
		c.Status(http.StatusBadRequest)
		return
	}

	err := rH.ucHandler.UrlPost(urlFromRequest(req))
	if err != nil {
		log(err)
		c.Errors = append(c.Errors, &gin.Error{Err:err,Type:gin.ErrorTypePublic})
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "successful"})
}
