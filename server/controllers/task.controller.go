package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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

	//get params
	cursor := c.DefaultQuery("cursor", "0")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	orderBy := c.DefaultQuery("orderBy", "created_at")
	sort := c.DefaultQuery("sort", "desc")
	search := c.DefaultQuery("search", "")

	tasks := data.Tasks{
		Cursor:  cursor,
		Limit:   limit,
		OrderBy: orderBy,
		Sort:    sort,
		Search:  search,
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
	if createTaskDto.Title == "" || createTaskDto.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title, Description or Due Date cannot be empty",
		})
		return
	}

	//init new task instance
	task := data.Task{
		Title:       createTaskDto.Title,
		Description: createTaskDto.Description,
		DueDate:     createTaskDto.DueDate,
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

func UpdateTask(c *gin.Context) {
	idParams := c.Param("id")
	if idParams == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Id not found",
		})
		return
	}

	//parse
	idInt, err := strconv.Atoi(idParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := uint(idInt)

	//get user id
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Print(uid)

	//get body parser
	updateTaskDto := dto.UpdateTaskDto{}
	if err := c.ShouldBindJSON(&updateTaskDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//validate
	if updateTaskDto.Title == "" || updateTaskDto.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title, Description or Due Date cannot be empty",
		})
		return
	}

	task, err := data.GetTask(id, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	task.Title = updateTaskDto.Title
	task.Description = updateTaskDto.Description
	task.DueDate = updateTaskDto.DueDate
	task.Update()

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func DeleteTask(c *gin.Context) {
	idParams := c.Param("id")
	if idParams == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Id not found",
		})
		return
	}

	//parse
	idInt, err := strconv.Atoi(idParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := uint(idInt)

	//get user id
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Print(uid)

	task, err := data.GetTask(id, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	task.Delete()

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}
