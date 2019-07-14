package mocks

import (
	"github.com/angelhack2019/lib/models"
)

type Users []models.User

var MockUsers = Users{
	{
		UUID:       "11e7d0a4-bdce-4423-92b8-c3a3fb0d55a6",
		Email:      "fakejdoe@uw.edu",
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
		UUID:       "c9405078-d394-40b3-9874-accdb3ccf1c6",
		Email:      "fakesfields@uw.edu",
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
