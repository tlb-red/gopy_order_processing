package service

import (
	"context"

	claim_model "tlb-red.com/gopy-order-processing/internal/models"
	claim_queue "tlb-red.com/gopy-order-processing/internal/repository"
)

type ClaimService struct {
	queue claim_queue.ClaimQueue
}

func NewClaimService(queue claim_queue.ClaimQueue) ClaimService {
	return ClaimService{queue: queue}
}

func (s ClaimService) Enqueue(ctx context.Context, claim claim_model.Claim) {

}
