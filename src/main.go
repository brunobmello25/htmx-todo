package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Task struct {
	Id    int
	Title string
	Done  bool
}

var tasks []Task

func ToggleTask(c *fiber.Ctx) error {
	fmt.Println("body: ", string(c.Body()))

	taskId := c.Params("id")

	var task Task
	for _, t := range tasks {
		if fmt.Sprint(t.Id) == taskId {
			task = t
		}
	}

	task.Done = !task.Done

	return c.Render("index", fiber.Map{
		"Tasks": tasks,
	})
}

func main() {
	tasks = append(tasks, Task{Id: 1, Title: "Learn Go", Done: false})

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/tasks")
	})

	app.Get("/tasks", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Tasks": tasks,
		})
	})

	app.Post("/tasks/:id/toggle", ToggleTask)

	app.Listen(":3000")
}
