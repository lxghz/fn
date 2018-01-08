package server

import (
	"net/http"
	"path"

	"github.com/fnproject/fn/api"
	"github.com/fnproject/fn/api/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleRouteDelete(c *gin.Context) {
	ctx := c.Request.Context()

	appIDorName := c.MustGet(api.App).(string)
	initApp := &models.App{Name: appIDorName, ID: appIDorName}
	app, err := s.datastore.GetApp(ctx, initApp)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}
	routePath := path.Clean(c.MustGet(api.Path).(string))

	if _, err := s.datastore.GetRoute(ctx, app, routePath); err != nil {
		handleErrorResponse(c, err)
		return
	}

	if err := s.datastore.RemoveRoute(ctx, app, routePath); err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Route deleted"})
}
