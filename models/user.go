package models

// User Model
type User struct {
	UUID        string `json:uuid`
	FirstName	string `json:first_name`
	LastName	string `json:last_name`
	Email       string `json:email`
	Password    string `json:password,omitempty`
	SumRatings  int    `json:sum_ratings`
	NumRatings  int    `json:num_ratings`
	Bio         string `json:bio`
	School      string `json:school`
	State       string `json:state`
	City        string `json:city`
	PhoneNumber string `json:phone_number`
	PicURL      string `json:pic_url`
}
