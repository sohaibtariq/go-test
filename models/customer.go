package models

import (
    "encoding/json"
)

type Customer struct {
    Id       *int64     `json:"id,omitempty"`
    Username *string    `json:"username,omitempty"`
    Address  *[]Address `json:"address,omitempty"`
}

func (c *Customer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

func (c *Customer) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.Username != nil {
        structMap["username"] = c.Username
    }
    if c.Address != nil {
        structMap["address"] = c.Address
    }
    return structMap
}

func (c *Customer) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id       *int64     `json:"id,omitempty"`
        Username *string    `json:"username,omitempty"`
        Address  *[]Address `json:"address,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Id = temp.Id
    c.Username = temp.Username
    c.Address = temp.Address
    return nil
}
