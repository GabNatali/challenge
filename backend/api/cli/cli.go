package cli

import (
	"github.com/alecthomas/kong"
	"github.com/codeableorg/weekend-challenge-13-GabNatali/config"
)

// Scheme

type scheme struct {
	EnvPath string `help:"Path to env config file" type:"path" optional:""`
}

// Parser

type Parser struct {
	scheme scheme
}

func NewParser() *Parser {
	return &Parser{scheme: scheme{}}
}

func (p *Parser) ParseConfig() (*config.Config, error) {
	kong.Parse(&p.scheme)
	return config.ParseEnv(p.scheme.EnvPath)
}
