package api

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

type DecodeRequestReq struct {
	AminoEncodedTx string `json:"amino_encoded_tx"`
}

type BatchDecodeRequestReq struct {
	AminoEncodedTx []string `json:"amino_encoded_tx"`
}

// Marshal - nolint
func (sb DecodeRequestReq) Marshal() []byte {
	out, err := json.Marshal(sb)
	if err != nil {
		panic(err)
	}
	return out
}

// DecodeHandler handles the /decode route
func (s *Server) DecodeTxHandler(w http.ResponseWriter, r *http.Request) {
	var req DecodeRequestReq

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = cdc.UnmarshalJSON(body, &req)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	txBytes, err := base64.StdEncoding.DecodeString(req.AminoEncodedTx)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var stdTx auth.StdTx
	err = cdc.UnmarshalBinaryLengthPrefixed(txBytes, &stdTx)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	bz, err := cdc.MarshalJSON(stdTx)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(bz)
}

// DecodeHandler handles the /batch_decode route
func (s *Server) BatchDecodeTxHandler(w http.ResponseWriter, r *http.Request) {
	var req BatchDecodeRequestReq

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = cdc.UnmarshalJSON(body, &req)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var stdTxs = make([]auth.StdTx, 0, 16)
	for _, tx := range req.AminoEncodedTx {
		txBytes, err := base64.StdEncoding.DecodeString(tx)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var stdTx auth.StdTx
		err = cdc.UnmarshalBinaryLengthPrefixed(txBytes, &stdTx)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		stdTxs = append(stdTxs, stdTx)
	}

	bz, err := cdc.MarshalJSON(stdTxs)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(bz)
}
