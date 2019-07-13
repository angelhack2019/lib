package mocks

import (
	"github.com/angelhack2019/lib/models"
)

type Users []models.User

var MockUsers = Users{
	{
		UUID:       "f30bf3a-8591-4034-950b-d611d3a93c47",
		Email:      "jdoe@uw.edu",
		Password:   "Unencrypted",
		SumRatings: 11,
		NumRatings: 3,
		Bio:        "I am generous",
		School:     "UW Bothel",
		State:      "WA",
		Phone:      "2061234567",
		PicURL:     "https://f4.bcbits.com/img/a0777435656_10.jpg",
	},
	{
		UUID:       "f30bf3a-8591-4034-950b-d611d3a93c47",
		Email:      "sfields@uw.edu",
		Password:   "NotSafe",
		SumRatings: 0,
		NumRatings: 0,
		Bio:        "I like food",
		School:     "UW Bothel",
		State:      "WA",
		Phone:      "2067654321",
		PicURL:     "https://upload.wikimedia.org/wikipedia/commons/thumb/1/18/Gidget_main_cast_1966_%28cropped%29.jpg/220px-Gidget_main_cast_1966_%28cropped%29.jpg",
	},
}
