package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode strings.Builder
	for _, ch := range s.name {

		nameEncode.WriteString(fmt.Sprintf("%d", ch))
	}
	s.nameEncode = nameEncode.String()
	return s.nameEncode
}

func (s *student) Decode() string {
	var nameDecode strings.Builder
	for i := 0; i < len(s.nameEncode); i += 3 {

		ch := s.nameEncode[i : i+3]
		fmt.Sscanf(ch, "%d", &ch)
		nameDecode.WriteString(fmt.Sprintf("%c", ch))
	}
	return nameDecode.String()
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + "is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student’s name " + a.name + "is : " + c.Decode())
	}
}
