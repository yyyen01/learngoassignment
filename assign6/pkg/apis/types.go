package apis

//type Circle struct {
//	ID           uint
//	Timestamp    int64
//	PersonID     uint
//	PersonName   string
//	Content      string
//	AtTimeHeight float32
//	AtTimeWeight float32
//	Visible      bool
//}

type TopPost struct {
	ID            uint32
	Timestamp     int64
	PersonID      uint32
	PersonName    string
	Content       string
	AtTimeHeight  float32
	AtTimeWeight  float32
	AtTimeFatRate float32
}

func (*Circle) TableName() string {
	return "testdb.circle"
}
