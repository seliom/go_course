//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
package main

import "testing"

type Player struct{
	name string
	maxHealth int 
	health int
	maxEnergy int 
	energy int 
}

func TestIncrement(t *testing.T){
	playerOne := Player{"Jack", 100, 100, 100, 100}

	playerOne.IncrementHealth(10)
	if playerOne.health > playerOne.maxHealth{
		t.errorf("Health can not go above its maximum health")
	}

	playerOne.IncrementEnergy(10)
	if playerOne.energy > playerOne.maxEnergy{
		t.errorf("Energy can not go above its maximum energy")
	}

}

func TestDecrement(t *testing.T){
	playerOne := Player{"Jack", 0, 100, 0, 100}

	playerOne.DecrementHealth(10)
	if playerOne.health < 0{
		t.errorf("Health can not go below 0")
	}

	playerOne.DecrementEnergy(10)
	if playerOne.energy < 0{
		t.errorf("Energy can not go below 0")
	}

}