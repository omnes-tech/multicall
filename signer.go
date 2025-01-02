package multicall

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type SignerInterface interface {
	SignTx(tx *types.Transaction, chainId *big.Int) (*types.Transaction, error)
	GetAddress() *common.Address
}

type GenericSigner struct {
	PrivateKey *ecdsa.PrivateKey
	Address    *common.Address
}

func NewSigner(privateKeyHex string) (SignerInterface, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &GenericSigner{
		PrivateKey: privateKey,
		Address:    &address,
	}, nil
}

func (s *GenericSigner) SignTx(tx *types.Transaction, chainId *big.Int) (*types.Transaction, error) {
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), s.PrivateKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (s *GenericSigner) GetAddress() *common.Address {
	return s.Address
}
