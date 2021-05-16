package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
)

type Person struct {
	Name      string
	Surname   string
	Birthdate string
	Address   string
	Phone     string
}

type Rules struct {
	LogEveryMessage    bool `json:",omitempty"`
	SleepBeforeSending int  `json:",omitempty"`
}

type Application struct {
	Name    string
	Workers int
	Dir     string `json:",omitempty"`
	Rules   []Rules
}

type Config struct {
	Debug        bool          `json:"debug"`
	Applications []Application `json:"applications"`
}

func setupCloseHandler(workersCount int, stopChan chan bool) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("\nCtrl-C pressed in terminal. Aborting...")
		for i := 0; i < workersCount; i++ {
			stopChan <- true
		}
	}()
}

func readConfig(conf *Config) {
	f, err := os.Open("./config.json")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &conf)

	if err != nil {
		log.Fatal(err)
	}
}

func startFileWatcher(dir string, bus chan string, stop chan bool) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("modified file:", event.Name, event.Op)
					bus <- event.Name
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			case <-stop:
				log.Println("File watcher quitting")
				return
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func producer(wg *sync.WaitGroup, id int, sleepTime int, filesChan chan string, dataChan chan []byte, stop chan bool) {
	fmt.Println("Gonna wait for file to read")

	defer wg.Done()

	select {
	case changedFile := <-filesChan:
		log.Printf("Producer #%d got file\n", id)
		f, err := os.Open(changedFile)
		if err != nil {
			log.Fatal(err)
		}

		data, err := io.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}

		if err = f.Close(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(data)

		if sleepTime > 0 {
			time.Sleep(time.Duration(int64(sleepTime)) * time.Second)
		}
		dataChan <- data
	case <-stop:
		log.Printf("Producer #%d quitting\n", id)
		return
	}
}

func consumer(wg *sync.WaitGroup, id int, logEveryMessage bool, debug bool, data chan []byte, stop chan bool) {
	defer wg.Done()
	logMessage := true

	select {
	case rawUser := <-data:
		person := Person{}

		log.Printf("Consumer #%d got data\n", id)
		fmt.Println(rawUser)

		if err := json.Unmarshal(rawUser, &person); err != nil {
			log.Fatal(err)
		}

		if debug && (logMessage || logEveryMessage) {
			log.Println(person)
		}

		logMessage = !logMessage
	case <-stop:
		log.Printf("Consumer #%d quitting\n", id)
		return
	}
}

func main() {
	config := Config{}
	readConfig(&config)

	workersCount := 1

	for _, app := range config.Applications {
		workersCount += app.Workers
	}

	wg := sync.WaitGroup{}

	changedFilesChan := make(chan string)
	dataChan := make(chan []byte)
	stopChan := make(chan bool, workersCount)

	setupCloseHandler(workersCount, stopChan)

	for _, app := range config.Applications {
		fmt.Println(app.Name, app.Workers)
		if app.Name == "Producer" {
			fmt.Println("Start fsnotify for", app.Dir)
			go startFileWatcher(app.Dir, changedFilesChan, stopChan)

			wg.Add(app.Workers)
			for i := 0; i < app.Workers; i++ {
				fmt.Println("Start producer worker #", i)
				go func(id int) {
					producer(&wg, id, app.Rules[0].SleepBeforeSending, changedFilesChan, dataChan, stopChan)
				}(i)
			}
		} else if app.Name == "Consumer" {
			for i := 0; i < app.Workers; i++ {
				wg.Add(1)
				fmt.Println("Start consumer worker #", i)
				go func(id int) {
					consumer(&wg, id, app.Rules[0].LogEveryMessage, config.Debug, dataChan, stopChan)
				}(i)
			}
		}
	}

	wg.Wait()
}
