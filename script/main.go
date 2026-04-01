package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run gen_problem.go \"Nome da Questão\"")
		return
	}

	// Captura todos os argumentos após o nome do script e junta com espaço
	input := strings.Join(os.Args[1:], " ")

	normalized := normalizeToSnakeCase(input)

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Erro ao capturar diretorio: %v\n", err)
		return
	}

	fileName := normalized + ".go"
	fullPath := filepath.Join(cwd, fileName)

	if _, err := os.Stat(fullPath); err == nil {
		fmt.Printf("O arquivo '%s' já existe\n", fullPath)
		return
	}

	err = os.WriteFile(fullPath, []byte("package algoritmosexercicios"), 0644)
	if err != nil {
		fmt.Printf("Erro ao criar o arquivo em %s: %v\n", fullPath, err)
		return
	}

	fmt.Printf("Arquivo '%s' criado com sucesso em: %s\n", fileName, fullPath)
}

func normalizeToSnakeCase(s string) string {
	// 1. Converter para minúsculo
	s = strings.ToLower(s)

	// 2. Substituir qualquer caractere que não seja letra ou número por espaço
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	s = reg.ReplaceAllString(s, " ")

	// 3. Trim de espaços nas extremidades
	s = strings.TrimSpace(s)

	// 4. Substituir espaços (simples ou múltiplos) por underscore
	regSpace := regexp.MustCompile(`\s+`)
	s = regSpace.ReplaceAllString(s, "_")

	return s
}
