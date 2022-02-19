package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	events "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

func main() {
	f, err := os.Open("./demos/003532904928626343963_0061388492.dem")
	if err != nil {
		log.Fatalf("err opening file: %v", err)
	}
	defer f.Close()

	p := dem.NewParser(f)
	defer p.Close()

	// Register handler on killer events
	p.RegisterEventHandler(func(e events.Kill) {

	})

	playerHurtHeader := []string{"player name", "userID", "team", "ViewDirectionX", "ViewDirectionY", "velocityX", "velocityY", "velocityZ", "Health", "Armor", "HealthDamage", "ArmorDamage", "HealthDamageTake", "ArmorDamageTaken", "IsAirborne", "Attacker Name", "attackerID", "AttackerWeapon", "HitGroup"}

	var playerHurtRecord [][]string
	playerHurtRecord = append(playerHurtRecord, playerHurtHeader)

	p.RegisterEventHandler(func(e events.PlayerHurt) {

		playerHurtRecord = getPlayerHurtInfo(e, playerHurtRecord)

	})
	p.RegisterEventHandler(func(e events.RoundStart) {})

	// Parse to end
	err = p.ParseToEnd()
	if err != nil {
		panic(err)
	}

	pf, err := os.Create("playerHurt.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer pf.Close()

	w := csv.NewWriter(pf)
	err = w.WriteAll(playerHurtRecord)
	if err != nil {
		log.Fatalln("failed to write", err)
	}
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func getPlayerHurtInfo(e events.PlayerHurt, r [][]string) [][]string {

	p := e.Player

	if p == nil {
		return r
	}

	if e.Attacker == nil {
		return r
	}

	if e.Weapon == nil {
		return r
	}

	// player
	playerName := p.Name
	userId := fmt.Sprintf("%v", p.UserID)
	team := fmt.Sprintf("%v", p.Team)
	viewDirectionX := fmt.Sprintf("%f", p.ViewDirectionX())
	viewDirectionY := fmt.Sprintf("%f", p.ViewDirectionY())

	velocityX := fmt.Sprintf("%f", p.Velocity().X)
	velocityY := fmt.Sprintf("%f", p.Velocity().Y)
	velocityV := fmt.Sprintf("%f", p.Velocity().Z)

	// isBot := p.IsBot
	// isConnected := p.IsConnected
	// isDefusing := p.IsDefusing
	// isPlanting := p.IsPlanting
	// isReloading := p.IsReloading

	// isCrouching := p.IsDucking()
	isAirborne := strconv.FormatBool(p.IsAirborne())
	// clipAmmoLeft := p.ActiveWeapon().AmmoInMagazine()
	// reserveAmmo := p.ActiveWeapon().AmmoReserve()

	// isScoped := p.IsScoped()
	// weaponZoom := p.ActiveWeapon().ZoomLevel()

	// attacker

	attacker := e.Attacker.Name
	attackerID := fmt.Sprintf("%v", e.Attacker.UserID)
	attackerWeapon := ""

	// weapon that hurt the player

	// weapon := e.Weapon.String()
	// weaponType := e.Weapon.Type

	// hit group
	hitGroup := fmt.Sprintf("%v", e.HitGroup)

	//
	health := fmt.Sprintf("%v", e.Health)
	armor := fmt.Sprintf("%v", e.Armor)
	healthDamage := fmt.Sprintf("%v", e.HealthDamage)
	ArmorDamage := fmt.Sprintf("%v", e.ArmorDamage)
	healthDamageTaken := fmt.Sprintf("%v", e.HealthDamageTaken)
	armorDamageTaken := fmt.Sprintf("%v", e.ArmorDamageTaken)

	newRow := []string{playerName, userId, team, viewDirectionX, viewDirectionY, velocityX, velocityY, velocityV, health, armor, healthDamage, ArmorDamage, healthDamageTaken, armorDamageTaken, isAirborne, attacker, attackerID, attackerWeapon, hitGroup}

	r = append(r, newRow)

	return r

}

// func getKillInfo(k *events.Kill) {
// 	weapon := k.Weapon.OriginalString
// 	victim := k.Victim.String()
// 	killer := k.Killer.String()
// 	assister := k.Assister.String()
// 	isHeadShot := k.IsHeadshot
// 	penetratedObjects := k.PenetratedObjects
// 	assistedFlash := k.AssistedFlash
// 	attackerBlind := k.AttackerBlind
// 	noScope := k.NoScope
// 	throughSmoke := k.ThroughSmoke
// 	distance := k.Distance

// 	isWallBang := k.IsWallBang()

// }
