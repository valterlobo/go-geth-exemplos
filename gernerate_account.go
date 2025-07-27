package main

import (
	"bufio"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	fmt.Println("Ethereum Account Creator (Geth-like)")
	fmt.Println("-----------------------------------")

	// Prompt for a password to encrypt the new account's private key.
	// This password will be required to unlock the account later.
	password, err := getPasswordFromUser()
	if err != nil {
		fmt.Printf("Error getting password: %v\n", err)
		return
	}

	// Generate a new private key.
	// This is the core of the account creation process.
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Printf("Failed to generate private key: %v\n", err)
		return
	}

	// Derive the public key from the private key.
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Error: Cannot assert public key to ECDSA public key")
		return
	}

	// Derive the Ethereum address from the public key.
	// This is the address you'll share to receive funds.
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// --- Save the private key to a keystore file ---
	// Create a temporary directory for the keystore.
	// In a real application, you'd specify a persistent directory (e.g., ~/.ethereum/keystore).
	keystoreDir, err := ioutil.TempDir("", "ethereum-keystore-")
	if err != nil {
		fmt.Printf("Failed to create temporary keystore directory: %v\n", err)
		return
	}
	defer os.RemoveAll(keystoreDir) // Clean up the temporary directory on exit

	fmt.Printf("Keystore directory created at: %s\n", keystoreDir)

	// Create a new keystore instance.
	// The `keystore.StandardScryptN` and `keystore.StandardScryptP` are default parameters
	// for strong encryption, similar to Geth's default.
	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

	// Encrypt the private key with the provided password and store it in the keystore.
	// This generates a JSON file in the keystore directory.
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		fmt.Printf("Failed to import private key into keystore: %v\n", err)
		return
	}

	// --- Display account information ---
	fmt.Println("\n--- New Ethereum Account Details ---")
	fmt.Printf("Address: %s\n", address.Hex())
	fmt.Printf("Keystore File: %s\n", filepath.Join(keystoreDir, account.URL.Path))
	fmt.Printf("Keystore Account ID: %s\n", account.Address.Hex())

	// WARNING: Displaying the raw private key is highly insecure for production use.
	// This is for educational purposes only to demonstrate the generated key.
	// Never expose your private key in a real application!
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Printf("Private Key (WARNING: DO NOT SHARE!): %x\n", privateKeyBytes)

	fmt.Println("\nAccount created successfully!")
	fmt.Println("Remember your password to unlock this account.")
	fmt.Println("The keystore file contains your encrypted private key.")
	fmt.Println("Keep the keystore file and password secure!")
}

// getPasswordFromUser prompts the user for a password and confirms it.
func getPasswordFromUser() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a password for your new account: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		fmt.Print("Confirm password: ")
		confirmPassword, _ := reader.ReadString('\n')
		confirmPassword = strings.TrimSpace(confirmPassword)

		if password == confirmPassword {
			if len(password) < 1 { // Basic check for empty password
				fmt.Println("Password cannot be empty. Please try again.")
				continue
			}
			return password, nil
		} else {
			fmt.Println("Passwords do not match. Please try again.")
		}
	}
}
