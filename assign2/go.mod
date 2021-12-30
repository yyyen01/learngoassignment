module assignmnet/learngoassignment/assign2

go 1.17

replace (
	assignmnet/learngoassignment/assign2/util => ./util
	github.com/armstrongli/go-bmi => ./go-bmi
)

require github.com/armstrongli/go-bmi v0.0.1

require assignmnet/learngoassignment/assign2/util v0.0.0-00010101000000-000000000000 // indirect
