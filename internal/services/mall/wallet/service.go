package wallet

import (
	"context"
	"gobee/ent"
	"gobee/ent/wallet"
	"gobee/internal/database"
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
)

type WalletService interface {
	QueryWallet(c *fiber.Ctx, userId int) (*ent.Wallet, error)
	QueryWalletPage(c *fiber.Ctx, pageQuery model.PageQuery) (int, []*ent.Wallet, error)
	CreateWallet(c context.Context, userId int) (*ent.Wallet, error)
	UpdateWalletBalance(c context.Context, walletId int, balance int) (*ent.Wallet, error)
	FreezeWallet(c context.Context, walletId int, amount int) (*ent.Wallet, error)
	UnfreezeWallet(c context.Context, walletId int, amount int) (*ent.Wallet, error)
}

type WalletServiceImpl struct {
	client *ent.Client
}

func NewWalletServiceImpl(client *ent.Client) *WalletServiceImpl {
	return &WalletServiceImpl{client: client}
}

func (s *WalletServiceImpl) QueryWallet(c *fiber.Ctx, userId int) (*ent.Wallet, error) {
	client := database.DB
	w, err := client.Wallet.Query().
		Where(wallet.UserIDEQ(userId)).
		Only(c.Context())
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (s *WalletServiceImpl) QueryWalletPage(c *fiber.Ctx, pageQuery model.PageQuery) (int, []*ent.Wallet, error) {
	client := database.DB
	count, err := client.Wallet.Query().Count(c.UserContext())
	if err != nil {
		return 0, nil, err
	}

	wallets, err := client.Wallet.Query().
		Order(ent.Desc(wallet.FieldID)).
		Offset((pageQuery.Page - 1) * pageQuery.Size).
		Limit(pageQuery.Size).
		All(c.Context())
	if err != nil {
		return 0, nil, err
	}

	return count, wallets, nil
}

func (s *WalletServiceImpl) CreateWallet(c context.Context, userId int) (*ent.Wallet, error) {
	client := database.DB
	w, err := client.Wallet.Create().
		SetUserID(userId).
		Save(c)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (s *WalletServiceImpl) UpdateWalletBalance(c context.Context, walletId int, balance int) (*ent.Wallet, error) {
	client := database.DB
	w, err := client.Wallet.UpdateOneID(walletId).
		SetBalance(balance).
		Save(c)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (s *WalletServiceImpl) FreezeWallet(c context.Context, walletId int, amount int) (*ent.Wallet, error) {
	client := database.DB
	w, err := client.Wallet.UpdateOneID(walletId).
		AddFrozenAmount(amount).
		AddBalance(-amount).
		Save(c)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (s *WalletServiceImpl) UnfreezeWallet(c context.Context, walletId int, amount int) (*ent.Wallet, error) {
	client := database.DB
	w, err := client.Wallet.UpdateOneID(walletId).
		AddFrozenAmount(-amount).
		AddBalance(amount).
		Save(c)
	if err != nil {
		return nil, err
	}
	return w, nil
}
