package commands


type Config struct {
	Commands map[string]CliCommand
	BaseUrl string
	NextUrl *string
	PreviousUrl *string
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"map": {
			Name:        "map",
			Description: "Display the next 20 location areas",
			Callback:    commandMap,
		},
		"mapp": {
			Name:        "mapp",
			Description: "Display the previous 20 location areas",
			Callback:    commandMapP,
		},
		"help": {
			Name:        "help",
			Description: "Display the Pokedex help text",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}
