package models

type Food struct {
	UUID 		int 	`json:uuid`
	PicUrl 		string 	`json:pic_url`
	ExpDate		int64	`json:exp_date`
	CreateDate	int64	`json:create_date`
}
