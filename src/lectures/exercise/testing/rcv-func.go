//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Name string
type Health int
type Energy int

type Player struct {
	name              Name
	health, maxHealth uint
	energy, maxEnergy uint
}

func (p *Player) addHealth(amount uint) {
	p.health += amount
	if p.health > p.maxHealth {
		p.health = p.maxHealth
	}
}

func (p *Player) applyDamage(amount uint) {
	if p.health-amount > p.health {
		p.health = 0
	} else {
		p.health -= amount
	}
}

func (p *Player) addEnergy(amount uint) {
	p.energy += amount
	if p.energy > p.maxEnergy {
		p.energy = p.maxEnergy
	}
}

func (p *Player) consumeEnergy(amount uint) {
	if p.energy-amount > p.energy {
		p.energy = 0
	} else {
		p.energy -= amount
	}
}

func print(p Player) {
	fmt.Println(p.name, "Health:", p.health, "Energy:", p.energy)
}

func main() {
	player := Player{"Player 1", 100, 100, 100, 100}
	print(player)

	player.addHealth(50)
	print(player)

	player.applyDamage(30)
	print(player)

	player.addEnergy(20)
	print(player)

	player.consumeEnergy(100)
	print(player)
}
