package models

import (
  "github.com/jinzhu/gorm"
)

type (
  // UserModel 
  User struct {
    gorm.Model
    FirstName string `json:"first_name" gorm:"type:varchar(30)"`
    LastName  string `json:"last_name" gorm:"type:varchar(30)"`
    Phone     string `json:"phone" gorm:"type:varchar(10);unique_index"`
    Email     string `json:"email" gorm:"type:varchar(255);unique_index"`
    Otp       string `json:"otp"`
    AuthToken string `json:"auth_token"`
  }

  // transformedUser represents a formatted user
  TransformedUser struct {
    ID        uint   `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Phone     string `json:"phone"`
    Email     string `json:"email"`
  }
)
