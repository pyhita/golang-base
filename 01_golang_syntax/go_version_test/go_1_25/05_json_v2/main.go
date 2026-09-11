package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

type Person struct {
	Name      string                 `json:"name"`
	BirthDate time.Time              `json:"birth_date,format:DateOnly"` // custom date format
	Address   `json:",inline"`       // inline address fields
	Extra     map[string]interface{} `json:",unknown"` // catch unknown JSON keys
}

func main() {
	src := []byte(`{
  "name": "Alice",
  "birth_date": "2001-07-15",
  "street": "123 Main St",
  "city": "Wonderland",
  "hobby": "Adventuring",
  "friends": [{"name": "Bob"}, {"name": "Cindy"}]
 }`)

	var p Person
	if err := json.Unmarshal(src, &p); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s\nBirthDate: %s\nAddress: %s, %s\n", p.Name, p.BirthDate.Format("2006-01-02"), p.Street, p.City)
	fmt.Println("Extra fields:", p.Extra)

	// Marshal with indentation option
	//out, err := json.Marshal(p, jsontext.WithIndent ("  "))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Marshaled JSON:")
	//fmt.Println(string(out))
}
