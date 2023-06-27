package models

import (
    "encoding/json"
)

type PetImage struct {
    Code    *int    `json:"code,omitempty"`
    Type    *string `json:"type,omitempty"`
    Message *string `json:"message,omitempty"`
}

func (p *PetImage) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

func (p *PetImage) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.Code != nil {
        structMap["code"] = p.Code
    }
    if p.Type != nil {
        structMap["type"] = p.Type
    }
    if p.Message != nil {
        structMap["message"] = p.Message
    }
    return structMap
}

func (p *PetImage) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Code    *int    `json:"code,omitempty"`
        Type    *string `json:"type,omitempty"`
        Message *string `json:"message,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Code = temp.Code
    p.Type = temp.Type
    p.Message = temp.Message
    return nil
}
