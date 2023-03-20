package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/octacian/timecalc/time"
	"github.com/peterh/liner"
)

var history_fn = filepath.Join(os.TempDir(), "timecalc.tmp")

func main() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	if f, err := os.Open(history_fn); err != nil {
		line.ReadHistory(f)
		f.Close()
	}

	for {
		input, err := line.Prompt(">>> ")
		if err != nil {
			if err == liner.ErrPromptAborted {
				return
			}
			log.Fatalf("error reading line: %s", err)
		}

		if input == "exit" {
			return
		}

		line.AppendHistory(input)

		result, err := time.FromString(input)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}
}
