package models

import (
    "encoding/json"
)

type Pet struct {
    Name      string         `json:"name"`
    PhotoUrls []string       `json:"photoUrls"`
    Id        *int64         `json:"id,omitempty"`
    Category  *Category      `json:"category,omitempty"`
    Tags      *[]Tag         `json:"tags,omitempty"`
    PetStatus *PetStatusEnum `json:"petStatus,omitempty"`
}

func (p *Pet) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

func (p *Pet) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["name"] = p.Name
    structMap["photoUrls"] = p.PhotoUrls
    if p.Id != nil {
        structMap["id"] = p.Id
    }
    if p.Category != nil {
        structMap["category"] = p.Category
    }
    if p.Tags != nil {
        structMap["tags"] = p.Tags
    }
    if p.PetStatus != nil {
        structMap["petStatus"] = p.PetStatus
    }
    return structMap
}

func (p *Pet) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Name      string         `json:"name"`
        PhotoUrls []string       `json:"photoUrls"`
        Id        *int64         `json:"id,omitempty"`
        Category  *Category      `json:"category,omitempty"`
        Tags      *[]Tag         `json:"tags,omitempty"`
        PetStatus *PetStatusEnum `json:"petStatus,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Name = temp.Name
    p.PhotoUrls = temp.PhotoUrls
    p.Id = temp.Id
    p.Category = temp.Category
    p.Tags = temp.Tags
    p.PetStatus = temp.PetStatus
    return nil
}
