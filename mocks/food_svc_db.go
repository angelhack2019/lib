package mocks

import (
	"github.com/angelhack2019/lib/models"
)

type Foods []models.Food

// Food 1 : Created 7/13 expires 7/19
// Food 2 : Created 7/5 expires 7/12
var MockFoods = Foods{
	{
		UUID:       "72a09081-359d-407f-8354-4e064070f1fa",
		PicURL:     "https://upload.wikimedia.org/wikipedia/commons/1/1d/Bagel-Plain-Alt.jpg",
		CreateDate: 1563030000,
		ExpDate:    1563494400,
	},
	{
		UUID:       "c9b4ac1a-5da9-46e9-af0f-2a583b11e50a",
		PicURL:     "https://www.impactcommunicationsinc.com/wp-content/uploads/2018/03/18-04_has_your_credibility_banana_turned_brown_1000.jpg",
		CreateDate: 1562319325,
		ExpDate:    1562889600,
	},
}
