package models

// User Model
type User struct {
	UUID        string `json:uuid`
	Email       string `json:email`
	Password    string `json:password`
	SumRatings  int    `json:sum_ratings`
	NumRatings  int    `json:num_ratings`
	Bio         string `json:bio`
	School      string `json:school`
	State       string `json:state`
	City        string `json:city`
	PhoneNumber string `json:phone_number`
	PicURL      string `json:pic_url`
}
