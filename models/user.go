package models

import (
    "encoding/json"
)

type User struct {
    Id         *int64  `json:"id,omitempty"`
    Username   *string `json:"username,omitempty"`
    FirstName  *string `json:"firstName,omitempty"`
    LastName   *string `json:"lastName,omitempty"`
    Email      *string `json:"email,omitempty"`
    Password   *string `json:"password,omitempty"`
    Phone      *string `json:"phone,omitempty"`
    UserStatus *int    `json:"userStatus,omitempty"`
}

func (u *User) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

func (u *User) toMap() map[string]any {
    structMap := make(map[string]any)
    if u.Id != nil {
        structMap["id"] = u.Id
    }
    if u.Username != nil {
        structMap["username"] = u.Username
    }
    if u.FirstName != nil {
        structMap["firstName"] = u.FirstName
    }
    if u.LastName != nil {
        structMap["lastName"] = u.LastName
    }
    if u.Email != nil {
        structMap["email"] = u.Email
    }
    if u.Password != nil {
        structMap["password"] = u.Password
    }
    if u.Phone != nil {
        structMap["phone"] = u.Phone
    }
    if u.UserStatus != nil {
        structMap["userStatus"] = u.UserStatus
    }
    return structMap
}

func (u *User) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id         *int64  `json:"id,omitempty"`
        Username   *string `json:"username,omitempty"`
        FirstName  *string `json:"firstName,omitempty"`
        LastName   *string `json:"lastName,omitempty"`
        Email      *string `json:"email,omitempty"`
        Password   *string `json:"password,omitempty"`
        Phone      *string `json:"phone,omitempty"`
        UserStatus *int    `json:"userStatus,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.Id = temp.Id
    u.Username = temp.Username
    u.FirstName = temp.FirstName
    u.LastName = temp.LastName
    u.Email = temp.Email
    u.Password = temp.Password
    u.Phone = temp.Phone
    u.UserStatus = temp.UserStatus
    return nil
}
