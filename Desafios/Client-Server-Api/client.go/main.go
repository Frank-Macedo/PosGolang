package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50000)
	defer cancel()

	file, err := os.Create("arquivo.txt")

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)

	if err != nil {
		panic(err)
	}

	done := make(chan error, 1)
	go func() {
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			done <- err
			return
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			done <- fmt.Errorf("status não OK: %s", res.Status)
			return
		}

		body, err := io.ReadAll(res.Body)

		if err != nil {
			done <- fmt.Errorf("erro ao ler body: %w", err)
			return
		}

		fmt.Println("Body recebido com sucesso:")
		_, err = file.Write(body)
		if err != nil {
			done <- fmt.Errorf("erro ao escrever no arquivo: %w", err)
		}

		done <- nil
	}()
	select {
	case err := <-done:
		if err != nil {
			fmt.Printf("erro na requisição: %v\n", err)
			return
		}
		fmt.Println("Requisição concluída com sucesso!")

	case <-ctx.Done():
		fmt.Println("timeout: requisição cancelada")
		return
	}

	if err != nil {
		panic(err)
	}

}
