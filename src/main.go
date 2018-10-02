package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type version struct {
	label  string
	number []int
}

type order struct {
	self     int
	children []*order
}

func readTags() *map[int]*order {

	list := map[int]*order{}

	reader := bufio.NewReader(os.Stdin)
	var ok bool
	for {

		tag, err := reader.ReadString('\n')

		if err != nil || len(tag) == 0 {
			break
		}
		//ToDO clean text on fits version
		//Parse on LastVersion if have labels lik aplha and beta
		s := strings.Split(tag, ".")
		n, _ := strconv.Atoi(s[0])
		o := order{}
		o.self = n

		if _, ok = list[n]; ok {

			list[n].children = append(list[n].children, &o)

		} else {

			list[n] = &o
		}

	}

	return &list
}

func main() {

	labels := map[string]int{
		"alpha": 0,
		"beta":  1,
		"rc":    2,
	}

	tags := readTags()
	fmt.Println("Done")
	fmt.Println(tags)
	fmt.Println(labels)

}
