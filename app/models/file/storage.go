package file

import (
	"crypto/sha1"
	"fmt"
	"path"
	"time"
)

var root = "storage"

func GetStoragePath(filename string) string {
	hash := sha1.New()
	hash.Write([]byte(time.Now().String() + filename))
	hashed := fmt.Sprintf("%x", hash.Sum(nil))

	return path.Join(root, hashed[:3], hashed[3:6], hashed[6:] + "-" + filename)
}
