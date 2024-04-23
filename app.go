package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"

	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	maxConcurrentDownloads = 5
)

var categories = map[string][]string{
	"Browsers": {
		"Chromium",
		"Brave",
		"Vivaldi",
		"Microsoft Edge",
		"Opera",
		"Opera GX",
		"Firefox",
		"LibreWolf",
		"Waterfox",
		"Floorp Browser",
		"Tor Browser",
		"Avast Secure Browser",
		"Thorium",
	},
	"Messaging": {
		"Discord",
		"Zoom",
		"Skype",
		"Telegram",
		"Viber",
		"Microsoft Teams",
		"Thunderbird",
	},
	"Media": {
		"VLC Media Player",
		"Spotify",
		"Audacity",
		"iTunes",
		"OBS",
		"Streamlabs",
		"Kdenlive",
	},
	"File Sharing": {
		"qBittorrent [X]",
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
		"Visual Studio Code",
		"Visual Studio",
		"Sublime Text",
	},
	"Imaging": {
		"Krita",
		"Blender",
		"GIMP",
		"Inkscape",
	},
	"Security": {
		"AVG AntiVirus Free",
		"Avast Free Antivirus",
	},
	"Office": {
		"OpenOffice",
		"OnlyOffice",
		"LibreOffice",
		"Foxit PDF Reader",
	},
	"Utilities": {
		"Anyburn",
		"MiniTool Partition Wizard",
		"BalenaEtcher (Installer)",
		"BalenaEtcher (Portable)",
		"Rufus 4.4",
	},
	"Windows Tweaks": {
		"Activate Windows LTSC",
	},
}

var applications = map[string]string{
	// Browsers
	"Chromium":             "https://github.com/Hibbiki/chromium-win64/releases/download/v123.0.6312.123-r1262506/mini_installer.sync.exe",
	"Brave":                "https://laptop-updates.brave.com/latest/winx64",
	"Vivaldi":              "https://downloads.vivaldi.com/stable/Vivaldi.3.8.2259.37.exe",
	"Microsoft Edge":       "https://go.microsoft.com/fwlink/?linkid=2108834&Channel=Stable&language=en&brand=M100",
	"Opera":                "https://net.geo.opera.com/opera/stable/windows?utm_tryagain=yes&utm_source=(direct)&utm_medium=doc&utm_campaign=(direct)&http_referrer=missing&utm_site=opera_com&&utm_lastpage=opera.com/",
	"Firefox":              "https://download.mozilla.org/?product=firefox-latest-ssl&os=win64&lang=en-US&_gl=1*12wy6k4*_ga*NzE1Njk2OTQ0LjE3MDk1NDU0NTQ.*_ga_MQ7767QQQW*MTcxMTM3MDk2MC4zLjEuMTcxMTM3MDk5Mi4wLjAuMA..",
	"LibreWolf":            "https://gitlab.com/api/v4/projects/44042130/packages/generic/librewolf/124.0.1-1/librewolf-124.0.1-1-windows-x86_64-setup.exe",
	"Waterfox":             "https://cdn1.waterfox.net/waterfox/releases/G6.0.11/WINNT_x86_64/Waterfox%20Setup%20G6.0.11.exe",
	"Floorp Browser":       "https://github.com/Floorp-Projects/Floorp/releases/download/v11.11.2/floorp-stub.installer.exe",
	"Tor Browser":          "https://www.torproject.org/dist/torbrowser/13.0.13/tor-browser-windows-x86_64-portable-13.0.13.exe",
	"Avast Secure Browser": "https://www.avast.com/download-thank-you.php?product=ASB&locale=en-ww&direct=1",
	"Thorium":              "https://github.com/Alex313031/Thorium-Win/releases/download/M122.0.6261.132/thorium_AVX2_mini_installer.exe",

	// Messaging
	"Discord":         "https://discord.com/api/download?platform=win",
	"Zoom":            "https://zoom.us/client/latest/ZoomInstaller.exe",
	"Skype":           "https://get.skype.com/go/getskype-full",
	"Telegram":        "https://telegram.org/dl/desktop/win64",
	"Viber":           "https://download.cdn.viber.com/desktop/windows/ViberSetup.exe",
	"Microsoft Teams": "https://go.microsoft.com/fwlink/?linkid=2187217&clcid=0x409&culture=en-us&country=us",
	"Thunderbird":     "https://download.mozilla.org/?product=thunderbird-115.9.0-SSL&os=win64&lang=en-US",

	// Media
	"VLC Media Player": "https://get.videolan.org/vlc/3.0.12/win64/vlc-3.0.12-win64.exe",
	"Spotify":          "https://download.scdn.co/SpotifyFullSetup.exe",
	"Audacity":         "https://github.com/audacity/audacity/releases/download/Audacity-3.4.2/audacity-win-3.4.2-64bit.exe",
	"iTunes":           "https://www.apple.com/itunes/download/win64",
	"OBS":              "https://cdn-fastly.obsproject.com/downloads/OBS-Studio-30.1.1-Full-Installer-x64.exe",
	"Streamlabs":       "https://streamlabs.com/streamlabs-desktop/download?sdb=0",
	"Kdenlive":         "https://download.kde.org/stable/kdenlive/24.02/windows/kdenlive-24.02.1.exe",
	// File Sharing
	"qBittorrent [X]": "https://www.fosshub.com/qBittorrent.html?dwl=qbittorrent_4.6.4_x64_setup.exe",

	// Compression
	"winrar": "https://www.rarlab.com/rar/wrar601.exe",
	"7-Zip":  "https://www.7-zip.org/a/7z2106-x64.exe",
	"PeaZip": "https://osdn.net/frs/redir.php?m=netix&f=peazip%2F7.10.0%2Fpeazip_portable-7.10.0.WINDOWS.exe",

	// Dev  Tools
	"WinSCP":             "https://winscp.net/download/files/202403251303916f4d129e8c7043e8537e47cb5d5f5b/WinSCP-6.3.2-Setup.exe",
	"Notepad++":          "https://github.com/notepad-plus-plus/notepad-plus-plus/releases/download/v8.6.4/npp.8.6.4.Installer.x64.exe",
	"PuTTY":              "https://the.earth.li/~sgtatham/putty/latest/w64/putty.exe",
	"FileZilla":          "https://download.filezilla-project.org/client/FileZilla_3.66.5_win64_sponsored2-setup.exe",
	"Eclipse":            "https://www.eclipse.org/downloads/download.php?file=/oomph/epp/2024-03/R/eclipse-inst-jre-win64.exe",
	"Visual Studio Code": "https://code.visualstudio.com/sha/download?build=stable&os=win32-x64-user",
	"Visual Studio":      "https://c2rsetup.officeapps.live.com/c2r/downloadVS.aspx?sku=community&channel=Release&version=VS2022&source=VSLandingPage&cid=2030:967bcd8b3df0dbedb8ebccd40b730d58",
	"Sublime Text":       "https://download.sublimetext.com/sublime_text_build_4169_x64_setup.exe",

	// Imaging
	"Krita":    "https://download.kde.org/stable/krita/5.2.2/krita-x64-5.2.2-setup.exe",
	"Blender":  "https://www.blender.org/download/release/Blender4.1/blender-4.1.0-windows-x64.msi/",
	"GIMP":     "https://download.gimp.org/gimp/v2.10/windows/gimp-2.10.36-setup-1.exe",
	"Inkscape": "https://inkscape.org/gallery/item/44617/inkscape-1.3.2_2023-11-25_091e20e-x64.exe",

	// Security
	"AVG AntiVirus Free":   "https://bits.avcdn.net/productfamily_ANTIVIRUS/insttype_FREE/platform_WIN_AVG/installertype_ONLINE/build_RELEASE/cookie_mmm_bav_tst_007_402_a?ref=clid_1296527067.1711656977--seid_1711656976--senu_1&alt=en-ww",
	"Avast Free Antivirus": "https://www.avast.com/download-thank-you.php?product=FAV-PPC&locale=en-ww&direct=1",

	// Utilities
	"Anyburn":                   "https://www.anyburn.com/anyburn_setup_x64.exe",
	"MiniTool Partition Wizard": "https://cdn2.minitool.com/?p=pw&e=pw-free",
	"BalenaEtcher (Installer)":  "https://github.com/balena-io/etcher/releases/download/v1.18.11/balenaEtcher-Setup-1.18.11.exe",
	"BalenaEtcher (Portable)":   "https://github.com/balena-io/etcher/releases/download/v1.18.11/balenaEtcher-Portable-1.18.11.exe",
	"Rufus (ver 4.4)":           "https://github.com/pbatard/rufus/releases/download/v4.4/rufus-4.4.exe",
	// Office
	"Foxit PDF Reader": "https://www.foxit.com/downloads/latest.html?product=Foxit-Reader&platform=Windows&version=&package_type=&language=English&distID=",
	"OpenOffice":       "https://sourceforge.net/projects/openofficeorg.mirror/files/4.1.15/binaries/en-US/Apache_OpenOffice_4.1.15_Win_x86_install_en-US.exe/download",
	"OnlyOffice":       "https://download.onlyoffice.com/install/desktop/editors/windows/distrib/onlyoffice/DesktopEditors_x64.exe",
	"LibreOffice":      "https://www.libreoffice.org/donate/dl/win-x86_64/24.2.2/ro/LibreOffice_24.2.2_Win_x86-64.msi",
}

type DownloadResult struct {
	Duration time.Duration
	Error    error
}

type WriteCounter struct {
	Total int64
	Count int64
	Ch    chan<- int
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Count += int64(n)
	progress := int(float64(wc.Count) / float64(wc.Total) * 100)
	wc.Ch <- progress
	return n, nil
}

func downloadFile(url, filename string, progress chan<- int) DownloadResult {
	start := time.Now()

	response, err := http.Get(url)
	if err != nil {
		return DownloadResult{Duration: 0, Error: err}
	}
	defer response.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return DownloadResult{Duration: 0, Error: err}
	}
	defer file.Close()

	totalSize, _ := strconv.Atoi(response.Header.Get("Content-Length"))

	counter := &WriteCounter{
		Total: int64(totalSize),
		Count: 0,
		Ch:    progress,
	}

	_, err = io.Copy(io.MultiWriter(file, counter), response.Body)
	if err != nil {
		return DownloadResult{Duration: 0, Error: err}
	}

	duration := time.Since(start)
	return DownloadResult{Duration: duration, Error: nil}
}

func runInstaller(filename string) error {
	ext := filepath.Ext(filename)
	var cmd *exec.Cmd
	switch ext {
	case ".exe":
		cmd = exec.Command("cmd", "/C", filename)
	case ".msi":
		cmd = exec.Command("msiexec", "/i", filename)
	default:
		return fmt.Errorf("unsupported file extension: %s", ext)
	}
	return cmd.Run()
}

func downloadApps(apps []string) {
	progress := make(chan int)
	defer close(progress)

	go func() {
		for p := range progress {
			fmt.Printf("\rDownloading... %d%%", p)
		}
	}()

	sem := make(chan struct{}, maxConcurrentDownloads)
	var wg sync.WaitGroup
	for _, appName := range apps {
		sem <- struct{}{}
		wg.Add(1)
		go func(name string) {
			defer func() {
				<-sem
				wg.Done()
			}()
			url, ok := applications[name]
			if !ok {
				fmt.Printf("Invalid application selection: %s\n", name)
				return
			}

			fmt.Printf("Downloading %s installer...\n", name)
			var filename string
			if name == "Blender" {
				filename = name + ".msi"
			} else {
				filename = name + ".exe"
			}
			result := downloadFile(url, filename, progress)
			if result.Error != nil {

				fmt.Printf("Error downloading installer for %s: %v\n", name, result.Error)
				return
			}
			green := "\033[32m"
			reset := "\033[0m"
			message := fmt.Sprintf("\nInstaller for %s downloaded successfully in %s.", name, result.Duration)

			fmt.Printf("%s%s%s\n", green, message, reset)

			//fmt.Printf("\nInstaller for %s downloaded successfully in %s.\n", name, result.Duration)
		}(appName)
	}

	wg.Wait()
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
				if strings.Contains(installer, "Blender") {
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

func downloadAndRunScript(url, filename string) error {
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
	if err != nil {
		return err
	}

	cmd := exec.Command("cmd", "/C", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
func removeMicrosoftEdge() {
	scriptURL := "https://github.com/ShadowWhisperer/Remove-MS-Edge/blob/main/Remove-Edge.exe"
	scriptFilename := "Remove-Edge.exe"
	err := downloadAndRunScript(scriptURL, scriptFilename)
	if err != nil {
		fmt.Printf("Error downloading file Remove-Edge.exe: %v\n", err)
		return
	}
}

func activateWindows() {
	cmd := exec.Command("powershell", "-Command", "irm https://massgrave.dev/get | iex")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error activating Windows: %v\n", err)
	}
}
