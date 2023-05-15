package postgres

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"udv-test-task-go/internal/usecase"
)

type pgDeleteDB struct {
	Conn *sql.DB
}

func NewDeleteDB(Conn *sql.DB) usecase.DeleteDBInterface {
	return &pgDeleteDB{
		Conn: Conn,
	}
}
func (p *pgDeleteDB) DeleteDatabase() error {
	fmt.Print("Введите кол-во БД для удаления: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	countStr := scanner.Text()
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return fmt.Errorf("неверное количество баз данных: %v", err)
	}

	for i := 0; i < count; i++ {
		fmt.Printf("Введите имя БД %d: ", i+1)
		scanner.Scan()
		dbname := scanner.Text()

		_, err := p.Conn.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbname))
		if err != nil {
			return fmt.Errorf("не удалось удалить базу данных %s: %v", dbname, err)
		}
		fmt.Printf("База данных %s удалена.\n", dbname)
	}
	return nil
}
