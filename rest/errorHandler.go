package rest

import (
	"encoding/json"
	log "erp/log"
	"fmt"
	"net/http"
)

type ErrorType uint32

type errorResponse struct {
	Error            string `json:"error"`
	CustomErrorCode  string `json:"error_code"`
	ErrorDescription string `json:"error_description,omitempty"`
	ErrorURI         string `json:"error_uri,omitempty"`
	ErrorCode        int    `json:"-"`
}

const (
	ErrorInvalid ErrorType = iota
	InternalError
	TokkenNotFound
	RequestReadError
	AuthcontextNotFound
	RequestUnmarshallError
	DeviceSerialNotFound
	IpOrFqdnNotFound
	Edge_not_ready
	Edge_not_compatible
	PostgresError
	FileDownloadError
	NATsError
	NxMsgMarshalFailed
	DiagServiceAlreadyRunning
	TxnIdNotFound
)

var errorMap map[ErrorType]errorResponse = map[ErrorType]errorResponse{

	RequestReadError: {
		ErrorCode:        http.StatusBadRequest,
		Error:            "invalid_request",
		ErrorDescription: "The request is malformed, a required parameter is missing or a parameter has an invalid value :  Request body read error",
	},

	TokkenNotFound: {
		ErrorCode:        http.StatusBadRequest,
		Error:            "tokken_not_found",
		ErrorDescription: "The request is malformed, tokken not found",
	},

	InternalError: {
		ErrorCode:        http.StatusBadRequest,
		Error:            "internal_server_error",
		ErrorDescription: "Something went wrong in server system",
	},

	Edge_not_ready: {
		ErrorCode:        http.StatusUpgradeRequired,
		Error:            "edge_not_ready",
		ErrorDescription: "No servers found for the device. Make sure edge is running seviceset 75 or above",
	},

	Edge_not_compatible: {
		ErrorCode:        http.StatusNotAcceptable,
		Error:            "Edge_not_compatible",
		ErrorDescription: "No servers found for the device",
	},

	AuthcontextNotFound: {
		ErrorCode:        http.StatusForbidden,
		Error:            "invalid_request",
		ErrorDescription: "The request is malformed: Failed to authenticate user",
	},

	RequestUnmarshallError: {
		ErrorCode:        http.StatusBadRequest,
		Error:            "invalid_request",
		ErrorDescription: "The request is malformed, a required parameter is missing or a parameter has an invalid value : Request unmarshall error",
	},

	DeviceSerialNotFound: {
		ErrorCode:        http.StatusBadRequest,
		Error:            "missing_device_serial",
		ErrorDescription: "The request is malformed: Device Serial is missing from request",
	},

	IpOrFqdnNotFound: {
		ErrorCode:        http.StatusBadRequest,
		Error:            "missing_ip_or_fqdn",
		ErrorDescription: "The request is malformed: Ip or Fqdn is missing from request",
	},

	TxnIdNotFound: {
		ErrorCode:        http.StatusBadRequest,
		Error:            "malformed_txnID",
		ErrorDescription: "The request is malformed: trxnid is missing or malformed",
	},

	FileDownloadError: {
		ErrorCode:        http.StatusInternalServerError,
		Error:            "internal_server_error",
		ErrorDescription: "File Download failed from bucket",
	},

	DiagServiceAlreadyRunning: {
		ErrorCode:        http.StatusBadRequest,
		Error:            "multiple_requests",
		ErrorDescription: "Same Diagnostic service already running for this device id",
	},

	PostgresError: {
		ErrorCode:        http.StatusInternalServerError,
		Error:            "internal_server_error",
		ErrorDescription: "Error processing response from PostgresError",
	},

	NATsError: {
		ErrorCode:        http.StatusInternalServerError,
		Error:            "internal_server_error",
		ErrorDescription: "Failed to send NATs message to device",
	},

	NxMsgMarshalFailed: {
		ErrorCode:        http.StatusBadRequest,
		Error:            "invalid_request",
		ErrorDescription: "Failed to marshal Nexus message",
	},
}

func WriteCustomErrorResp(w http.ResponseWriter, r *http.Request, errorDesc string, errorCode ErrorType) {

	respError, ok := errorMap[errorCode]
	if !ok {
		http.Error(w, "internal_server_error", http.StatusInternalServerError)
		return
	}

	if errorDesc != "" {
		respError.ErrorDescription += errorDesc
	}
	writeError(w, r, respError.ErrorCode, respError)
}

func writeError(w http.ResponseWriter, r *http.Request, code int, errresp interface{}) {

	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(code)
	b, err := json.Marshal(errresp)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Error in marshaling JSON resp err[%s]", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func WriteSuccessMessage(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Error in marshaling JSON resp err[%s]", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Logger.Debugf("WriteSuccessMessage: %+v", data)

	w.Header().Set("Content-Type", "application/json;")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(code)
	w.Write(b)
}
