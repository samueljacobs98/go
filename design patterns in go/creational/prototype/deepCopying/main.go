package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

// func (p *Person) DeepCopy() *Person {
// 	q := *p
// 	q.Address = p.Address.DeepCopy()
// 	copy(q.Friends, p.Friends)
// 	return &q
// }

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

func main() {
	john := Person{"John", &Address{"123 London Road", "London", "UK"}, []string{"Chris", "Matt"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")
	// jane := john
	// jane.Name = "Jane" // ok
	// jane.Address.StreetAddress = "321 Baker Street"

	// jane := john
	// jane.Address = &Address{
	// 	john.Address.StreetAddress,
	// 	john.Address.City,
	// 	john.Address.Country,
	// }

	// jane.Name = "Jane" // ok
	// jane.Address.StreetAddress = "321 Baker Street"

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
