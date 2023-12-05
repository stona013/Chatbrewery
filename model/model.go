package model

import (
	"fmt"
	"os"
	"time"
)

// Monster struct für die Daten des Monsters
type Monster struct {
	Save      Save     `json:"save"`
	Skill     Skill    `json:"skill"`
	HP        HP       `json:"hp"`
	Source    string   `json:"source"`
	CR        string   `json:"cr"`
	Type      string   `json:"type"`
	Name      string   `json:"name"`
	DamageRes []string `json:"damageResistances"`
	Traits    []Trait  `json:"trait"`
	AC        []AC     `json:"ac"`
	Alignment []string `json:"alignment"`
	Senses    []string `json:"senses"`
	Languages []string `json:"languages"`
	Size      []string `json:"size"`
	Actions   []Action `json:"action"`
	Speed     Speed    `json:"speed"`
	Str       int      `json:"str"`
	Dex       int      `json:"dex"`
	Con       int      `json:"con"`
	Int       int      `json:"int"`
	Wis       int      `json:"wis"`
	Cha       int      `json:"cha"`
}

type AC struct {
	From []string `json:"from"`
	AC   int      `json:"ac"`
}

type HP struct {
	Formula string `json:"formula"`
	Average int    `json:"average"`
}

type Speed struct {
	Walk int `json:"walk"`
}

type Save struct {
	Dex string `json:"dex"`
	Con string `json:"con"`
	Wis string `json:"wis"`
}

type Skill struct {
	Perception string `json:"perception"`
	Stealth    string `json:"stealth"`
}

type Trait struct {
	Name    string   `json:"name"`
	Entries []string `json:"entries"`
}

type Action struct {
	Name    string   `json:"name"`
	Entries []string `json:"entries"`
}

// Character struct für die Daten des Charakters
type Character struct {
	Meta    Meta      `json:"_meta"`
	Monster []Monster `json:"monster"`
}

// Meta struct für Meta-Informationen
type Meta struct {
	DateLastModifiedHash string   `json:"_dateLastModifiedHash"`
	Sources              []Source `json:"sources"`
	DateAdded            int64    `json:"dateAdded"`
	DateLastModified     int64    `json:"dateLastModified"`
}

type Source struct {
	Json         string   `json:"json"`
	Abbreviation string   `json:"abbreviation"`
	Version      string   `json:"version"`
	Authors      []string `json:"authors"`
	ConvertedBy  []string `json:"convertedBy"`
}

// writeToFile schreibt Daten in eine Datei
func WriteToFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// getOrCreateCharacter gibt das aktuelle Charakterobjekt zurück oder erstellt ein neues
func GetOrCreateCharacter(filename string, chars []Character) Character {
	for _, char := range chars {
		if char.Meta.DateLastModified == 0 {
			// Ein leeres Charakterobjekt wurde gefunden
			return char
		}
	}

	// Erstelle ein neues Charakterobjekt
	now := time.Now().Unix()
	newChar := Character{
		Meta: Meta{
			Sources: []Source{
				{
					Json:         "Malgorgon",
					Abbreviation: "MG",
					Authors:      []string{"Krzysztof"},
					ConvertedBy:  []string{"Krzysztof"},
					Version:      "unknown",
				},
			},
			DateAdded:            now,
			DateLastModified:     now,
			DateLastModifiedHash: fmt.Sprintf("%x", now),
		},
		Monster: []Monster{},
	}

	chars = append(chars, newChar)

	return newChar
}
