package models

type Subscription struct {
	ID           int64  `json:"id"`
	Service_name string `json:"service_name"`
	Price        int64  `json:"price"`
	User_id      string `json:"user_id"`
	Start_date   string `json:"start_date"`
	End_date     string `json:"end_date"`
}
