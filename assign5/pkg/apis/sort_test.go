package apis

import (
	"reflect"
	"testing"
)

func TestBubbleSorter(t *testing.T) {
	machine := NewMutexFatRateRankMachine(1, nil, BubbleSort{})
	p1 := Person{Name: "p1", CurrentFatRate: 2.1}
	p2 := Person{Name: "p2", CurrentFatRate: 1.1}
	p3 := Person{Name: "p3", CurrentFatRate: 3.1}
	p4 := Person{Name: "p4", CurrentFatRate: 1.5}
	p5 := Person{Name: "p5", CurrentFatRate: 7.1}
	p6 := Person{Name: "p6", CurrentFatRate: 2.2}
	p := []*Person{&p1, &p2, &p3, &p4, &p5, &p6}
	machine.Scoreboard = p
	machine.Registrants = map[string]*Person{
		"p1": &p1,
		"p2": &p2,
		"p3": &p3,
		"p4": &p4,
		"p5": &p5,
		"p6": &p6,
	}

	BubbleSort{}.Sort(machine)

	f := []*Person{
		&Person{Name: "p2", CurrentFatRate: 1.1, rank: 0},
		&Person{Name: "p4", CurrentFatRate: 1.5, rank: 1},
		&Person{Name: "p1", CurrentFatRate: 2.1, rank: 2},
		&Person{Name: "p6", CurrentFatRate: 2.2, rank: 3},
		&Person{Name: "p3", CurrentFatRate: 3.1, rank: 4},
		&Person{Name: "p5", CurrentFatRate: 7.1, rank: 5},
	}
	//machine.PrintScoreboard()
	//fmt.Println(f)

	if !reflect.DeepEqual(machine.Scoreboard, f) {
		t.Fatalf("Fail. Sequence in array not expected.")
	}

	if machine.Registrants["p1"].rank != 2 || machine.Registrants["p2"].rank != 0 || machine.Registrants["p3"].rank != 4 || machine.Registrants["p4"].rank != 1 || machine.Registrants["p5"].rank != 5 || machine.Registrants["p6"].rank != 3 {
		t.Fatalf("Fail. Sequence in map not expected.")
	}
}

func TestQuickSort(t *testing.T) {
	machine := NewMutexFatRateRankMachine(1, nil, QuickSort{})
	p1 := Person{Name: "p1", CurrentFatRate: 2.1}
	p2 := Person{Name: "p2", CurrentFatRate: 1.1}
	p3 := Person{Name: "p3", CurrentFatRate: 3.1}
	p4 := Person{Name: "p4", CurrentFatRate: 1.5}
	p5 := Person{Name: "p5", CurrentFatRate: 7.1}
	p6 := Person{Name: "p6", CurrentFatRate: 2.2}
	p := []*Person{&p1, &p2, &p3, &p4, &p5, &p6}
	machine.Scoreboard = p
	machine.Registrants = map[string]*Person{
		"p1": &p1,
		"p2": &p2,
		"p3": &p3,
		"p4": &p4,
		"p5": &p5,
		"p6": &p6,
	}

	QuickSort{}.Sort(machine)

	f := []*Person{
		&Person{Name: "p2", CurrentFatRate: 1.1, rank: 0},
		&Person{Name: "p4", CurrentFatRate: 1.5, rank: 1},
		&Person{Name: "p1", CurrentFatRate: 2.1, rank: 2},
		&Person{Name: "p6", CurrentFatRate: 2.2, rank: 3},
		&Person{Name: "p3", CurrentFatRate: 3.1, rank: 4},
		&Person{Name: "p5", CurrentFatRate: 7.1, rank: 5},
	}
	machine.PrintScoreboard()
	//fmt.Println(f)

	if !reflect.DeepEqual(machine.Scoreboard, f) {
		t.Fatalf("Fail. Sequence in array not expected.")
	}

	if machine.Registrants["p1"].rank != 2 || machine.Registrants["p2"].rank != 0 || machine.Registrants["p3"].rank != 4 || machine.Registrants["p4"].rank != 1 || machine.Registrants["p5"].rank != 5 || machine.Registrants["p6"].rank != 3 {
		t.Fatalf("Fail. Sequence in map not expected.")
	}
}
