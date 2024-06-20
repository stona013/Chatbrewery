package handlers

import (
	"bytes"
	"ddServer/model"
	"embed"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// Function to handle file download
func downloadFile(w http.ResponseWriter, r *http.Request, filePath, fileType string) {
	w.Header().Set("Content-Disposition", "attachment; filename="+filePath)
	if fileType == "json" {
		w.Header().Set("Content-Type", "application/json")
	} else {
		w.Header().Set("Content-Type", "text/plain")
	}
	http.ServeFile(w, r, filePath)
}

// Function to generate the JSON structure based on the parameters
func generateJsonStructure(choosenFactor string) string {
	baseJson := ""

	if choosenFactor == "monster" {
		baseJson += `and use the following json structure as a sample but put your variable into the fields: 
    {
        "_meta": {
            "_dateLastModifiedHash": "66658f76",
            "sources": [
                {
                    "json": "chatbrewery",
                    "abbreviation": "CGPT",
                    "version": "0.1.6",
                    "authors": [
                        "Chat GPT"
                    ],
                    "convertedBy": [
                        "Krzysztof"
                    ]
                }
            ],
            "dateAdded": 1717931894,
            "dateLastModified": 1717931894
        },
        "monster": [
            {
                "source": "chatbrewery",
                "cr": "MonsterCR",
                "type": "MonsterType",
                "name": "MonsterName",
                "ac": [
                    {
                        "from": [
                            "ArmorType"
                        ],
                        "ac": ArmorClass
                    }
                ],
                "str": Strength,
                "dex": Dexterity,
                "con": Constitution,
                "int": Intelligence,
                "wis": Wisdom,
                "cha": Charisma,
                "size": [
                    "MonsterSize"
                ],
                "speed": {
                    "walk": WalkSpeed,
                    "burrow": BurrowSpeed,
                    "climb": ClimbSpeed,
                    "fly": FlySpeed,
                    "swim": SwimSpeed
                },
                "save": {
                    "dex": DexteritySave,
                    "con": ConstitutionSave,
                    "wis": WisdomSave,
                    "cha": CharismaSave,
                    "str": StrengthSave,
                    "int": IntelligenceSave
                },
                "skill": {
                    "stealth": StealthSkill,
                    "acrobatics": AcrobaticsSkill,
                    "animalHandling": AnimalHandlingSkill,
                    "arcana": ArcanaSkill,
                    "athletics": AthleticsSkill,
                    "deception": DeceptionSkill,
                    "history": HistorySkill,
                    "insight": InsightSkill,
                    "intimidation": IntimidationSkill,
                    "investigation": InvestigationSkill,
                    "medicine": MedicineSkill,
                    "nature": NatureSkill,
                    "perception": PerceptionSkill,
                    "performance": PerformanceSkill,
                    "persuasion": PersuasionSkill,
                    "sleightOfHand": SleightOfHandSkill,
                    "survival": SurvivalSkill,
                    "religion": ReligionSkill
                },
                "hp": {
                    "formula": "HPFormula",
                    "average": AverageHP
                },
                "senses": [
                    "MonsterSenses"
                ],
                "languages": [
                    "MonsterLanguages"
                ],
                "vulnerable": [
                    "Vulnerabilities"
                ],
                "conditionImmune": [
                    "ConditionImmunities"
                ],
                "resist": [
                    "Resistances"
                ],
                "immune": [
                    "Immunities"
                ],
                "trait": [
                    {
                        "name": "TraitName",
                        "entries": [
                            "TraitDescription"
                        ]
                    }
                ],
                "action": [
                    {
                        "name": "ActionName",
                        "entries": [
                            "{@atk mw} {@hit AttackBonus} to hit, reach Reach ft., one target. {@h}Damage ({@damage}) DamageType damage."
                        ]
                    }
                ],
                "fluff": {
                    "entries": [
                        "Description",
                        "DescriptionText",
                    ]
                }
                {{if .IsLegendary}},
                "legendaryActions": NumberOfLegendaryActions,
                "legendaryHeader": [
                    ""
                ],
                "legendary": [
                    {
                        "name": "LegendaryActionName (ActionCost)",
                        "entries": [
                            "LegendaryActionEntries"
                        ]
                    },
                    {
                        "name": "LegendaryActionName (ActionCost)",
                        "entries": [
                            "LegendaryActionEntries"
                        ]
                    },
                    {
                        "name": "LegendaryActionName (ActionCost)",
                        "entries": [
                            "LegendaryActionEntries"
                        ]
                    }
                ]
                {{end}}
                {{if .IsSpellcaster}},
	{
				"spellcasting": [
    {
      "name": "Spellcasting",
      "headerEntries": [
        "SpellcastingHeader"
      ],
     	"spells": {
						"0": {
							"spells": [
								"{@spell spellname}",
								"{@spell spellname}"
							],
						},
						"1": {
							"spells": [
								"{@spell spellname}",
								"{@spell spellname}"
							],
							"slots": 4
						},
						"2": {
							"spells": [
								"{@spell spellname}",
								"{@spell spellname}"
							],
							"slots": 3
						},
						"3": {
							"spells": [
								"{@spell spellname}",
								"{@spell spellname}"
							],
							"slots": 2
						},
						"4": {
							"spells": [
								"{@spell spellname}",
								"{@spell spellname}"
							],
							"slots": 3
						},
						"5": {
							"spells": [
								"{@spell spellname}",
								"{@spell spellname}"
							],
							"slots": 3
						},
						"6": {
							"spells": [
								"{@spell spellname}",
								"{@spell spellname}"
							],
							"slots": 3
						},
						"7": {
							"spells": [
								"{@spell spellname}",
								"{@spell spellname}"
							],
							"slots": 3
						},
						"8": {
							"spells": [
								"{@spell spellname}",
								"{@spell spellname}"
							],
							"slots": 3
						},
							"9": {
							"spells": [
								"{@spell spellname}",
								"{@spell spellname}"
							],
							"slots": 3
						}
					},
      "ability": "SpellcastingAbility",
      "type": "spellcasting"
    }
  ]
                {{end}}
            }
        ]
    }
	
	readable format without backticks at the beginning and end`
	} else {
		baseJson += `  and use the following json structure as a sample but put your variable into the fields: 
{
    "_meta": {
        "sources": [
            {
                "json": "chatbrewery",
                "abbreviation": "CGPT",
                "full": "Chatgpt",
                "authors": [
                    "Krzysztof Stankiewicz",
                    "Chat GPT"
                ],
                "convertedBy": [
                    "Krzysztof Stankiewicz",
                    "Chat GPT"
                ],
                "version": "0.1.6"
            }
        ]
    },
    "item": [
        {
            "name": "Item Name",
            "source": "chatbrewery",
            "baseItem": "Base Item Name (optional)",
            "type": "Item Type (optional)", // Types: A (Ammunition), HA (Heavy Armor), MA (Medium Armor), LA (Light Armor), INS (Instrument), P (Potion), RD (Rod), M (Melee Weapon), R (Ranged Weapon), WD (Wand), SC (Scroll), "" (Anything else)
            "tier": "Item Tier (optional)", // major, minor
            "rarity": "Rarity", // Possible Rarities: "rare", "none", "uncommon", "very rare", "legendary", "artifact", "unknown", "common", "unknown (magic)", "varies"
            "reqAttune": true,
            "weight": Weight, 
            "stealth": true, if disadvanage on Stealth
            "ac": 12,  (optional) 
            "acSpecial": "12 AC, 15 AC during the night", (optional) 
            "weaponCategory": "category", // possible category simple / martial (optional) 
            "ammunition": true, (optional) 
            "range": "80/320", (optional) 
            "wondrous": true, (optional)
            "bonusSpellAttack": "+1", (optional)
            "bonusAC": "+1", (optional)
            "bonusWeapon": "+1", (optional)
            "bonusWeaponAttack": "+1", (optional)
            "bonusSpellAttack": "+1", (optional)
            "bonusSavingThrow": "+1", (optional)
            "charges": number, (optional)
            "recharge": "time", // posiliby time: "round", "restShort", "restLong", "dawn", "dusk" or "midnight". (optional)
            "ability": { (optional)
			"static"{ (optional)
                } 
			 "str": number
			 "dex": number
			 "con": number
			 "wis": number
			 "cha": number
			 "int": number
            }
            "resist": [ (optional)
                "damage type"
            ],
			"attachedSpells": [ (optional)
                "spell name"
            ],			
			 "curse": true/false,
            "dmg1:": "how many damage die" /possible damage dice: xd2, xd4, xd6, xd8, xd10, xd12
			"dmgtype": "Damage type" // Possible Types: "A" (Acid), "B" (Bludgeoning), "C" (Cold), "F" (Fire), "O" (Force), "L" (Lightning), "N" (Necrotic), "P" (Piercing), "I" (Poison), "Y" (Psychic), "R" (Radiant), "S" (Slashing), "T" (Thunder)
            "property": [
                "Item Property (optional)" / Weapon Properties: "T" (Thrown),
                "V" (Versatile),
                "H" (Heavy),
                "2H" (Two-Handed),
                "F" (Finesse),
                "L" (Light),
                "R" (Reach),
                "A" (Ammunition),
                "LD" (Loading),
                "S" (Special),
                "AF" (Ammunition (futuristic)),
                "RLD" (Reload)
            ],
            "entries": [
                "Description or Flavor Text",
                {
                    "type": "entries",
                    "name": "Entry Name (optional)",
                    "entries": [
                        "Specific details or abilities"
                    ]
                }
            ]
        }
    ]
}

			readable format without backticks at the beginning and end`

	}
	return baseJson
}

// Function to generate the text prompt based on the parameters
func generateMonsterPrompt(name, choosenFactor, cr, monsterType, monsterInfo string, isJson, isRandom, isLegendary, isSpellcaster bool) string {
	if isJson {
		if isRandom {
			prompt := `Create a detailed and  random DnD monster. You choose the monster name(Think of a Name for the Monster and make it suits its Type features and other Factors), 
			type (choose a type from: Fiends, Undead, Beast, Monstrosity, Celestial, Abberation, Humanoid, Giant, Elemental, Dragon, Construct, Ooze, Fey or Plant), 
			and CR (choose a number 1 and 30). also write a detailed discription and backstory for the monster into the DescriptionTextfield in the json. `
			if isLegendary {
				prompt += ` Make it a legendary monster with appropriate legendary actions and resistances.`
			}
			if isSpellcaster {
				prompt += ` Make it a spellcaster with appropriate spells.`
			}
			prompt += generateJsonStructure(choosenFactor)
			return prompt
		} else {
			prompt := `Create a detailed and unique Homebrew DnD monster. If not given Choose the monster name, type, and CR (between 1 and 30). ALso  
			write a detailed discription and backstory for the monster into the DescriptionTextfield in the json.`
			if name != "" {
				prompt += "\nName: " + name
			}
			if cr != "" {
				prompt += "\nCR: " + cr
			}
			if monsterType != "" {
				prompt += "\nType: " + monsterType
			}
			if monsterInfo != "" {
				prompt += "\nDetails: " + monsterInfo
			}
			if isLegendary {
				prompt += `\nMake it a legendary monster with appropriate legendary actions and resistances.`
			}
			if isSpellcaster {
				prompt += `\nMake it a spellcaster with appropriate spells.`
			}
			prompt += generateJsonStructure(choosenFactor)
			return prompt
		}
	}

	if isRandom {
		prompt := `Create a detailed, random and unique DnD monster. Include and choose all necessary details such as name, type, abilities, CR, and stats.`
		if isLegendary {
			prompt += ` Make it a legendary monster with appropriate legendary actions and resistances.`
		}
		if isSpellcaster {
			prompt += ` Make it a spellcaster with appropriate spells.`
		}
		prompt += ` Respond with the monster sheet in a clear, readable format without backticks at the beginning and end.`
		return prompt
	} else {
		prompt := `Create a detailed and unique DnD monster. Include all necessary details such as name, type, abilities, CR, and stats and use the following details for the creation.`
		if name != "" {
			prompt += "\nName: " + name
		}
		if cr != "" {
			prompt += "\nCR: " + cr
		}
		if monsterType != "" {
			prompt += "\nType: " + monsterType
		}
		if monsterInfo != "" {
			prompt += "\nDetails: " + monsterInfo
		}
		if isLegendary {
			prompt += `\nMake it a legendary monster with appropriate legendary actions and resistances.`
		}
		if isSpellcaster {
			prompt += `\nMake it a spellcaster with appropriate spells.`
		}
		prompt += ` Respond with the monster sheet in a clear, readable format without backticks at the beginning and end.`
		return prompt
	}
}

func generateItemPrompt(name, choosenFactor, itemType, itemRarity, itemInfo string, isJson, isRandom bool) string {
	if isJson {
		prompt := `Create a detailed DnD magic item.`
		if name != "" {
			prompt += "\nName: " + name
		}
		if itemType != "" {
			prompt += "\nType: " + itemType
		}
		if itemRarity != "" {
			prompt += "\nRarity: " + itemRarity
		}
		if itemInfo != "" {
			prompt += "\nDetails: " + itemInfo
		}
		prompt += generateJsonStructure(choosenFactor)
		return prompt
	}

	prompt := `Create a detailed DnD magic item.`
	if name != "" {
		prompt += "\nName: " + name
	}
	if itemType != "" {
		prompt += "\nType: " + itemType
	}
	if itemRarity != "" {
		prompt += "\nRarity: " + itemRarity
	}
	if itemInfo != "" {
		prompt += "\nDetails: " + itemInfo
	}
	prompt += `\nRespond with the item details in a clear, readable format.`
	return prompt
}

func AIHandler(content embed.FS, monsters *[]model.Monster) http.HandlerFunc {
	log.Print("AIHandler called")

	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("AIHandler request received")

		if r.Method == http.MethodPost {
			apiKey := r.FormValue("apikey")
			generationType := r.FormValue("generationType")
			fileFormat := r.FormValue("fileFormat")
			isLegendary := r.FormValue("legendary") == "legendary"
			isSpellcaster := r.FormValue("spellcaster") == "spellcaster"

			isRandom := generationType == "random"
			isJson := fileFormat == "json"
			choosenFactor := r.FormValue("choosenFactor")

			var prompt string

			if choosenFactor == "monster" {
				name := r.FormValue("monsterName")
				cr := r.FormValue("monsterCR")
				monsterType := r.FormValue("monsterType")
				monsterInfo := r.FormValue("monsterinfo")
				prompt = generateMonsterPrompt(name, choosenFactor, cr, monsterType, monsterInfo, isJson, isRandom, isLegendary, isSpellcaster)
			} else {
				name := r.FormValue("itemName")
				itemType := r.FormValue("itemType")
				itemRarity := r.FormValue("itemRarity")
				itemInfo := r.FormValue("iteminfo")
				prompt = generateItemPrompt(name, choosenFactor, itemType, itemRarity, itemInfo, isJson, isRandom)
			}

			log.Print(prompt)

			data := map[string]interface{}{
				"model":       "gpt-4-turbo",
				"messages":    []map[string]string{{"role": "user", "content": prompt}},
				"max_tokens":  2500,
				"temperature": 0.6,
			}

			requestData, err := json.Marshal(data)
			if err != nil {
				log.Printf("Error marshalling JSON: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestData))
			if err != nil {
				log.Printf("Error creating request: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+apiKey)

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Printf("Error sending request: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()

			responseBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Error reading response: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			var aiResponse OpenAIResponse
			if err := json.Unmarshal(responseBody, &aiResponse); err != nil {
				log.Printf("Error unmarshalling JSON: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if len(aiResponse.Choices) > 0 {
				messageContent := aiResponse.Choices[0].Message.Content

				var fileExtension string
				if fileFormat == "json" {
					fileExtension = ".json"
				} else {
					fileExtension = ".txt"
				}

				filename := "generated_" + choosenFactor + fileExtension
				err := ioutil.WriteFile(filename, []byte(messageContent), 0644)
				if err != nil {
					log.Printf("Error writing to file: %v\n", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}

				downloadFile(w, r, filename, fileFormat)
				return
			}
		}

		tmplFiles := []string{"templates/base.html", "templates/header.html", "templates/main.html", "templates/footer.html", "templates/ai.html"}
		tmpl, err := template.ParseFS(content, tmplFiles...)
		if err != nil {
			log.Printf("Template parsing error: %v\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := map[string]interface{}{
			"Title": "Dungeons & Dragons Monster Generator",
		}
		err = tmpl.ExecuteTemplate(w, "ai", data)
		if err != nil {
			log.Printf("Template execution error: %v\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
