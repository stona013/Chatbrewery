package main

import (
	"ddServer/handlers"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// Test case 1: Check if the root route ("/") returns the expected response.
	t.Run("Root Route", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.FormHandler(content, &Monsters))
		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	// Test case 2: Check if the "/submit" route returns the expected response.
	t.Run("Submit Route", func(t *testing.T) {
		dir, err := filepath.Abs("test_data")
		if err != nil {
			log.Fatal(err)
		}

		filename := filepath.Join(dir, "monster.json")
		EnsureDirExists(dir)

		formData := url.Values{
			"filename":    {filename},
			"name":        {"Monster Name"},
			"source":      {"Monster Source"},
			"size":        {"Monster Size"},
			"type":        {"Monster Type"},
			"alignment":   {"Monster Alignment"},
			"ac":          {"15"},            // Beispielwert für AC
			"acFrom":      {"Natural Armor"}, // Beispielwert für AC From
			"hpAverage":   {"30"},            // Beispielwert für HP Average
			"hpFormula":   {"2d10+5"},        // Beispielwert für HP Formula
			"speed":       {"30"},            // Beispielwert für Speed
			"str":         {"16"},            // Beispielwert für Str
			"dex":         {"14"},            // Beispielwert für Dex
			"con":         {"18"},            // Beispielwert für Con
			"int":         {"10"},            // Beispielwert für Int
			"wis":         {"12"},            // Beispielwert für Wis
			"cha":         {"8"},             // Beispielwert für Cha
			"saveDex":     {"+2"},            // Beispielwert für Save Dex
			"saveCon":     {"+4"},            // Beispielwert für Save Con
			"saveWis":     {"+1"},            // Beispielwert für Save Wis
			"perception":  {"+3"},            // Beispielwert für Perception
			"stealth":     {"+2"},            // Beispielwert für Stealth
			"damageRes":   {"Fire, Cold"},    // Beispielwert für Damage Resistances
			"senses":      {"Darkvision"},    // Beispielwert für Senses
			"languages":   {"Common"},        // Beispielwert für Languages
			"cr":          {"2"},             // Beispielwert für CR
			"traitName":   {"Trait Name"},    // Beispielwert für Trait Name
			"traitEntry":  {"Trait Entry"},   // Beispielwert für Trait Entry
			"actionName":  {"Action Name"},   // Beispielwert für Action Name
			"actionEntry": {"Action Entry"},  // Beispielwert für Action Entry
		}

		log.Println("Writing data to file:", filename)

		req, _ := http.NewRequest("POST", "/submit", strings.NewReader(formData.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.SubmitHandler(content, &chars, &Monsters, filename))
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	// Test case 3: Check if the "/images/" route returns the expected response.
	t.Run("Images Route", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/images/", nil)
		rr := httptest.NewRecorder()
		handler := http.StripPrefix("/images/", http.FileServer(http.FS(content)))
		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	// Test case 4: Check if the "/addMonster" route returns the expected response.
	t.Run("AddMonster Route", func(t *testing.T) {
		dir, err := filepath.Abs("test_data")
		if err != nil {
			log.Fatal(err)
		}

		filename := filepath.Join(dir, "monster.json")
		EnsureDirExists(dir)
		formData := url.Values{
			"filename":    {filename},
			"name":        {"Monster Name"},
			"source":      {"Monster Source"},
			"size":        {"Monster Size"},
			"type":        {"Monster Type"},
			"alignment":   {"Monster Alignment"},
			"ac":          {"15"},            // Beispielwert für AC
			"acFrom":      {"Natural Armor"}, // Beispielwert für AC From
			"hpAverage":   {"30"},            // Beispielwert für HP Average
			"hpFormula":   {"2d10+5"},        // Beispielwert für HP Formula
			"speed":       {"30"},            // Beispielwert für Speed
			"str":         {"16"},            // Beispielwert für Str
			"dex":         {"14"},            // Beispielwert für Dex
			"con":         {"18"},            // Beispielwert für Con
			"int":         {"10"},            // Beispielwert für Int
			"wis":         {"12"},            // Beispielwert für Wis
			"cha":         {"8"},             // Beispielwert für Cha
			"saveDex":     {"+2"},            // Beispielwert für Save Dex
			"saveCon":     {"+4"},            // Beispielwert für Save Con
			"saveWis":     {"+1"},            // Beispielwert für Save Wis
			"perception":  {"+3"},            // Beispielwert für Perception
			"stealth":     {"+2"},            // Beispielwert für Stealth
			"damageRes":   {"Fire, Cold"},    // Beispielwert für Damage Resistances
			"senses":      {"Darkvision"},    // Beispielwert für Senses
			"languages":   {"Common"},        // Beispielwert für Languages
			"cr":          {"2"},             // Beispielwert für CR
			"traitName":   {"Trait Name"},    // Beispielwert für Trait Name
			"traitEntry":  {"Trait Entry"},   // Beispielwert für Trait Entry
			"actionName":  {"Action Name"},   // Beispielwert für Action Name
			"actionEntry": {"Action Entry"},  // Beispielwert für Action Entry
		}
		req, _ := http.NewRequest("POST", "/addMonster", strings.NewReader(formData.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.SubmitHandler(content, &chars, &Monsters, filename))
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	// Test case 5: Check if the "/main" route returns the expected response.
	t.Run("Main Route", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/main", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.MainHandler(content, &Monsters))
		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	// Test case 6: Check if the "/about" route returns the expected response.
	t.Run("About Route", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/about", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.AboutHandler(content))
		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	// Test case 7: Check if the "/contact" route returns the expected response.
	t.Run("Contact Route", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/contact", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.ContactHandler(content))
		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func EnsureDirExists(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Println("Error creating directory:", err)
	}
	return err
}
