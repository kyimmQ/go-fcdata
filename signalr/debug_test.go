package signalr

import (
	"fmt"
	"testing"
)

func TestDebugNegotiate(t *testing.T) {
	token := getAuthenticatedToken(t)
	fmt.Printf("Token: Bearer %s\n", token)
	c := NewClient("https://fc-datahub.ssi.com.vn/v2.0/signalr", token)
	header := c.getAuthHeader()
	fmt.Printf("Auth Header: %v\n", header)
	negResp, err := c.negotiate()
	if err != nil {
		t.Fatalf("negotiate error: %v", err)
	}
	fmt.Printf("Negotiate Response: %+v\n", negResp)
}
