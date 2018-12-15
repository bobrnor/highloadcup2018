package inmemory

type Premium struct {
	Start  int64 `json:"start"`
	Finish int64 `json:"finish"`
}

type Likes struct {
	ID int64 `json:"id"`
	TS int64 `json:"ts"`
}

type Account struct {
	ID        int64    `json:"id"`
	Email     string   `json:"email"`
	Fname     *string  `json:"fname"`
	Sname     *string  `json:"sname"`
	Phone     *string  `json:"phone"`
	Sex       string   `json:"sex"`
	Birth     int64    `json:"birth"`
	Country   *string  `json:"country"`
	City      *string  `json:"city"`
	Joined    int64    `json:"joined"`
	Status    string   `json:"status"`
	Interests []string `json:"interests"`
	Premium   Premium  `json:"premium"`
	Likes     []Likes  `json:"likes"`
}

type Accounts struct {
	Accounts []Account `json:"accounts"`
}
