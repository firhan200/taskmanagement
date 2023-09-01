package controllers

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/firhan200/taskmanagement/data"
	"github.com/firhan200/taskmanagement/dto"
	"github.com/firhan200/taskmanagement/utils"
	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	//get params
	cursor := c.Query("cursor", "")
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	orderBy := c.Query("orderBy", "created_at")
	sort := c.Query("sort", "desc")
	search := c.Query("search", "")

	db := data.NewConnection()
	taskManager := data.NewTaskManager(db)

	tasks := taskManager.GetTasks(
		uid,
		cursor,
		limit,
		orderBy,
		sort,
		search,
	)

	//get body parser
	c.Status(http.StatusOK).JSON(tasks)

	return nil
}

func GetTaskById(c *fiber.Ctx) error {
	//get user id
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	idParams := c.Params("id")
	//parse
	idInt, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	id := uint(idInt)

	db := data.NewConnection()
	taskManager := data.NewTaskManager(db)
	task, err := taskManager.GetSingleTask(uid, id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//get body parser
	return c.Status(http.StatusOK).JSON(task)
}

func CreateTask(c *fiber.Ctx) error {
	//get user id
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//get body parser
	createTaskDto := dto.CreateTaskDto{}
	if err := c.BodyParser(&createTaskDto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//validate
	if createTaskDto.Title == "" || createTaskDto.Description == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Title, Description or Due Date cannot be empty",
		})
	}

	db := data.NewConnection()
	taskManager := data.NewTaskManager(db)
	createdId, saveErr := taskManager.Save(
		uid,
		createTaskDto.Title,
		createTaskDto.Description,
		createTaskDto.DueDate,
	)

	if saveErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": saveErr.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"task_id": createdId,
	})
}

func UpdateTask(c *fiber.Ctx) error {
	idParams := c.Params("id")
	if idParams == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Id not found",
		})
	}

	//parse
	idInt, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	id := uint(idInt)

	//get user id
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//get body parser
	updateTaskDto := dto.UpdateTaskDto{}
	if err := c.BodyParser(&updateTaskDto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//validate
	if updateTaskDto.Title == "" || updateTaskDto.Description == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Title, Description or Due Date cannot be empty",
		})
	}

	db := data.NewConnection()
	taskManager := data.NewTaskManager(db)
	updatedTask, saveErr := taskManager.Update(
		uid,
		id,
		updateTaskDto.Title,
		updateTaskDto.Description,
		updateTaskDto.DueDate,
	)

	if saveErr != nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"task": updatedTask,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"task": updatedTask,
	})
}

func DeleteTask(c *fiber.Ctx) error {
	idParams := c.Params("id")
	if idParams == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Id not found",
		})
	}

	//parse
	idInt, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	id := uint(idInt)

	//get user id
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	db := data.NewConnection()
	taskManager := data.NewTaskManager(db)
	deleteErr := taskManager.Remove(uid, id)

	if deleteErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": deleteErr,
		})
	}

	return c.SendStatus(http.StatusOK)
}

func GenerateRandomData(c *fiber.Ctx) error {
	//get user id
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}

	db := data.NewConnection()
	tm := data.NewTaskManager(db)

	tasks := []*data.Task{}
	taskchan := make(chan *data.Task)

	go func() {
		wg := sync.WaitGroup{}

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				task, err := tm.Save(
					uid,
					gofakeit.Sentence(gofakeit.IntRange(5, 10)),
					gofakeit.Sentence(gofakeit.IntRange(10, 50)),
					gofakeit.DateRange(time.Now().UTC().AddDate(0, 0, -5), time.Now().UTC().AddDate(0, 0, 14)),
				)
				if err != nil {
					return
				}

				taskchan <- task
			}()
		}
		wg.Wait()
		close(taskchan)
	}()

	for tc := range taskchan {
		tasks = append(tasks, tc)
	}

	c.Status(http.StatusOK).JSON(&tasks)

	return nil
}
