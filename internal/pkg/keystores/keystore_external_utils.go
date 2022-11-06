/*
MIT License

Copyright (c) 2020 @protolambda

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package keystores

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/NethermindEth/sedge/configs"
	"github.com/google/uuid"
	hbls "github.com/herumi/bls-eth-go-binary/bls"
	"github.com/protolambda/go-keystorev4"
	"github.com/protolambda/zrnt/eth2/beacon"
	eth2 "github.com/protolambda/zrnt/eth2/configs"
	"github.com/protolambda/zrnt/eth2/util/hashing"
	"github.com/protolambda/ztyp/tree"
	"github.com/tyler-smith/go-bip39"
	e2types "github.com/wealdtech/go-eth2-types/v2"
	util "github.com/wealdtech/go-eth2-util"
	"golang.org/x/sync/errgroup"
)

func init() {
	hbls.Init(hbls.BLS12_381)
	hbls.SetETHmode(hbls.EthModeLatest)
}

type WalletOutput interface {
	InsertAccount(priv e2types.PrivateKey, keyPath string, insecure bool, idx uint64) error
}

// Following EIP 2335
type KeyFile struct {
	id        uuid.UUID
	name      string
	publicKey e2types.PublicKey
	secretKey e2types.PrivateKey
}

type KeyEntry struct {
	KeyFile
	path       string
	passphrase string
	insecure   bool
}

func NewKeyEntry(priv e2types.PrivateKey, keyPath string, ww *WalletWriter, insecure bool) (*KeyEntry, error) {
	if ww.generalPassphrase == "" {
		var pass [32]byte
		n, err := rand.Read(pass[:])
		if err != nil {
			return nil, err
		}
		if n != 32 {
			return nil, errors.New(configs.KeyEntryGenerationError)
		}
		// Convert it to human readable characters, to keep it manageable
		ww.generalPassphrase = base64.URLEncoding.EncodeToString(pass[:])
	}
	return &KeyEntry{
		KeyFile: KeyFile{
			id:        uuid.New(),
			name:      "val_" + hex.EncodeToString(priv.PublicKey().Marshal()),
			publicKey: priv.PublicKey(),
			secretKey: priv,
		},
		path:       keyPath,
		passphrase: ww.generalPassphrase,
		insecure:   insecure,
	}, nil
}

func (ke *KeyEntry) MarshalJSON() ([]byte, error) {
	var salt [32]byte
	if _, err := rand.Read(salt[:]); err != nil {
		return nil, err
	}
	kdfParams := &keystorev4.PBKDF2Params{
		Dklen: 32,
		C:     262144,
		Prf:   "hmac-sha256",
		Salt:  salt[:],
	}
	if ke.insecure { // INSECURE but much faster, this is useful for ephemeral testnets
		kdfParams.C = 2
	}
	cipherParams, err := keystorev4.NewAES128CTRParams()
	if err != nil {
		return nil, fmt.Errorf(configs.AESParamsCreationError, err)
	}
	crypto, err := keystorev4.Encrypt(ke.secretKey.Marshal(), []byte(ke.passphrase),
		kdfParams, keystorev4.Sha256ChecksumParams, cipherParams)
	if err != nil {
		return nil, fmt.Errorf(configs.SecretEncryptionError, err)
	}
	keystore := &keystorev4.Keystore{
		Crypto:      *crypto,
		Description: fmt.Sprintf("0x%x", ke.publicKey.Marshal()),
		Pubkey:      ke.publicKey.Marshal(),
		Path:        ke.path,
		UUID:        ke.id,
		Version:     4,
	}
	return json.Marshal(keystore)
}

func (ke *KeyEntry) PubHex() string {
	return "0x" + hex.EncodeToString(ke.publicKey.Marshal())
}

func (ke *KeyEntry) Path() string {
	return ke.path
}

func (ke *KeyEntry) PubHexBare() string {
	return hex.EncodeToString(ke.publicKey.Marshal())
}

type WalletWriter struct {
	sync.RWMutex
	generalPassphrase string
	entries           []*KeyEntry
}

func NewWalletWriter(entries uint64, passphrase string) *WalletWriter {
	return &WalletWriter{
		entries:           make([]*KeyEntry, entries),
		generalPassphrase: passphrase,
	}
}

func (ww *WalletWriter) InsertAccount(priv e2types.PrivateKey, keyPath string, insecure bool, idx uint64) error {
	key, err := NewKeyEntry(priv, keyPath, ww, insecure)
	if err != nil {
		return err
	}
	ww.RWMutex.Lock()
	defer ww.RWMutex.Unlock()
	ww.entries[idx] = key
	return nil
}

func (ww *WalletWriter) WriteOutputs(fpath string) error {
	if _, err := os.Stat(fpath); !os.IsNotExist(err) {
		return errors.New(configs.KeystoreOutputExistingError)
	}
	valKeysPath := filepath.Join(fpath, "validator_keys")
	if err := os.MkdirAll(valKeysPath, os.ModePerm); err != nil {
		return err
	}

	var g errgroup.Group
	// For all: write JSON keystore files, each in their own directory (lighthouse requirement)
	for _, k := range ww.entries {
		e := k
		g.Go(func() error {
			dat, err := e.MarshalJSON()
			if err != nil {
				return err
			}
			keystoreFilename := fmt.Sprintf("keystore-%s.json", strings.ReplaceAll(e.path, "/", "_"))
			{
				if err := os.WriteFile(filepath.Join(valKeysPath, keystoreFilename), dat, 0o644); err != nil {
					return err
				}
			}
			return nil
		})
	}
	// Write passphrase file
	passFileName := "keystore_password.txt"
	if err := os.WriteFile(filepath.Join(fpath, passFileName), []byte(ww.generalPassphrase), 0o644); err != nil {
		return err
	}

	return g.Wait()
}

func CreateKeystores(
	vkgd ValidatorKeysGenData,
) error {
	ww := NewWalletWriter(vkgd.MaxIndex-vkgd.MinIndex, vkgd.Passphrase)
	err := selectVals(vkgd.Mnemonic, vkgd.MinIndex, vkgd.MaxIndex, ww, vkgd.Insecure)
	if err != nil {
		return fmt.Errorf(configs.KeystoreGenerationError, err)
	}

	err = ww.WriteOutputs(vkgd.OutputPath)
	if err != nil {
		return fmt.Errorf(configs.KeystoreGenerationError, err)
	}
	return nil
}

// Narrow pubkeys: we don't want 0xAb... to be different from ab...
func narrowedPubkey(pub string) string {
	return strings.TrimPrefix(strings.ToLower(pub), "0x")
}

func selectVals(sourceMnemonic string,
	minAcc uint64, maxAcc uint64,
	output WalletOutput,
	insecure bool,
) error {
	valSeed, err := mnemonicToSeed(sourceMnemonic)
	if err != nil {
		return err
	}

	var g errgroup.Group
	// Try look for unassigned accounts in the wallet
	for i := minAcc; i < maxAcc; i++ {
		idx := i
		g.Go(func() error {
			valAccPath := fmt.Sprintf("m/12381/3600/%d/0/0", idx)
			a, err := util.PrivateKeyFromSeedAndPath(valSeed, valAccPath)
			if err != nil {
				return fmt.Errorf(configs.KeystoreDerivationError, valAccPath)
			}
			pubkey := narrowedPubkey(hex.EncodeToString(a.PublicKey().Marshal()))
			if err := output.InsertAccount(a, valAccPath, insecure, idx-minAcc); err != nil {
				if err.Error() == fmt.Sprintf(configs.KeystoreExistingInWalletError, pubkey) {
					// Account with this pubkey already exists in output wallet, skipping it
					return nil
				} else {
					return fmt.Errorf(configs.KeystoreImportingError, pubkey, err)
				}
			}
			return nil
		})

	}
	return g.Wait()
}

func CreateMnemonic() (string, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return "", fmt.Errorf(configs.MnemonicGenerationError, err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf(configs.MnemonicGenerationError, err)
	}
	return mnemonic, nil
}

func mnemonicToSeed(mnemonic string) (seed []byte, err error) {
	mnemonic = strings.TrimSpace(mnemonic)
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New(configs.InvalidMnemonicError)
	}
	return bip39.NewSeed(mnemonic, ""), nil
}

func CreateDepositData(
	vkgd ValidatorKeysGenData,
) error {
	var genesisForkVersion beacon.Version
	if err := genesisForkVersion.UnmarshalText([]byte(vkgd.ForkVersion)); err != nil {
		return fmt.Errorf(configs.ForkVersionDecodeError, err)
	}

	valSeed, err := mnemonicToSeed(vkgd.Mnemonic)
	if err != nil {
		return fmt.Errorf(configs.BadMnemonicError, err)
	}
	withdrSeed, err := mnemonicToSeed(vkgd.Mnemonic)
	if err != nil {
		return fmt.Errorf(configs.BadMnemonicError, err)
	}

	depositPath := filepath.Join(vkgd.OutputPath, "deposit_data.json")
	if _, err := os.Stat(depositPath); !os.IsNotExist(err) {
		return errors.New(configs.KeystoreOutputExistingError)
	}

	depositData := new(bytes.Buffer)
	if vkgd.AsJsonList {
		depositData.WriteString("[")
	}
	for i := vkgd.MinIndex; i < vkgd.MaxIndex; i++ {
		valAccPath := fmt.Sprintf("m/12381/3600/%d/0/0", i)
		val, err := util.PrivateKeyFromSeedAndPath(valSeed, valAccPath)
		if err != nil {
			return fmt.Errorf(configs.KeystoreSecretKeyCreationError, valAccPath, err)
		}
		withdrAccPath := fmt.Sprintf("m/12381/3600/%d/0", i)
		withdr, err := util.PrivateKeyFromSeedAndPath(withdrSeed, withdrAccPath)
		if err != nil {
			return fmt.Errorf(configs.WithdrawalSecretKeyCreationError, withdrAccPath, err)
		}

		var pub beacon.BLSPubkey
		copy(pub[:], val.PublicKey().Marshal())

		var withdrPub beacon.BLSPubkey
		copy(withdrPub[:], withdr.PublicKey().Marshal())
		withdrCreds := hashing.Hash(withdrPub[:])
		withdrCreds[0] = eth2.Mainnet.BLS_WITHDRAWAL_PREFIX[0]

		data := beacon.DepositData{
			Pubkey:                pub,
			WithdrawalCredentials: withdrCreds,
			Amount:                beacon.Gwei(vkgd.AmountGwei),
			Signature:             beacon.BLSSignature{},
		}
		msgRoot := data.ToMessage().HashTreeRoot(tree.GetHashFn())
		var secKey hbls.SecretKey
		if err := secKey.Deserialize(val.Marshal()); err != nil {
			return fmt.Errorf(configs.KeystoreSecretKeyConvertionError, err)
		}

		dom := beacon.ComputeDomain(eth2.Mainnet.DOMAIN_DEPOSIT, genesisForkVersion, beacon.Root{})
		msg := beacon.ComputeSigningRoot(msgRoot, dom)
		sig := secKey.SignHash(msg[:])
		copy(data.Signature[:], sig.Serialize())

		dataRoot := data.HashTreeRoot(tree.GetHashFn())
		dataMessageRoot := data.MessageRoot()
		jsonData := map[string]interface{}{
			"pubkey":                 hex.EncodeToString(data.Pubkey[:]),
			"withdrawal_credentials": hex.EncodeToString(data.WithdrawalCredentials[:]),
			"amount":                 uint64(data.Amount),
			"signature":              hex.EncodeToString(data.Signature[:]),
			"deposit_message_root":   hex.EncodeToString(dataMessageRoot[:]),
			"deposit_data_root":      hex.EncodeToString(dataRoot[:]),
			"fork_version":           genesisForkVersion.String()[2:],
			"network_name":           vkgd.NetworkName,
			"account":                valAccPath, // for ease with tracking where it came from.
			"version":                1,          // ethereal cli requirement
			"deposit_cli_version":    "2.3.0",    // for launchpad to work
		}
		jsonStr, err := json.Marshal(jsonData)
		if vkgd.AsJsonList && i+1 < vkgd.MaxIndex {
			jsonStr = append(jsonStr, ',')
		}
		if err != nil {
			return fmt.Errorf(configs.DepositDataEncodingError, err)
		}
		depositData.WriteString(string(jsonStr))
	}
	if vkgd.AsJsonList {
		depositData.WriteString("]")
	}

	err = os.WriteFile(depositPath, depositData.Bytes(), 0o644)
	if err != nil {
		return fmt.Errorf(configs.DepositFileWriteError, err)
	}
	return nil
}
