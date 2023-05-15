package main

import (
	"bufio"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"udv-test-task-go/config"
	"udv-test-task-go/internal/database"
	"udv-test-task-go/internal/usecase/postgres"
)

func main() {
	cfg := config.Config{
		Driver: "postgres",
	}
	// Чтение данных конфигурации от пользователя
	fmt.Println("Введите данные конфигурации:")
	fmt.Print("Username: ")
	fmt.Scan(&cfg.Username)
	fmt.Print("Password: ")
	fmt.Scan(&cfg.Password)
	fmt.Print("Host: ")
	fmt.Scan(&cfg.Host)
	fmt.Print("Port: ")
	fmt.Scan(&cfg.Port)

	db, err := database.ConnectSQL(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Выводим список команд и ждем ввода
	fmt.Println("Выберите команду:")
	fmt.Println("1. Вывести список баз данных")
	fmt.Println("2. Удалить базу данных")
	fmt.Println("3. Создать бэкап базы данных")
	fmt.Println("4. Восстановить базу данных из бэкапа")
	fmt.Println("5. Выход")
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "1":
			ListDB := postgres.NewGetListDB(db.Conn)
			names, err := ListDB.GetListDatabases()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("=== Список всех БД ===")
			for _, name := range names {
				fmt.Println(name)
			}
		case "2":
			DeleteDB := postgres.NewDeleteDB(db.Conn)
			err = DeleteDB.DeleteDatabase()
			if err != nil {
				log.Fatal(err)
			}
		case "3":
			Backup := postgres.NewBackup()
			err = Backup.BackupDatabase(cfg)
			if err != nil {
				log.Fatal(err)
			}
		case "4":
			RestoreDB := postgres.NewRestoreDB(db.Conn)
			err = RestoreDB.RestoreDatabase(cfg)
			if err != nil {
				log.Fatal(err)
			}
		case "5":
			os.Exit(0)
		}
	}
}
