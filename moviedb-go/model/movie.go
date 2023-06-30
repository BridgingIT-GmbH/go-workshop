package model

type Movie struct {
	Id                string   `bson:"id,omitempty" json:"id,omitempty"`
	Title             string   `bson:"title,omitempty" json:"title,omitempty"`
	Description       string   `bson:"description,omitempty" json:"description,omitempty"`
	Actors            []string `bson:"actors" json:"actors"`
	DurationInMinutes int      `bson:"durationInMinutes,omitempty" json:"durationInMinutes,omitempty"`
	Base64Cover       string   `bson:"base64Cover,omitempty" json:"-,omitempty"`
}
