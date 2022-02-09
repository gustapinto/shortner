package controller

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Holder struct {
	Url string `json:"url"`
}

type ShortnerController struct {
	controller *Controller
	urls       map[string]string
}

func NewShortnerController() *ShortnerController {
	return &ShortnerController{
		urls: make(map[string]string),
	}
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

func getRandomUrl(host string, tlsConn *tls.ConnectionState) string {
	scheme := "http"
	if tlsConn != nil {
		scheme = "https"
	}

	rand.Seed(time.Now().UnixMicro())
	random := strconv.Itoa(rand.Intn(10000))

	shortned := fmt.Sprintf("%s://%s/%s", scheme, host, random)

	return shortned
}

func (c *ShortnerController) ShortUrl(ctx *gin.Context) {
	shortned := getRandomUrl(ctx.Request.Host, ctx.Request.TLS)

	var holder Holder

	if err := ctx.BindJSON(&holder); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{})
	}

	c.urls[shortned] = holder.Url

	ctx.IndentedJSON(http.StatusCreated, gin.H{})
}
