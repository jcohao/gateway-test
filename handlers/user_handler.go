package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gateway-test/dao"
	"gateway-test/models"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Vars 只能读取路径中的参数
	// vars := mux.Vars(r)

	// fmt.Printf("vars: %v\n", vars)

	// age, _ := strconv.Atoi(vars["age"])
	// height, _ := strconv.Atoi(vars["height"])

	// user := &models.User{
	// 	Code:    vars["code"],
	// 	Age:     age,
	// 	Name:    vars["name"],
	// 	Height:  height,
	// 	Address: vars["address"],
	// }

	// jsonData, _ := json.Marshal(user)

	// fmt.Printf("user: %v\n", string(jsonData))
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	user := models.User{}
	fmt.Printf("body: %v\n", string(body))
	json.Unmarshal(body, &user)
	fmt.Printf("user: %v\n", user)

	dao.CreateUser(&user)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	users, _ := dao.GetUsers()
	data, _ := json.Marshal(users)
	fmt.Fprintf(w, "Users: %v\n", string(data))
}
