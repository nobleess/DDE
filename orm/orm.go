package orm

import (
	"DDE/analyze"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
)

type DB struct {
	Conn *gorm.DB
}

func (db DB) Migration() {
	err := db.Conn.AutoMigrate(&Prom{}, &Like{}, &View{}, &Comment{}, &Activity{}, &Passion{}, &Count{})
	if err != nil {
		log.Fatal(err)
	}
}

var DBLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second,   // Slow SQL threshold
		LogLevel:                  logger.Silent, // Log level
		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      true,          // Don't include params in the SQL log
		Colorful:                  false,         // Disable color
	},
)

func NewDB(name string) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{Logger: DBLogger})
	if err != nil {
		return nil, err
	}
	return &DB{Conn: db}, nil
}

type Prom struct {
	ID    uint   `gorm:"primaryKey"`
	Promm string `gorm:"unique"`
}

func StatYearsAdapter(s *analyze.StatYears, PromID uint) ([]Like, []View, []Comment, []Count) {
	like := make([]Like, 0, len(s.Likes))
	view := make([]View, 0, len(s.Views))
	comment := make([]Comment, 0, len(s.Comments))
	count := make([]Count, 0, len(s.Counts))

	for y, c := range s.Likes {
		like = append(like, Like{
			PromID: PromID,
			Year:   y,
			Count:  c,
		})
	}

	for y, c := range s.Views {
		view = append(view, View{
			PromID: PromID,
			Year:   y,
			Count:  c,
		})
	}

	for y, c := range s.Comments {
		comment = append(comment, Comment{
			PromID: PromID,
			Year:   y,
			Count:  c,
		},
		)
	}

	for y, c := range s.Counts {
		count = append(count, Count{
			PromID: PromID,
			Year:   y,
			Count:  c,
		})
	}
	return like, view, comment, count
}

type Like struct {
	ID     uint `gorm:"primaryKey"`
	PromID uint
	Year   int
	Count  int
}

type View struct {
	ID     uint `gorm:"primaryKey"`
	PromID uint
	Year   int
	Count  int
}

type Comment struct {
	ID     uint `gorm:"primaryKey"`
	PromID uint
	Year   int
	Count  int
}

type Count struct {
	ID     uint `gorm:"primaryKey"`
	PromID uint
	Year   int
	Count  int
}

func BIAdapter(bi analyze.BIMetrics, PromID uint) ([]Activity, []Passion) {

	activity := make([]Activity, 0, len(bi.Activity))
	passion := make([]Passion, 0, len(bi.Passion))

	for y, c := range bi.Activity {
		activity = append(activity, Activity{
			PromID: PromID,
			Year:   y,
			Proc:   c,
		})
	}

	for y, c := range bi.Passion {
		passion = append(passion, Passion{
			PromID: PromID,
			Year:   y,
			Proc:   c,
		})
	}

	return activity, passion

}

type Activity struct {
	ID     uint `gorm:"primaryKey"`
	PromID uint
	Year   int
	Proc   int
}

type Passion struct {
	ID     uint `gorm:"primaryKey"`
	PromID uint
	Year   int
	Proc   int
}
