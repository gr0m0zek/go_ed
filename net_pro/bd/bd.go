package bd

type Vehicle struct {
	ID       int64  `gorm:"primaryKey"`
	Mark     string `gorm:"type varchar(25)"`
	Model    string `gorm:"type varchar(25)"`
	Number   string `gorm:"type varchar(25)"`
	Distance uint64
	Year     uint16 //add check to production year. Exactly datatypes.Date ...
}
