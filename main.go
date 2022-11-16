package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	path = "problems.csv"
)

func exit(msg string) {
	fmt.Print("\n")
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	//init
	userPath := flag.String("path", path, "Use -path filepath(string) or -p=filepath(string) for customizing filepath")
	limit := flag.Int("limit", 30, "Use -limit=duration(int) or -limit duration(int) for customizing limit")

	flag.Parse()

	file, err := newFile(*userPath)
	if err != nil {
		exit("Error opening file")
	}
	dictionary, err := newDictionary(file)
	if err != nil {
		exit("Error creating dictionary")
	}

	// game process
	fmt.Println("Welcome to the Quiz Game!")
	fmt.Println("When you are ready press enter to activate game and timer!")
	fmt.Scanln()
	fmt.Println("Input answer on every question: ")
	correctCount := 0
	done := make(chan struct{})
	go func() {
		for question, answer := range dictionary {
			fmt.Print(question, " = ")
			userAnswer := 0
			fmt.Scan(&userAnswer)
			if userAnswer == answer {
				correctCount++
			}
		}
		close(done)
	}()

	select {
	case <-done:
		exit(fmt.Sprintf("Done!\nYou've got %d correct answer out of %d", correctCount, len(dictionary)))
	case <-time.After(time.Duration(*limit) * time.Second):
		exit(fmt.Sprintf("Timer is over!\nYou've got %d correct answer out of %d", correctCount, len(dictionary)))
	}
}
