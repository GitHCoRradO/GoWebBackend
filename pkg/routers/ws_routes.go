package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

func AddGroupWSRoutes(r *gin.Engine) {
	group := r.Group("/")

	group.GET("/text", textApi)
	group.GET("/json", jsonApi)
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func textApi(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	defer ws.Close()
	mt, message, err := ws.ReadMessage()
	if err != nil {
		log.Println("error read message")
		log.Fatal(err)
	}

	var count = 0
	for {
		count++
		if count > 10 {
			break
		}
		message = []byte(string(message) + " " + strconv.Itoa(count))
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("error write message: " + err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}

func jsonApi(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	defer ws.Close()
	var data struct {
		A string `json:"a"`
		B int `json:"b"`
	}
	err = ws.ReadJSON(&data)
	if err != nil {
		log.Println("error read message")
		log.Fatal(err)
	}
	var count = 0
	for {
		count++
		if count > 10 {
			break
		}
		err = ws.WriteJSON(struct {
			A string `json:"a"`
			B int `json:"b"`
			C int `json:"c"`
		}{
			A: data.A,
			B: data.B,
			C: count,
		})
		if err != nil {
			log.Println("error write json: " + err.Error())
		}
		time.Sleep(1 * time.Second)

	}
}

