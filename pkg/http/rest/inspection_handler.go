package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
	"github.com/samuael/Project/CarInspection/pkg/adding"
	"github.com/samuael/Project/CarInspection/pkg/deleting"
	"github.com/samuael/Project/CarInspection/pkg/http/rest/auth"
	"github.com/samuael/Project/CarInspection/pkg/listing"
	"github.com/samuael/Project/CarInspection/pkg/userpolicy"
)

// PostHandler provides access to Post api methods.
type InspectionHandler interface {
	// CreateInspaction( response http.ResponseWriter,request *http.Request)
	GetInspections(w http.ResponseWriter, r *http.Request  , params httprouter.Params )
	AddInspection(w http.ResponseWriter, r *http.Request  , params httprouter.Params)
	DeleteInspection(w http.ResponseWriter, r *http.Request , params httprouter.Params)
	EditInspection(w http.ResponseWriter, r *http.Request , params httprouter.Params)
}

type inspectionHandler struct {
	l listing.Service
	a adding.Service
	d deleting.Service
	u userpolicy.Service
}

// 
func NewInspectionHandler(l listing.Service, a adding.Service, d deleting.Service, u userpolicy.Service) InspectionHandler {
	return &inspectionHandler{
		l: l,
		a: a,
		d: d,
		u: u,
	}
}

func (h inspectionHandler) GetInspections(w http.ResponseWriter, r *http.Request , params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	list, err := h.l.GetMyInspections(9)
	if err != nil {
		http.Error(w, "Failed to get posts", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(list)
}

// AddPost handler for POST /api/post requests
func (h inspectionHandler) AddInspection(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var post adding.Inspection

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&post); err != nil {
		http.Error(w, "Failed to parse post", http.StatusBadRequest)
		return
	}

	credentials := r.Context().Value("credentials").(*auth.AppClaims)
	post.AuthorID = credentials.ID

	if err := h.a.AddInspection(post); err != nil {
		http.Error(w, "Failed to add post", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("New post added.")
}

func (h inspectionHandler) DeleteInspection(w http.ResponseWriter, r *http.Request , params httprouter.Params ) {

	postID, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil || postID == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	credentials := r.Context().Value("credentials").(*auth.AppClaims)
	userID := credentials.ID

	if allowed := h.u.IsOwnerOfPost(userID, uint(postID)); allowed == false {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if err := h.d.DeleteInspection(uint(postID)); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("post deleted.")
}

func (h inspectionHandler) EditInspection(w http.ResponseWriter, r *http.Request  , params httprouter.Params) {
	value := params.ByName("id")
	if value =="" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}
