package hash
import (
	"io"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/blake2b"
	"encoding/hex"
	"sort"
)

//define function type
var hash_algorithms map[string]func(string) string = map[string]func(string) string{"sha256":sha256sum,"sha384":sha384sum,"sha512":sha512sum,"ripemd160":ripemd160sum,"sha3_224":sha3sha224sum,"sha3_256":sha3sha256sum,"sha3_384":sha3sha384sum,"sha3_512":sha3sha512sum,"sha512_224":sha512_224sum,"sha512_256":sha512_256sum,"black2s_256":black2s_256sum,"black2s_384":black2b_384sum,"black2s_512":black2b_512sum}

func sha256sum(data string) string {
	bytes := sha256.Sum256([] byte(data))
	return hex.EncodeToString(bytes[:])
}

func sha384sum(data string) string {
	bytes := sha512.Sum384([] byte(data))
	return hex.EncodeToString(bytes[:])
}

func sha512sum(data string) string {
	bytes := sha512.Sum512([] byte(data))
	return hex.EncodeToString(bytes[:])
}

func ripemd160sum(data string) string {
	rip := ripemd160.New()
	io.WriteString(rip, data)
	return hex.EncodeToString(rip.Sum(nil))
}

func sha3sha224sum(data string) string {
	bytes := sha3.Sum224([] byte(data))
	return hex.EncodeToString(bytes[:])
}

func sha3sha256sum(data string) string {
	bytes := sha3.Sum256([] byte(data))
	return hex.EncodeToString(bytes[:])
}

func sha3sha384sum(data string) string {
	bytes := sha3.Sum384([] byte(data))
	return hex.EncodeToString(bytes[:])
}

func sha3sha512sum(data string) string {
	bytes := sha3.Sum512([] byte(data))
	return hex.EncodeToString(bytes[:])
}

func sha512_224sum(data string) string {
	bytes := sha512.Sum512_224([] byte(data))
	return hex.EncodeToString(bytes[:])
}

func sha512_256sum(data string) string {
	bytes := sha512.Sum512_256([] byte(data))
	return hex.EncodeToString(bytes[:])
}

func black2s_256sum(data string) string {
	bytes := blake2s.Sum256([] byte(data))
	return hex.EncodeToString(bytes[:])
}

func black2b_256sum(data string) string {
	bytes := blake2b.Sum256([] byte(data))
	return hex.EncodeToString(bytes[:])
}

func black2b_384sum(data string) string {
	bytes := blake2b.Sum384([] byte(data))
	return hex.EncodeToString(bytes[:])
}

func black2b_512sum(data string) string {
	bytes := blake2b.Sum512([] byte(data))
	return hex.EncodeToString(bytes[:])
}

func HashSum(algorithm string, data string) string {
	fname, ok := hash_algorithms[algorithm];
	if ok {
		return fname(data)
	} else {
		return ""
	}
}

func AlgorithmList() []string{
	var response []string = make([]string, 0)
	for key := range hash_algorithms {
        response = append(response, key)
    }
	sort.SliceStable(response, func(i, j int) bool { return response[i] < response[j] })
	return response
}
