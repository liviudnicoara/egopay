package accounts

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/pkg/errors"

	"github.com/liviudnicoara/egopay/pkg/filesystem"
)

type AccountRepository interface {
	CreateAccount(password string) (string, error)
	GetAccount(address string, password string) (Account, error)
}

type accountRepository struct {
	keyStorePath string
	keyStore     *keystore.KeyStore
}

func NewAccountRepository(keyStorePath string) AccountRepository {
	_ = filesystem.CreateDirectoryIfNotExists(keyStorePath)

	keyStore := keystore.NewKeyStore(keyStorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	return &accountRepository{
		keyStorePath: keyStorePath,
		keyStore:     keyStore,
	}
}

func (r *accountRepository) CreateAccount(password string) (string, error) {
	account, err := r.keyStore.NewAccount(password)

	return account.Address.String(), err
}

func (r *accountRepository) GetAccount(address string, password string) (Account, error) {
	accountFilePath, found, err := searchFile(r.keyStorePath, strings.TrimPrefix(address, "0x"))

	if err != nil {
		return Account{}, errors.WithMessagef(err, "could not retrieve account: %s", address)
	}

	if !found {
		return Account{}, errors.New(fmt.Sprintf("account not found: %s", address))
	}

	f, err := os.ReadFile(accountFilePath)
	if err != nil {
		return Account{}, errors.WithMessagef(err, "could not retrieve account: %s", address)
	}

	ks, err := keystore.DecryptKey(f, password)

	if err != nil {
		return Account{}, errors.WithMessage(err, fmt.Sprintf("could not retrieve account: %s", address))
	}

	binaryPVK := crypto.FromECDSA(ks.PrivateKey)
	encodedPVK := hexutil.Encode(binaryPVK)

	binaryPBK := crypto.FromECDSAPub(&ks.PrivateKey.PublicKey)
	encodedPBK := hexutil.Encode(binaryPBK)

	return Account{
		Address:           address,
		PrivateKey:        ks.PrivateKey,
		PrivateKeyEncoded: encodedPVK,
		PublicKeyEncoded:  encodedPBK,
	}, nil
}

func searchFile(directory, searchString string) (string, bool, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return "", false, err
	}

	searchString = strings.ToLower(searchString)

	for _, file := range files {
		if strings.Contains(file.Name(), searchString) {
			return filepath.Join(directory, file.Name()), true, nil
		}
	}

	return "", false, nil
}
