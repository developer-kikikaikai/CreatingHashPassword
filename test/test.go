package main
import (
	"../hash"
	"fmt"
	"sort"
)

func test_HashSum() bool {
	teststring := "teststring"
	var teststring_hash map[string]string = map[string]string{}
	teststring_hash["sha256"]="3c8727e019a42b444667a587b6001251becadabbb36bfed8087a92c18882d111"
	teststring_hash["sha384"]="edd2ab3262b6c0121d706087045b60d51d6dc2f7419987ba12c983053a70c1057f15f58608ee07e1225266df36ba2c9c"
	teststring_hash["sha512"]="6253b39071e5df8b5098f59202d414c37a17d6a38a875ef5f8c7d89b0212b028692d3d2090ce03ae1de66c862fa8a561e57ed9eb7935ce627344f742c0931d72"
	teststring_hash["ripemd160"]="cd566972b5e50104011a92b59fa8e0b1234851ae"
	teststring_hash["sha3_224"]="17fce94d5da9ed5ce89fb6aaa884ecc94daf3ec35946cc39cbefbda6"
	teststring_hash["sha3_256"]="a52ef463edb03558be0598fb9a3ae67e0fc41c44bd0b1d4392b354d1e817fea8"
	teststring_hash["sha3_384"]="39660f7772dcc83fc8b5b67dc25537f23bb1bd34967d1e7ba8ff7585967553404b3114ac74ff6ecf7219200a1533c20c"
	teststring_hash["sha3_512"]="80938d92fe5d216884786ab8c23b87dc1c572e603c5b5d84d2454fa9bebbc84afa1d7a33ed35ed0196a023043e29c8f67b452102571e0f1944b8381b213a81fb"
	teststring_hash["sha512_224"]="da91a185e8c947f031419df9b9fc76accd56bb62ef8024c21046c364"
	teststring_hash["sha512_256"]="107b9a6123f289a3b14d7ce109c9b94feda782bf5bcaa09ce9e6f00773e21430"
	teststring_hash["black2s_256"]="73fb0bb92d259f9c9b13398aec62c73f204279661ea51b37ab72c11acb27aa3a"
	teststring_hash["black2s_384"]="9078c57a4ba97ea1b1bbbdc94df0f0cfd64246d154d7bdd730520737f20c970570e4e87ca17799242fa85d0fc63f56a3"
	teststring_hash["black2s_512"]="eb2c1252ac1d7d684f1b474b207610aa58c62d2b335b9adeac269d4eab5bd2e7bfe2f6a7a51be61b4f55e5c7dfced6922ca66af2f5eddf8d539f4ca1a28c1232"
	teststring_hash["other"]=""

	for alg, result := range teststring_hash {
		res := hash.HashSum(alg, teststring)
		if result != res {
			fmt.Printf("%s\n", res)
			fmt.Printf("%s\n", result)
			fmt.Printf("Failed to do hash %s\n", alg)
			return false
		}
	}
	return true
}

func test_AlgorithmList() bool {
	var result []string = []string{"sha256", "sha384", "sha512", "ripemd160", "sha3_224", "sha3_256", "sha3_384", "sha3_512", "sha512_224", "sha512_256", "black2s_256", "black2s_384", "black2s_512"}
	sort.SliceStable(result, func(i, j int) bool { return result[i] < result[j] })
	res := hash.AlgorithmList()
	if(len(result) != len(res)) {
		fmt.Printf("Failed to do hash algorithm %d %d\n", len(result), len(res))
		return false
	}

	for i := range res {
		if res[i] != result[i] {
			fmt.Printf("Failed to do hash algorithm %s and %s\n", res[i], result[i])
			return false
		}
	}

	return true
}

func main() {
	if !test_HashSum() {
		fmt.Printf("Failed to do HashSum\n")
		return
	}
	if !test_AlgorithmList() {
		fmt.Printf("Failed to do AlgorithmList\n")
		return
	}

	fmt.Printf("Finish to do all test\n")
}
