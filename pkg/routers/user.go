package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/githcorrado/fake-admagic/pkg/dao"
	"github.com/githcorrado/fake-admagic/pkg/request"
	"github.com/githcorrado/fake-admagic/pkg/response"
	"net/http"
)

func AddGroupUser(r *gin.Engine) {
	group := r.Group("/user")

	group.POST("/login", loginJSON)
	group.POST("/newLogin", createNewUser)
	group.POST("/createUserWithExistingName", createUserWithExistingName)
}


// @Summary Super User login TODO?
// @Produce json
// @Success 200 {object} response.LoginResponse "success"
// @Failure 404 {object} response.LoginResponse "unauthorized"
// @Router /user/login [post]
func loginJSON(c *gin.Context) {
	var json request.LoginForm
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(404, response.LoginResponse{Status: "error", Err: err.Error() })
		return
	}
	fmt.Println(json)
	if json.User != "Corrado" || json.Password != "321" {
		c.JSON(http.StatusUnauthorized, response.LoginResponse{Status: "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, response.LoginResponse{Status: "Welcome!"})
}

// @Summary New user creation
// @Produce json
// @Success 200 {object} response.LoginResponse "success"
// @Failure 404 {object} response.LoginResponse "DB error"
// @Failure 404 {object} response.LoginResponse "User name already exists"
// @Failure 404 {object} response.LoginResponse "JSON form error"
// @Router /user/newLogin [post]
func createNewUser(c *gin.Context) {
	var json request.LoginForm
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(404, response.LoginResponse{Status: "LoginForm JSON error", Err: err.Error() })
		return
	}
	currentUsers, dbErr := dao.GetAllUserNames()
	if dbErr != nil {
		c.JSON(404, response.LoginResponse{Status: "error.", Err: dbErr.Error() })
		return
	}
	exists := false
	for _, user := range currentUsers {
		fmt.Println(user)
		if user.User == json.User {
			exists = true
			break
		}
	}
	if exists == true {
		c.JSON(404, response.LoginResponse{Status: "User Name exists." })
		return
	}
	newUser, dbErr1 := dao.CreateUser(json)
	if dbErr1 != nil {
		c.JSON(404, response.LoginResponse{Status: "error.", Err: dbErr1.Error() })
		return
	}
	c.JSON(http.StatusOK, response.LoginResponse{Status: fmt.Sprintf("Welcome %s!", newUser.User)})
}

// @Summary Test New user creation
// @Produce json
// @Success 200 {object} response.LoginResponse "success"
// @Failure 404 {object} response.LoginResponse "User name already exists violates username primary key constraint"
// @Failure 404 {object} response.LoginResponse "JSON form error"
// @Router /user/createUserWithExistingName [post]
func createUserWithExistingName(c *gin.Context) {
	var json request.LoginForm
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(404, response.LoginResponse{Status: "LoginForm JSON error", Err: err.Error() })
		return
	}
	newUser, err1 := dao.CreateUser(json)
	if err1 != nil {
		c.JSON(404, response.LoginResponse{Status: "error.", Err: err1.Error() })
		return
	}
	c.JSON(http.StatusOK, response.LoginResponse{Status: fmt.Sprintf("Welcome %s!", newUser.User)})


}
