package app

import (
	"assignment/learngoassignment/assign5/pkg/apis"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type FatRateClient struct {
	outputFilePath   string
	timeout          time.Duration
	totalRegistrants int
	sorter           apis.Sorter
	writer           apis.Writer
}

//const totalRegistrants int = 100

func (client FatRateClient) execute() {

	//define a channel to catch Control C in order to let all go routine before terminating
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	file := client.writer.CreateOutputFile(client.outputFilePath)
	defer file.Close()

	manager := apis.NewMutexFatRateRankMachine(client.totalRegistrants, client.writer, client.sorter)
	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup
	wg.Add(client.totalRegistrants)

	//register for Registrants
	for i := 0; i < client.totalRegistrants; i++ {
		go func(i int) {
			defer wg.Done()
			person := apis.NewPerson(fmt.Sprintf("Person%d", i))
			person.Register(manager)
		}(i)
	}
	wg.Wait()
	manager.PrintScoreboardWithLock()
	//registration done

	//simulate a timeout situation. In real case, it should be some terminating condition
	reqCtx, cancel := context.WithTimeout(context.Background(), client.timeout)
	defer func() {
		fmt.Println("[Main]: canceling context")
		cancel()
	}()

	for _, person := range manager.Registrants {
		wg.Add(2)
		go func(person *apis.Person) {
			defer wg.Done()
			person.UpdateTallWeightAge(reqCtx, manager)
		}(person)

		go func(person *apis.Person) {
			defer wg.Done()
			person.GetRank(reqCtx, manager)
		}(person)

	}

	select {
	case <-done:
		log.Println("[Main] syscall.SIGINT or syscall.SIGTERM received. Terminating")
		cancel()
		wg.Wait()
	case <-reqCtx.Done():
		log.Println("[Main] Timeout Reached. Terminating")
		wg.Wait()
	}

	log.Println("---------------------------------------")
	log.Println("Final Scoreboard!")
	log.Println("---------------------------------------")
	manager.PrintScoreboardWithLock()

	log.Println("Program ended successfully!")
}
