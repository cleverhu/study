package services

import (
	"context"
	"jtthinkStudy/primary/grpcTest/pbfiles"
)

type ProdService struct {
}

func (this *ProdService) GetProd(c context.Context, in *pbfiles.ProdRequest) (*pbfiles.ProdResponse, error) {

	return nil, nil
}
