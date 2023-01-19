package models

import "github.com/lib/pq"

type Recipe struct {
  Id int
  Name string
  Description string
  PrepMinutes int
  Author int
  Steps pq.StringArray `gorm:"type:string[]"`
  Nutrition pq.Float64Array `gorm:"type:float64[]"`
}

