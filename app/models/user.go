package models

type User struct {
  Id                  int
  Name                string  `sql:"type:varchar(100)"`
}

