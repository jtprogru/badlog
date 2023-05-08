package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("/var/log/bad.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		r := r.Intn(2147483647)
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		line := fmt.Sprintf("%s token: %d\n", timestamp, r)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
		_ = file.Sync()
		time.Sleep(1 * time.Second)
	}
}
