package structure

type Row struct {
	Id           int
	OwnerId      int
	Name         string
	Description  string
	Category     string
	WaitingState bool
	Detail       string
	IsOpen       bool
	NowPersonnel int
}

type UserInfo struct {
	Id      int
	StoreId int
	Tocken  string
}
