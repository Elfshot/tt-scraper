package mongo_models

import "time"

type SotdCollModel struct {
	Timestamp time.Time `bson:"timestamp"`
	Bonus     uint8     `bson:"bonus"`
	Skill     string    `bson:"skill"`
	Aptitude  string    `bson:"aptitude,omitempty"`
}
