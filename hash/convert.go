package hash
import (
	_ "fmt"
	"encoding/hex"
)

/*use z and q twice*/
var wordsWithoutSymbol map[int][]string
var wordsWithSymbol map[int][]string

func initializeHashList() {
	wordsWithoutSymbol = make(map[int][]string, 0)
	wordsWithSymbol = make(map[int][]string, 0)
	/** 1/1 scale **/
	/*without symbol string list*/
	tmpWithoutSymbol := make([]string, 4);
	tmpWithoutSymbol[0] = "0123456789abcdef";
	tmpWithoutSymbol[1] = "ghijklmnopqrstuv";
	tmpWithoutSymbol[2] = "wxyzABCDEFGHIJKL";
	tmpWithoutSymbol[3] = "MNOPQRSTUVWXYZqz";
	wordsWithoutSymbol[1] = tmpWithoutSymbol

	/*with symbol string list*/
	tmpWithSymbol := make([]string, 6);
	_ = copy(tmpWithSymbol, tmpWithoutSymbol);
	tmpWithSymbol[4] = "!\"#$%&'()*+,-./:";
	tmpWithSymbol[5] = ";<=>?@[]^_`{|}~, ";
	wordsWithSymbol[1] = tmpWithSymbol

	/** 1/2 scale. other pattern is same **/
	wordsWithoutSymbol[0] = make([]string, 1)
	symbollen := len(tmpWithoutSymbol);
	looplen := 0x10/symbollen;
	for i:=0 ; i< looplen; i++ {
		for j:=0 ; j< symbollen; j++ {
			wordsWithoutSymbol[0][0] += tmpWithoutSymbol[j];
		}
	}

	wordsWithSymbol[0] = make([]string, 1)
	looplen = 2//loop 2 times
	for i:=0 ; i< looplen; i++ {
		for j:=0 ; j< symbollen; j++ {
			wordsWithSymbol[0][0] += tmpWithoutSymbol[j];
		}
	}
	wordsWithSymbol[0][0] += "0123456789abcdef";
	wordsWithSymbol[0][0] += "MNOPQRSTUVWXYZqz";
	wordsWithSymbol[0][0] += "!\"#$%&'()*+,-./:";
	wordsWithSymbol[0][0] += ";<=>?@[]^_`{|}~, ";
	//fmt.Printf("tmpWithSymbol:%v\n", wordsWithSymbol);
	//fmt.Printf("tmpWithoutSymbol:%v\n", wordsWithoutSymbol);
}

type hashTable struct {
	table []string
	scaledown int
}

func (this hashTable) convert(hashString string) string {

	hexvalue := 0
	result := ""
	hashStringLength := len(hashString)
	for i := 0; i < hashStringLength ; i++ {
		hexvalue = hexvalue << 4//1 byte
		tmpval, _ := hex.DecodeString("0" + string(hashString[i]))
		hexvalue |= int(tmpval[0])
		if (i % this.scaledown) == (this.scaledown - 1) {
			index_table := (hexvalue + i) % len(this.table)
			index_value := hexvalue % len(this.table[index_table])
			result += string(this.table[index_table][index_value])
			hexvalue = 0
		}
	}
	return result
}

func generateHashTable(scaleDownDenominator int, useSymbol bool) hashTable {
	var res hashTable

	//scaleDownDenominator: 1, 2, 3, 4, 8
	if useSymbol {
		res.table = wordsWithSymbol[scaleDownDenominator%2];
	} else {
		res.table = wordsWithoutSymbol[scaleDownDenominator%2];
	}
	res.scaledown = scaleDownDenominator
	return res
}

//base hash is 
func CompressHash(hashString string, scaleDownDenominator int, useSymbol bool) string {
	if wordsWithoutSymbol == nil {
		initializeHashList();
	}
	table := generateHashTable(scaleDownDenominator, useSymbol);
	return table.convert(hashString)
}
