package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

// Sua struct de domínio
type Dominio struct {
	ID    int
	Nome  string
	Valor float64
}

// go get github.com/xuri/excelize/v2
func lerExcelParaStructs(caminho string) ([]Dominio, error) {
	file, err := excelize.OpenFile(caminho)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir arquivo Excel: %w", err)
	}
	defer file.Close()

	// Assumindo que os dados estão na primeira planilha
	sheet := file.GetSheetName(0)

	// Ler todas as linhas da planilha
	rows, err := file.GetRows(sheet)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter linhas do Excel: %w", err)
	}

	// Criar slice para armazenar structs
	var dominios []Dominio

	// Processar cada linha (pulando o cabeçalho)
	for i, row := range rows {
		if i == 0 {
			continue // Pula o cabeçalho
		}

		// Tratar os valores de cada célula
		id, _ := strconv.Atoi(row[0])
		valor, _ := strconv.ParseFloat(row[2], 64)

		dominios = append(dominios, Dominio{
			ID:    id,
			Nome:  row[1],
			Valor: valor,
		})
	}

	return dominios, nil
}

func lerCSVParaStructs(caminho string) ([]Dominio, error) {
	file, err := os.Open(caminho)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir arquivo: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Ler todas as linhas
	linhas, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao ler linhas do CSV: %w", err)
	}

	// Criar slice para armazenar structs
	var dominios []Dominio

	// Processar cada linha (pulando o cabeçalho)
	for i, linha := range linhas {
		if i == 0 {
			continue // Pula o cabeçalho
		}
		// Converter os valores e preencher a struct
		id, _ := strconv.Atoi(linha[0])
		valor, _ := strconv.ParseFloat(linha[2], 64)

		dominios = append(dominios, Dominio{
			ID:    id,
			Nome:  linha[1],
			Valor: valor,
		})
	}

	return dominios, nil
}

func calculatePercentComposition(input map[string][]float64) map[string][]float64 {
	// Descobrir o tamanho máximo dos slices nos valores do map
	var maxLen int
	for _, values := range input {
		if len(values) > maxLen {
			maxLen = len(values)
		}
	}

	// Inicializar o slice de totais
	totals := make([]float64, maxLen)

	// Calcular o total para cada índice
	for _, values := range input {
		for i, value := range values {
			totals[i] += value
		}
	}

	// Montar o map de saída com os percentuais
	output := make(map[string][]float64)
	for key, values := range input {
		percentages := make([]float64, len(values))
		for i, value := range values {
			if totals[i] != 0 {
				percentages[i] = (value / totals[i]) * 100
			} else {
				percentages[i] = 0
			}
		}
		output[key] = percentages
	}

	return output
}

func main() {
	// Exemplo de uso
	dominios, err := lerCSVParaStructs("exemplo.csv")
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Dados csv:", dominios[:3])

	dominios2, err := lerExcelParaStructs("exemplo.xlsx")
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Dados excel:", dominios2[:3])
}
