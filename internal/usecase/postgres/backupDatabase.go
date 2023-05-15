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

	cmd := exec.Command("pg_dump",
		"-h", cfg.Host,
		"-p", strconv.Itoa(cfg.Port),
		"-U", cfg.Username,
		"-d", dbname)

	fmt.Println("Введите путь, по которому сохранить бэкап (.dump):")
	reader := bufio.NewReader(os.Stdin)
	var filepath string
	fmt.Fscan(reader, &filepath)
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create backup file: %v", err)
	}
	defer file.Close()

	cmd.Stdout = file

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err, "stdinPipe")
		return err
	}
	defer stdin.Close()

	// запускаем команду
	err = cmd.Start()
	if err != nil {
		fmt.Println(err, "start cmd")
		return err
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err, "wait cmd")
		return err
	}
	fmt.Printf("Бэкап %s был создан по пути: %s \n", dbname, filepath)
	return nil
}
