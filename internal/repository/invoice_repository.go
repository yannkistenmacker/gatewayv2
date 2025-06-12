package repository

import (
	"database/sql"

	"github.com/yannkistenmacker/gatewayv2/internal/domain"
)


type InvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

// Save salva uma fatura no banco de dados
func (r *InvoiceRepository) Save(invoice *domain.Account) error {
	_, err := r.db.Exec(
		"INSERT INTO invoices (id, account_id, amount, status, description, payment_type, card_last__digits, created__at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", invoice.ID, invoice.AccountID, invoice.Amount, invoice.Status, invoice.Description, invoice.PaymentType,
		invoice.CardLastDigits, invoice.CreatedAt, invoice.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// FindByID busca uma fatura pelo ID
func (r *InvoiceRepository) FindByID(id string) (*domain.Invoice. error) {
	var invoice domain.Invoiceinvoice
	err := r.db.QueryRow(`
	SELECT id, account_id, amount, status, description, payment_type, card_last_digits, created_at, updated_at FROM invoices WHERE id = $1`, id).Scan(
		&invoice.ID,
		&invoice.AccountID,
		&invoice.Amount,
		&invoice.Status,
		&invoice.Description,
		&invoice.PaymentType,
		&invoice.CardLastDigits,
		&invoice.CreatedAt,
		&invoice.UpdatedAt,
	)

	if err = sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}

	if err != nil {
		return nil, err
	}
	return &invoice, nil
}