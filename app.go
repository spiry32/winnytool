package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"sort"
)

var categories = map[string][]string{
	"Browsers": {
		"Google Chrome",
		"Brave",
		"Vivaldi",
		"Microsoft Edge",
		"Opera",
		"Opera GX",
		"Firefox",
	},
	"Messaging": {
		"Discord",
		"Zoom",
		"Skype",
	},
	"Media": {
		"VLC Media Player",
		"Spotify",
	},
	"File Sharing": {
		"qBittorrent",
	},
	"Compression": {
		"winrar",
		"7-Zip",
		"PeaZip",
	},
}

var applications = map[string]string{
	// Browsers
	"Google Chrome":  "https://dl.google.com/chrome/install/8795.700/chrome_installer.exe",
	"Brave":          "https://laptop-updates.brave.com/latest/winx64",
	"Vivaldi":        "https://downloads.vivaldi.com/stable/Vivaldi.3.8.2259.37.exe",
	"Microsoft Edge": "https://go.microsoft.com/fwlink/?linkid=2136346",
	"Opera":          "https://www.opera.com/download/get/?id=47899&location=360&nothanks=yes&sub=marine",
	"Opera GX":       "https://www.opera.com/download/get/?id=47899&location=360&nothanks=yes&sub=marine",
	"Firefox":        "https://www.mozilla.org/firefox/download/thanks/",
	// Messaging
	"Discord": "https://discord.com/api/download?platform=win",
	"Zoom":    "https://zoom.us/client/latest/ZoomInstaller.exe",
	"Skype":   "https://get.skype.com/go/getskype-full",
	// Media
	"VLC Media Player": "https://get.videolan.org/vlc/3.0.12/win64/vlc-3.0.12-win64.exe",
	"Spotify":          "https://download.scdn.co/SpotifyFullSetup.exe",
	// File Sharing
	"qBittorrent": "https://sourceforge.net/projects/qbittorrent/files/qbittorrent-win32/qbittorrent_4.3.8_setup.exe/download",
	// Compression
	"winrar": "https://www.rarlab.com/rar/wrar601.exe",
	"7-Zip":  "https://www.7-zip.org/a/7z2106-x64.exe",
	"PeaZip": "https://osdn.net/frs/redir.php?m=netix&f=peazip%2F7.10.0%2Fpeazip_portable-7.10.0.WINDOWS.exe",
}

func downloadFile(url, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	return err
}

func runInstaller(filename string) error {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/C", filename)
		return cmd.Run()
	default:
		fmt.Println("Currently supported only on Windows.")
		return nil
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
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
			catNum := choice[0] - '0'
			if catNum < 1 || int(catNum) > len(categories) {
				fmt.Println("Invalid category number.")
				continue
			}
			selectedCategory := categoriesList[int(catNum)-1]

			clearScreen()
			fmt.Printf("Available Applications in %s:\n", selectedCategory)
			apps := categories[selectedCategory]
			for i, app := range apps {
				fmt.Printf("[%d] %s\n", i+1, app)
			}

			fmt.Print("Enter the application number or 0 to go home: ")
			scanner.Scan()
			appChoice := scanner.Text()
			switch appChoice {
			case "0":
				continue
			default:
				appNum := appChoice[0] - '0'
				if appNum < 1 || int(appNum) > len(apps) {
					fmt.Println("Invalid application number.")
					continue
				}
				selectedApp := apps[int(appNum)-1]
				url, ok := applications[selectedApp]
				if !ok {
					fmt.Println("Invalid application selection.")
					continue
				}

				fmt.Printf("Downloading %s installer...\n", selectedApp)
				err := downloadFile(url, "installer.exe")
				if err != nil {
					fmt.Println("Error downloading installer:", err)
					continue
				}
				fmt.Println("Installer downloaded successfully.")

				fmt.Println("Running installer...")
				err = runInstaller("installer.exe")
				if err != nil {
					fmt.Println("Error running installer:", err)
					continue
				}
				fmt.Println("Installation completed.")
			}
		}
	}
}
