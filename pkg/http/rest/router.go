package rest

import (
	"encoding/json"
	_ "github.com/samuael/Project/CarInspection/api"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Route returns an http handler for the api.
func Route( adminhandler IAdminHandler ) http.Handler {
	router := httprouter.New()


	router.POST("/api/admin/login/", adminhandler.AdminLogin  )

	http.ListenAndServe(":8080", router)
	return router
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func respHandler(h func(http.ResponseWriter, *http.Request) (interface{}, int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status, err := h(w, r)
		if err != nil {
			data = struct {
				Error string `json: "error"`
			}{
				Error: err.Error(),
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Printf("could not encode response to output: %v", err)
		}
	}
}
