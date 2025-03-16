package rid

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func GenerateInstanceID(hashName string, uid uint64, prefix string) string {
	hash := sha256.New()
	hash.Write([]byte(hashName + strconv.Itoa(int(uid))))

	hashStr := hex.EncodeToString(hash.Sum(nil))

	return prefix + hashStr[:8]
}
