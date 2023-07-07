package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
func main(){
  fmt.Println("Lets Count Those Calories")
  var calorieText = readCalories()
  elfs := sumCalories(calorieText)
  fatElf := findMostDenseElf(elfs)
  fmt.Printf("elfID:",fatElf.sum)
  top3Fatties := findTop3Elfs(elfs)
  fmt.Printf("Top 3 Calorie Sum",top3Fatties)
}

func readCalories() string{
	byteSlice, error := os.ReadFile("input.txt")
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	calories := string(byteSlice)
	return calories
}

func sumCalories(calorieText string) []elf {
  elfs := []elf{}
  var sum = 0
  var elfNumber = 1
  for _, c := range strings.Split(calorieText, "\n"){
    if(c == ""){
      e := elf{id: elfNumber, sum:sum}
      elfs = append(elfs,e)
      fmt.Println(fmt.Sprintf("Elf Number:%d  Calories:%d",elfNumber,sum))
      sum = 0
      elfNumber++
    } else {
      cal, err:= strconv.Atoi(c)
      if(err != nil){
        fmt.Println("Error:", err)
      }
      sum = sum + cal
    }

  }
  return elfs 
}

func findMostDenseElf(elfs []elf) elf{
  sort.Slice(elfs, func(i,j int) bool{
    return elfs[i].sum > elfs[j].sum
  })
  return elfs[0]
}

func findTop3Elfs(elfs []elf) int{
  sort.Slice(elfs, func(i,j int) bool{
    return elfs[i].sum > elfs[j].sum
  })
  return elfs[0].sum + elfs[1].sum + elfs[2].sum
}

type elf struct {
  id int
  sum int
}

