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
	version                = "1.5.4"
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
		printHeader()
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
				printHeader()
				fmt.Println("Windows Tweaks Menu:")
				fmt.Println("[1] Remove Microsoft Edge")
				fmt.Println("[2] Activate Windows")
				fmt.Println("[3] Chris Titus Tech's Windows Utility")
				fmt.Println("[4] Delete temp files")
				fmt.Println("[5] Install Windows Sandbox")
				fmt.Println("[6] Uninstall Windows Sandbox")
				fmt.Println("[7] Run Disk Cleanup")

				fmt.Print("Enter the tweak number or 0 to go home: ")
				scanner.Scan()
				tweakChoice := scanner.Text()
				if tweakChoice == "0" {
					continue
				}

				tweakNum, err := strconv.Atoi(tweakChoice)
				if err != nil || (tweakNum != 1 && tweakNum != 2 && tweakNum != 3 && tweakNum != 4 && tweakNum != 5 && tweakNum != 6 && tweakNum != 7) {
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
				case 4:
					deleteTemporaryFiles()
				case 5:
					installWindowsSandbox()
				case 6:
					uninstallWindowsSandbox()
				case 7:
					openDiskCleanup()
				}
				fmt.Println("Press Enter to continue...")
				scanner.Scan()
				continue
			}

			clearScreen()
			printHeader()
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
			clearScreen()
			printHeader()
			fmt.Printf("Are you sure you want to download the following apps?? (yes/no)\n%s\n", strings.Join(selectedApps, ", "))
			scanner.Scan()
			confirmation := scanner.Text()

			if confirmation != "yes" {
				fmt.Println("Operation cancelled. Back to menu...")
				continue
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

func printHeader() {
	red := "\033[31m"
	green := "\033[32m"
	yellow := "\033[33m"
	blue := "\033[34m"
	resetc := "\033[0m"

	logo := []string{
		"        __        _____ _   _ _   ___   _______ ___   ___  _       _ ",
		"        \\ \\      / /_ _| \\ | | \\ | \\ \\ / /_   _/ _ \\ / _ \\| |     | |",
		"         \\ \\ /\\ / / | ||  \\| |  \\| |\\ V /  | || | | | | | | |     | |",
		"          \\ V  V /  | || |\\  | |\\  | | |   | || |_| | |_| | |___  |_|",
		"           \\_/\\_/  |___|_| \\_|_| \\_| |_|   |_| \\___/ \\___/|_____| (_)",
	}

	fmt.Println(blue, logo[0])
	fmt.Println(green, logo[1])
	fmt.Println(red, logo[2])
	fmt.Println(yellow, logo[3])
	fmt.Println(blue, logo[4]) // Using blue again for the last line
	fmt.Printf("%s%s%s%s\n", yellow, "				Version: ", version, resetc)
	fmt.Println(green, "-----------------------------------------------------------------------", resetc)
}
