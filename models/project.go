package models

import "time"

type Project struct {
  Id int64 `json:"id"`

  Name        string `json:"name"`
  Description string `json:"description"`
  Link        string `json:"link"`

  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
}
