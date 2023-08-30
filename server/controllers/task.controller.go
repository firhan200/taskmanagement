package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/firhan200/taskmanagement/data"
	"github.com/firhan200/taskmanagement/dto"
	"github.com/firhan200/taskmanagement/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	c.Status(http.StatusOK).JSON(&data.Tasks{})
	return nil

	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	//get params
	cursor := c.Query("cursor", "0")
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	orderBy := c.Query("orderBy", "created_at")
	sort := c.Query("sort", "desc")
	search := c.Query("search", "")

	c.Status(http.StatusOK).JSON(&data.Tasks{})
	return nil

	tasks := data.Tasks{
		Cursor:  cursor,
		Limit:   limit,
		OrderBy: orderBy,
		Sort:    sort,
		Search:  search,
	}
	tasks.GetByUserId(uid)

	//get body parser
	c.Status(http.StatusOK).JSON(tasks)

	return nil
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	//get body parser
	c.JSON(http.StatusOK, fiber.Map{
		"id": id,
	})
}

func CreateTask(c *fiber.Ctx) {
	//get user id
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return
	}

	//get body parser
	createTaskDto := dto.CreateTaskDto{}
	if err := c.JSON(&createTaskDto); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return
	}

	//validate
	if createTaskDto.Title == "" || createTaskDto.Description == "" {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
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
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(fiber.Map{
		"task_id": taskId,
	})
}

func UpdateTask(c *fiber.Ctx) {
	idParams := c.Params("id")
	if idParams == "" {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Id not found",
		})
		return
	}

	//parse
	idInt, err := strconv.Atoi(idParams)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return
	}
	id := uint(idInt)

	//get user id
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return
	}
	fmt.Print(uid)

	//get body parser
	updateTaskDto := dto.UpdateTaskDto{}
	if err := c.JSON(&updateTaskDto); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return
	}

	//validate
	if updateTaskDto.Title == "" || updateTaskDto.Description == "" {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Title, Description or Due Date cannot be empty",
		})
		return
	}

	task, err := data.GetTask(id, uid)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
		return
	}

	task.Title = updateTaskDto.Title
	task.Description = updateTaskDto.Description
	task.DueDate = updateTaskDto.DueDate
	task.Update()

	c.Status(http.StatusOK).JSON(fiber.Map{
		"task": task,
	})
}

func DeleteTask(c *fiber.Ctx) {
	idParams := c.Params("id")
	if idParams == "" {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Id not found",
		})
		return
	}

	//parse
	idInt, err := strconv.Atoi(idParams)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return
	}
	id := uint(idInt)

	//get user id
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return
	}
	fmt.Print(uid)

	task, err := data.GetTask(id, uid)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
		return
	}

	task.Delete()

	c.Status(http.StatusOK).JSON(fiber.Map{
		"task": task,
	})
}
