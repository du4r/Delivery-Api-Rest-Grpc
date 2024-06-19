package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"mega_api/models"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

// @Summary Create 
// @Description Cria um novo usuario
// @Tags Create costumer
// @Accept  json
// @Produce  json
// @Param costumer body models.Costumer true "Costumer"
// @Success 201 {object} models.Costumer
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /costumer [post]
func Create(w http.ResponseWriter, r *http.Request){
	var costumer models.Costumer

	err := json.NewDecoder(r.Body).Decode(&costumer)

	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(costumer)

	//var resp map[string]any
	var resp int

	if err != nil {
		if id == -1{
			fmt.Print(err)
			resp = http.StatusBadRequest
			http.Error(w,http.StatusText(http.StatusBadRequest),http.StatusBadRequest)
		}else{
			resp = http.StatusInternalServerError
			http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		}
	}else {
		resp = http.StatusCreated
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// @Summary Get todos
// @Description retorna uma lista com todos os usuarios cadastrados
// @Tags List costumers
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Costumer
// @Router /costumer [get]
func List(w http.ResponseWriter, r *http.Request){
	costumers, err := models.GetAllCostumers() 
	if err != nil {
		log.Printf("Erro ao obter registros >  %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(costumers)
}

// @Summary Get por id
// @Description retorna apenas um cliente pelo parametro passado do id
// @Tags get costumer by id
// @Accept  json
// @Produce  json
// @Param id path int true "Costumer ID"
// @Success 200 {object} models.Costumer
// @Failure 404 {object} string "Not Found"
// @Router /costumer/{id} [get]
func Get(w http.ResponseWriter ,r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r,"id"))
	if err != nil {
		log.Printf("erro ao fazer parse do ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	costumer, err := models.GetCostumer(int64(id))
	if err != nil {
		log.Printf("Erro ao Obter o Registro:  %v", err)
		http.Error(w,http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(costumer)
} 

// @Summary Update
// @Description atualiza de cliente no banco de dados
// @Tags update costumer
// @Accept  json
// @Produce  json
// @Param id path int true "Costumer ID"
// @Param costumer body models.Costumer true "Costumer"
// @Success 200 {object} models.Costumer
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /costumer/{id} [put]
func Update(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(chi.URLParam(r,"id"))
	if err != nil {
		log.Printf("erro ao fazer parse do ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var costumer models.Costumer

	err = json.NewDecoder(r.Body).Decode(&costumer) 
	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	
	rows, err := models.Update(int64(id), costumer)
	if err != nil {
		log.Printf("Erro ao atualizar registro:  %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram atualizados %d registros", rows)
	}

	resp := map[string]any{
		"Error": false,
		"Message": "dados atualizados com sucesso!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// @Summary Delete
// @Description deleta um cliente por id
// @Tags costumer delete
// @Accept  json
// @Produce  json
// @Param id path int true "Costumer ID"
// @Success 200 {object} string "OK"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /costumer/{id} [delete]
func Delete(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(chi.URLParam(r,"id"))
	if err != nil {
		log.Printf("erro ao fazer parse do ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao atualizar registro:  %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram removidos %d registros", rows)
	}

	resp := map[string]any{
		"Error": false,
		"Message": "registro removido com sucesso!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}