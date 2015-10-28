package certgen

import (
	"testing"
	"time"
)

func TestCertGen(t *testing.T) {
	Generate("localhost", "Jan 1 15:04:05 2011", 365*24*time.Hour, false, 2048, "P256")
}
