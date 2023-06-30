package utils

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func StartIngest(filePaths []string, collectionName string) error {
	input := collectionName + "|" + strings.Join(filePaths, "|")

	cmd := exec.Command("python3", "/Users/irvyn/work/chat-pdf/src/ingest2.py")
	cmd.Dir = "/Users/irvyn/work/chat-pdf/src" // <-- set the working directory otherwise the process spawns outwith context

	cmd.Stdin = strings.NewReader(input)

	var out bytes.Buffer
	cmd.Stdout = &out

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
		log.Printf("stderr: %q\n", stderr.String())
		return err
	}

	log.Printf("stdout: %q\n", out.String())
	return nil
}
