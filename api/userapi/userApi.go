package userapi

import (
	"DemoProject/entities"
	"DemoProject/models"
	"encoding/json"
	"net/http"
	"fmt"
)

func FindUser(response http.ResponseWriter, request *http.Request){
	ids, ok := request.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "Url Param is missing")
		return
	}
	user, err := models.FindUser(ids[0])
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
		return
	}
	responseWithJSON(response, http.StatusOK, user)
}

func GetAll(response http.ResponseWriter, request *http.Request){
	users := models.GetAllUser()
	responseWithJSON(response, http.StatusOK, users)
}

func CreateUser(response http.ResponseWriter, request *http.Request){
	var user entities.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil{
		responseWithError(response, http.StatusBadRequest, err.Error())
	}else{
		result := models.CreateUser(&user)
		if !result {
			responseWithError(response, http.StatusBadRequest, err.Error())
			return
		}
		responseWithJSON(response, http.StatusOK, user)
	}
}
func UpdateUser(response http.ResponseWriter, request *http.Request) {
	var user entities.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		result := models.UpdateUser(&user)
		if !result {
			responseWithError(response, http.StatusBadRequest, "Couldn't update user")
			return
		}
		responseWithJSON(response, http.StatusOK, "Update user successfully")
	}
}
func DeleteUser(response http.ResponseWriter, request *http.Request){
	ids, ok := request.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "Urt Param id is missing")
		return
	}
	result := models.DeleteUser(ids[0])
	if !result {
		responseWithError(response, http.StatusBadRequest, "Could not delete this user")
		return
	}
	responseWithJSON(response, http.StatusOK, "Delete is okay")
}

func responseWithError(response http.ResponseWriter, statusCode int, msg string){
	responseWithJSON(response, statusCode, map[string]string{
		"error":msg,
	})
}

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}){
	 result, _ := json.Marshal(data)
	 response.Header().Set("Content-Type", "application/json")
	 response.WriteHeader(statusCode)
	 response.Write(result)
}

func TestThu(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "welcom to port 5000")
}



