package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
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
			if name == "GoLang" || name == "Blender" || name == "EpicGames" {
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
		}(appName)
	}

	wg.Wait()
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

func ChristTitusTechsWinutil() {
	cmd := exec.Command("powershell", "-Command", "irm https://christitus.com/win | iex")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error opening app: %v\n", err)
	}
}
func deleteTemporaryFiles() {
	userTempDir := os.TempDir()
	fmt.Println("Deleting temporary files from:", userTempDir)

	deletedCount := 0
	err := filepath.Walk(userTempDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return nil
		}
		if !info.IsDir() {
			fmt.Println("Deleting file:", path)
			err := os.Remove(path)
			if err != nil {
				fmt.Printf("Error deleting file %s: %v\n", path, err)
			} else {
				deletedCount++
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error deleting temporary files: %v\n", err)
	} else {
		if deletedCount > 0 {
			green := "\033[32m"
			reset := "\033[0m"
			fmt.Printf("%sTemporary files deleted successfully.%s\n", green, reset)

			if deletedCount < len(os.TempDir()) {
				fmt.Println("Some files could not be deleted because they are being used by other applications.")

				fmt.Print("Do you want to see the files that could not be deleted? (yes/no): ")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				choice := strings.ToLower(scanner.Text())

				if choice == "yes" {
					fmt.Println("Files that could not be deleted:")
					err := filepath.Walk(userTempDir, func(path string, info os.FileInfo, err error) error {
						if err != nil {
							fmt.Printf("Error accessing path %s: %v\n", path, err)
							return nil
						}
						if !info.IsDir() && !os.IsNotExist(err) {
							fmt.Println(path)
						}
						return nil
					})
					if err != nil {
						fmt.Printf("Error accessing files: %v\n", err)
					}
				}
			}
		} else {
			red := "\033[31m"
			reset := "\033[0m"
			fmt.Printf("%sNo temporary files were deleted.%s\n", red, reset)
		}
	}
}

func installWindowsSandbox() error {

	cmd := exec.Command("powershell", "-Command", "Enable-WindowsOptionalFeature -FeatureName 'Containers-DisposableClientVM' -All -Online")

	cmd.Run()

	fmt.Println("Windows Sandbox installed successfully.")
	return nil
}

func uninstallWindowsSandbox() error {
	cmd := exec.Command("powershell", "-Command", "Disable-WindowsOptionalFeature -FeatureName 'Containers-DisposableClientVM' -Online")

	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("Windows Sandbox uninstalled successfully.")
	return nil
}

func openDiskCleanup() error {
	cmd := exec.Command("cleanmgr")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
