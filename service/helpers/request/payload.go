package request

import (
	"net/http"

	"github.com/DeniesKresna/gohelper/utint"
	"github.com/gorilla/mux"
)

func GetVar(r *http.Request, key string, def ...string) (res string) {
	vars := mux.Vars(r)
	res, ok := vars[key]
	if !ok {
		if len(def) > 0 {
			return def[0]
		}
	}
	return
}

func GetInt64Var(r *http.Request, key string, def ...int64) (res int64) {
	vars := mux.Vars(r)
	resStr, ok := vars[key]
	if !ok {
		if len(def) > 0 {
			return def[0]
		}
		return
	}
	return utint.Convert64FromString(resStr, res)
}

func GetIntVar(r *http.Request, key string, def ...int) (res int) {
	vars := mux.Vars(r)
	resStr, ok := vars[key]
	if !ok {
		if len(def) > 0 {
			return def[0]
		}
		return
	}
	return utint.ConvertFromString(resStr, res)
}
