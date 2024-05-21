package datatransferobjects

import "time"

type ClientParams struct {
	Nested bool `json:"nested"`
}

type OrderParams struct {
	Client    string `json:"client"`
	Nested    bool   `json:"nested"`
	Calculate bool   `json:"calculate"`
	Costrate  string `json:"costrate"`
}

type AdvertisementParams struct {
	OrderID   int
	FromDate  time.Time `json:"fromDate"`
	ToDate    time.Time `json:"toDate"`
	Actual    bool      `json:"actual"`
	Calculate bool      `json:"calculate"`
	Costrate  string    `json:"costrate"`
}

type SubscribeParams struct {
	UserID string `json:"userID"`
	URL    string `json:"url"`
}

type FileParams struct {
	Format     string `json:"format,omitempty"`
	Size       string `json:"size,omitempty"`
	Miniatures bool   `json:"miniatures"`
	Token      string `json:"token"`
}

type UserToken struct {
	Log  string `json:"log"`
	Perm int32  `json:"perm"`
	Exp  int64  `json:"exp"`
}

type Files struct {
	FileNames []string `json:"fileNames"`
}

type File struct {
	Name string `json:"name"`
	Size int64  `json:"size,omitempty"`
	Data []byte `json:"data,omitempty"`
}
