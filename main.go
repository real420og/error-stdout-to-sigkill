package main

import (
	"bufio"
	"errors"
	"github.com/creack/pty"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

func errorExists(s string) error {
	errorIndex := strings.Index(s, "error")
	lostIndex := strings.Index(s, "lost")
	if errorIndex > 0 || lostIndex > 0 {
		return errors.New("")
	}

	return nil
}
func main() {
	args := strings.Join(os.Args[1:], " ")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		cmd := exec.Command("sh", "-c", args)
		tty, err := pty.Start(cmd)
		if err != nil {
			log.Fatalln(err)
		}
		defer tty.Close()

		go func() {
			scanner := bufio.NewScanner(tty)
			for scanner.Scan() {
				text := scanner.Text()

				log.Println("[" + args + "] " + scanner.Text())

				if err := errorExists(text); err != nil {

					signals <- os.Kill
				}
			}
		}()

		err = cmd.Wait()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	log.Printf("signal: %s", <-signals)
}
