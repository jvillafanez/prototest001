// Code generated by protoc-gen-microweb. DO NOT EDIT.
// source: v1.proto

package v1

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/golang/protobuf/jsonpb"
)

type webConfigServiceHandler struct {
	r chi.Router
	h ConfigServiceHandler
}

func (h *webConfigServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func (h *webConfigServiceHandler) GetConfig(w http.ResponseWriter, r *http.Request) {

	req := &GetConfigRequest{}

	resp := &GetConfigResponse{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.GetConfig(
		r.Context(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func RegisterConfigServiceWeb(r chi.Router, i ConfigServiceHandler, middlewares ...func(http.Handler) http.Handler) {
	handler := &webConfigServiceHandler{
		r: r,
		h: i,
	}

	r.MethodFunc("POST", "/api/v1/config/get-config", handler.GetConfig)
}

// GetConfigRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of GetConfigRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetConfigRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *GetConfigRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := GetConfigRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*GetConfigRequest)(nil)

// GetConfigRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of GetConfigRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetConfigRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *GetConfigRequest) UnmarshalJSON(b []byte) error {
	return GetConfigRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*GetConfigRequest)(nil)

// GetConfigResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of GetConfigResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetConfigResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *GetConfigResponse) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := GetConfigResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*GetConfigResponse)(nil)

// GetConfigResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of GetConfigResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetConfigResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *GetConfigResponse) UnmarshalJSON(b []byte) error {
	return GetConfigResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*GetConfigResponse)(nil)