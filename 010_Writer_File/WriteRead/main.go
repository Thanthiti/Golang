package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// readFile อ่านไฟล์ทีละบรรทัดแล้ว return slice ของ string
func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return lines, nil
}

// processLines ทำ logic กับข้อมูล เช่น แปลงเป็นพิมพ์ใหญ่
func processLines(lines []string) []string {
	var processed []string
	for _, line := range lines {
		processed = append(processed, strings.ToUpper(line))
	}
	return processed
}

// writeFile เขียน slice ของ string ลงไฟล์
func writeFile(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}

	return nil
}

func main() {
	// 1. อ่านไฟล์ input.txt
	lines, err := readFile("output.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// 2. ประมวลผลข้อมูล (ทำให้เป็นพิมพ์ใหญ่)
	processed := processLines(lines)

	// 3. เขียนไฟล์ output.txt
	if err := writeFile("output.txt", processed); err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	log.Println("File processing completed successfully.")
}
