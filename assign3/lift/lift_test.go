package lift

import (
	"log"
	"reflect"
	"testing"
)

func Test_LifeOperation1(t *testing.T) {
	lift := Lift{
		targets:      make([]int, 0, 5),
		currentLevel: -1,
		movingDir:    0,
	}
	totalLevel, err := lift.getTotalLevel()
	if err != nil {
		log.Printf("Expect no error but hit error %v", err)
		t.Fail()
	}

	if totalLevel != 5 {
		log.Println("Expect total number of Floors is 5, but returning ", totalLevel)
		t.Fail()
	}

	isStopped := lift.isStopped()

	if !isStopped {
		log.Println("Expect lift not moving but it is moving")
		t.Fail()
	}

	nextLevel, err := lift.moveToNextLevel()
	if err == nil || err != ErrNoMoreTargets {
		log.Println("Expect to throw ErrNoMoreTargets but the error is ", ErrNoMoreTargets)
		t.Fail()
	}
	log.Println("nextlevel is ", nextLevel)

}

func Test_LifeOperation2(t *testing.T) {
	lift := Lift{
		targets:      make([]int, 0, 5),
		currentLevel: -1,
		movingDir:    0,
	}
	totalLevel, err := lift.getTotalLevel()
	if err != nil {
		log.Printf("Expect no error but hit error %v", err)
		t.Fail()
	}

	if totalLevel != 5 {
		log.Println("Expect total number of Floors is 5, but returning ", totalLevel)
		t.Fail()
	}

	err = lift.setCurrentLevel(1)
	if err != nil {
		log.Printf("Setting Current Level: Expect no error but hit error %v", err)
		t.Fail()
	}

	currentLevel := lift.getCurrentLevel()

	if currentLevel != 1 {
		log.Println("Expect current level is 1 but it is at ", currentLevel)
		t.Fail()
	}

	err = lift.addTargetFloor(3)
	if err != nil {
		log.Printf("Add target Floor : Expect no error but hit error %v", err)
		t.Fail()
	}

	nextLevelToGo, err := lift.moveToNextLevel()

	if nextLevelToGo != 3 {
		log.Println("Expect next level is 3 but it is going to  ", currentLevel)
		t.Fail()
	}

	isStopped := lift.isStopped()

	if !isStopped {
		log.Println("Expect lift to stop but it is moving")
		t.Fail()
	}

	currentLevel = lift.getCurrentLevel()

	if currentLevel != 3 {
		log.Println("Expect current level is 3 but it is at ", currentLevel)
		t.Fail()
	}

}

func Test_LifeOperation3(t *testing.T) {
	lift := Lift{
		targets:      make([]int, 0, 5),
		currentLevel: -1,
		movingDir:    0,
	}
	totalLevel, err := lift.getTotalLevel()
	if err != nil {
		log.Printf("Expect no error but hit error %v", err)
		t.Fail()
	}

	if totalLevel != 5 {
		log.Println("Expect total number of Floors is 5, but returning ", totalLevel)
		t.Fail()
	}

	err = lift.setCurrentLevel(3)
	if err != nil {
		log.Printf("Setting Current Level: Expect no error but hit error %v", err)
		t.Fail()
	}

	currentLevel := lift.getCurrentLevel()

	if currentLevel != 3 {
		log.Println("Expect current level is 3 but it is at ", currentLevel)
		t.Fail()
	}

	err = lift.addTargetFloor(4, 2)
	if err != nil {
		log.Printf("Add target Floor : Expect no error but hit error %v", err)
		t.Fail()
	}

	nextLevelToGo, err := lift.moveToNextLevel()

	if nextLevelToGo != 4 {
		log.Println("Expect next level is 4 but it is going to  ", currentLevel)
		t.Fail()
	}

	nextLevelToGo, err = lift.moveToNextLevel()

	if nextLevelToGo != 2 {
		log.Println("Expect next level is 2 but it is going to  ", currentLevel)
		t.Fail()
	}

	isStopped := lift.isStopped()
	if err != nil {
		log.Printf("Expect no error but hit error %v", err)
		t.Fail()
	}

	if !isStopped {
		log.Println("Expect lift not moving but it is moving")
		t.Fail()
	}

	currentLevel = lift.getCurrentLevel()

	if currentLevel != 2 {
		log.Println("Expect current level is 3 but it is at ", currentLevel)
		t.Fail()
	}

}

func Test_LifeOperation4(t *testing.T) {
	lift := Lift{
		targets:      make([]int, 0, 5),
		currentLevel: -1,
		movingDir:    0,
	}
	totalLevel, err := lift.getTotalLevel()
	if err != nil {
		log.Printf("Expect no error but hit error %v", err)
		t.Fail()
	}

	if totalLevel != 5 {
		log.Println("Expect total number of Floors is 5, but returning ", totalLevel)
		t.Fail()
	}

	err = lift.setCurrentLevel(3)
	if err != nil {
		log.Printf("Setting Current Level: Expect no error but hit error %v", err)
		t.Fail()
	}

	currentLevel := lift.getCurrentLevel()

	if currentLevel != 3 {
		log.Println("Expect current level is 3 but it is at ", currentLevel)
		t.Fail()
	}

	err = lift.addTargetFloor(4, 2, 5)
	if err != nil {
		log.Printf("Add target Floor : Expect no error but hit error %v", err)
		t.Fail()
	}

	nextLevelToGo, err := lift.moveToNextLevel()

	if err != nil {
		log.Printf("Expect no error but hit error %v", err)
		t.Fail()
	}

	if nextLevelToGo != 4 {
		log.Println("Expect next level is 4 but it is going to  ", currentLevel)
		t.Fail()
	}

	nextLevelToGo, err = lift.moveToNextLevel()

	if nextLevelToGo != 5 {
		log.Println("Expect next level is 5 but it is going to  ", currentLevel)
		t.Fail()
	}

	nextLevelToGo, err = lift.moveToNextLevel()

	if nextLevelToGo != 2 {
		log.Println("Expect next level is 2 but it is going to  ", currentLevel)
		t.Fail()
	}

	isStopped := lift.isStopped()

	if !isStopped {
		log.Println("Expect lift not moving but it is moving")
		t.Fail()
	}

	currentLevel = lift.getCurrentLevel()

	if currentLevel != 2 {
		log.Println("Expect current level is 3 but it is at ", currentLevel)
		t.Fail()
	}

}

func Test_LifeOperation5(t *testing.T) {
	lift := Lift{
		targets:      make([]int, 0, 5),
		currentLevel: -1,
		movingDir:    0,
	}
	totalLevel, err := lift.getTotalLevel()
	if err != nil {
		log.Printf("Expect no error but hit error %v", err)
		t.Fail()
	}

	if totalLevel != 5 {
		log.Println("Expect total number of Floors is 5, but returning ", totalLevel)
		t.Fail()
	}

	err = lift.setCurrentLevel(3)
	if err != nil {
		log.Printf("Setting Current Level: Expect no error but hit error %v", err)
		t.Fail()
	}

	currentLevel := lift.getCurrentLevel()

	if currentLevel != 3 {
		log.Println("Expect current level is 3 but it is at ", currentLevel)
		t.Fail()
	}

	err = lift.addTargetFloor(4, 2, 5)
	if err != nil {
		log.Printf("Add target Floor : Expect no error but hit error %v", err)
		t.Fail()
	}

	nextLevelToGo, err := lift.moveToNextLevel()

	if err != nil {
		log.Printf("Expect no error but hit error %v", err)
		t.Fail()
	}

	if nextLevelToGo != 4 {
		log.Println("Expect next level is 4 but it is going to  ", currentLevel)
		t.Fail()
	}

	nextLevelToGo, err = lift.moveToNextLevel()

	if err != nil {
		log.Printf("Expect no error but hit error %v", err)
		t.Fail()
	}

	if nextLevelToGo != 5 {
		log.Println("Expect next level is 5 but it is going to  ", currentLevel)
		t.Fail()
	}

	err = lift.addTargetFloor(3, 1, 4, 1, 3)

	if err != nil {
		log.Printf("Expect no error but hit error %v", err)
		t.Fatal()
	}

	nextLevels, err := lift.moveTillStop()

	if err != nil {
		log.Printf("Expect no error but hit error %v", err)
		t.Fail()
	}

	expectedLevelSeq := []int{4, 3, 2, 1}
	if !reflect.DeepEqual(expectedLevelSeq, nextLevels) {
		log.Println("Expected sequence is ", expectedLevelSeq, "but shown ", nextLevels)
		t.Fail()
	}

	isStopped := lift.isStopped()

	if !isStopped {
		log.Println("Expect lift not moving but it is moving")
		t.Fail()
	}

	currentLevel = lift.getCurrentLevel()

	if currentLevel != 1 {
		log.Println("Expect current level is 1 but it is at ", currentLevel)
		t.Fail()
	}

}

func Test_LifeOperation6(t *testing.T) {
	lift := Lift{
		targets:      make([]int, 0, 5),
		currentLevel: -1,
		movingDir:    0,
	}
	totalLevel, err := lift.getTotalLevel()
	if err != nil {
		log.Printf("Expect no error but hit error %v", err)
		t.Fail()
	}

	if totalLevel != 5 {
		log.Println("Expect total number of Floors is 5, but returning ", totalLevel)
		t.Fail()
	}

	err = lift.setCurrentLevel(3)
	if err != nil {
		log.Printf("Setting Current Level: Expect no error but hit error %v", err)
		t.Fail()
	}

	currentLevel := lift.getCurrentLevel()

	if currentLevel != 3 {
		log.Println("Expect current level is 3 but it is at ", currentLevel)
		t.Fail()
	}

	err = lift.addTargetFloor(2, 5, 4)
	if err != nil {
		log.Printf("Add target Floor : Expect no error but hit error %v", err)
		t.Fail()
	}

	nextLevelToGo, err := lift.moveToNextLevel()

	if err != nil {
		log.Printf("Expect no error but hit error %v", err)
		t.Fail()
	}

	if nextLevelToGo != 2 {
		log.Println("Expect next level is 2 but it is going to  ", nextLevelToGo)
		t.Fail()
	}

	err = lift.addTargetFloor(1, 3)

	if err != nil {
		log.Printf("Expect no error but hit error %v", err)
		t.Fatal()
	}

	nextLevels, err := lift.moveTillStop()

	if err != nil {
		log.Printf("Expect no error but hit error %v", err)
		t.Fail()
	}

	expectedLevelSeq := []int{1, 3, 4, 5}
	if !reflect.DeepEqual(expectedLevelSeq, nextLevels) {
		log.Println("Expected sequence is ", expectedLevelSeq, "but shown ", nextLevels)
		t.Fail()
	}

	isStopped := lift.isStopped()

	if !isStopped {
		log.Println("Expect lift not moving but it is moving")
		t.Fail()
	}

	currentLevel = lift.getCurrentLevel()

	if currentLevel != 5 {
		log.Println("Expect current level is 5 but it is at ", currentLevel)
		t.Fail()
	}

}
