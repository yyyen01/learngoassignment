package apis

import (
	"log"
	"sync"
)

type MutexFatRateRankMachine struct {
	Scoreboard  []*Person
	Registrants map[string]*Person
	L           sync.RWMutex
	writer      Writer
	sorter      Sorter
}

func NewMutexFatRateRankMachine(maxRegistrant int, writer Writer, sorter Sorter) *MutexFatRateRankMachine {
	return &MutexFatRateRankMachine{
		Scoreboard:  make([]*Person, 0, maxRegistrant),
		Registrants: map[string]*Person{},
		writer:      writer,
		sorter:      sorter,
	}
}

func (machine *MutexFatRateRankMachine) register(p Person) {
	machine.L.Lock()
	defer machine.L.Unlock()
	_, exist := machine.Registrants[p.Name]

	if !exist {
		machine.Scoreboard = append(machine.Scoreboard, &p)
		machine.Registrants[p.Name] = &p
		machine.sorter.Sort(machine)
		if machine.writer != nil {
			machine.writer.Write(p)
		}
	}

}

func (machine *MutexFatRateRankMachine) PrintScoreboardWithLock() {
	machine.L.RLock()
	defer machine.L.RUnlock()
	machine.PrintScoreboard()
}

func (machine *MutexFatRateRankMachine) PrintScoreboard() {
	for i, person := range machine.Scoreboard {
		log.Printf("%d : %+v\n", i+1, person)
	}
}

func (machine *MutexFatRateRankMachine) getRank(name string) (int, bool) {
	machine.L.RLock()
	defer machine.L.RUnlock()

	person, exist := machine.Registrants[name]
	if exist {
		return person.rank + 1, true
	}
	return 0, false
}

func (machine *MutexFatRateRankMachine) updateRank() {
	machine.L.Lock()
	defer machine.L.Unlock()

	machine.sorter.Sort(machine)

}
