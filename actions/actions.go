package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)

var currentMonsterHealth int = MONSTER_HEALTH
var currentPlayerHealth int = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) int {
	minAttackVal := PLAYER_MIN_ATTACK_VALUE
	maxAttackVal := PLAYER_MAX_ATTACK_VALUE
	if isSpecialAttack {
		minAttackVal = PLAYER_MIN_SPECIAL_ATTACK_VALUE
		maxAttackVal = PLAYER_MAX_SPECIAL_ATTACK_VALUE
	}
	damageValue := generateRandomBetween(minAttackVal, maxAttackVal)
	currentMonsterHealth -= damageValue
	return damageValue
}

func HealPlayer() int {
	healValue := generateRandomBetween(PLAYER_MIN_HEAL_VALUE, PLAYER_MAX_HEAL_VALUE)
	healthDiff := PLAYER_HEALTH - currentPlayerHealth
	if healthDiff >= healValue {
		currentPlayerHealth += healValue
		return healValue
	}
	currentPlayerHealth = PLAYER_HEALTH
	return healthDiff
}

func AttackPlayer() int {
	damageValue := generateRandomBetween(MONSTER_MIN_ATTACK_VALUE, MONSTER_MAX_ATTACK_VALUE)
	currentPlayerHealth -= damageValue
	return damageValue
}

func GetHealthAmounts() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}

func generateRandomBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
