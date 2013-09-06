package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func ReadMap(jsonBlob string) map[string]string {
	a := make(map[string]string)
	if err := json.Unmarshal([]byte(jsonBlob), &a); err == nil {
		return a
	}
	return a
}

func WriteMap(a map[string]string) {
	if b, err := json.Marshal(a); err == nil {
		fmt.Println(string(b))
	}
}

func IsInt(number string) bool {
	r, _ := regexp.Compile("^[0-9]+$")
	return r.MatchString(number)
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	switch os.Args[1] {
	case "new":
		m := make(map[string]string)
		if len(os.Args) == 2 {
			in := bufio.NewReader(os.Stdin)
			for i := 0; ; i++ {
				raw, _, err := in.ReadLine()
				if err != nil {
					break
				}
				m[strconv.Itoa(i)] = string(raw)
			}
		} else {
			for key, value := range os.Args[2:] {
				m[strconv.Itoa(key)] = value
			}
		}
		WriteMap(m)
	case "get":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		m := ReadMap(os.Args[3])
		if value, ok := m[os.Args[2]]; ok {
			fmt.Println(value)
		}
	case "set":
		if len(os.Args) != 5 {
			os.Exit(1)
		}
		m := ReadMap(os.Args[4])
		key := os.Args[3]
		if !IsInt(key) {
			os.Exit(1)
		}
		m[key] = os.Args[3]
		WriteMap(m)
	case "delete":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		m := ReadMap(os.Args[3])
		delete(m, os.Args[2])
		WriteMap(m)
	case "list":
		if len(os.Args) != 3 {
			os.Exit(1)
		}
		m := ReadMap(os.Args[2])
		max := 0
		for key := range m {
			if num, _ := strconv.Atoi(key); num > max {
				max = num
			}
		}
		for i := 0; i <= max; i++ {
			if value, ok := m[strconv.Itoa(i)]; ok {
				fmt.Println(value)
			}
		}
	case "len":
		if len(os.Args) != 3 {
			os.Exit(1)
		}
		fmt.Println(len(ReadMap(os.Args[2])))
	case "has":
		if len(os.Args) != 5 {
			os.Exit(1)
		}
		m := ReadMap(os.Args[4])
		switch os.Args[2] {
		case "key":
			if _, ok := m[os.Args[3]]; ok {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		case "value":
			search := os.Args[3]
			for _, value := range m {
				if value == search {
					os.Exit(0)
				}
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	default:
		fmt.Println("not supported:", os.Args[1])
		os.Exit(1)
	}
	os.Exit(0)
}
