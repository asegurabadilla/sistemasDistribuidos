/*
 * GoT API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"net/http"
	"encoding/json"
	"path"
	s "strings"
	//"fmt"
	data "../csvData"
	model "../models"
)
//_______________________________________________________
func findHouse(x string) int {
	x = x[5:len(x)]
	for i, house := range data.Houses {
		if x == house.HouseId {
			return i
		}
	}
	return -1
}
func findCharacter(x string) int {
	x = x[9:len(x)]
	for i, character := range data.Characters {
		if x == character.CharacterId {
			return i
		}
	}
	return -1
}
func findBattle(x string) int {
	x = x[6:len(x)]
	for i, battle := range data.Battles {
		if x ==battle.BattleId {
			return i
		}
	}
	return -1
}
//_______________________________________________________
func BattleBattleIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	index := findBattle(id)
	newBattles := append(data.Battles[:index], data.Battles[index+1:]...)
	data.Battles = newBattles
	data.WriteData("csvFiles/battles.csv")
	data.ReadData("csvFiles/battles.csv")//refresh data
	data.ReadData("csvFiles/houses.csv")//refresh data
	w.WriteHeader(http.StatusOK)
}

func BattleBattleIdGet(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findBattle(id)
	if i == -1 {
		return
	}
	dataJson, _ := json.Marshal(data.Battles[i])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func BattleBattleIdPut(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	for i, battle := range data.Battles {
		if s.Replace(id, "battle", "", -1) == battle.BattleId {
			err := json.NewDecoder(r.Body).Decode(&battle)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			data.Battles[i] = battle
		}
	}
	data.WriteData("csvFiles/battles.csv")
	data.ReadData("csvFiles/battles.csv")//refresh data
	data.ReadData("csvFiles/houses.csv")//refresh data
	w.WriteHeader(http.StatusOK)
}

func BattlePost(w http.ResponseWriter, r *http.Request) {
	var battle model.Battle
	err := json.NewDecoder(r.Body).Decode(&battle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data.Battles = append(data.Battles, battle)
	data.WriteData("csvFiles/battles.csv")
	data.ReadData("csvFiles/battles.csv")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func CharacterCharacterIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	index := findCharacter(id)
	newCharacters := append(data.Characters[:index], data.Characters[index+1:]...)
	data.Characters = newCharacters
	data.WriteData("csvFiles/characters.csv")
	data.ReadData("csvFiles/characters.csv")//refresh data
	data.ReadData("csvFiles/houses.csv")//refresh data
	w.WriteHeader(http.StatusOK)
}

func CharacterCharacterIdGet(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findCharacter(id)
	if i == -1 {
		return
	}
	dataJson, _ := json.Marshal(data.Characters[i])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func CharacterCharacterIdPut(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	for i, character := range data.Characters {
		if s.Replace(id, "character", "", -1) == character.CharacterId {
			err := json.NewDecoder(r.Body).Decode(&character)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			data.Characters[i] = character
		}
	}
	data.WriteData("csvFiles/characters.csv")
	data.ReadData("csvFiles/characters.csv")//refresh data
	data.ReadData("csvFiles/houses.csv")//refresh data
	w.WriteHeader(http.StatusOK)
}

func CharacterPost(w http.ResponseWriter, r *http.Request) {
	var character model.Character
	err := json.NewDecoder(r.Body).Decode(&character)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data.Characters = append(data.Characters, character)
	data.WriteData("csvFiles/characters.csv")
	data.ReadData("csvFiles/characters.csv")//refresh data
	w.WriteHeader(http.StatusOK)
}

func HouseHouseIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	index := findHouse(id)
	newHouses := append(data.Houses[:index], data.Houses[index+1:]...)
	data.Houses = newHouses
	data.WriteData("csvFiles/houses.csv")
	data.ReadData("csvFiles/houses.csv")//refresh data
	w.WriteHeader(http.StatusOK)
}

func HouseHouseIdGet(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findHouse(id)
	if i == -1 {
		return
	}
	dataJson, _ := json.Marshal(data.Houses[i])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func HouseHouseIdPut(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	for i, house := range data.Houses {
		if s.Replace(id, "house", "", -1) == house.HouseId {
			err := json.NewDecoder(r.Body).Decode(&house)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			data.Houses[i] = house
		}
	}
	data.WriteData("csvFiles/houses.csv")
	data.ReadData("csvFiles/houses.csv")//refresh data
	w.WriteHeader(http.StatusOK)
}

func HousePost(w http.ResponseWriter, r *http.Request) {
	var house model.House
	err := json.NewDecoder(r.Body).Decode(&house)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data.Houses = append(data.Houses, house)
	data.WriteData("csvFiles/houses.csv")
	data.ReadData("csvFiles/houses.csv")//refresh data
	w.WriteHeader(http.StatusOK)
}
