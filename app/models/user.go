package models

import (
  "github.com/jinzhu/gorm"
)

type (
  // UserModel 
  User struct {
    gorm.Model
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Phone     string `json:"phone"`
    Email     string `json:"email"`
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
