package usecases

import (
	"context"
	"errors"
	"log"

	"github.com/RyaWcksn/chibiko/entities"
	"github.com/RyaWcksn/chibiko/forms"
	"github.com/RyaWcksn/chibiko/pkgs/encryptions"
	ierror "github.com/RyaWcksn/chibiko/pkgs/errors"
)

// Encode implements IUsecase
func (uc *UsecaseImpl) Encode(ctx context.Context, payload *forms.EncodeRequest) (resp string, err error) {

	sqlPayload := entities.SaveDatabase{
		Url: payload.Url,
	}

	// TODO if temporary insert to redis

	sqlResp, err := uc.dbPort.Save(ctx, &sqlPayload)
	if err != nil {
		log.Printf("error := %v", err)
		return "", errors.New(ierror.InternalError)
	}

	res := encryptions.Encode(sqlResp)
	return res, nil

}
