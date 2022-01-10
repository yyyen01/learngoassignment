package lift

import (
	"log"
	"sort"
	"time"
)

type Lift struct {
	targets      []int //hold the targeted floors details pressed by passengers
	currentLevel int   //current floor level
	movingDir    int   //0 - stop, 1 - up , 2- down
}

func (l Lift) IsEmpty() bool {
	//log.Println("len(l.targets) =", len(l.targets))
	return len(l.targets) == 0
}

func (l *Lift) addTargetFloor(targetFloor ...int) error {
	upLevels := make([]int, 0, cap(l.targets))
	downLevels := make([]int, 0, cap(l.targets))

	if len(targetFloor) == 0 {
		return TargetsFloorCannotBeEmpty

	}
	if l.IsEmpty() {
		if targetFloor[0] > l.currentLevel {
			(*l).movingDir = Up
		} else {
			(*l).movingDir = Down
		}
	} else {
		targetFloor = append(targetFloor, (*l).targets...)

	}

	for _, val := range targetFloor {
		log.Println("--val :", val, " currentLevel :", l.currentLevel)

		if val > l.currentLevel {
			upLevels = append(upLevels, val)
		} else if val < l.currentLevel {
			downLevels = append(downLevels, val)
		} else {
			return CurrentLevelNotAllowed
		}
	}

	//remove duplicate floor
	upLevels = unique(upLevels)
	downLevels = unique(downLevels)

	if l.movingDir == Up {
		sort.Ints(upLevels)
		sort.Sort(sort.Reverse(sort.IntSlice(downLevels)))
		(*l).targets = append(upLevels, downLevels...)
	} else if l.movingDir == Down {
		sort.Ints(downLevels)
		sort.Ints(upLevels)
		(*l).targets = append(downLevels, upLevels...)
	}
	log.Println("L targets :", l.targets)
	return nil
}

func (l Lift) getTotalLevel() (totalLevel int, err error) {
	totalLevel = cap(l.targets)
	return totalLevel, nil
}

func (l Lift) isStopped() bool {
	return l.movingDir == Stop
}

func (l *Lift) setCurrentLevel(target int) error {
	l.currentLevel = target
	return nil
}

func (l Lift) getCurrentLevel() int {
	return l.currentLevel
}

func (l *Lift) moveToNextLevel() (nextLevel int, err error) {
	err = nil
	if !l.IsEmpty() {
		nextLevel = l.targets[0]
		(*l).targets = l.targets[1:]

		if len(l.targets) == 0 {
			(*l).movingDir = Stop
		} else {
			if nextLevel > l.currentLevel {
				(*l).movingDir = Up
			} else {
				(*l).movingDir = Down
			}
		}
		(*l).currentLevel = nextLevel
		time.Sleep(1)
	} else {
		err = ErrNoMoreTargets

	}

	return nextLevel, err
}

func (l *Lift) moveTillStop() (nextLevel []int, err error) {
	nextLevel = make([]int, 0, cap(l.targets))

	for {
		if !l.IsEmpty() {
			nextLevelToGo, err := l.moveToNextLevel()
			if err != nil {
				return nil, err
			}
			nextLevel = append(nextLevel, nextLevelToGo)
		} else {
			break
		}
	}

	return nextLevel, nil
}
