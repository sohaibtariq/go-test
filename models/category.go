package models

import (
    "encoding/json"
)

type Category struct {
    Id   *int64  `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
}

func (c *Category) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

func (c *Category) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.Name != nil {
        structMap["name"] = c.Name
    }
    return structMap
}

func (c *Category) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id   *int64  `json:"id,omitempty"`
        Name *string `json:"name,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Id = temp.Id
    c.Name = temp.Name
    return nil
}
