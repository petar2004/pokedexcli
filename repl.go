package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

func startRepl(cfg *config, args ...string) {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()

		input := strings.Split(strings.ToLower(text), " ")

		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		args := input[1:]

		command, exists := cfg.commands[commandName]
		if !exists {
			fmt.Printf("Unknown command: %s\n", commandName)
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getMap(cfg *config, args ...string) (locationAreaResponse, error) {
	var body []byte

	cachedData, found := cfg.cache.Get(cfg.next)
	if found {
		body = cachedData
	} else {
		res, err := http.Get(cfg.next)
		if err != nil {
			return locationAreaResponse{}, err
		}
		defer res.Body.Close()

		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return locationAreaResponse{}, err
		}

		cfg.cache.Add(cfg.next, body)
	}
	var locations locationAreaResponse

	err := json.Unmarshal(body, &locations)
	if err != nil {
		return locationAreaResponse{}, err
	}

	return locations, nil
}

func commandMap(cfg *config, args ...string) error {
	locations, err := getMap(cfg)
	if err != nil {
		return err
	}

	cfg.next = locations.Next
	cfg.previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for commandName, command := range cfg.commands {
		fmt.Printf("%s: %s\n", commandName, command.description)
	}
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + args[0])
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var response locationAreaDetails

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	for _, encounter := range response.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	fmt.Printf("Throwing a Pokeball at " + args[0] + "...\n")

	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + args[0])
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var pokemonInfo pokemonDetails

	err = json.Unmarshal(body, &pokemonInfo)
	if err != nil {
		return err
	}

	if pokemonInfo.BaseExperience <= 0 {
		return fmt.Errorf("invalid base experience")
	}

	roll := rand.Intn(pokemonInfo.BaseExperience)

	const catchChance = 50

	if roll < catchChance {
		fmt.Printf(args[0] + " was caught!\n")
		cfg.pokedex[pokemonInfo.Name] = pokemonInfo
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf(args[0] + " escaped!\n")
	}

	return nil
}

func commandInspect(cfg *config, args ...string) error {
	inspectedPokemon, found := cfg.pokedex[args[0]]
	if !found {
		return fmt.Errorf("You have not caught: %s", args[0])
	}

	fmt.Printf("Name: %s \n", inspectedPokemon.Name)
	fmt.Printf("Height: %d \n", inspectedPokemon.Height)
	fmt.Printf("Weight: %d \n", inspectedPokemon.Weight)

	fmt.Printf("Stats:\n")
	for _, stat := range inspectedPokemon.Stats {
		fmt.Printf("  -%s: %d \n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, typ := range inspectedPokemon.Types {
		fmt.Printf("  - %s\n", typ.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
