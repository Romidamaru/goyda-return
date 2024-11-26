package ent

import "sync"

// Task представляет задачу
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var (
	Tasks   = []Task{} // Хранилище задач
	IdCount = 1        // Счётчик для генерации ID
	Mu      sync.Mutex // Мьютекс для защиты данных
)
