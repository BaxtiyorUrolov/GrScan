package service

import (
	"context"
	"grscan/api/models"
	"grscan/pkg/logger"
	"grscan/storage"
)

type registerService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewRegisterService(storage storage.IStorage, log logger.ILogger) registerService {
	return registerService{
		storage: storage,
		log:     log,
	}
}

func (r registerService) Verify(ctx context.Context, verify models.CreateRegister) (bool, error) {
	r.log.Info("category create service layer", logger.Any("verify", verify))


	code, err := r.storage.Register().GetByID(ctx, verify.Phone)
	if err != nil {
		r.log.Error("ERROR in service layer while get code by phone", logger.Error(err))
		return false, err
	}

	if code.Code == verify.Code {
		err := r.storage.Register().UpdateStatus(ctx, verify.Phone)
		if err != nil {
			r.log.Error("ERROR in service layer while updating verify status by phone", logger.Error(err))
			return false, err
		}
	}

	return true, nil
}
