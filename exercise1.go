package main

import (
	"os"
	"flag"
	"fmt"
	"encoding/csv"
	"time"
)

func main(){
	file := flag.String("Filecsv", "problems.csv", "Enter the input test file")
	timelimit := flag.Int("timelimit", 3, "Test period")
	flag.Parse()
	f, err := os.Open(*file)
	if err != nil {
		fmt.Println( err)
		os.Exit(1)
	}
	defer f.Close()
 	record := csv.NewReader(f)
 	count:= 0

	r, _ := record.ReadAll()
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
	for _,row := range r{
		fmt.Println(row[0])
		chane := make(chan string)
		go func() {
			ans := ""
			fmt.Scanln(&ans)
			if ans == row[1] {
				count++
			}
			chane <- ans
		}()
		select {
		case <-timer.C:
				fmt.Println("Time's up!..", count)
				return
		case <-chane:
			continue
		}
	}
}
