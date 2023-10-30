package usecase

import (
	"math/rand"

	"github.com/uzimbra/orch-process-transactions/internal/entity"
)

func ProcessTransactionUseCase(t *entity.Transaction) (*entity.Transaction, error) {
	x := rand.Intn(2)

	if x == 0 {
		t.Status = string(entity.Refused)
		return t, nil
	}

	t.Status = string(entity.Approved)
	return t, nil
}
