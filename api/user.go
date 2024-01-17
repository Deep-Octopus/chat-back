package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-chat/config"
	resp "go-chat/middleware"
	"go-chat/models"
	"go-chat/service"
	"go-chat/utils"
	"net/http"
	"strconv"
	"time"
)

type UserApi struct {
}

// GetUserList
// @Tags 获取用户列表
// @Produce json
// @Success 200
// @Router /user/getUserList [get]
func (u UserApi) GetUserList(c *gin.Context) {
	userList := service.GetUserList()
	c.JSON(http.StatusOK, resp.OK.WithData(userList))
}

// CreateUser
// @Tags 新增用户
// @Success 200
// @Router /user/createUser [post]
func (u UserApi) CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, resp.ErrParam)
		return
	}
	if err := service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, resp.Err.WithMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.OK.WithMsg("新增成功"))
}

// Login
// @Tags 登录
// @Success 200
// @Router /user/login [post]
func (u UserApi) Login(c *gin.Context) {
	user := models.UserBasic{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, resp.ErrParam)
		return
	}
	tmpUser := service.GetUserByName(user.Name)
	if tmpUser.ID == 0 {
		c.JSON(http.StatusNotFound, resp.Err.WithMsg("用户不存在"))
		return
	}
	if !utils.CheckPasswordHash(user.Password, tmpUser.Password) {
		c.JSON(http.StatusBadRequest, resp.Err.WithMsg("密码错误"))
		return
	}
	identity, _ := utils.GenToken(int64(tmpUser.ID))
	tmpUser.Identity = identity
	if err := models.UpdateUser(tmpUser); err != nil {
		c.JSON(http.StatusInternalServerError, resp.Err.WithMsg("Token保存错误"))
		return
	}
	c.JSON(http.StatusOK, resp.OK.WithData(tmpUser.Identity))
}

// DeleteUser
// @Tags 删除用户
// @Success 200
// @Router /user/deleteUser [get]
func (u UserApi) DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Err.WithMsg(err.Error()))
		return
	}
	user.ID = uint(id)
	if err := service.DeleteUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, resp.Err.WithMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.OK.WithMsg("删除成功"))
}

// UpdateUser
// @Tags 修改用户
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param email formData string false "email"
// @Success 200
// @Router /user/updateUser [post]
func (u UserApi) UpdateUser(c *gin.Context) {
	user := models.UserBasic{}

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, resp.ErrParam)
		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, resp.ErrParam.WithMsg("用户Id不能为空"))
		return
	}
	if err := service.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, resp.Err.WithMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.OK.WithMsg("修改成功"))
}

func (u UserApi) SendMsg(ctx *gin.Context) {
	ws, err := config.UpGrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 关闭websocket连接
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)

	MsgHandler(ws, ctx)
}

func MsgHandler(ws *websocket.Conn, ctx *gin.Context) {
	for {
		msg, err := utils.Subscribe(ctx, "Octopus")
		if err != nil {
			fmt.Println(err)
		}
		tm := time.Now().Format(config.DateTimeFormat)
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		if err := ws.WriteMessage(1, []byte(m)); err != nil {
			fmt.Println(err)
		}
	}
}
