package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func createFile(fileName string) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

func writeFile(fileName string, data []byte) error {
	err := ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func readFile(fileName string) ([]byte, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func addItemToFile(fileName string, newItem Item) error {
	data, err := readFile(fileName)
	if err != nil {
		return err
	}

	var items []Item
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}

	items = append(items, newItem)

	newData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}

	if err := writeFile(fileName, newData); err != nil {
		return err
	}

	return nil
}

func main() {
	fileName := "items.json"

	if err := createFile(fileName); err != nil {
		fmt.Println("Error creating file:", err)
		return
	}


	newItem := Item{ID: 1, Name: "Item 1"}
	if err := addItemToFile(fileName, newItem); err != nil {
		fmt.Println("Error adding item to file:", err)
		return
	}

	data, err := readFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File contents:")
	fmt.Println(string(data))
}
