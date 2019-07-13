package mocks

import (
	"github.com/angelhack2019/lib/models"
)

type Foods []models.Food

// Food 1 : Created 7/13 expires 7/19
// Food 2 : Created 7/5 expires 7/12
var MockFoods = Foods{
	{
		UUID:       "688c0daf-2833-485c-a419-b84a152a75a6",
		PicURL:     "https://upload.wikimedia.org/wikipedia/commons/1/1d/Bagel-Plain-Alt.jpg",
		CreateDate: 1563030000,
		ExpDate:	1563494400,
	},
	{
		UUID:       "688c0daf-2833-485c-a419-b84a152a75a6",
		PicURL:     "https://www.impactcommunicationsinc.com/wp-content/uploads/2018/03/18-04_has_your_credibility_banana_turned_brown_1000.jpg",
		CreateDate: 1562319325,
		ExpDate:	1562889600,
	},
}
