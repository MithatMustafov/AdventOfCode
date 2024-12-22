package main

import (
	"aoc/utils/fileparser"
	"fmt"
	"strconv"
)

type SecretNumber int
type Buyer struct {
    TheSecretNumber SecretNumber
}

type BuyersInfo  struct {
    Buyers []Buyer
	JUMP int
}

func (buyersInfo *BuyersInfo) parseData(input []string) {
    buyersInfo.Buyers = make([]Buyer, 0)
    for _, line := range input{
        secretNumber, _ := strconv.Atoi(line)
        buyersInfo.Buyers = append(buyersInfo.Buyers, Buyer{
			TheSecretNumber: SecretNumber(secretNumber),
		})
    }
}

func (secret SecretNumber) multiplyBy64() SecretNumber { return secret * 64 }
func (secret SecretNumber) multiplyBy2048() SecretNumber { return secret * 2048 }
func (secret SecretNumber) divideBy32() SecretNumber { return secret / 32 }
func (result SecretNumber) mix(secret SecretNumber) SecretNumber { return secret ^ result }
func (secret SecretNumber) prune() SecretNumber { return secret % 16777216 }

func (secret SecretNumber) calcNext() SecretNumber {
    secret = secret.multiplyBy64().mix(secret).prune()
	secret = secret.divideBy32().mix(secret).prune()
	secret = secret.multiplyBy2048().mix(secret).prune()
    return secret
}

func PartOne(buyersInfo BuyersInfo) int {
	var sumSecretNumbers int = 0
	for idxBuyer, buyer := range buyersInfo.Buyers{
		for idxNext := 0; idxNext < buyersInfo.JUMP; idxNext++ {
        	buyer.TheSecretNumber = buyer.TheSecretNumber.calcNext()
		}
		sumSecretNumbers += int(buyer.TheSecretNumber)
		buyersInfo.Buyers[idxBuyer] = buyer
	}
	return sumSecretNumbers
}

type Pattern [4]int

func PartTwo(buyersInfo BuyersInfo) int {
	total := make(map[Pattern]int)

	for _, buyer := range buyersInfo.Buyers{
		last := buyer.TheSecretNumber % 10
		patternList := make([][2]int, 0, 2000)

		for idxNext := 0; idxNext < buyersInfo.JUMP; idxNext++ {
        	buyer.TheSecretNumber = buyer.TheSecretNumber.calcNext()
			temp := buyer.TheSecretNumber % 10
			patternList = append(patternList, [2]int{int(temp) - int(last), int(temp)})
			last = temp
		}

		seen := make(map[Pattern]bool)

		for i := 0; i < len(patternList)-4; i++ {
			var pat Pattern

			for j := 0; j < 4; j++ { pat[j] = patternList[i+j][0] }

			val := patternList[i+3][1]
			if !seen[pat] {
				seen[pat] = true
				total[pat] += val
			}
		}
	}
	
	maxVal := 0
	for _, n := range total {if n > maxVal {maxVal = n}}

	return maxVal
}

func main() {
	fmt.Printf("--- Day 22: Monkey Market ---\n")
	var buyersInfo BuyersInfo
	input := fileparser.ReadFileLines("input", false)
	buyersInfo.parseData(input)
	buyersInfo.JUMP = 2000;
	fmt.Printf("PART[1]: %v\n", PartOne(buyersInfo))
	buyersInfo.parseData(input)
	fmt.Printf("PART[2]: %v\n", PartTwo(buyersInfo))
}
