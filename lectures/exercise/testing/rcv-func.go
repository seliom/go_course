package main

import "fmt"

//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
type Player struct{
	name string
	maxHealth int 
	health int
	maxEnergy int 
	energy int 
}

func (player *Player) IncrementHealth(heal int){
	player.health += heal
	if player.health >= player.maxHealth{
		fmt.Println(player.name, "is at Max Health")
		fmt.Println("Went from", player.health - heal, "health to", player.maxHealth, "health")
		player.health = player.maxHealth
	} else {
	fmt.Println(player.name, "went from", player.health - heal, "health to", player.health, "health")
	}
}

func (player *Player) DecrementHealth(damage int){
	player.health -= damage
	if player.health <= 0{
		fmt.Println(player.name, "is Dead")
		fmt.Println("Went from", player.health + damage, "to 0 health")
		player.health = 0
	} else {
	fmt.Println(player.name, "went from", player.health + damage, "health to", player.health, "health")
	}
}

func (player *Player) IncrementEnergy(heal int){
	player.energy += heal
	if player.energy >= player.maxEnergy{
		fmt.Println(player.name, "is at Max Energy")
		fmt.Println("Went from", player.energy - heal, "energy to", player.maxEnergy, "energy")
		player.energy = player.maxEnergy
	} else {
	fmt.Println(player.name, "went from", player.energy - heal, "energy to", player.energy, "energy")
	}
}

func (player *Player) DecrementEnergy(damage int){
	player.energy -= damage
	if player.energy <= 0{
		fmt.Println(player.name, "is tired")
		fmt.Println("Went from", player.energy + damage, "energy to 0 energy")
		player.energy = 0
	} else {
	fmt.Println(player.name, "went from", player.energy + damage, "energy to", player.energy, "energy")
	}
}

func main() {

}
