package main

import (
	"github.com/hanstanawi/monster-slayer/actions"
	"github.com/hanstanawi/monster-slayer/interactions"
)

var currentRound int = 0
var gameRounds = []interactions.RoundStats{}

func main() {
	startGame()

	winner := "" // "Player || "Monster" || ""

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interactions.PrintGreeting()
}

func executeRound() string {
	currentRound += 1
	isSpecialround := currentRound%3 == 0
	interactions.PrintAvailableActions(isSpecialround)
	userChoice := interactions.GetPlayerChoice(isSpecialround)

	var playerAttackVal int
	var playerHealVal int

	switch userChoice {
	case "ATTACK":
		playerAttackVal = actions.AttackMonster(isSpecialround)
	case "HEAL":
		playerHealVal = actions.HealPlayer()
	case "SPECIAL ATTACK":
		playerAttackVal = actions.AttackMonster(true)
	default:
		break
	}

	monsterAttackVal := actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthAmounts()

	// Gather round data
	roundStatistics := interactions.RoundStats{
		ActionType:       userChoice,
		PlayerAttackVal:  playerAttackVal,
		PlayerHealVal:    playerHealVal,
		MonsterAttackVal: monsterAttackVal,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
	}

	// Print round data
	roundStatistics.PrintStats()

	gameRounds = append(gameRounds, roundStatistics)

	if playerHealth <= 0 {
		return "Monster"
	}
	if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interactions.DeclareWinner(winner)
	interactions.WriteLogFile(&gameRounds)
}
