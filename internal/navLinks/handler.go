package navLinks

import (
	"cms-project/pkg/response"
	"net/http"
)

// GetMenusHandler handles retrieving all menus
// @Summary      Get navigation links
// @Description  Retrieve a list of all navigation links (menus)
// @Tags         NavLink
// @Success      200 {object} response.APIResponse
// @Failure      500  {object}  response.APIResponse  "Internal Server Error"
// @Router       /api/nav_links [get]
func GetNavLinksHandler(w http.ResponseWriter, r *http.Request) {
	menus, err := GetNavLinks()
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, false, "Failed to fetch menus", nil)
		return
	}
	response.JSON(w, http.StatusOK, true, "Users retrieved successfully", menus)
}

// CreateMenuHandler handles creating a new menu
// func CreateMenuHandler(w http.ResponseWriter, r *http.Request) {
// 	var menu Menu
// 	if err := json.NewDecoder(r.Body).Decode(&menu); err != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}
// 	if err := CreateMenu(menu); err != nil {
// 		http.Error(w, "Failed to create menu", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// }

// GetMenuByIDHandler handles retrieving a single menu by ID
// func GetMenuByIDHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid menu ID", http.StatusBadRequest)
// 		return
// 	}
// 	menu, err := GetMenuByID(id)
// 	if err != nil {
// 		http.Error(w, "Menu not found", http.StatusNotFound)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(menu)
// }

// DeleteMenuHandler handles deleting a menu by ID
// func DeleteMenuHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid menu ID", http.StatusBadRequest)
// 		return
// 	}
// 	if err := DeleteMenu(id); err != nil {
// 		http.Error(w, "Failed to delete menu", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusNoContent)
// }
