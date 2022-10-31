package main

import "testing"

//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
func newPlayer() Player {
	return Player{
		name:      "test",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}
}

//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//

func TestHealth(t *testing.T) {
	player := newPlayer()
	player.addHealth(999)
	if player.health > player.maxHealth {
		t.Fatalf("health went over the limit:%v,want:%v", player.health, player.maxHealth)
	}
	player.applyDamage(player.maxHealth + 1)
	// if player.health < 0 {
	// 	t.Fatalf("Health:%v Minimum:0", player.health)
	// }
	if player.health > player.maxHealth {
		t.Fatalf("Health:%v Minimum:%v", player.health, player.maxHealth)
	}
}
func TestEnergy(t *testing.T) {
	player := newPlayer()
	player.addEnergy(999)
	if player.energy > player.maxEnergy {
		t.Fatalf("health went over the limit:%v,want:%v", player.energy, player.maxEnergy)
	}
	player.applyDamage(player.maxHealth + 1)
	// if player.energy < 0 {
	// 	t.Fatalf("Health:%v Minimum:0", player.energy)
	// }
	if player.energy > player.maxEnergy {
		t.Fatalf("Health:%v Minimum:%v", player.energy, player.maxEnergy)
	}
}

//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
