package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

//People is a slice of type person
type People []Person

//Len returns the length of input value
func (p People) Len() int {
	return len(p)
}

//Less reports when it is necessary to sort element i before element j
func (p People) Less(i, j int) bool {
	if p[i].birthDay.Unix() == p[j].birthDay.Unix() {
		if p[i].firstName == p[j].firstName {
			return p[i].lastName < p[j].lastName
		}
		return p[i].firstName < p[j].firstName
	}
	return p[i].birthDay.Unix() > p[j].birthDay.Unix()
}

//Swap swaps the elements with indexes i and j.
func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//StrToDate convert date from string to time.Time format
func StrToDate(date string) time.Time {
	res, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return res
}

func main() {
	p := People{
		{"Ivan", "Ivanov", StrToDate("2005-08-10")},
		{"Ivan", "Ivanov", StrToDate("2003-08-05")},
		{"Artiom", "Ivanov", StrToDate("2005-08-10")},
	}

	sort.Sort(p)
	for i:=range(p){
		fmt.Println(p[i].firstName,p[i].lastName, p[i].birthDay.Year(),p[i].birthDay.Month(),p[i].birthDay.Day())
	}
}
