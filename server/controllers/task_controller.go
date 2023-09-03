package controllers

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/firhan200/taskmanagement/data"
	"github.com/firhan200/taskmanagement/dto"
	"github.com/firhan200/taskmanagement/repositories"
	"github.com/firhan200/taskmanagement/services"
	"github.com/firhan200/taskmanagement/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TaskHandler struct {
	db *gorm.DB
}

func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{
		db: db,
	}
}

func (th *TaskHandler) GetTasks() fiber.Handler {
	return func(c *fiber.Ctx) error {
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

		repo := repositories.NewTaskRepository(th.db)
		service := services.NewTaskService(repo)

		tasks, err := service.GetTasksByUserId(
			uid,
			cursor,
			limit,
			orderBy,
			sort,
			search,
		)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		//get body parser
		return c.Status(http.StatusOK).JSON(tasks)
	}
}

func (th *TaskHandler) GetTaskById() fiber.Handler {
	return func(c *fiber.Ctx) error {
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

		repo := repositories.NewTaskRepository(th.db)
		service := services.NewTaskService(repo)
		task, err := service.GetByIdAuthorize(uid, id)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		//get body parser
		return c.Status(http.StatusOK).JSON(task)
	}
}

func (th *TaskHandler) CreateTask() fiber.Handler {
	return func(c *fiber.Ctx) error {
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

		repo := repositories.NewTaskRepository(th.db)
		service := services.NewTaskService(repo)
		createdId, saveErr := service.Create(
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
}

func (th *TaskHandler) UpdateTask() fiber.Handler {
	return func(c *fiber.Ctx) error {
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

		repo := repositories.NewTaskRepository(th.db)
		service := services.NewTaskService(repo)
		updatedTask, saveErr := service.Update(
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
}

func (th *TaskHandler) DeleteTask() fiber.Handler {
	return func(c *fiber.Ctx) error {
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

		repo := repositories.NewTaskRepository(th.db)
		service := services.NewTaskService(repo)
		deleteErr := service.Delete(uid, id)

		if deleteErr != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": deleteErr.Error(),
			})
		}

		return c.SendStatus(http.StatusOK)
	}
}

func (th *TaskHandler) GenerateRandomData() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//get user id
		uid, err := utils.ExtractTokenID(c)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
			return nil
		}

		repo := repositories.NewTaskRepository(th.db)
		service := services.NewTaskService(repo)

		tasks := []*data.Task{}
		taskchan := make(chan *data.Task)

		go func() {
			wg := sync.WaitGroup{}

			for i := 0; i < 100; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()

					task, err := service.Create(
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
}
