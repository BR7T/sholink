package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/BR7T/sholink/database/functions"
)

type shortResponse struct{
	OriginalURL string `json:"url"`
}

func main(){
	r := chi.NewRouter()
	// necessário fazer middleware

	r.Get("/" , func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bem vindo ao Sholink"))
	})
	r.Get("/{nanoid}" , func(w http.ResponseWriter, r *http.Request) {
		nanoId := chi.URLParam(r,"nanoid")

		originalLink := functions.FindLinkByNano(nanoId)
		if originalLink == ""{
			fmt.Fprint(w , "Não foi possível achar o seu link")
			return
		}
		http.Redirect(w,r,originalLink,http.StatusFound)
	})

	r.Post("/" , func(w http.ResponseWriter, r *http.Request) {
		var req shortResponse
		if err := json.NewDecoder(r.Body).Decode(&req);err != nil{
			http.Error(w , "JSON Inválido" , http.StatusBadRequest)
			return

		}
		if req.OriginalURL == ""{
			http.Error(w , "URL obrigtaória" , http.StatusBadRequest)
			return
		}
		
		result , nano , err := functions.InsertLink(req.OriginalURL)
		if err != nil{
			fmt.Fprintf(w,"Não foi possível inserir o link\nError:%v",err)
		}
		if result{
			fmt.Printf("Sucesso! Url encurtada:%v","localhost:8000/"+nano)
		}

	})
	
	
	fmt.Println("✅ Servidor iniciado na porta 8000")
	http.ListenAndServe(":8000" , r)
}