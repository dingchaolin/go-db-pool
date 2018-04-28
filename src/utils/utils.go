package utils

import (
	"strconv"
	"time"
	"math/rand"
	"path/filepath"
	"os"
	"log"
	"strings"
	"crypto/md5"
	"encoding/hex"
)

var RootDir string

func init()  {
	RootDir = GetParentDirectory(GetCurrentDirectory())
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func Md5(Ori string) string {

	h := md5.New()
	h.Write([]byte(Ori))
	return hex.EncodeToString(h.Sum(nil))

}

func GetParentDirectory(dirctory string) string {
	return Substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}




func RandomTargetId()  string{
	return strconv.Itoa(int(time.Now().Unix())) + "-" + strconv.Itoa(RandInt(100000,999999)) + "-"+ Krand(5,1)
}

func RandomNo(pre string)  string{
	return pre + strconv.Itoa(int(time.Now().Unix())) + "-" + strconv.Itoa(RandInt(100000,999999)) + "-"+ Krand(10,1)
}

func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}
/**
 * size 随机码的位数
 * kind 0    // 纯数字
        1    // 小写字母
        2    // 大写字母
        3    // 数字、大小写字母
*/
func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}


