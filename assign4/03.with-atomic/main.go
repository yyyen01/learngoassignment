package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

var v atomic.Value

type Person struct {
	//assuming there is no duplicate name
	name           string
	baseFatRate    float64
	currentFatRate float64
}

type MutexFatRateRankMachine struct {
	scoreboard  []Person
	registrants map[string]Person
	//l           sync.RWMutex
}

func NewMutexFatRateRankMachine(maxRegistrant int) *MutexFatRateRankMachine {
	machine := &MutexFatRateRankMachine{
		scoreboard:  make([]Person, 0, maxRegistrant),
		registrants: map[string]Person{},
	}
	v.Store(machine)
	return machine
}

func (machine *MutexFatRateRankMachine) register(p Person) {
	//machine.l.Lock()
	//defer machine.l.Unlock()
	m1 := v.Load().(*MutexFatRateRankMachine)
	defer v.Store(m1)

	_, ok := m1.registrants[p.name]
	//if it does not exist in scoreboard, then register as new user
	if ok == false {
		for i, pl := range m1.scoreboard {
			if p.currentFatRate <= pl.currentFatRate {
				m1.scoreboard = insert(m1.scoreboard, p, i)
				m1.registrants[p.name] = p
				return
			}
		}
		m1.scoreboard = append(m1.scoreboard, p)
		m1.registrants[p.name] = p
	}
}

func insert(a []Person, p Person, i int) []Person {

	return append(a[:i], append([]Person{p}, a[i:]...)...)
}

//func (machine *MutexFatRateRankMachine) PrintScoreboardWithLock() {
//	machine.l.RLock()
//	defer machine.l.RUnlock()
//	machine.PrintScoreboard()
//}

func (machine *MutexFatRateRankMachine) PrintScoreboard() {
	m1 := v.Load().(*MutexFatRateRankMachine)
	for i, person := range m1.scoreboard {
		fmt.Println(i+1, ": ", person.name, " ", person.currentFatRate)
	}
}

func (machine *MutexFatRateRankMachine) getRank(name string) (int, bool) {
	//machine.l.RLock()
	//defer machine.l.RUnlock()
	m1 := v.Load().(*MutexFatRateRankMachine)

	for i, person := range m1.scoreboard {
		if person.name == name {
			return i + 1, true
		}
	}
	return 0, false
}

func (machine *MutexFatRateRankMachine) updateFatRate(p Person) (int, bool) {
	//machine.l.Lock()
	//defer machine.l.Unlock()
	defer v.Store(machine)
	_, ok := machine.registrants[p.name]
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
				return i2 + 1, true
			}
		}

		//apend to the end
		machine.scoreboard = append(machine.scoreboard, p)
		return len(machine.scoreboard), true

	} else {
		return 0, false
	}
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
	totalRegistrants := 1000

	//define a channel to catch Control C in order to let all go routine before terminating
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	manager := NewMutexFatRateRankMachine(totalRegistrants)
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
	reqCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer func() {
		fmt.Println("[Main]: canceling context")
		cancel()
	}()

	//finish := false

	for _, person := range manager.registrants {
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
					rank, ok := manager.updateFatRate(person)
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
					rank, ok := manager.getRank(person.name)
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
	case <-done:
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

	fmt.Println("Program ended successfully!")
}
