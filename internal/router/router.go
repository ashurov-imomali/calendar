package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/internal/service"
	"main/models"
	"net/http"
	"strconv"
)

type Handlers struct {
	Srv *service.Service
}

func InitHandler(srv *service.Service) *Handlers {
	return &Handlers{Srv: srv}
}

func GetHandlers(h *Handlers) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // Добавьте эту строку, если ваше приложение разрешает авторизацию с учетом учетных данных
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	events := r.Group("/event")
	events.GET("/", h.getEvents)
	events.DELETE("/:id", h.deleteEvents)
	events.PUT("/", h.updateEvents)
	return r
}

func (h *Handlers) getEvents(c *gin.Context) {
	events, err := h.Srv.GetEvents()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": events})
}

func (h *Handlers) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (h *Handlers) deleteEvents(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if err := h.Srv.DeleteEvent(id); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}

func (h *Handlers) updateEvents(c *gin.Context) {
	var events []models.Events
	err := c.ShouldBindJSON(&events)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	updateEvents, err := h.Srv.UpdateEvents(events)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updateEvents})
}
