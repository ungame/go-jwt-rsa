package keypair

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	pemPrivateKey []byte
	pemPublicKey  []byte
)

func init() {

	var (
		_, keysPath, _, _ = runtime.Caller(0)
		pvtKeyFilePath    = filepath.Join(filepath.Dir(keysPath), "key.pem")
		pubKeyFilePath    = filepath.Join(filepath.Dir(keysPath), "key.pem.pub")
		err               error
	)

	pemPrivateKey, err = os.ReadFile(pvtKeyFilePath)
	if err != nil {
		log.Fatalln("error on read private key: ", err.Error())
	}

	pemPublicKey, err = os.ReadFile(pubKeyFilePath)
	if err != nil {
		log.Fatalln("error on read public key: ", err.Error())
	}
}

func GetPrivateKey() []byte {
	pvtKey := make([]byte, len(pemPrivateKey))
	copy(pvtKey, pemPrivateKey)
	return pvtKey
}

func GetPublicKey() []byte {
	pubKey := make([]byte, len(pemPublicKey))
	copy(pubKey, pemPublicKey)
	return pubKey
}
