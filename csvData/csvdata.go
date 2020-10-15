package csv

import (
    "encoding/csv"
    "log"
    "os"
    s "strings"
    model "../models"
)

//________________________________________________________
var Battles []model.Battle
var Characters []model.Character
var Houses []model.House
//________________________________________________________
var charactersRecords string
var battleRecorsds string
var housesRecords string
//________________________________________________________
func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}

//________________________________________________________
func searchCharacter(characterName string) []model.Character{
    characterMatched := []model.Character{}
    for _, character := range Characters {
      if characterName == character.Name {
         characterMatched = append(characterMatched, character)
      }
    }
    return characterMatched
}
func searchBattle(location string) []model.Battle{
    battleMatched := []model.Battle{}
    for _, battle := range Battles {
      if location == battle.Location {
         battleMatched = append(battleMatched, battle)
      }
    }
    return battleMatched
}
//________________________________________________________
func ReadData(filePath string) {
    file, err1 := os.Open(filePath)
    checkError("Unable to read input file "+filePath, err1)
    defer file.Close()
    filePath = filePath[9:len(filePath)]
    switch filePath {
    case "battles.csv":
        csvReader := csv.NewReader(file)
        BattleRecorsds, err2 := csvReader.ReadAll()
        checkError("Unable to parse file as CSV for "+filePath, err2)
        defer file.Close()
        Battles = []model.Battle{}
        for _, record := range BattleRecorsds {
            battle := model.Battle{
                 BattleId        : record[0],
                 BattleName      : record[1],
                 AttackerKing    : record[2],
                 Defenderking    : record[3],
                 AttackerHouse   : record[4],
                 ParticipateHouse: record[5],
                 BattleType      : record[6],
                 Year            : record[7],
                 Location        : record[8],
                 Region          : record[9]}
            Battles = append(Battles, battle)
        }
    case "characters.csv":
        csvReader := csv.NewReader(file)
        charactersRecords, err2 := csvReader.ReadAll()
        checkError("Unable to parse file as CSV for "+filePath, err2)
        defer file.Close()
        Characters = []model.Character{}
        for _, record := range charactersRecords {
            character := model.Character{
                  CharacterId: record[0],
                  Name       : record[1],
                  Title      : record[2],
                  Male       : record[3],
                  Culture    : record[4],
                  Mother     : record[5],
                  Father     : record[6],
                  House      : record[7],
                  Age        : record[8]}
            Characters = append(Characters, character)
        }
    case "houses.csv":
        csvReader := csv.NewReader(file)
        housesRecords, err2 := csvReader.ReadAll()
        checkError("Unable to parse file as CSV for "+filePath, err2)
        defer file.Close()
        Houses = []model.House{}
        for index, record := range housesRecords {
            if index != 0 {//skip headers
              house := model.House{
                     HouseId       : record[0],
                     CharacterName : record[1],
                     Location      : record[2],
                     HouseName     : record[3],
                     DeathYear     : record[4],
                     BookOfDeath   : record[5],
                     Gender        : record[6],
                     Nobility      : record[7],
                     Battle        : searchBattle(record[2]),
                 	   Character     :searchCharacter(record[1])}
              Houses = append(Houses, house)
            }
        }
    file.Close()
  }
}
//________________________________________________________
func WriteData(filePath string) {
    file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
    checkError("Cannot create file", err)
    defer file.Close()
    filePath = filePath[9:len(filePath)]
    switch filePath {
    case "battles.csv":
      file.Seek(0, 0)
      writer := csv.NewWriter(file)
      defer writer.Flush()
      for _, battle := range Battles {
          record := []string{s.Replace(battle.BattleId, "battle", "", -1), battle.BattleName, battle.AttackerKing,
              battle.Defenderking, battle.AttackerHouse, battle.ParticipateHouse,
              battle.BattleType, battle.Year,battle.Location,battle.Region}
          err := writer.Write(record)
          checkError("Cannot write to file", err)
      }
      writer.Flush()
    case "characters.csv":
      file.Seek(0, 0)
      writer := csv.NewWriter(file)
      defer writer.Flush()
      for _, character := range Characters {
          record := []string{s.Replace(character.CharacterId, "character", "", -1), character.Name, character.Title,
              character.Male, character.Culture, character.Mother,
              character.Father, character.House,character.Age}
          err := writer.Write(record)
          checkError("Cannot write to file", err)
      }
      writer.Flush()

    case "houses.csv":
      file.Seek(0, 0)
      writer := csv.NewWriter(file)
      defer writer.Flush()
      for _, house := range Houses {
          record := []string{s.Replace(house.HouseId, "house", "", -1), house.CharacterName, house.Location,
              house.HouseName, house.DeathYear, house.BookOfDeath,
              house.Gender, house.Nobility}
          err := writer.Write(record)
          checkError("Cannot write to file", err)
      }
      writer.Flush()

    file.Close()
  }
}
