package controllers

import (
	"net/http"

	"github.com/firhan200/taskmanagement/data"
	"github.com/firhan200/taskmanagement/dto"
	"github.com/firhan200/taskmanagement/utils"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	tasks := data.Tasks{
		Cursor: 0,
	}
	tasks.GetByUserId(uid)

	//get body parser
	c.JSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	//get body parser
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func CreateTask(c *gin.Context) {
	//get user id
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//get body parser
	createTaskDto := dto.CreateTaskDto{}
	if err := c.ShouldBindJSON(&createTaskDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//validate
	if createTaskDto.Title == "" || createTaskDto.Description == "" || createTaskDto.DueDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title, Description or Due Date cannot be empty",
		})
		return
	}

	//parse due date to time format
	date, err := utils.ParseDateString(createTaskDto.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//init new task instance
	task := data.Task{
		Title:       createTaskDto.Title,
		Description: createTaskDto.Description,
		DueDate:     date,
		UserId:      uid,
	}

	//save and check if error
	taskId, err := task.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task_id": taskId,
	})
}
