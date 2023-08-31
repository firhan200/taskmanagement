package controllers

import (
	"errors"
	"fmt"
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

	// c.Status(http.StatusOK).JSON(&data.Tasks{})
	// return nil

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

func GetTaskById(c *fiber.Ctx) error {
	//get user id
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	idParams := c.Params("id")
	//parse
	idInt, err := strconv.Atoi(idParams)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	id := uint(idInt)

	task, err := data.GetTask(id, uid)

	//get body parser
	c.Status(http.StatusOK).JSON(task)
	return nil
}

func CreateTask(c *fiber.Ctx) error {
	//get user id
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}

	//get body parser
	createTaskDto := dto.CreateTaskDto{}
	if err := c.BodyParser(&createTaskDto); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}

	//validate
	if createTaskDto.Title == "" || createTaskDto.Description == "" {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Title, Description or Due Date cannot be empty",
		})
		return nil
	}

	//init new task instance
	task := data.Task{
		Title:       createTaskDto.Title,
		Description: createTaskDto.Description,
		DueDate:     createTaskDto.DueDate.UTC(),
		UserId:      uid,
	}

	//save and check if error
	taskId, err := task.Save()
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(fiber.Map{
		"task_id": taskId,
	})
	return nil
}

func UpdateTask(c *fiber.Ctx) error {
	idParams := c.Params("id")
	if idParams == "" {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Id not found",
		})
		return errors.New("id not found")
	}

	//parse
	idInt, err := strconv.Atoi(idParams)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
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
	fmt.Print(uid)

	//get body parser
	updateTaskDto := dto.UpdateTaskDto{}
	if err := c.BodyParser(&updateTaskDto); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	//validate
	if updateTaskDto.Title == "" || updateTaskDto.Description == "" {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Title, Description or Due Date cannot be empty",
		})
		return errors.New("Title, Description or Due Date cannot be empty")
	}

	task, err := data.GetTask(id, uid)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	task.Title = updateTaskDto.Title
	task.Description = updateTaskDto.Description
	task.DueDate = updateTaskDto.DueDate.UTC()
	task.Update()

	c.Status(http.StatusOK).JSON(fiber.Map{
		"task": task,
	})
	return nil
}

func DeleteTask(c *fiber.Ctx) error {
	idParams := c.Params("id")
	if idParams == "" {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Id not found",
		})
		return errors.New("Id not found")
	}

	//parse
	idInt, err := strconv.Atoi(idParams)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
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
	fmt.Print(uid)

	task, err := data.GetTask(id, uid)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	task.Delete()

	c.Status(http.StatusOK).JSON(fiber.Map{
		"task": task,
	})
	return nil
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

	tasks := []*data.Task{}
	taskchan := make(chan *data.Task)

	go func() {
		wg := sync.WaitGroup{}

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				task := &data.Task{
					UserId:      uid,
					Title:       gofakeit.Sentence(gofakeit.IntRange(5, 10)),
					Description: gofakeit.Sentence(gofakeit.IntRange(10, 50)),
					DueDate:     gofakeit.DateRange(time.Now().UTC().AddDate(0, 0, -5), time.Now().UTC().AddDate(0, 0, 7)),
				}

				_, err := task.Save()
				if err != nil {
					return
				}

				taskchan <- task
			}()
		}

		fmt.Println("waiting")
		wg.Wait()
		close(taskchan)

		fmt.Println("finish wait")
	}()

	for tc := range taskchan {
		tasks = append(tasks, tc)
	}

	c.Status(http.StatusOK).JSON(&tasks)

	return nil
}
