package models

import (
    "encoding/json"
)

type Tag struct {
    Id   *int64  `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
}

func (t *Tag) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

func (t *Tag) toMap() map[string]any {
    structMap := make(map[string]any)
    if t.Id != nil {
        structMap["id"] = t.Id
    }
    if t.Name != nil {
        structMap["name"] = t.Name
    }
    return structMap
}

func (t *Tag) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id   *int64  `json:"id,omitempty"`
        Name *string `json:"name,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    t.Id = temp.Id
    t.Name = temp.Name
    return nil
}
