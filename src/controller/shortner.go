package controller

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
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

func (c *ShortnerController) ShortUrl(ctx *gin.Context) {
	var holder Holder

	if err := ctx.BindJSON(&holder); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{})
	}

	if !strings.Contains(holder.Url, "http://") && !strings.Contains(holder.Url, "https://") {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "The url must be valid, containing either http:// or https://",
		})
		return
	}

	shortned := c.getRandomUrl(ctx.Request.Host, ctx.Request.TLS)
	c.urls[shortned] = holder.Url

	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"newUrl": shortned,
	})
}

func (c *ShortnerController) GetUrls(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, c.urls)
}

func (c *ShortnerController) Redirect(ctx *gin.Context) {
	scheme := c.getScheme(ctx.Request.TLS)
	requestUrl := fmt.Sprintf("%s://%s%s", scheme, ctx.Request.Host, ctx.Request.URL.Path)
	redirectTo := c.urls[requestUrl]

	ctx.Redirect(http.StatusMovedPermanently, redirectTo)
}

func (c *ShortnerController) getRandomUrl(host string, tlsConn *tls.ConnectionState) string {
	rand.Seed(time.Now().UnixMicro())
	random := strconv.Itoa(rand.Intn(1000000))
	scheme := c.getScheme(tlsConn)

	shortned := fmt.Sprintf("%s://%s/r/%s", scheme, host, random)

	for _, url := range c.urls {
		if shortned == url {
			return c.getRandomUrl(host, tlsConn)
		}
	}

	return shortned
}

func (c *ShortnerController) getScheme(tlsConn *tls.ConnectionState) string {
	if tlsConn != nil {
		return "https"
	}

	return "http"
}
