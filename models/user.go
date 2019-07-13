package models

type User struct {
	UUID       int 		`json:uuid`
	Email      string	`json:email`
	Password   string	`json:password`
	SumRatings int		`json:sum_ratings`
	NumRatings int		`json:num_ratings`
	Bio        string	`json:bio`
	School     string	`json:school`
	State      string	`json:state`
	Phone      string	`json:phone`
	PicUrl     string	`json:pic_url`
}
