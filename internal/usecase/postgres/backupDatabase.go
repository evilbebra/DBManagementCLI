package postgres

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"udv-test-task-go/config"
	"udv-test-task-go/internal/usecase"
)

func NewBackup() usecase.BackupInterface {
	return &pgConnBackup{}
}

type pgConnBackup struct{}

func (p *pgConnBackup) BackupDatabase(cfg config.Config) error {
	fmt.Print("Введите имя БД, которую бэкапим: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dbname := scanner.Text()

	fmt.Println("Введите путь, по которому сохранить бэкап (.dump):")
	reader := bufio.NewReader(os.Stdin)
	var filepath string
	fmt.Fscan(reader, &filepath)

	cmd := exec.Command("pg_dump",
		"-h", cfg.Host,
		"-p", strconv.Itoa(cfg.Port),
		"-U", cfg.Username,
		"-d", dbname,
		"-f", filepath)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("run cmd %v", err)
	}

	fmt.Printf("Бэкап %s был создан по пути: %s \n", dbname, filepath)
	return nil
}
