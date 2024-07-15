package api

import (
	"encoding/json"
	"log"
	"sync"
)

// DocumentProcessor хранит кэшированные документы и обеспечивает потокобезопасность
type DocumentProcessor struct {
	mu     sync.Mutex
	cached map[string]map[string]Document
}

// NewDocumentProcessor создает новый экземпляр DocumentProcessor с инициализированным кэшем
func NewDocumentProcessor() *DocumentProcessor {
	return &DocumentProcessor{
		cached: make(map[string]map[string]Document),
	}
}

// ФУнкция ProcessDocuments обрабатывает документы из входной очереди и отправляет их в выходную очередь
func ProcessDocuments(inputQueue <-chan Document, outputQueue chan<- Document, done chan<- bool) {
	processor := NewDocumentProcessor()
	for doc := range inputQueue { // обработка каждого документа из входной очереди
		processedDoc, err := processor.Process(&doc) // обработка документа
		if err != nil {                              // если возникла ошибка, переходим к следующему документу
			continue
		}
		if processedDoc != nil {
			outputQueue <- *processedDoc // отправляем обработанный документ в выходную очередь
		}
	}
	close(outputQueue) // закрытие выходную очередь после завершения обработки
	done <- true       // сигнал о завершении обработки
}

// WriterService принимает документы из выходной очереди и записывает их в лог в формате JSON
func WriterService(outputQueue <-chan Document, done chan<- bool) {
	for doc := range outputQueue { // проходимся по документам из выходной очереди
		output, err := json.Marshal(doc)
		if err != nil {
			log.Printf("Ошибки в формировании json: %v", err) // запись ошибка в журнал
			continue
		}
		log.Printf("Документ: %s", output) // запись документа в лог
	}
	done <- true // сигнал о завершении
}
