package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type mapData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.Next != "" {
		url = c.Next
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return err
	}
	if err != nil {
		return err
	}

	var data mapData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	c.Next = data.Next
	c.Previous = data.Previous

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(c *config) error {
	if c.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	res, err := http.Get(c.Previous)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return err
	}
	if err != nil {
		return err
	}
	var data mapData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	c.Next = data.Next
	c.Previous = data.Previous

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
	return nil
}
