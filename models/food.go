package models

// Food Model
type Food struct {
	UUID       string `json:uuid`
	PicURL     string `json:pic_url`
	ExpDate    int64  `json:exp_date`
	CreateDate int64  `json:create_date`
}
