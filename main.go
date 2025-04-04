package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

var (
	version string = "dev"
)

func main() {
	configFile := flag.String("config", "", "The buf.yaml file or data to use for configuration.")
	displayVersion := flag.Bool("version", false, "Display the version of the tool.")
	flag.Parse()

	if *displayVersion {
		fmt.Println("Version:", version)
		var buffer bytes.Buffer

		cmd := exec.Command("buf", "--version")
		cmd.Stdout = &buffer
		err := cmd.Run()
		if err != nil {
			log.Error(err)
		}

		bufVersion := buffer.String()
		bufVersion = strings.TrimSpace(bufVersion)
		if bufVersion == "" {
			bufVersion = "unknown"
		}

		buffer.Reset()

		cmd = exec.Command("which", "buf")
		cmd.Stdout = &buffer
		err = cmd.Run()
		if err != nil {
			log.Debug(err)
		}

		bufPath := buffer.String()
		bufPath = strings.TrimSpace(bufPath)

		locatedAt := ""
		if bufPath != "" {
			locatedAt = fmt.Sprintf("located at %s", bufPath)
		}

		fmt.Println("Your current `buf` version:", bufVersion, locatedAt)
		return
	}

	// generate a temporary file pivot
	tmpFile := fmt.Sprintf("buf-original-%d.proto", time.Now().UnixNano())
	tmpDir, err := os.MkdirTemp("", "buf-")
	if err != nil {
		log.Fatal(err)
	}

	tmpFilePath := fmt.Sprintf("%s/%s", tmpDir, tmpFile)
	tmpFileHandle, err := os.Create(tmpFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// ensure the temporary file and directory are removed on exit
	defer func() {
		os.RemoveAll(tmpDir)
	}()

	// read the input from STDIN as provided by a text editor
	file, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	tmpFileHandle.Write(file)
	tmpFileHandle.Close()

	// here we need to yield any parameters, for now we rely on the
	// php-cs-fixer CLI to handle any validations and errors
	args := []string{"format", tmpFilePath}
	if *configFile != "" {
		args = append(args, fmt.Sprintf("--config=%s", *configFile))
	}

	cmd := exec.Command("buf", args...)

	// the formatted file will be written to the stdout
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// make it rain!
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
