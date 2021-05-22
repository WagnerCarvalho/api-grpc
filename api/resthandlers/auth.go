package resthandlers

import (
	"api-grpc/api/restutils"
	"api-grpc/pb"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type AuthHandlers interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	PutUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type authHandlers struct {
	authSvcClient pb.AuthServiceClient
}

func NewAuthHandlers(authSvcClient pb.AuthServiceClient) AuthHandlers {
	return &authHandlers{authSvcClient: authSvcClient}
}

func (h *authHandlers) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutils.WriteError(w, http.StatusBadRequest, restutils.ErrEmptyBody)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user := new(pb.User)
	err = json.Unmarshal(body, user)
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user.Created = time.Now().Unix()
	user.Updated = user.Created
	user.Id = bson.NewObjectId().Hex()
	resp, err := h.authSvcClient.SignUp(r.Context(), user)
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	restutils.WriteAsJson(w, http.StatusCreated, resp)
}

func (h *authHandlers) PutUser(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutils.WriteError(w, http.StatusBadRequest, restutils.ErrEmptyBody)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user := new(pb.User)
	err = json.Unmarshal(body, user)
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	vars := mux.Vars(r)
	user.Id = vars["id"]
	resp, err := h.authSvcClient.UpdateUser(r.Context(), user)
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	restutils.WriteAsJson(w, http.StatusOK, resp)
}

func (h *authHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resp, err := h.authSvcClient.GetUser(r.Context(), &pb.GetUserRequest{Id: vars["id"]})
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	restutils.WriteAsJson(w, http.StatusOK, resp)
}

func (h *authHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Start GetUsers")
	stream, err := h.authSvcClient.ListUsers(r.Context(), &pb.ListUserRequest{})
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var users []*pb.User

	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			restutils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		users = append(users, user)
	}
	restutils.WriteAsJson(w, http.StatusOK, users)
}

func (h *authHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resp, err := h.authSvcClient.DeleteUser(r.Context(), &pb.GetUserRequest{Id: vars["id"]})
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", resp.Id)
	restutils.WriteAsJson(w, http.StatusNoContent, nil)
}
