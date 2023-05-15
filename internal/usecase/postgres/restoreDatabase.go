package postgres

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"udv-test-task-go/config"
	"udv-test-task-go/internal/usecase"
)

type pgRestoreDB struct {
	Conn *sql.DB
}

func NewRestoreDB(Conn *sql.DB) usecase.RestoreInterface {
	return &pgRestoreDB{
		Conn: Conn,
	}
}

func (p *pgRestoreDB) RestoreDatabase(cfg config.Config) error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите имя базы данных для восстановления: ")
	scanner.Scan()
	dbname := scanner.Text()
	rows, err := p.Conn.Query(fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname='%s'", dbname))
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		_, err = p.Conn.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Printf("База данных %s создана.\n", dbname)
	}

	fmt.Print("Введите путь до бекапа (.dump): ")
	scanner.Scan()
	filepath := scanner.Text()

	// Восстанавливаем базу данных из бэкапа
	cmd := exec.Command("psql",
		"-h", cfg.Host,
		"-p", strconv.Itoa(cfg.Port),
		"-U", cfg.Username,
		"-d", dbname,
		"-f", filepath)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("=== База данных %s восстановлена из бэкапа. ===\n", dbname)
	return nil
}
