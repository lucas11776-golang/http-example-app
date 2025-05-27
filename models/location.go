package models

type Geolocation struct {
	Connection string `connection:"sqlite"`
	ID         int64  `column:"id" type:"primary_key"`
	Address    int64  `column:"Address" type:"string"`
	Latitude   int64  `column:"latitude" type:"float"`
	Longitude  int64  `column:"longitude" type:"float"`
}
