// Package utils
// @author ufec https://github.com/ufec
// @date 2022/5/11
package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
)

// BuildThumbnailWithVideo
//  @Description: 通过视频生成缩略图
//  @param videoPath string	video视频地址
//  @param outputPath string 生成缩略图的路径(含文件名、文件后缀)
//  @return error 错误
func BuildThumbnailWithVideo(videoPath, outputPath string) error {
	cmd := "ffmpeg -i " + "./" + videoPath + " -f image2 -t 0.001 " + "./" + outputPath
	fmt.Println(cmd)
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		fmt.Println("sssssss")
		return err
	}
	return nil
}

// MakeDir
//  @Description: 创建目录
//  @param dir string 目录路径
//  @return error 创建目录error
func MakeDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}

// PathExists
//  @Description: 检测指定的path是否存在
//  @param path string 路径
//  @return bool 是否存在
//  @return error 不存在的错误
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// HmacSha256
//  @Description: 使用secret对message进行hmac_sha256散列
//  @param message string 要散列的文本内容
//  @param secret string 密钥
//  @param outputType string 结果类型 hex / base64
//  @return string 散列后的字符串
func HmacSha256(message, secret, outputType string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	if outputType == "hex" {
		return hex.EncodeToString(h.Sum(nil))
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
