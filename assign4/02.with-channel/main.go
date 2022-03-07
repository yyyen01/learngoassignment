package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const totalRegistrants int = 1000

type Person struct {
	//assuming there is no duplicate name
	name           string
	baseFatRate    float64
	currentFatRate float64
}

type FatRateRankMachine struct {
	scoreboard  []Person
	registrants map[string]Person
}

func fatRateMachineOperator(in <-chan func(*FatRateRankMachine), done <-chan struct{}) {
	machine := &FatRateRankMachine{
		scoreboard:  make([]Person, 0, totalRegistrants),
		registrants: map[string]Person{},
	}

	for {
		select {
		case <-done:
			return
		case f := <-in:
			f(machine)
		}
	}
}

type ChannelFatRateRankMachine chan func(*FatRateRankMachine)

func NewChannelFatRateRankMachine() (ChannelFatRateRankMachine, func()) {
	ch := make(ChannelFatRateRankMachine, 1000)
	done := make(chan struct{})
	go fatRateMachineOperator(ch, done)
	return ch, func() {
		close(done)
	}
}

func (c ChannelFatRateRankMachine) getRegistrants() map[string]Person {
	var registrants map[string]Person
	done := make(chan struct{})
	c <- func(machine *FatRateRankMachine) {
		registrants = machine.registrants
		close(done)
		//fmt.Println("inside1 : reg1=", registrants)
	}
	<-done
	//fmt.Println("inside : reg=", registrants)
	return registrants
}

func (c ChannelFatRateRankMachine) UpdateFatrate(p Person) (int, bool) {
	var val int
	var ok1 bool
	done := make(chan struct{})
	c <- func(machine *FatRateRankMachine) {
		defer close(done)
		_, ok := machine.registrants[p.name]
		//fmt.Println("[update]name=", p.name, ok)
		//if it does not exist in scoreboard, then return as false
		if ok == true {
			for i, pl := range machine.scoreboard {
				if p.name == pl.name {
					machine.scoreboard = remove(machine.scoreboard, i)
					break
				}
			}

			for i2, pl2 := range machine.scoreboard {
				if p.currentFatRate <= pl2.currentFatRate {
					machine.scoreboard = insert(machine.scoreboard, p, i2)
					val = i2 + 1
					ok1 = true
					//fmt.Println("[update]update,removee, name=", p.name, val, " ", ok1)
					return
				}
			}

			//apend to the end
			machine.scoreboard = append(machine.scoreboard, p)
			val = len(machine.scoreboard)
			ok1 = true
			return

		} else {
			val = 0
			ok1 = false
			return
		}
	}
	<-done
	return val, ok1

}

func (c ChannelFatRateRankMachine) getRank(p Person) (int, bool) {
	var val int
	var ok bool
	done := make(chan struct{})

	c <- func(machine *FatRateRankMachine) {
		defer close(done)
		for i, person := range machine.scoreboard {
			if person.name == p.name {
				val = i + 1
				ok = true
				return
			}
		}
		val = 0
		ok = false
	}

	<-done

	return val, ok
}

func (c ChannelFatRateRankMachine) register(p Person) {

	c <- func(machine *FatRateRankMachine) {
		_, ok := machine.registrants[p.name]
		//if it does not exist in scoreboard, then register as new user
		if ok == false {
			for i, pl := range machine.scoreboard {
				if p.currentFatRate <= pl.currentFatRate {
					machine.scoreboard = insert(machine.scoreboard, p, i)
					machine.registrants[p.name] = p
					return
				}
			}
			machine.scoreboard = append(machine.scoreboard, p)
			machine.registrants[p.name] = p
		}
	}

}

func insert(a []Person, p Person, i int) []Person {

	return append(a[:i], append([]Person{p}, a[i:]...)...)
}

func (c ChannelFatRateRankMachine) PrintScoreboard() {
	done := make(chan struct{})
	c <- func(machine *FatRateRankMachine) {
		for i, person := range machine.scoreboard {
			fmt.Println(i+1, ": ", person.name, " ", person.currentFatRate)
		}
		close(done)
	}
	<-done

}

func remove(a []Person, i int) []Person {
	return append(a[:i], a[i+1:]...)
}

func randFloat(min, max float64) float64 {
	if min < 0 {
		min = 0.0
	}
	return min + rand.Float64()*(max-min)
}

func main() {

	//define a channel to catch Control C in order to let all go routine before terminating
	signaldone := make(chan os.Signal, 1)
	signal.Notify(signaldone, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	manager, finish := NewChannelFatRateRankMachine()
	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup
	wg.Add(totalRegistrants)

	//register for Registrants
	for i := 0; i < totalRegistrants; i++ {
		go func(i int) {
			defer wg.Done()
			name := fmt.Sprintf("Person%d", i)
			base := randFloat(0, 0.4)
			manager.register(Person{name: name, baseFatRate: base, currentFatRate: base})
		}(i)
	}
	wg.Wait()
	manager.PrintScoreboard()
	//registration done

	//simulate a timeout situation. In real case, it should be some terminating condition
	reqCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		fmt.Println("[Main]: canceling context")
		cancel()
	}()

	registrants := manager.getRegistrants()
	//fmt.Println("registrants: =", registrants)

	for _, person := range registrants {
		wg.Add(2)
		go func(person Person) {
			defer wg.Done()

			//Loop:
			for {
				select {
				case <-reqCtx.Done():
					fmt.Println("[", person.name, "]", "[Update FatRate]", "Timeout! Exiting")
					//wait till all current go routine to finish
					return
				default:
					//update
					minFatRate := person.baseFatRate - 0.2
					person.currentFatRate = randFloat(minFatRate, person.baseFatRate+0.2)
					rank, ok := manager.UpdateFatrate(person)
					if ok {
						fmt.Println("[", person.name, "]", "[Update FatRate]", ",", rank, ",", person.currentFatRate, " ", time.Now())
					} else {
						fmt.Println("[", person.name, "]", "[Update FatRate]", "  not found!")
					}

				}

			}
		}(person)

		go func(person Person) {
			defer wg.Done()

			//Loop:
			for {
				select {
				case <-reqCtx.Done():
					fmt.Println("[", person.name, "]", "[Query rank]", "Timeout! Exiting")
					return
				default:
					rank, ok := manager.getRank(person)
					if ok {
						fmt.Println("[", person.name, "]", "[Query rank]", rank, ",", person.currentFatRate, " ", time.Now())
					} else {
						fmt.Println("[", person.name, "]", "[Query rank]", "  not found!")
					}

				}

			}
		}(person)

	}

	select {
	case <-signaldone:
		fmt.Println("[Main] syscall.SIGINT or syscall.SIGTERM received. Terminating")
		cancel()
		wg.Wait()

	case <-reqCtx.Done():
		fmt.Println("[Main] Timeout Reached. Terminating")
		wg.Wait()

	}

	fmt.Println("---------------------------------------")
	fmt.Println("Final Scoreboard!")
	fmt.Println("---------------------------------------")
	manager.PrintScoreboard()
	//close the channel for scoreboard
	finish()

	fmt.Println("Program ended successfully!")
}
