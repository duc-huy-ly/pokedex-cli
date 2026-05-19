# Pokedex CLI

A command-line interface (CLI) application for interacting with the world of Pokémon. This application allows you to explore different locations, encounter and catch Pokémon, and view the details of the Pokémon you have caught.

## Features

- **Explore Locations**: Navigate through different areas in the Pokémon world using the `map` and `mapb` commands.
- **Discover Pokémon**: Use the `explore` command in a specific location to find wild Pokémon.
- **Catch Pokémon**: Attempt to catch wild Pokémon with the `catch` command.
- **Inspect Pokémon**: View the details of Pokémon you have caught, including their stats and types, using the `inspect` command.
- **View Pokedex**: See a list of all the Pokémon you have caught with the `pokedex` command.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.15 or later)

### Installation & Running the application

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/duc-huy-ly/pokedex-cli.git
    cd pokedex-cli
    ```

2.  **Run the application:**
    ```sh
    go run .
    ```

## Commands

The Pokedex CLI provides a set of commands to interact with the Pokémon world.

- `help`: Displays a help message with all available commands.
- `map`: Displays the next 20 locations in the Pokémon world.
- `mapb`: Displays the previous 20 locations.
- `explore <location_name>`: Shows the Pokémon that can be found in the specified location.
- `catch <pokemon_name>`: Attempts to catch the specified Pokémon.
- `inspect <pokemon_name>`: Displays the details of a caught Pokémon.
- `pokedex`: Lists all the Pokémon you have caught.
- `exit`: Exits the Pokedex CLI.
