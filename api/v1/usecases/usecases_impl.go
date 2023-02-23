package usecases

import (
	"context"
	"log"

	"github.com/RyaWcksn/chibiko/entities"
	"github.com/RyaWcksn/chibiko/forms"
	"github.com/RyaWcksn/chibiko/pkgs/encryptions"
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
		return "", nil
	}

	res := encryptions.Encode(sqlResp)
	return res, nil

}
