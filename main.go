package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func genFile(name string, count int, ntype int) {
	const maxVal = 1_000_000_000
	var (
		sb strings.Builder
		s  string
	)
	f, e := os.Create(name)
	if e != nil {
		log.Fatal(e)
		return
	}
	defer f.Close()
	for i := 0; i < count; i++ {
		switch ntype {
		case 0:
			s = fmt.Sprintf("%10d", i /* rand.Intn(maxVal)*/)
			//s = fmt.Sprintf("%8d", rand.Int63n(maxVal))
		}
		sb.WriteString(s)
		if i < count-1 {
			sb.WriteString(" ")
		}
	}
	f.WriteString(sb.String())
}

func main() {
	fmt.Println("io...")
	genFile("input.txt", 1000000, 0)

}
