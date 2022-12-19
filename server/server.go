package main

import (
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

type Cotacao struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {

	http.HandleFunc("/cotacao", BuscaCotacaoHandler)

	http.ListenAndServe("localhost:8080", nil)

}

func BuscaCotacaoHandler(w http.ResponseWriter, r *http.Request) {

	ctxAPI, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	ctxDB, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	cotacao, error := BuscaCotacao(ctxAPI, "USD-BRL")
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	// retorna a cotacao para o cliente
	RetornaCotacao(w, r, cotacao)

	// armazenar no banco de dados a cotação
	PreparaDB := AbreConexaoDB()
	err := insertCotacao(PreparaDB, cotacao.Usdbrl.Code+"-"+cotacao.Usdbrl.Codein, cotacao.Usdbrl.Bid, ctxDB)
	if err != nil {
		panic(err)
	}
}

func RetornaCotacao(w http.ResponseWriter, r *http.Request, cotacao Cotacao) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// devolver apenas o bid
	err := json.NewEncoder(w).Encode(cotacao.Usdbrl.Bid)
	if err != nil {
		panic(err)
	}
}

func BuscaCotacao(ctx context.Context, moeda string) (Cotacao, error) {
	var cotacao Cotacao
	url := "https://economia.awesomeapi.com.br/json/last/" + moeda

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return cotacao, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return cotacao, err
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&cotacao); err != nil {
		return cotacao, err
	}
	return cotacao, nil
}

func AbreConexaoDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	return db
}

func insertCotacao(db *sql.DB, moeda string, valorcotacao string, ctx context.Context) error {

	stmt, err := db.Prepare("INSERT INTO Cotacao (moeda, valorcotacao, datahora) values (?, ?, ?)")
	if err != nil {
		return err
	}
	defer db.Close()
	defer stmt.Close()

	stmt.Exec(moeda, valorcotacao, time.Now())

	return err
}
