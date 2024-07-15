package api

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestProcessDocuments это тестовая функция для провер	ки обработки документов
func TestProcessDocuments(t *testing.T) {
	// создаем каналы для входных и выходных данных и сигнализации о завершении
	inputQueue := make(chan Document, 100)
	outputQueue := make(chan Document, 100)
	done := make(chan bool)

	// запускаем горутину для обработки документов
	go ProcessDocuments(inputQueue, outputQueue, done)

	// создаем набор тестовых документов
	documents := []Document{
		{Url: "http://example.com", PubDate: 20240601, FetchTime: 20240714, Text: "Первый документ"},
		{Url: "http://example.com", PubDate: 20240602, FetchTime: 20240715, Text: "Обновленный документ 1"},
		{Url: "http://example.com", PubDate: 20240603, FetchTime: 20240713, Text: "Обновленный документ 2"},
	}

	// отправляем документы в входную очередь
	for _, doc := range documents {
		inputQueue <- doc
	}
	// закрываем входную очередь
	// сигнализируем, что данные отправлены
	close(inputQueue)

	// ждем завершения обработки документов
	<-done

	// ожидаемые результаты обработки
	expectedOutputs := []Document{
		{Url: "http://example.com", PubDate: 20240601, FetchTime: 20240714, Text: "Первый документ", FirstFetchTime: 20240714},
		{Url: "http://example.com", PubDate: 20240601, FetchTime: 20240715, Text: "Обновленный документ 1", FirstFetchTime: 20240714},
		{Url: "http://example.com", PubDate: 20240603, FetchTime: 20240715, Text: "Обновленный документ 1", FirstFetchTime: 20240713},
	}

	// Проверка соответствия полученных результатов с ожидаемым
	i := 0
	for got := range outputQueue {
		assert.Equal(t, expectedOutputs[i], got)
		i++
	}
}

// TestWriterService это тестовая функция для проверки записи после обработки документов
func TestWriterService(t *testing.T) {
	outputQueue := make(chan Document, 100) // создаем канал для выходных документов
	done := make(chan bool)                 // канал для сигнализации о завершении записи

	var logOutput bytes.Buffer
	log.SetOutput(&logOutput) // перенаправляем вывод журнала в буфер

	go WriterService(outputQueue, done) // запускаем запись документов в журнал в отдельной горутине

	documents := []Document{
		{Url: "http://example.com", PubDate: 20240601, FetchTime: 20240714, Text: "Первый документ", FirstFetchTime: 20240714},
		{Url: "http://example.com", PubDate: 20240601, FetchTime: 20240715, Text: "Обновленный документ 1", FirstFetchTime: 20240714},
		{Url: "http://example.com", PubDate: 20240603, FetchTime: 20240715, Text: "Обновленный документ 1", FirstFetchTime: 20240713},
	}

	for _, doc := range documents {
		outputQueue <- doc // отправляем документы в выходную очередь
	}
	close(outputQueue) // закрываем выходную очередь после отправки всех документов

	<-done // ждем завершения записи

	for _, doc := range documents { // проверяем, что все документы записаны в журнал
		expectedOutput, _ := json.Marshal(doc)                         // преобразуем документ в формат JSON
		assert.Contains(t, logOutput.String(), string(expectedOutput)) // проверяем что JSON присутствует в журнале
	}
}
