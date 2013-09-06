package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	switch os.Args[1] {
	case "new":
		m := make(map[string]string)
		if len(os.Args) == 2 {
			in := bufio.NewReader(os.Stdin)
			for {
				raw, _, err := in.ReadLine()
				if err != nil {
					break
				}
				pair := string(raw)
				if strings.Contains(pair, ":") {
					split := strings.SplitN(pair, ":", 2)
					m[split[0]] = split[1]
				}
			}
		} else {
			for _, pair := range os.Args[2:] {
				if strings.Contains(pair, ":") {
					split := strings.SplitN(pair, ":", 2)
					m[split[0]] = split[1]
				}
			}
		}
		WriteMap(m)
	case "get":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		m := ReadMap(os.Args[3])
		key := os.Args[2]
		if value, ok := m[key]; ok {
			fmt.Println(value)
		}
	case "set":
		if len(os.Args) != 5 {
			os.Exit(1)
		}
		m := ReadMap(os.Args[4])
		key := os.Args[2]
		value := os.Args[3]
		m[key] = value
		WriteMap(m)
	case "delete":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		m := ReadMap(os.Args[3])
		key := os.Args[2]
		delete(m, key)
		WriteMap(m)
	case "list":
		if len(os.Args) != 3 {
			os.Exit(1)
		}
		for key, value := range ReadMap(os.Args[2]) {
			fmt.Println(key + ":" + value)
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
		search := os.Args[3]
		m := ReadMap(os.Args[4])
		switch os.Args[2] {
		case "key":
			if _, ok := m[search]; ok {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		case "value":
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
