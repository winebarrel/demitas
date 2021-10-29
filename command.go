package demitas

import (
	"io"
	"os"
	"os/exec"
	"os/signal"
	"sync"
)

func runCommand(cmdWithArgs []string) error {
	cmd := exec.Command(cmdWithArgs[0], cmdWithArgs[1:]...)

	outReader, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	errReader, err := cmd.StderrPipe()

	if err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig)

	go func() {
		for {
			s := <-sig
			_ = cmd.Process.Signal(s)
		}
	}()

	go func() {
		_, _ = io.Copy(os.Stdout, outReader)
		wg.Done()
	}()

	go func() {
		_, _ = io.Copy(os.Stderr, errReader)
		wg.Done()
	}()

	err = cmd.Start()

	if err != nil {
		return err
	}

	err = cmd.Wait()

	if err != nil {
		return err
	}

	wg.Wait()

	return nil
}
