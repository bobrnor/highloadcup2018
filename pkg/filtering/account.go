package filtering

type Premium struct {
	Start  int64 `json:"start"`
	Finish int64 `json:"finish"`
}

type Likes struct {
	ID int64 `json:"id"`
	TS int64 `json:"ts"`
}

type Account struct {
	ID        int64    `json:"id,omitempty"`
	Sex       string   `json:"sex,omitempty"`
	Email     string   `json:"email,omitempty"`
	Status    string   `json:"status,omitempty"`
	Fname     *string  `json:"fname,omitempty"`
	SName     *string  `json:"sname,omitempty"`
	Phone     *string  `json:"phone,omitempty"`
	Country   *string  `json:"country,omitempty"`
	City      *string  `json:"city,omitempty"`
	Birth     int64    `json:"birth,omitempty"`
	Interests []string `json:"interests,omitempty"`
	Likes     []Likes  `json:"likes,omitempty"`
	Premium   *Premium `json:"premium,omitempty"`
}
