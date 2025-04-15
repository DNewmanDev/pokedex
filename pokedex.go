package main

type Pokedex struct {
	pokedex map[string]PokemonInfo
}

func (p *Pokedex) Add(name string, info PokemonInfo) {
	if _, exists := p.pokedex[name]; !exists {
		p.pokedex[name] = info
	}
}

func (p *Pokedex) Get(name string) (PokemonInfo, bool) {

	val, ok := p.pokedex[name]
	if !ok {
		return PokemonInfo{}, ok
	}
	return val, ok
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		pokedex: make(map[string]PokemonInfo),
	}
}
