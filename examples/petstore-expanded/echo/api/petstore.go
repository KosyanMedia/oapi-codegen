//go:generate go run github.com/KosyanMedia/oapi-codegen/v2/cmd/oapi-codegen --config=types.cfg.yaml ../../petstore-expanded.yaml
//go:generate go run github.com/KosyanMedia/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yaml ../../petstore-expanded.yaml

package api

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

type PetStore struct {
	Pets   map[int64]Pet
	NextId int64
	Lock   sync.Mutex
}

func NewPetStore() *PetStore {
	return &PetStore{
		Pets:   make(map[int64]Pet),
		NextId: 1000,
	}
}

// This function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendPetstoreError(ctx echo.Context, code int, message string) error {
	petErr := Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, petErr)
	return err
}

// Here, we implement all of the handlers in the ServerInterface
func (p *PetStore) FindPets(ctx echo.Context, params FindPetsParams) (*FindPetsResponse, error) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	result := make([]Pet, 0)

	for _, pet := range p.Pets {
		if params.Tags != nil {
			// If we have tags,  filter pets by tag
			for _, t := range params.Tags {
				if pet.Tag != nil && (*pet.Tag == t) {
					result = append(result, pet)
				}
			}
		} else {
			// Add all pets if we're not filtering
			result = append(result, pet)
		}

		if params.Limit != nil {
			l := int(*params.Limit)
			if len(result) >= l {
				// We're at the limit
				break
			}
		}
	}
	return &FindPetsResponse{
		JSON200: result,
	}, nil
}

func (p *PetStore) AddPet(ctx echo.Context, requestBody AddPetJSONBody) (*AddPetResponse, error) {
	newPet := NewPet(requestBody)

	// We're always asynchronous, so lock unsafe operations below
	p.Lock.Lock()
	defer p.Lock.Unlock()

	// We handle pets, not NewPets, which have an additional ID field
	var pet Pet
	pet.Name = newPet.Name
	pet.Tag = newPet.Tag
	pet.Id = p.NextId
	p.NextId = p.NextId + 1

	// Insert into map
	p.Pets[pet.Id] = pet

	// Now, we have to return the NewPet
	return &AddPetResponse{
		JSON201: &pet,
	}, nil
}

func (p *PetStore) FindPetByID(ctx echo.Context, petId int64) (*FindPetByIDResponse, error) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	pet, found := p.Pets[petId]
	if !found {
		return &FindPetByIDResponse{
			Code: http.StatusNotFound,
			JSONDefault: &Error{
				Code:    int32(http.StatusNotFound),
				Message: fmt.Sprintf("Could not find pet with ID %d", petId),
			},
		}, nil
	}

	return &FindPetByIDResponse{
		JSON200: &pet,
	}, nil
}

func (p *PetStore) DeletePet(ctx echo.Context, petId int64) (resp *DeletePetResponse, err error) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	_, found := p.Pets[petId]
	if !found {
		return &DeletePetResponse{
			Code: http.StatusNotFound,
			JSONDefault: &Error{
				Code:    int32(http.StatusNotFound),
				Message: fmt.Sprintf("Could not find pet with ID %d", petId),
			},
		}, nil
	}
	delete(p.Pets, petId)
	return &DeletePetResponse{
		Code: http.StatusNoContent,
	}, nil
}
