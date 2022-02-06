package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (c *Controller) Render(ctx *gin.Context, template string) {
	c.RenderWithData(ctx, template, gin.H{})
}

func (c *Controller) RenderWithData(ctx *gin.Context, template string, data gin.H) {
	if !strings.Contains(template, ".tmpl") {
		template = fmt.Sprintf("%s.tmpl", template)
	}

	ctx.HTML(http.StatusOK, template, data)
}
