package api

import (
	"github.com/gin-gonic/gin"
	resp "go-chat/middleware"
	"go-chat/models"
	"go-chat/service"
	"go-chat/utils"
	"net/http"
	"strconv"
)

// GetUserList
// @Tags 获取用户列表
// @Produce json
// @Success 200
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	userList := service.GetUserList()
	c.JSON(http.StatusOK, resp.OK.WithData(userList))
}

// GetUser
// @Tags 获取用户
// @Produce json
// @Success 200
// @Router /user/getUser [get]
func GetUser(c *gin.Context) {
	username := c.Query("username")
	user := service.GetUserByUsername(username)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, resp.Err.WithMsg("用户不存在"))
		return
	}
	c.JSON(http.StatusOK, resp.OK.WithData(&user))
}

// CreateUser
// @Tags 新增用户
// @Success 200
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, resp.ErrParam)
		return
	}

	if err := service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, resp.Err.WithMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.OK.WithData(user.Username))
}

// Login
// @Tags 登录
// @Success 200
// @Router /user/login [post]
func Login(c *gin.Context) {
	user := models.UserBasic{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, resp.ErrParam)
		return
	}
	tmpUser := service.GetUserByUsername(user.Username)
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
	tmpUser.Password = "******"
	c.JSON(http.StatusOK, resp.OK.WithData(tmpUser))
}

// DeleteUser
// @Tags 删除用户
// @Success 200
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
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
func UpdateUser(c *gin.Context) {
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

func GetFriends(ctx *gin.Context) {
	username := ctx.Query("username")
	user := models.TakeUserByUsername(username)
	friends := service.GetFriendByUsername(user.ID)
	ctx.JSON(http.StatusOK, resp.OK.WithData(friends))
}
func GetGroups(ctx *gin.Context) {
	username := ctx.Query("username")
	user := models.TakeUserByUsername(username)
	groups := service.GetGroupByUsername(user.ID)
	ctx.JSON(http.StatusOK, resp.OK.WithData(groups))
}
func GetListMessage(ctx *gin.Context) {
	username := ctx.Query("username")
	user := models.TakeUserByUsername(username)
	lms := service.GetListMessageByUsername(user.ID)
	ctx.JSON(http.StatusOK, resp.OK.WithData(lms))
}
