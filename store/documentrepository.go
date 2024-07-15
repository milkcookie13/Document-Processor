package store

import (
	"context"
	api "dcproccer/app"
)

// GetDocument извлекает документ из базы данных по его URL
// если документ не найден, возвращает ошибку, указывающую на это
// в случае любой другой ошибки возвращает эту ошибку
func (store *Store) GetDocument(ctx context.Context, url string) (api.Document, error) {
	// SQL-запрос для извлечения документа по URL
	// SELECT url, pub_date, fetch_time, text, first_fetch_time FROM documents WHERE url = $1

	return api.Document{}, nil
}

// SaveDocument вставляет новый документ в базу данных
// если возникает ошибка во время вставки, то возвращает эту ошибку
func (store *Store) SaveDocument(ctx context.Context, doc api.Document) error {
	// SQL для вставки нового документа
	// INSERT INTO documents (url, pub_date, fetch_time, text, first_fetch_time) VALUES ($1, $2, $3, $4, $5)

	return nil
}

// UpdateDocument обновляет существующий документ в базе данных по его URL
// если возникает ошибка во время обновления, то возвращает эту ошибку
func (store *Store) UpdateDocument(ctx context.Context, doc api.Document) error {
	// SQL запрос для обновления существующего документа
	// UPDATE documents SET pub_date = $1, fetch_time = $2, text = $3, first_fetch_time = $4 WHERE url = $5

	return nil
}
