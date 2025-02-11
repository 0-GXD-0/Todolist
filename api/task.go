package api

import (
	"demo1/pkg/utils"
	"demo1/service"
	"strings"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService

	// 获取并处理 Authorization 头部
	token := c.GetHeader("Authorization")
	// 去除 Bearer 前缀
	token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer "))
	logging.Println("Processed token:", token)

	claim, _ := utils.ParseToken(token)
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService

	if err := c.ShouldBind(&showTask); err == nil {
		res := showTask.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func ListTask(c *gin.Context) {
	var ListTask service.ListTaskService

	// 获取并处理 Authorization 头部
	token := c.GetHeader("Authorization")
	// 去除 Bearer 前缀
	token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer "))
	logging.Println("Processed token:", token)

	claim, _ := utils.ParseToken(token)

	if err := c.ShouldBind(&ListTask); err == nil {
		res := ListTask.List(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func UpdateTask(c *gin.Context) {
	var UpdateTask service.UpdateTaskService

	if err := c.ShouldBind(&UpdateTask); err == nil {
		res := UpdateTask.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func SearchTask(c *gin.Context) {
	var SearchTask service.SearchTaskService

	// 获取并处理 Authorization 头部
	token := c.GetHeader("Authorization")
	// 去除 Bearer 前缀
	token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer "))
	logging.Println("Processed token:", token)

	claim, _ := utils.ParseToken(token)

	if err := c.ShouldBind(&SearchTask); err == nil {
		res := SearchTask.Search(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func DeleteTask(c *gin.Context) {
	var DeleteTask service.DeleteTaskService

	if err := c.ShouldBind(&DeleteTask); err == nil {
		res := DeleteTask.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}
