package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

const firebaseCertURL = "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"

func VerifyFirebaseJWT(idToken string) (jwt.MapClaims, error) {
	// Fetch Firebase public keys
	certificates, err := getFirebasePublicKeys()
	if err != nil {
		return nil, err
	}

	// Parse the ID token
	token, err := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		// Get the public key for token verification
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("kid header not found in the token")
		}

		key, ok := certificates[kid]
		if !ok {
			return nil, fmt.Errorf("public key not found for kid: %s", kid)
		}

		return jwt.ParseRSAPublicKeyFromPEM([]byte(key))
	})

	if err != nil {
		return nil, err
	}

	// Verify the token
	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unable to extract claims")
	}

	return claims, nil
}

func getFirebasePublicKeys() (map[string]string, error) {
	resp, err := http.Get(firebaseCertURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch Firebase public keys, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var keys map[string]string
	err = json.Unmarshal(body, &keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}
