package router

import (
	"JumpCat-Server/internal/handler"
	"JumpCat-Server/internal/util"
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/create-room", methodHandler(handler.CreateRoom, http.MethodPost))
	mux.HandleFunc("/get-room", methodHandler(handler.GetRoom, http.MethodGet))
	mux.HandleFunc("/update-status", methodHandler(handler.SubmitPrepare, http.MethodPost))
	mux.HandleFunc("/join-room",methodHandler(handler.JoinRoom,http.MethodPost))

	mux.HandleFunc("/game-connect", handler.GameConnectHandler)

	return mux
}

// methodHandler HTTP 方法处理
func methodHandler(h http.HandlerFunc, allowedMethods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, method := range allowedMethods {
			if r.Method == method {
				h.ServeHTTP(w, r)
				return
			}
		}
		util.WriteResponse(w, http.StatusMethodNotAllowed, nil)
	}
}
