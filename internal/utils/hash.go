package utils

import "crypto/md5"

func Hash(req string) string {
	hash := md5.New()

	hash.Write([]byte(req))

	return string(hash.Sum(nil))
}
