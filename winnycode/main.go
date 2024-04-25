package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"strconv"
	"strings"
	"time"
)

const (
	maxConcurrentDownloads = 5
)

type DownloadResult struct {
	Duration time.Duration
	Error    error
}

type WriteCounter struct {
	Total int64
	Count int64
	Ch    chan<- int
}

func main() {
	for {
		clearScreen()
		fmt.Println("Available Categories:")
		categoriesList := make([]string, 0, len(categories))
		for category := range categories {
			categoriesList = append(categoriesList, category)
		}
		sort.Strings(categoriesList)
		for i, category := range categoriesList {
			fmt.Printf("[%d] %s\n", i+1, category)
		}

		fmt.Print("Enter the category number or 0 to go home: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "0":
			continue
		default:
			catNum, err := strconv.Atoi(choice)
			if err != nil || catNum < 1 || catNum > len(categoriesList) {
				fmt.Println("Invalid category number.")
				continue
			}
			selectedCategory := categoriesList[catNum-1]

			if selectedCategory == "Windows Tweaks" {
				clearScreen()
				fmt.Println("Windows Tweaks Menu:")
				fmt.Println("[1] Remove Microsoft Edge")
				fmt.Println("[2] Activate Windows")
				fmt.Println("[3] Chris Titus Tech's Windows Utility")
				fmt.Print("Enter the tweak number or 0 to go home: ")
				scanner.Scan()
				tweakChoice := scanner.Text()
				if tweakChoice == "0" {
					continue
				}

				tweakNum, err := strconv.Atoi(tweakChoice)
				if err != nil || (tweakNum != 1 && tweakNum != 2 && tweakNum != 3) {
					fmt.Println("Invalid tweak number.")
					continue
				}

				clearScreen()
				switch tweakNum {
				case 1:
					removeMicrosoftEdge()
				case 2:
					activateWindows()
				case 3:
					ChristTitusTechsWinutil()
				}
				fmt.Println("Press Enter to continue...")
				scanner.Scan()
				continue
			}

			clearScreen()

			clearScreen()

			clearScreen()
			fmt.Printf("Available Applications in %s:\n", selectedCategory)
			apps := categories[selectedCategory]
			for i, app := range apps {
				fmt.Printf("[%d] %s\n", i+1, app)
			}

			fmt.Print("Enter the application numbers to download (comma-separated, ex: 1,2,3 or just 1) or 0 to go home: ")
			scanner.Scan()
			appChoices := scanner.Text()
			if appChoices == "0" {
				continue
			}

			appNumsStr := strings.Split(appChoices, ",")
			var appNums []int
			for _, numStr := range appNumsStr {
				num, err := strconv.Atoi(strings.TrimSpace(numStr))
				if err != nil || num < 1 || num > len(apps) {
					fmt.Printf("Invalid application number: %s\n", numStr)
					continue
				}
				appNums = append(appNums, num)
			}

			if len(appNums) == 0 {
				fmt.Println("No valid application numbers selected.")
				continue
			}

			selectedApps := make([]string, len(appNums))
			for i, num := range appNums {
				selectedApps[i] = apps[num-1]
			}
			downloadApps(selectedApps)

			for _, installer := range selectedApps {
				rosu := "\033[31m"
				reset := "\033[0m"
				fmt.Printf("Running installer for %s...\n", installer)
				fmt.Printf("%s%s%s\n", rosu, "!!!!!!!DO NOT CLOSE THIS WINDOW TERMINAL!!!!", reset)

				ext := ".exe"
				if strings.Contains(installer, "Blender") || strings.Contains(installer, "GoLang") || strings.Contains(installer, "EpicGames") {
					ext = ".msi"
				}
				err := runInstaller(installer + ext)
				if err != nil {
					fmt.Printf("Error running installer for %s: %v\n", installer, err)
					continue
				}
				fmt.Printf("Installation of %s completed.\n", installer)
			}
		}
	}
}
