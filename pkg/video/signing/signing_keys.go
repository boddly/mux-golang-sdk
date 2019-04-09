package signing

// URLSigningKey is used as the secret when signing any Mux URL. Mux requires a JSON Web Token as the value of the token query parameter.
// The token query parameter must be set for URLs that reference a playback ID with a signed playback policy.
type URLSigningKey struct {
	ID         string `json:"id"`
	PrivateKey string `json:"private_key"`
	CreatedAt  int    `json:"created_at"`
}
