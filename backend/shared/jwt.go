package shared

import "github.com/golang-jwt/jwt"

type Payload struct {
	jwt.MapClaims        // ExpiryAt, IssueAt
	Session       string `json:"session"`
}
