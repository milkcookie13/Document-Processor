package store

import (
	"context"
	"database/sql"
	api "dcproccer/app"
)

// DocumentStore определяет интерфейс для работы с хранилищем документов
type DocumentStore interface {
	Open() error
	Close()
	GetDocument(ctx context.Context, url string) (api.Document, error)
	SaveDocument(ctx context.Context, doc api.Document) error
	UpdateDocument(ctx context.Context, doc api.Document) error
}

// Store представляет хранилище документов с подключением к базе данных
type Store struct {
	config *Config
	db     *sql.DB
}

// New создает новый экземпляр Store с заданной конфигурацией
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open открывает соединение с базой данных PostgreSQL
// если подключение не удается, то возвращает ошибку
func (s *Store) Open() error {
	// открытие соединения с базой данных
	db, err := sql.Open("postgres", s.config.DataBaseURL)
	if err != nil {
		return err
	}

	// проверка подключения к базе данных
	if err := db.Ping(); err != nil {
		return err
	}

	// сохранение соединения в структуре Store
	s.db = db

	return nil
}

// Close закрывает соединение с базой данных.
func (s *Store) Close() {
	// закрытие соединения с базой данных

	s.db.Close()
}
