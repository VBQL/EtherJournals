package keys

import (
  "crypto/rand"
  "crypto/rsa"
  "crypto/x509"
  "crypto/sha256"
  "encoding/pem"
  "os"
  "fmt"
  "bufio"
)


func GeneratePrivKey(){
  // Generate key
  privKey, err := rsa.GenerateKey(rand.Reader, 4096)
  if err != nil {
      fmt.Println(err.Error)
      os.Exit(1)
  }

  pemPrivFile, err := os.Create("privkey.pem")
  if err != nil {
      fmt.Println(err)
      os.Exit(1)
  }

  var pemPrivBlock = &pem.Block{
    Type:  "RSA PRIVATE KEY",
    Bytes: x509.MarshalPKCS1PrivateKey(privKey),
  }

  err = pem.Encode(pemPrivFile, pemPrivBlock)
  if err != nil {
      fmt.Println(err)
      os.Exit(1)
  }
  pemPrivFile.Close()

}

func ImportPrivKey() (*rsa.PrivateKey){
  privKeyFile, err := os.Open("privkey.pem")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  pemfileinfo, _ := privKeyFile.Stat()
  var size int64 = pemfileinfo.Size()
  pembytes := make([]byte, size)
  buffer := bufio.NewReader(privKeyFile)
  _, err = buffer.Read(pembytes)
  data, _ := pem.Decode([]byte(pembytes))
  privKeyFile.Close()

  privKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return privKeyImported
}


func EncryptData([]byte) []byte{
  privkey := ImportPrivKey()
  pubkey := &privkey.PublicKey
  message := ([]byte("testing message in byte array"))
  label := []byte("")
  hash := sha256.New()
  ciphertext, err := rsa.EncryptOAEP(
      hash, 
      rand.Reader, 
      pubkey, 
      message, 
      label)
  if err != nil {
      fmt.Println(err)
      os.Exit(1)
  }
  return ciphertext
}

func DecryptData(ciphertext []byte) []byte{

  privkey := ImportPrivKey()
  hash := sha256.New()
  label := []byte("")
  plainText, err := rsa.DecryptOAEP(
      hash, 
      rand.Reader, 
      privkey, 
      ciphertext, 
      label)
  if err != nil {
      fmt.Println(err)
      os.Exit(1)
  }
  return plainText

}

func signMessage(data []byte) []byte{
  return([]byte("Will be implemented later, working on client side for now"))
}
