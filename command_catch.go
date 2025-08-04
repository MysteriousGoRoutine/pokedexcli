package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a pokemon name")
	}
	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	// TODO: Implement catch logic
	pokemon, err := config.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Рассчитываем вероятность поимки на основе BaseExperience
	// Покемоны с большим BaseExperience сложнее поймать
	// Используем формулу: чем больше BaseExperience, тем меньше шанс поимки
	// Максимальный шанс поимки 90%, минимальный 10%
	catchChance := 90.0
	if pokemon.BaseExperience > 0 {
		// Нормализуем BaseExperience (обычно от 30 до 600+)
		difficulty := float64(pokemon.BaseExperience) / 10.0
		catchChance = 90.0 - (difficulty * 0.8)
		if catchChance < 10.0 {
			catchChance = 10.0
		}
	}

	// Генерируем случайное число от 0 до 100
	roll := rand.Float64() * 100.0

	if roll <= catchChance {
		fmt.Printf("Caught %s! (Base Experience: %d, Catch Chance: %.1f%%)\n",
			pokemon.Name, pokemon.BaseExperience, catchChance)
		return nil
	} else {
		fmt.Printf("%s broke free! Better luck next time... (Base Experience: %d, Catch Chance: %.1f%%)\n",
			pokemon.Name, pokemon.BaseExperience, catchChance)
		return nil
	}
}
