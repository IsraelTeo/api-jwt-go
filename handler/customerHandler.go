package handler

import (
	"encoding/json"
	"net/http"

	"github.com/IsraelTeo/api-jwt-go/db"
	"github.com/IsraelTeo/api-jwt-go/model"
	"github.com/gorilla/mux"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := model.NewResponse(model.MessageTypeError, "Invalid Method", nil)
		model.ResponseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	user := model.User{}
	if err := db.GDB.First(&user, id); err != nil {
		response := model.NewResponse(model.MessageTypeError, "User was not found", nil)
		model.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	response := model.NewResponse(model.MessageTypeSuccess, "User found", user)
	model.ResponseJSON(w, http.StatusOK, response)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := model.NewResponse(model.MessageTypeError, "Method get not permit", nil)
		model.ResponseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	var users []model.User
	if err := db.GDB.Find(&users); err != nil {
		response := model.NewResponse(model.MessageTypeError, "Users not found", nil)
		model.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	if len(users) == 0 {
		response := model.NewResponse(model.MessageTypeSuccess, "Users List empty", nil)
		model.ResponseJSON(w, http.StatusNoContent, response)
		return
	}

	response := model.NewResponse(model.MessageTypeSuccess, "Users found", users)
	model.ResponseJSON(w, http.StatusNoContent, response)
}

func SaveUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := model.NewResponse(model.MessageTypeError, "Method post not permit", nil)
		model.ResponseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	user := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response := model.NewResponse(model.MessageTypeError, "Bad request: invalid JSON data", nil)
		model.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	if err := db.GDB.Create(&user); err != nil {
		response := model.NewResponse(model.MessageTypeError, "Internal Server Error", nil)
		model.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := model.NewResponse(model.MessageTypeSuccess, "User created successfusly", nil)
	model.ResponseJSON(w, http.StatusCreated, response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := model.NewResponse(model.MessageTypeError, "Method put not permit", nil)
		model.ResponseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	user := model.User{}
	if err := db.GDB.First(&user, id); err != nil {
		response := model.NewResponse(model.MessageTypeError, "User not found", nil)
		model.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response := model.NewResponse(model.MessageTypeError, "Bad request", nil)
		model.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	db.GDB.Save(&user)
	response := model.NewResponse(model.MessageTypeSuccess, "User updated successfull", user)
	model.ResponseJSON(w, http.StatusOK, response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := model.NewResponse(model.MessageTypeError, "Method delete not permit", nil)
		model.ResponseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	role := model.Role{}
	if err := db.GDB.First(&role, id); err != nil {
		response := model.NewResponse(model.MessageTypeError, "User not found", nil)
		model.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	db.GDB.Delete(&role)
	response := model.NewResponse(model.MessageTypeSuccess, "User deleted successfull", nil)
	model.ResponseJSON(w, http.StatusOK, response)
}
