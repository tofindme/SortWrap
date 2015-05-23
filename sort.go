package main

import (
	"fmt"
	"sort"
)

type People struct {
	name   string
	age    int
	career string
	hobby  string
}

type Peoples []People

func (s Peoples) Len() int {
	return len(s)
}

func (s Peoples) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ByNameSort struct {
	Peoples
	sort string //asc or desc?
}

func (by ByNameSort) Less(i, j int) bool {
	if by.sort == "asc" {
		return by.Peoples[i].name > by.Peoples[j].name
	} else if by.sort == "desc" {
		return by.Peoples[i].name < by.Peoples[j].name
	}
	return false
}

type ByAgeSort struct {
	Peoples
	sort string //asc or desc?
}

func (by ByAgeSort) Less(i, j int) bool {
	if by.sort == "asc" {
		return by.Peoples[i].age > by.Peoples[j].age
	} else if by.sort == "desc" {
		return by.Peoples[i].age < by.Peoples[j].age
	}
	return false
}

func main() {

	stu := []People{
		{"Roce", 20, "teacher", "reader"},
		{"Anderow", 35, "coder", "sport"},
		{"Lily", 18, "stu", "eat"},
	}

	sort.Sort(ByNameSort{stu, "asc"})
	fmt.Println("sort by name asc:")
	fmt.Println(stu)

	sort.Sort(ByNameSort{stu, "desc"})
	fmt.Println("sort by name desc:")
	fmt.Println(stu)

	sort.Sort(ByAgeSort{stu, "desc"})
	fmt.Println("sort by age desc:")
	fmt.Println(stu)

	// output:
	// sort by name asc:
	// [{Roce 20 teacher reader} {Lily 18 stu eat} {Anderow 35 coder sport}]
	// sort by name desc:
	// [{Anderow 35 coder sport} {Lily 18 stu eat} {Roce 20 teacher reader}]
	// sort by age desc:
	// [{Lily 18 stu eat} {Roce 20 teacher reader} {Anderow 35 coder sport}]

}
