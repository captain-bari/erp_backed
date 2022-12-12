package rest

import (
	auth "erp/Auth"
	"erp/postgres"
	types "erp/types"

	"encoding/json"
	log "erp/log"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func login(w http.ResponseWriter, req *http.Request) {
	loginReq := types.LoginReq{}
	err := json.NewDecoder(req.Body).Decode(&loginReq)
	if err != nil {
		log.Logger.Errorf("Login : RequestUnmarshallError: %s", err.Error())
		WriteCustomErrorResp(w, req, "", RequestUnmarshallError)
		return
	}

	log.Logger.Debugf("login: Request[%+v]", loginReq)

	notFound, name, role, err := auth.AuthenticateUser(loginReq.UserID, loginReq.UserHash)
	if err != nil {
		log.Logger.Errorf("Login : auth.AuthenticateUser: %s", err.Error())
		WriteCustomErrorResp(w, req, err.Error(), PostgresError)
		return
	}

	if notFound {
		WriteCustomErrorResp(w, req, err.Error(), AuthcontextNotFound)
		return
	}

	tokken, err := auth.GenerateJWT(loginReq.UserID, role)
	if err != nil {
		log.Logger.Errorf("Login|GenerateJWT : ERROR:[%s]", err.Error())
		WriteCustomErrorResp(w, req, err.Error(), InternalError)
		return
	}

	WriteSuccessMessage(w, req, http.StatusOK, types.LoginResp{
		UserRole: role,
		UserName: name,
		Tokken:   tokken,
	})
}

func getMaterials(w http.ResponseWriter, req *http.Request) {
	tokkenStr := req.Header.Get("tokken")
	if tokkenStr == "" {
		log.Logger.Errorf("REST: Tokken not found")
		WriteCustomErrorResp(w, req, "", TokkenNotFound)
		return
	}
	valid, _, err := auth.CheckJwt(tokkenStr)
	if err != nil {
		log.Logger.Errorf("REST|AUTH|CheckJwt: Error while checking tokken [%s]", err.Error())
		WriteCustomErrorResp(w, req, "", AuthcontextNotFound)
		return
	}

	if !valid {
		log.Logger.Infof("REST|AUTH|CheckJwt: Invalid tokken:[%s]", tokkenStr)
		WriteCustomErrorResp(w, req, "", AuthcontextNotFound)
		return
	}

	materials, err := postgres.GetMaterialsDB()
	if err != nil {
		log.Logger.Errorf("REST|POSTGRES|GetMaterialsDB: ERROR[%s]", err.Error())
		WriteCustomErrorResp(w, req, "", InternalError)
		return
	}

	WriteSuccessMessage(w, req, http.StatusOK, types.GetMaterialsResp{
		Materials: materials,
	})
}
