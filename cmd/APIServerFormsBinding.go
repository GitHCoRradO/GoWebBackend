package main

import (
	"context"
	"errors"
	"github.com/githcorrado/fake-admagic/pkg/dao"
	_ "github.com/githcorrado/fake-admagic/pkg/database"
	"github.com/githcorrado/fake-admagic/pkg/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	/*
	router := gin.Default()

	router.POST("/loginJSON", func(c *gin.Context) {
		var json request.LoginForm
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(404, gin.H { "error": err.Error() })
			return
		}
		fmt.Println(json)
		if json.User != "Corrado" || json.Password != "321" {
			c.JSON(http.StatusUnauthorized, gin.H {"status": "unauthorized" })
			return
		}
		c.JSON(http.StatusOK, gin.H {"status": "welcome!"})
	})

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", func(fl validator.FieldLevel) bool {
			date, ok := fl.Field().Interface().(time.Time)
			if ok {
				today := time.Now()
				if today.After(date) {
					return false
				}
			}
			return true
		})
	}

	router.GET("/booking", func(c *gin.Context) {
		var b request.Booking
		if err := c.ShouldBindWith(&b, binding.Query); err == nil {
			c.JSON(200, gin.H {"message":"Booking dates are valid"})
		} else {
			c.JSON(404, gin.H {"error": err.Error() })
		}
	})

	router.GET("/test-header", func(c *gin.Context) {
		var h request.HeadForm
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(404, gin.H { "error": err.Error() })
			return
		}
		fmt.Println("%#v", h)
		c.File("/Users/luolingxiao/projects/myrepos/fake-ADMagic/log/test-file.log")
	})
	 */

	dao.InitTables()
	/*
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=fake_admagic sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&request.LoginForm{})
	db.Create(&request.LoginForm{User: "Corrado", Password: "321"})
	 */


	router := routers.NewRouter()

	srv := &http.Server {
		Addr: ":8000",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server ....")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")

}
