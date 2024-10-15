package controllers

import (
	"compartilhatech/internal/application/dto"
	"compartilhatech/internal/application/service_interface"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type regrasVencimentoController struct {
	service service_interface.RegraService
}

func NewRegraController(mux *http.ServeMux, s service_interface.RegraService) *regrasVencimentoController {
	c := regrasVencimentoController{
		service: s,
	}
	mux.HandleFunc("POST /regrasVencimento", c.handleCreate)
	mux.HandleFunc("GET /regrasVencimento", c.handleList)
	mux.HandleFunc("GET /regrasVencimento/{regrasVencimentoID}", c.handleGetById)
	mux.HandleFunc("DELETE /regrasVencimento/{regrasVencimentoID}", c.handleDelete)
	mux.HandleFunc("PATCH /regrasVencimento/{regrasVencimentoID}", c.handleUpdate)

	return &c
}

func (c *regrasVencimentoController) handleCreate(w http.ResponseWriter, r *http.Request) {
	payload := dto.CreateRegra{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	regrasVencimento, err := c.service.InsertRegra(payload)
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responseJson(w, http.StatusCreated, regrasVencimento)
}

func (c *regrasVencimentoController) handleList(w http.ResponseWriter, r *http.Request) {
	regrasVencimentos, err := c.service.ListRegra()
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseJson(w, http.StatusOK, regrasVencimentos)
}

func (c *regrasVencimentoController) handleGetById(w http.ResponseWriter, r *http.Request) {
	CodvencStr := r.PathValue("regrasVencimentoID")
     
	CodvencInt, err := strconv.Atoi(CodvencStr)
	if err != nil {
		fmt.Println("Erro ao converter RegrasVencimentoID para int", err)
		w.WriteHeader(http.StatusBadRequest)  // Retorna um erro 400 Bad Request
        return
	}

	Codvenc := int32(CodvencInt)

	regrasVcencimento, err := c.service.GetByIdRegra(Codvenc)
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if regrasVcencimento == nil {
		fmt.Println("regrasVcencimento not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responseJson(w, http.StatusOK, regrasVcencimento)
}

func (c *regrasVencimentoController) handleUpdate(w http.ResponseWriter, r *http.Request) {
	payload := dto.UpdateRegra{}

	CodvencStr := r.PathValue("regrasVencimentoID")

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	CodvencInt, err := strconv.Atoi(CodvencStr)
	if err != nil {
		fmt.Println("Erro ao converter RegrasVencimentoID para int", err)
		w.WriteHeader(http.StatusBadRequest)  // Retorna um erro 400 Bad Request
        return
	}

	Codvenc := int32(CodvencInt)

	regrasVencimento, err := c.service.UpdateRegra(Codvenc, payload)
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responseJson(w, http.StatusOK, regrasVencimento)
}

func (c *regrasVencimentoController) handleDelete(w http.ResponseWriter, r *http.Request) {
	CodvencStr := r.PathValue("regrasVencimentoID")

	CodvencInt, err := strconv.Atoi(CodvencStr)
	if err != nil {
		fmt.Println("Erro ao converter RegrasVencimentoID para int", err)
		w.WriteHeader(http.StatusBadRequest)  // Retorna um erro 400 Bad Request
        return
	}

	Codvenc := int32(CodvencInt)

	err = c.service.DeleteRegra(Codvenc)
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responseJson(w, http.StatusNoContent, nil)
}

/*func responseJsonRegra(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(v)
}*/
