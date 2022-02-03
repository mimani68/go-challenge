package domain

type Segment struct {
	Id      string `json:"id,omitempty" bson:"_id"`
	UNAME   string `json:"uname,omitempty" bson:"uname"`
	Segment string `json:"segment,omitempty" bson:"segment"`
}
