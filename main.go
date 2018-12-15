package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bobrnor/highloadcup2018/pkg/account"

	"github.com/bobrnor/highloadcup2018/pkg/filtering"

	"github.com/bobrnor/highloadcup2018/pkg/http/rest"

	"github.com/bobrnor/highloadcup2018/pkg/storage/inmemory"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

var accounts account.Accounts

func main() {
	log.Println("Hello, highload cup!")

	flag.Parse()

	setupCPUProfiling()
	setupMemProfiling()

	parseData()
	run()
}

func setupCPUProfiling() {
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-signalChan
		if *cpuprofile != "" {
			pprof.StopCPUProfile()
		}
		os.Exit(0)
	}()
}

func setupMemProfiling() {
	go func() {
		i := 0
		for range time.Tick(10 * time.Second) {
			if *memprofile != "" {
				filename := *memprofile + strconv.Itoa(i)
				f, err := os.Create(filename)
				if err != nil {
					log.Fatal("could not create memory profile: ", err)
				}
				runtime.GC() // get up-to-date statistics
				if err := pprof.WriteHeapProfile(f); err != nil {
					log.Fatal("could not write memory profile: ", err)
				}
				f.Close()
				i++
			}
		}
	}()
}

func parseData() {
	dataPath, ok := os.LookupEnv("DATA_PATH")
	if !ok {
		log.Println("no data path")
		os.Exit(1)
	}

	err := filepath.Walk(dataPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if !strings.HasPrefix(info.Name(), "accounts") {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		if err := json.NewDecoder(f).Decode(&accounts); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() {
	repository := inmemory.New(accounts)
	filteringService := filtering.New(repository)
	handler := rest.New(filteringService)

	log.Println("server is on tap: http://localhost:80")
	log.Fatal(http.ListenAndServe(":80", handler))
}
