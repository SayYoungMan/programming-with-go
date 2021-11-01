package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type name struct {
		Fname string
		Lname string
	}

	var (
		path  string
		names []name
	)

	fmt.Println("Please enter the path to text file")
	fmt.Scan(&path)

	dat, err := os.Open(path)
	if err != nil {
		fmt.Println("Error occurred while reading file: ", err.Error())
	}

	sc := bufio.NewScanner(dat)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		n := strings.Fields(sc.Text())
		fn := name{
			Fname: n[0],
			Lname: n[1],
		}

		names = append(names, fn)
	}

	dat.Close()

	for _, n := range names {
		fmt.Printf("First name: %s\n Last name: %s\n\n", n.Fname, n.Lname)
	}
}
