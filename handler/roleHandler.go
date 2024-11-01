package handler

import (
	"encoding/json"
	"net/http"

	"github.com/IsraelTeo/api-jwt-go/db"
	"github.com/IsraelTeo/api-jwt-go/model"
	"github.com/gorilla/mux"
)

func GetRoleById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := model.NewResponse(model.MessageTypeError, "Method get are not permit", nil)
		model.ResponseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	role := model.Role{}
	err := db.GDB.First(&role, id)
	if err != nil {
		response := model.NewResponse(model.MessageTypeError, "Role are not found", nil)
		model.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	response := model.NewResponse(model.MessageTypeSuccess, "Role found", role)
	model.ResponseJSON(w, http.StatusOK, response)
}

func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := model.NewResponse(model.MessageTypeSuccess, "Method get not permit", nil)
		model.ResponseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	var roles []model.Role
	err := db.GDB.Find(&roles)
	if err != nil {
		response := model.NewResponse(model.MessageTypeError, "Failed to fetch roles", nil)
		model.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	if len(roles) == 0 {
		response := model.NewResponse(model.MessageTypeError, "Roles List is empty", roles)
		model.ResponseJSON(w, http.StatusNoContent, response)
		return
	}

	response := model.NewResponse(model.MessageTypeSuccess, "Roles found", roles)
	model.ResponseJSON(w, http.StatusOK, response)
}

func SaveRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := model.NewResponse(model.MessageTypeError, "Method post not permit", nil)
		model.ResponseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	role := model.Role{}
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		response := model.NewResponse(model.MessageTypeError, "Bad Request: Invalid JSON data", nil)
		model.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	if err := db.GDB.Create(&role).Error; err != nil {
		response := model.NewResponse(model.MessageTypeError, "Internal Server Error", nil)
		model.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := model.NewResponse(model.MessageTypeSuccess, "Role created successfusly", nil)
	model.ResponseJSON(w, http.StatusCreated, response)
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := model.NewResponse(model.MessageTypeError, "Method put not permit", nil)
		model.ResponseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	role := model.Role{}
	err := db.GDB.First(&role, id)
	if err != nil {
		response := model.NewResponse(model.MessageTypeError, "Role not found", nil)
		model.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		response := model.NewResponse(model.MessageTypeError, "Error decoding request body", nil)
		model.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	db.GDB.Save(&role)
	response := model.NewResponse(model.MessageTypeError, "Role updated successfully", role)
	model.ResponseJSON(w, http.StatusNotFound, response)

}

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := model.NewResponse(model.MessageTypeError, "Method delete not permit", nil)
		model.ResponseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	role := model.Role{}
	if err := db.GDB.First(&role, id); err != nil {
		response := model.NewResponse(model.MessageTypeError, "Role not found", nil)
		model.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	db.GDB.Delete(&role)
	response := model.NewResponse(model.MessageTypeSuccess, "Role deleted successfull", nil)
	model.ResponseJSON(w, http.StatusOK, response)
}
