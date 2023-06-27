package models

import (
    "encoding/json"
)

type Address struct {
    Street *string `json:"street,omitempty"`
    City   *string `json:"city,omitempty"`
    State  *string `json:"state,omitempty"`
    Zip    *string `json:"zip,omitempty"`
}

func (a *Address) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

func (a *Address) toMap() map[string]any {
    structMap := make(map[string]any)
    if a.Street != nil {
        structMap["street"] = a.Street
    }
    if a.City != nil {
        structMap["city"] = a.City
    }
    if a.State != nil {
        structMap["state"] = a.State
    }
    if a.Zip != nil {
        structMap["zip"] = a.Zip
    }
    return structMap
}

func (a *Address) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Street *string `json:"street,omitempty"`
        City   *string `json:"city,omitempty"`
        State  *string `json:"state,omitempty"`
        Zip    *string `json:"zip,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    a.Street = temp.Street
    a.City = temp.City
    a.State = temp.State
    a.Zip = temp.Zip
    return nil
}
