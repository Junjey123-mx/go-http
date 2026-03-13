package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	ID      int    `json:"id"`
	Cancion string `json:"cancion"`
	Album   string `json:"album"`
	Autor   string `json:"autor"`
	Genero  string `json:"genero"`
	Anio    int    `json:"anio"`
	Sello   string `json:"sello"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

var items []Item

func main() {
	loadItems()

	http.HandleFunc("/api/items", itemsHandler)
	http.HandleFunc("/api/items/", itemByIDHandler)

	log.Println("Ejercicio 4 API running on :24584")
	log.Fatal(http.ListenAndServe(":24584", nil))
}

func loadItems() {
	file, err := os.ReadFile("./data/items.json")
	if err != nil {
		log.Fatal("error reading file:", err)
	}

	err = json.Unmarshal(file, &items)
	if err != nil {
		log.Fatal("error parsing JSON:", err)
	}
}

func saveItems() {
	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		log.Println("error marshaling JSON:", err)
		return
	}

	err = os.WriteFile("./data/items.json", data, 0644)
	if err != nil {
		log.Println("error writing file:", err)
	}
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetItems(w, r)
	case http.MethodPost:
		handleCreateItem(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func itemByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		handleUpdateItem(w, r)
	case http.MethodPatch:
		handlePatchItem(w, r)
	case http.MethodDelete:
		handleDeleteItem(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func handleGetItems(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	idParam := query.Get("id")
	autorParam := strings.TrimSpace(query.Get("autor"))
	generoParam := strings.TrimSpace(query.Get("genero"))

	if idParam != "" {
		id, err := strconv.Atoi(idParam)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid id parameter")
			return
		}

		for _, item := range items {
			if item.ID == id {
				writeJSON(w, http.StatusOK, item)
				return
			}
		}

		writeError(w, http.StatusNotFound, "item not found")
		return
	}

	if autorParam == "" && generoParam == "" {
		writeJSON(w, http.StatusOK, items)
		return
	}

	var filtered []Item
	for _, item := range items {
		matchAutor := autorParam == "" || strings.EqualFold(item.Autor, autorParam)
		matchGenero := generoParam == "" || strings.EqualFold(item.Genero, generoParam)

		if matchAutor && matchGenero {
			filtered = append(filtered, item)
		}
	}

	writeJSON(w, http.StatusOK, filtered)
}

func handleCreateItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item

	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	if err := validateItem(newItem); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	newItem.ID = generateNextID()
	items = append(items, newItem)
	saveItems()

	writeJSON(w, http.StatusCreated, newItem)
}

func handleUpdateItem(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id in path")
		return
	}

	var updated Item
	err = json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	if err := validateItem(updated); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	for i, item := range items {
		if item.ID == id {
			updated.ID = id
			items[i] = updated
			saveItems()
			writeJSON(w, http.StatusOK, updated)
			return
		}
	}

	writeError(w, http.StatusNotFound, "item not found")
}

func handlePatchItem(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id in path")
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	for i, item := range items {
		if item.ID == id {
			if v, ok := updates["cancion"].(string); ok && strings.TrimSpace(v) != "" {
				item.Cancion = v
			}
			if v, ok := updates["album"].(string); ok && strings.TrimSpace(v) != "" {
				item.Album = v
			}
			if v, ok := updates["autor"].(string); ok && strings.TrimSpace(v) != "" {
				item.Autor = v
			}
			if v, ok := updates["genero"].(string); ok && strings.TrimSpace(v) != "" {
				item.Genero = v
			}
			if v, ok := updates["anio"].(float64); ok && int(v) > 0 {
				item.Anio = int(v)
			}
			if v, ok := updates["sello"].(string); ok && strings.TrimSpace(v) != "" {
				item.Sello = v
			}

			items[i] = item
			saveItems()
			writeJSON(w, http.StatusOK, item)
			return
		}
	}

	writeError(w, http.StatusNotFound, "item not found")
}

func handleDeleteItem(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id in path")
		return
	}

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			saveItems()
			writeJSON(w, http.StatusOK, MessageResponse{Message: "item deleted"})
			return
		}
	}

	writeError(w, http.StatusNotFound, "item not found")
}

func validateItem(item Item) error {
	if strings.TrimSpace(item.Cancion) == "" {
		return newValidationError("cancion is required")
	}
	if strings.TrimSpace(item.Album) == "" {
		return newValidationError("album is required")
	}
	if strings.TrimSpace(item.Autor) == "" {
		return newValidationError("autor is required")
	}
	if strings.TrimSpace(item.Genero) == "" {
		return newValidationError("genero is required")
	}
	if item.Anio <= 0 {
		return newValidationError("anio must be greater than 0")
	}
	if strings.TrimSpace(item.Sello) == "" {
		return newValidationError("sello is required")
	}

	return nil
}

func newValidationError(msg string) error {
	return &validationError{message: msg}
}

type validationError struct {
	message string
}

func (e *validationError) Error() string {
	return e.message
}

func generateNextID() int {
	maxID := 0
	for _, item := range items {
		if item.ID > maxID {
			maxID = item.ID
		}
	}
	return maxID + 1
}

func getIDFromPath(path string) (int, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) != 3 {
		return 0, strconv.ErrSyntax
	}

	return strconv.Atoi(parts[2])
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorResponse{Error: message})
}
