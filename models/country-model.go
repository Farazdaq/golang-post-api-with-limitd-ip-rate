// This file the main data handler and it hadels the data logic
// it can be connect the data base

package models

// the a struct defnine the country info
type Country struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Population int    `json:"population"`
	City       string `json:"city"`
}

// Data
var Counts = []Country{
	{ID: 1, Name: "Afghanistan", Population: 37172386, City: "Kabul"},
	{ID: 2, Name: "Albania", Population: 2866376, City: "Tirana"},
	{ID: 3, Name: "Algeria", Population: 42228429, City: "Alger"},
	{ID: 3, Name: "American Samoa", Population: 55465, City: "Fagatogo"},
	{ID: 4, Name: "Andorra", Population: 77006, City: "Andorra la Vella"},
	{ID: 5, Name: "Angola", Population: 30809762, City: "Luanda"},
	{ID: 6, Name: "Anguilla", Population: 15094, City: "The Valley"},
	{ID: 7, Name: "Antarctica", Population: 1106, City: "null"},
	{ID: 8, Name: "Antigua and Barbuda", Population: 96286, City: "Saint John's"},
	{ID: 9, Name: "Argentina", Population: 44494502, City: "Buenos Aires"},
	{ID: 10, Name: "Armenia", Population: 2951776, City: "Yerevan"},
}
