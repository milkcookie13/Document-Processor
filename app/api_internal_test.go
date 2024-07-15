package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Функция для теста обработки документа
func TestProcess(t *testing.T) {
	// Создаем тест кейсы
	testCases := []struct {
		name           string   // Название теста
		input          Document // Входные данные
		expectedOutput Document // Выходные данные
		expectError    bool     // Должна ли произойти ошибка
	}{
		{
			name: "тест 1",
			input: Document{
				Url:       "http://example.com",
				PubDate:   20240601,
				FetchTime: 20240714,
				Text:      "Первый документ",
			},
			expectedOutput: Document{
				Url:            "http://example.com",
				PubDate:        20240601,
				FetchTime:      20240714,
				Text:           "Первый документ",
				FirstFetchTime: 20240714,
			},
			expectError: false,
		},
		{
			name: "тест 2",
			input: Document{
				Url:       "http://example.com",
				PubDate:   20240602,
				FetchTime: 20240715,
				Text:      "Обновленный документ 1",
			},
			expectedOutput: Document{
				Url:            "http://example.com",
				PubDate:        20240601,
				FetchTime:      20240715,
				Text:           "Обновленный документ 1",
				FirstFetchTime: 20240714,
			},
			expectError: false,
		},
		{
			name: "тест 3",
			input: Document{
				Url:       "http://example.com",
				PubDate:   20240603,
				FetchTime: 20240713,
				Text:      "Обновленный документ 2",
			},
			expectedOutput: Document{
				Url:            "http://example.com",
				PubDate:        20240603,
				FetchTime:      20240715,
				Text:           "Обновленный документ 1",
				FirstFetchTime: 20240713,
			},
			expectError: false,
		},
		{
			name: "тест 4",
			input: Document{
				Url:       "",
				PubDate:   20240604,
				FetchTime: 20240716,
				Text:      "Документ без URL",
			},
			expectedOutput: Document{},
			expectError:    true,
		},
		{
			name: "тест 5",
			input: Document{
				Url:       "http://example.com",
				PubDate:   20240605,
				FetchTime: 20240715,
				Text:      "Документ с одинаковым FetchTime",
			},
			expectedOutput: Document{
				Url:            "http://example.com",
				PubDate:        20240603,
				FetchTime:      20240715,
				Text:           "Документ с одинаковым FetchTime",
				FirstFetchTime: 20240713,
			},
			expectError: false,
		},
	}

	dp := NewDocumentProcessor()   // создаем новый обработчик документов
	for _, tc := range testCases { // перебираем созданные тест кейсы
		t.Run(tc.name, func(t *testing.T) { // запускаем проверку тест кейсов
			got, err := dp.Process(&tc.input) // запускаем документ на обработку
			if tc.expectError {               // если предусмотрена ошибка
				assert.Error(t, err) // проверка появилась ли ошибка
			} else {
				assert.NoError(t, err)                   // проверка что ошибка не появилась
				assert.Equal(t, tc.expectedOutput, *got) // Проверка соответствия результата ожидаемому
			}
		})
	}
}
