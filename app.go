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
		"Audacity",
		"iTunes",
	},
	"File Sharing": {
		"qBittorrent",
	},
	"Compression": {
		"winrar",
		"7-Zip",
		"PeaZip",
	},
	"Developer Tools": {
		"WinSCP",
		"Notepad++",
		"PuTTY",
		"Eclipse",
		"FileZilla",
	},
	"Imaging": {
		"Krita",
		"Blender",
		"GIMP",
		"Inkscape",
	},
}

var applications = map[string]string{
	// Browsers
	"Google Chrome [X]":  "https://dl.google.com/chrome/install/8795.700/chrome_installer.exe",
	"Brave":              "https://laptop-updates.brave.com/latest/winx64",
	"Vivaldi":            "https://downloads.vivaldi.com/stable/Vivaldi.3.8.2259.37.exe",
	"Microsoft Edge [X]": "https://go.microsoft.com/fwlink/?linkid=2136346",
	"Opera":              "https://www.opera.com/download/get/?id=47899&location=360&nothanks=yes&sub=marine",
	"Opera GX":           "https://www.opera.com/download/get/?id=47899&location=360&nothanks=yes&sub=marine",
	"Firefox":            "https://download.mozilla.org/?product=firefox-latest-ssl&os=win64&lang=en-US&_gl=1*12wy6k4*_ga*NzE1Njk2OTQ0LjE3MDk1NDU0NTQ.*_ga_MQ7767QQQW*MTcxMTM3MDk2MC4zLjEuMTcxMTM3MDk5Mi4wLjAuMA..",
	// Messaging
	"Discord": "https://discord.com/api/download?platform=win",
	"Zoom":    "https://zoom.us/client/latest/ZoomInstaller.exe",
	"Skype":   "https://get.skype.com/go/getskype-full",
	// Media
	"VLC Media Player": "https://get.videolan.org/vlc/3.0.12/win64/vlc-3.0.12-win64.exe",
	"Spotify":          "https://download.scdn.co/SpotifyFullSetup.exe",
	"Audacity":         "https://github.com/audacity/audacity/releases/download/Audacity-3.4.2/audacity-win-3.4.2-64bit.exe",
	"iTunes":           "https://www.apple.com/itunes/download/win64",
	// File Sharing
	"qBittorrent": "https://sourceforge.net/projects/qbittorrent/files/qbittorrent-win32/qbittorrent_4.3.8_setup.exe/download",
	// Compression
	"winrar": "https://www.rarlab.com/rar/wrar601.exe",
	"7-Zip":  "https://www.7-zip.org/a/7z2106-x64.exe",
	"PeaZip": "https://osdn.net/frs/redir.php?m=netix&f=peazip%2F7.10.0%2Fpeazip_portable-7.10.0.WINDOWS.exe",
	// Dev  Tools
	"WinSCP":    "https://winscp.net/download/files/202403251303916f4d129e8c7043e8537e47cb5d5f5b/WinSCP-6.3.2-Setup.exe",
	"Notepad++": "https://github.com/notepad-plus-plus/notepad-plus-plus/releases/download/v8.6.4/npp.8.6.4.Installer.x64.exe",
	"PuTTY":     "https://the.earth.li/~sgtatham/putty/latest/w64/putty.exe",
	"FileZilla": "https://download.filezilla-project.org/client/FileZilla_3.66.5_win64_sponsored2-setup.exe",
	"Eclipse":   "https://www.eclipse.org/downloads/download.php?file=/oomph/epp/2024-03/R/eclipse-inst-jre-win64.exe",
	// Imaging
	"Krita":       "https://download.kde.org/stable/krita/5.2.2/krita-x64-5.2.2-setup.exe",
	"Blender [X]": "https://www.blender.org/download/release/Blender4.0/blender-4.0.2-windows-x64.msi/",
	"GIMP":        "https://download.gimp.org/gimp/v2.10/windows/gimp-2.10.36-setup-1.exe",
	"Inkscape":    "https://inkscape.org/gallery/item/44617/inkscape-1.3.2_2023-11-25_091e20e-x64.exe",
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
