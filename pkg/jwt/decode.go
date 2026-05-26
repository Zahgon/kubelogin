// Package jwt provides JWT manipulations.
// See https://tools.ietf.org/html/rfc7519#section-4.1.3
package jwt

// DecodeWithoutVerify decodes the JWT string and returns the claims.
// Note that this method does not verify the signature and always trust it.
func DecodeWithoutVerify(s string) (*Claims, error) { _ = "STUB: not implemented"; return nil, nil }

// DecodePayloadAsPrettyJSON decodes the JWT string and returns the pretty JSON string.
func DecodePayloadAsPrettyJSON(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// DecodePayloadAsRawJSON extracts the payload and returns the raw JSON.
func DecodePayloadAsRawJSON(s string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func decodePayload(payload string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
