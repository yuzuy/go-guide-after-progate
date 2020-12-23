package server

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/yuzuy/go-guide-after-progate/db"
	"github.com/yuzuy/go-guide-after-progate/todo"
)

type Server struct {
	r  *gin.Engine
	db *db.DB
}

func New() *Server {
	return &Server{
		r:  gin.Default(),
		db: db.New(),
	}
}

func (s *Server) findTasks(c *gin.Context) {
	tasks := s.db.FindTasks()

	c.JSON(200, tasks)
}

func (s *Server) addTask(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		name = "untitled"
	}
	task := &todo.Task{
		ID:        uuid.New().String(), // uuidの生成
		Name:      name,
		CreatedAt: time.Now(), // 現在時刻取得
	}

	err := s.db.AddTask(task)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{})
}

func (s *Server) updateTask(c *gin.Context) {
	id := c.Param("id")
	name := c.PostForm("name")
	isDoneStr := c.PostForm("is_done")

	if name == "" {
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	isDone, err := strconv.ParseBool(isDoneStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "could not parse is_done to bool"})
		return
	}

	err = s.db.UpdateTask(id, name, isDone)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

func (s *Server) removeTask(c *gin.Context) {
	id := c.Param("id")
	err := s.db.RemoveTask(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

func (s *Server) Start() error {
	s.r.GET("/tasks", s.findTasks)
	s.r.POST("/tasks", s.addTask)
	s.r.PATCH("/tasks/:id", s.updateTask)
	s.r.DELETE("/tasks/:id", s.removeTask)

	return s.r.Run()
}
