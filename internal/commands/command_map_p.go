package commands

import (
	"fmt"
	"strings"
)

func commandMapP(cfg *Config) error {
	if cfg.PreviousUrl == nil {
		fmt.Printf("you're on the first page\n")
		return nil
	}
	s := *cfg.PreviousUrl
	if !strings.HasPrefix(s, "http") {
		s = cfg.BaseUrl + s
	}
	err := print_locations(s, cfg)
	if err != nil {
		return err
	}
	return nil
}
