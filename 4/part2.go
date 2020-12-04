package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// This code to remain largely unaltered from the state it was used to solve,
// mostly for historical purposes.

func main() {
	f, err := ioutil.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}

	var fields = []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		"cid",
	}

	var valid int

	var passports = strings.Split(string(f), "\n\n")

	for _, passport := range passports {
		words := strings.Fields(passport)

		var present int
		var hasCID bool

		for _, word := range words {
			parts := strings.Split(word, ":")
			k, v := parts[0], parts[1]

			for _, field := range fields {
				if field != k {
					continue
				}

				switch field {
				case "cid":
					hasCID = true

				case "byr":
					if len(v) != 4 {
						continue
					}
					yr, err := strconv.Atoi(v)
					if err != nil || yr < 1920 || yr > 2002 {
						continue
					}

				case "iyr":
					if len(v) != 4 {
						continue
					}
					yr, err := strconv.Atoi(v)
					if err != nil || yr < 2010 || yr > 2020 {
						continue
					}

				case "eyr":
					if len(v) != 4 {
						continue
					}
					yr, err := strconv.Atoi(v)
					if err != nil || yr < 2020 || yr > 2030 {
						continue
					}

				case "hgt":
					switch {
					case strings.HasSuffix(v, "cm"):
						n, err := strconv.Atoi(strings.TrimSuffix(v, "cm"))
						if err != nil || n < 150 || n > 193 {
							continue
						}

					case strings.HasSuffix(v, "in"):
						n, err := strconv.Atoi(strings.TrimSuffix(v, "in"))
						if err != nil || n < 59 || n > 76 {
							continue
						}

					default:
						continue
					}

				case "hcl":
					if !strings.HasPrefix(v, "#") || len(v) != 7 {
						continue
					}
					_, err := strconv.ParseInt(strings.TrimPrefix(v, "#"), 16, 32)
					if err != nil {
						continue
					}

				case "ecl":
					switch v {
					case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
					default:
						continue
					}

				case "pid":
					_, err := strconv.Atoi(v)
					if err != nil || len(v) != 9 {
						continue
					}
				}

				present++
			}
		}

		if present == 8 || (present == 7 && !hasCID) {
			valid++
		}
	}

	fmt.Println(valid)
}
