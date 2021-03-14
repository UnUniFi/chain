package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers auction-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
    r.HandleFunc("/auction/auctions/{id}", getAuctionHandler(clientCtx)).Methods("GET")
    r.HandleFunc("/auction/auctions", listAuctionHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
    r.HandleFunc("/auction/auctions", createAuctionHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/auction/auctions/{id}", updateAuctionHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/auction/auctions/{id}", deleteAuctionHandler(clientCtx)).Methods("POST")

}
