package useraccount

import (
	"net/http"
	"encoding/json"
	
	"go.mongodb.org/mongo-driver/mongo"
	
	"github.com/Davidmnj91/myrentals/api"
	"github.com/Davidmnj91/myrentals/validators"
)

/*NewUserAccountHandler creates a pointer to a new Useraccount Handler*/
func NewUserAccountHandler(db *mongo.Database) *Handler {
	return &Handler {
		repo: newUserAccountRepo(db),
	}
}

/*Handler to manage user endpoint*/
type Handler struct {
	repo Repo
}

/*Register function to add a useraccount to the system*/
func (userHandler *Handler) Register(w http.ResponseWriter, r *http.Request) {
	user := UserAccount{}
	json.NewDecoder(r.Body).Decode(&user)

	errors := map[string]string{}

	if (validators.IsEmpty(user.Password)) {
		errors["password"] = "Passord cannot be empty"
	}

	if (validators.IsEmpty(user.Username)) {
		errors["username"] = "Username cannot be empty"
	}

	if len(errors) > 0 {
		api.RespondwithJSON(w, 422, errors)
		return
	}

	newID, err := userHandler.repo.Create(r.Context(), &user)

	if err != nil {
		api.RespondWithError(w, 422, err.Error())
		return
	}

	api.RespondwithJSON(w, 201, newID)
}