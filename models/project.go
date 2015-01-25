package models

import "time"

type Project struct {
  Id int64 `json:"id"`

  Name        string `json:"name",schema:"name"`
  Description string `json:"description",schema:"description"`
  Link        string `json:"link",schema:"link"`

  CreatedAt time.Time `json:"createdAt",schema:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt",schema:"updatedAt"`

  User       User
  UserId     int64 `json:"userId,schema:-"`
  Category   Category
  CategoryId int64 `json:"categoryId,schema:-"`
}
