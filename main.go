/*
 * GoT API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"log"
	"net/http"
	"os"
	sw "github.com/asegurabadilla/sistemasDistribuidos/go"
	data "github.com/asegurabadilla/sistemasDistribuidos/csvData"
)

func readCsvFiles(){
	data.ReadData("csvFiles/battles.csv")
	data.ReadData("csvFiles/characters.csv")
	data.ReadData("csvFiles/houses.csv")
}

func main() {
	log.Printf("Server started")
	readCsvFiles()
	router := sw.NewRouter()
	//log.Fatal(http.ListenAndServe(":", router))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
