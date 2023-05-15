package database

import (
	_ "github.com/lib/pq"
	"testing"
	"udv-test-task-go/config"
)

func TestConnectSQL(t *testing.T) {
	// Подготавливаем тестовые данные
	cfg := config.Config{
		Host:     "localhost",
		Port:     5433,
		Username: "postgres",
		Password: "root",
		Driver:   "postgres",
	}

	// Запускаем функцию ConnectSQL
	dbConn, err := ConnectSQL(cfg)
	if err != nil {
		t.Errorf("Ошибка при подключении к базе данных: %v", err)
	}

	// Проверяем, что соединение с базой данных не является nil
	if dbConn.Conn == nil {
		t.Errorf("Соединение с базой данных не было установлено")
	}

	// Закрываем соединение с базой данных
	err = dbConn.Conn.Close()
	if err != nil {
		t.Errorf("Ошибка при закрытии соединения с базой данных: %v", err)
	}
}
