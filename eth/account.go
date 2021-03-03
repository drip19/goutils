package eth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// 根据私钥得出公钥和地址
func PriKeyToAddress(privateKey string) (*ecdsa.PrivateKey, common.Address, error) {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, [20]byte{}, err
	}
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, [20]byte{}, err
	}
	return pk, crypto.PubkeyToAddress(*publicKeyECDSA), nil
}
