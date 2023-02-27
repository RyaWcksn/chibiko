package usecases

import (
	"context"
	"log"

	"github.com/RyaWcksn/chibiko/entities"
	"github.com/RyaWcksn/chibiko/forms"
	"github.com/RyaWcksn/chibiko/pkgs/encryptions"
	"github.com/RyaWcksn/chibiko/pkgs/errors"
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
		return "", ierror.GetError(ierror.InternalServer, err)
	}

	res := encryptions.Encode(sqlResp)
	return res, nil

}

// Decode implements IUsecase
func (uc *UsecaseImpl) Decode(ctx context.Context, payload *forms.DecodePayload) (resp string, err error) {
	decode, err := encryptions.Decode(payload.Param)
	if err != nil {
		log.Printf("error := %v", err)
		return "", errors.GetError(errors.InternalServer, err)
	}

	sqlPayload := entities.GetDatabase{
		Id: int(decode),
	}

	sqlRes, err := uc.dbPort.Get(ctx, &sqlPayload)
	if err != nil {
		return "", err
	}

	return sqlRes, nil
}
