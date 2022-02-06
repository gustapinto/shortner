package controller

import (
	"github.com/gin-gonic/gin"
)

type ShortnerController struct {
	controller *Controller
}

func NewShortnerController() *ShortnerController {
	return &ShortnerController{}
}

func (c *ShortnerController) Index(ctx *gin.Context) {
	c.controller.Render(ctx, "index")
}

func (c *ShortnerController) Create(ctx *gin.Context) {
	c.controller.Render(ctx, "create")
}

func (c *ShortnerController) List(ctx *gin.Context) {
	c.controller.RenderWithData(ctx, "list", gin.H{})
}
