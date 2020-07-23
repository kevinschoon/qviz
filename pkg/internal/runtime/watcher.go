package runtime

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gopkg.in/fsnotify.v1"
)

func Watch(scriptPath string, evalInCh chan evalOpts, errCh chan error) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		errCh <- err
		return
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					errCh <- nil
					return
				}
				// NOTE: vim is weird
				if isWrite(event) || isRemove(event) {
					raw, err := ioutil.ReadFile(scriptPath)
					if err != nil {
						errCh <- err
						return
					}
					evalInCh <- evalOpts{ScriptContents: string(raw)}
					if isRemove(event) {
						// wait for the file to move back again.
						err = wait(scriptPath, 5*time.Second)
						if err != nil {
							errCh <- err
							return
						}
						err = watcher.Add(scriptPath)
						if err != nil {
							errCh <- err
							return
						}
					}
				}
			case err, ok := <-watcher.Errors:
				log.Println("caught error")
				if !ok {
					errCh <- nil
					return
				}
				errCh <- err
			}
		}
	}()
	err = watcher.Add(scriptPath)
	if err != nil {
		errCh <- err
		return
	}
	<-done
}

func isWrite(evt fsnotify.Event) bool {
	return evt.Op&fsnotify.Write == fsnotify.Write
}

func isRemove(evt fsnotify.Event) bool {
	return evt.Op&fsnotify.Remove == fsnotify.Remove
}

func wait(path string, timeout time.Duration) error {
	start := time.Now()
	for {
		if time.Since(start) >= timeout {
			return fmt.Errorf("timeout")
		}
		_, err := os.Stat(path)
		if err == nil {
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}
}
