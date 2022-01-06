package common

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

// resPath 资源执行目录
var resPath = "./"

func GetResPath() string {
	return resPath
}

// CopyFormRes 从 res 中拷贝资源到指定路径
// uuid res中的索引
// dst 目标文件夹
func CopyFormRes(uuid string, dstPath string) error {
	if !VerifyUUID(uuid) {
		return nil
	}
	var srcFile, dstFile *os.File
	srcFile, err := os.Open(GetPath(uuid))
	if err != nil {
		return err
	}
	// 覆盖写入
	dstFile, err = os.OpenFile(dstPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	defer func() {
		if srcFile != nil {
			srcFile.Close()
		}
		if dstFile != nil {
			dstFile.Close()
		}
	}()
	return nil
}

func SaveBytesToRes(uuid string, v []byte) error {
	if !VerifyUUID(uuid) {
		uuid = CreateUUID()
	}
	fp, err := os.Create(GetPath(uuid))
	if err != nil {
		return err
	}
	defer fp.Close()
	_, err = fp.Write(v)
	if err != nil {
		return err
	}
	return err
}

// SaveJsonToRes 保存一个json对象到res
// 处理命名和未命名对象
func SaveJsonToRes(uuid string, v interface{}) error {
	if !VerifyUUID(uuid) {
		uuid = CreateUUID()
	}
	fp, err := os.Create(GetPath(uuid))
	if err != nil {
		return err
	}
	defer fp.Close()
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = fp.Write(data)
	if err != nil {
		return err
	}
	return err
}

func OpenBytesFromRes(uuid string) ([]byte, error) {
	filePtr, err := os.Open(GetPath(uuid))
	if err != nil {
		return nil, err
	}
	defer filePtr.Close()

	return ioutil.ReadAll(filePtr)
}

func OpenJsonFromRes(uuid string, v interface{}) error {
	filePtr, err := os.Open(GetPath(uuid))
	if err != nil {
		return err
	}
	defer filePtr.Close()

	decoder := json.NewDecoder(filePtr)
	return decoder.Decode(v)
}

var sized int

func InitResSystem(root string, size int) {
	resPath = root + "res/"

	if !Exists(resPath) {
		_ = os.Mkdir(resPath, os.ModePerm)
	}

	DirMap = make([]string, size)
	sized = size
	for i := 0; i < size; i++ {
		m := GetMD5("r3inbowari" + strconv.Itoa(i))
		DirMap[i] = m + "/"
		if !Exists(resPath + m) {
			err := os.Mkdir(resPath+m, os.ModePerm)
			if err != nil {
				println(err.Error())
			}

		}
	}
}

var DirMap []string

func GetPath(uuid string) string {
	return resPath + DirMap[calc(uuid)] + uuid
}

func calc(uuid string) int {
	var ret int
	for _, v := range GetMD5(uuid) {
		ret += int(v)
	}
	return ret % sized
}
