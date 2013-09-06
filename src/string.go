package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadMap(jsonBlob string) map[string]string {
	a := make(map[string]string)
	if err := json.Unmarshal([]byte(jsonBlob), &a); err == nil {
		return a
	}
	return a
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	switch os.Args[1] {
	case "contains":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		if strings.Contains(os.Args[2], os.Args[3]) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	case "count":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		fmt.Println(strings.Count(os.Args[2], os.Args[3]))
	case "has-prefix":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		if strings.HasPrefix(os.Args[2], os.Args[3]) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	case "has-suffix":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		if strings.HasSuffix(os.Args[2], os.Args[3]) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	case "index":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		fmt.Println(strings.Index(os.Args[2], os.Args[3]))
	case "last-index":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		fmt.Println(strings.LastIndex(os.Args[2], os.Args[3]))
	case "repeat":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		if count, err := strconv.Atoi(os.Args[3]); err == nil && count > 0 {
			fmt.Println(strings.Repeat(os.Args[2], count))
		} else {
			os.Exit(1)
		}
	case "replace":
		switch len(os.Args) {
		case 5:
			fmt.Println(strings.Replace(os.Args[2], os.Args[3], os.Args[4], -1))
		case 6:
			if n, err := strconv.Atoi(os.Args[5]); err == nil && n > 0 {
				fmt.Println(strings.Replace(os.Args[2], os.Args[3], os.Args[4], n))
			} else {
				os.Exit(1)
			}
		default:
			os.Exit(1)
		}
	case "to-lower":
		if len(os.Args) != 3 {
			os.Exit(1)
		}
		fmt.Println(strings.ToLower(os.Args[2]))
	case "to-upper":
		if len(os.Args) != 3 {
			os.Exit(1)
		}
		fmt.Println(strings.ToUpper(os.Args[2]))
	case "trim":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		fmt.Println(strings.Trim(os.Args[2], os.Args[3]))
	case "trim-left":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		fmt.Println(strings.TrimLeft(os.Args[2], os.Args[3]))
	case "trim-prefix":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		fmt.Println(strings.TrimPrefix(os.Args[2], os.Args[3]))
	case "trim-right":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		fmt.Println(strings.TrimRight(os.Args[2], os.Args[3]))
	case "trim-suffix":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		fmt.Println(strings.TrimSuffix(os.Args[2], os.Args[3]))
        case "len":
		if len(os.Args) != 3 {
			os.Exit(1)
		}
		fmt.Println(len(os.Args[2]))
	case "split":
		switch len(os.Args) {
		case 4:
			for _, value := range strings.Split(os.Args[2], os.Args[3]) {
				fmt.Println(value)
			}
		case 5:
			if n, err := strconv.Atoi(os.Args[4]); err == nil && n > 0 {
				for _, value := range strings.SplitN(os.Args[2], os.Args[3], n) {
					fmt.Println(value)
				}
			} else {
				os.Exit(1)
			}
		default:
			os.Exit(1)
		}
	case "join":
		if len(os.Args) != 4 {
			os.Exit(1)
		}
		newArray := []string{}
		for _, value := range ReadMap(os.Args[2]) {
			newArray = append(newArray, value)
		}
		fmt.Println(strings.Join(newArray, os.Args[3]))
	case "line":
		switch len(os.Args) {
		case 3:
			n, err := strconv.Atoi(os.Args[2])
			if err != nil {
				os.Exit(1)
			}
			in := bufio.NewReader(os.Stdin)
			for i := 0; ; i++ {
				raw, _, err := in.ReadLine()
				if err != nil {
					break
				}
				if i+1 == n {
					fmt.Println(string(raw))
					break
				}

			}
		case 4:
			n, err1 := strconv.Atoi(os.Args[2])
			length, err2 := strconv.Atoi(os.Args[3])
			if err1 != nil || err2 != nil {
				os.Exit(1)
			}
			in := bufio.NewReader(os.Stdin)
			for i := 0; ; i++ {
				raw, _, err := in.ReadLine()
				if err != nil {
					break
				}
				if i+1 >= n+length {
					break
				}

				if n <= i+1 && i+1 < n+length {
					fmt.Println(string(raw))
				}
			}
		default:
			os.Exit(1)
		}
	default:
		fmt.Println("not supported:", os.Args[1])
		os.Exit(1)
	}
	os.Exit(0)
}
