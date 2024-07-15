package api

import (
	"errors"
	"strings"
)

type Document struct {
	Url            string // URL документа
	PubDate        uint64 // Дата публикации
	FetchTime      uint64 // Время извлечения
	Text           string // Текст документа
	FirstFetchTime uint64 // Время первого извлечения
}

type Processor interface {
	Process(d *Document) (*Document, error) // метод для обработки документа
}

// var (
// 	unprocessed_queue []Document
// 	cached            map[string]map[string]Document = make(map[string]map[string]Document)
// 	processed_queue   []Document
// 	min_fetchtime_doc Document
// 	max_fetchtime_doc Document
// )

// Process обрабатывает документ, обновляя его в зависимости от времени извлечения и других данных
func (dp *DocumentProcessor) Process(d *Document) (*Document, error) {
	// проверяем, что URL документа не пустой
	if d.Url == "" {
		return nil, errors.New("у документа отсутствует url")
	}

	url := strings.ToLower(d.Url) // url полученного документа

	dp.mu.Lock()         // блокируем мьютекс для записи
	defer dp.mu.Unlock() // разблокируем мьютекс по завершении работы функции

	docs, ok := dp.cached[url] // Проверка был ли документ с таким url раньше
	if ok {                    // если документ найден в кэше
		// Обновляем документы с минимальным и максимальным FetchTime
		if d.FetchTime <= docs["min_fetchtime_doc"].FetchTime {
			docs["min_fetchtime_doc"] = *d
		} else {
			d.PubDate = docs["min_fetchtime_doc"].PubDate
		}

		if d.FetchTime >= docs["max_fetchtime_doc"].FetchTime {
			docs["max_fetchtime_doc"] = *d
		} else {
			d.FetchTime = docs["max_fetchtime_doc"].FetchTime
			d.Text = docs["max_fetchtime_doc"].Text
		}
	} else {
		// если документа с таким URL нет в кэше, создаем новую запись
		dp.cached[url] = map[string]Document{
			"min_fetchtime_doc": *d,
			"max_fetchtime_doc": *d,
		}
		docs = dp.cached[url]
	}

	// установка FirstFetchTime документа
	d.FirstFetchTime = docs["min_fetchtime_doc"].FetchTime

	return d, nil
}
