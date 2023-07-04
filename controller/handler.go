package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zileyuan/go-working-calendar/response"
	"github.com/zileyuan/go-working-calendar/service"
)

type StatusType int

func HolidayAction(c *gin.Context) {
	date := c.Param("date")
	holiday, status := service.Holiday(date)
	resp := response.Message{
		Status: status,
		Data:   holiday,
	}
	c.JSON(http.StatusOK, resp)
}

func CountAction(c *gin.Context) {
	from := c.Param("from")
	to := c.Param("to")
	length, status := service.Count(from, to)
	resp := response.Message{
		Status: status,
		Data:   length,
	}
	c.JSON(http.StatusOK, resp)
}

func CalcAction(c *gin.Context) {
	from := c.Param("from")
	amount := c.Param("amount")
	length, status := service.Calc(from, amount)
	resp := response.Message{
		Status: status,
		Data:   length,
	}
	c.JSON(http.StatusOK, resp)
}
