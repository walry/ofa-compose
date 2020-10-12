package tools

import (
	md52 "crypto/md5"
	"encoding/hex"
	json2 "encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"mime/multipart"
	"path/filepath"
	"time"
)

func CreateSaveFileName(file *multipart.FileHeader) string {
	//文件扩展名
	suffix := file.Header["Content-Type"][0][12:]
	dir := viper.GetString("uploadsPath")
	json,_ := json2.Marshal(file.Header)
	t,_ := time.Now().MarshalBinary()
	name := append(json,t...)
	md5 := md52.New()
	md5.Write(name)
	cipherStr := md5.Sum(nil)
	prefix := fmt.Sprintf("%s",hex.EncodeToString(cipherStr))
	fileName := filepath.Join(dir,prefix + "." + suffix)
	return fileName
}
