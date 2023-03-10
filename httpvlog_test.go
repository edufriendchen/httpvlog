package httpvlog

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/route"
	"testing"
)

func TestHttpVlogConsoleColor(t *testing.T) {
	router := route.NewEngine(config.NewOptions([]config.Option{}))
	//The default console output has colors
	ForceConsoleColor()
	router.Use(Logger())
	router.GET("/ping/:user", func(ctx context.Context, c *app.RequestContext) {
		user := c.Param("user")
		assert.DeepEqual(t, "close", c.Request.Header.Get("Connection"))
		c.Response.SetConnectionClose()
		c.JSON(201, map[string]string{"hi": user})
	})
	w := ut.PerformRequest(router, "GET", "/ping/FriendChen", &ut.Body{Body: bytes.NewBufferString("1"), Len: 1},
		ut.Header{Key: "Connection", Value: "close"})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "{\"hi\":\"FriendChen\"}", string(resp.Body()))
}

func TestHttpVlogDisableConsoleColor(t *testing.T) {
	router := route.NewEngine(config.NewOptions([]config.Option{}))
	//Turn off console output color
	DisableConsoleColor()
	router.Use(Logger())
	router.GET("/ping/:user", func(ctx context.Context, c *app.RequestContext) {
		user := c.Param("user")
		assert.DeepEqual(t, "close", c.Request.Header.Get("Connection"))
		c.Response.SetConnectionClose()
		c.JSON(201, map[string]string{"hi": user})
	})
	w := ut.PerformRequest(router, "GET", "/ping/FriendChen", &ut.Body{Body: bytes.NewBufferString("1"), Len: 1},
		ut.Header{Key: "Connection", Value: "close"})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "{\"hi\":\"FriendChen\"}", string(resp.Body()))
}
