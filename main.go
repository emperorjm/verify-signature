package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Message that was signed
	msg := os.Getenv("MSG")
	// Public key of the private key used to sign the message above
	pubKeyHex := os.Getenv("PUB_KEY_HEX")
	// Signature generated after signing the message above
	sigHex := os.Getenv("SIG_HEX")

	pubKeyBytes, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		log.Fatalf("Failed to decode hex: %v", err)
	}
	pubKey := secp256k1.PubKey{Key: pubKeyBytes}

	sigBytes, err := hex.DecodeString(sigHex)
	if err != nil {
		log.Fatalf("Failed to decode hex: %v", err)
	}

	bites := []byte(msg)

	if pubKey.VerifySignature(bites, sigBytes) {
		fmt.Println("Signature is valid!")
	} else {
		fmt.Println("Signature is not valid.")
	}
}
