package store

import (
	"fmt"
	"strings"
	"testing"
)

// TestStore инициализирует Store для тестирования с заданным URL базы данных
// функция возвращает экземпляр Store и функцию для очистки таблиц и закрытия соединения
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper() // указывает, что это вспомогательная функция теста

	// создаем новый конфиг и устанавливаем URL базы данных
	config := NewConfig()
	config.DataBaseURL = databaseURL
	s := New(config)

	// открываем соединение с базой данных
	if err := s.Open(); err != nil {
		t.Fatal(err) // если не удается открыть соединение, завершаем тест с ошибкой
	}

	// возвращаем Store и функцию для очистки таблиц и закрытия соединения
	return s, func(tables ...string) {
		// если указаны таблицы, очищаем их
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err) // если не удается очистить таблицы, завершаем тест с ошибкой
			}
		}
		// закрываем соединение с базой данных
		s.Close()
	}
}
