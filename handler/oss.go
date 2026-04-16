package handler

import (
	"bytes"

	"nail/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const avatarSignedExpireSec int64 = 1800

/*oss上传文件,path:云端位置,name:本地文件位置。账号参数从 config.ini [oss] 读取*/
func putOssObject(path, name string) error {
	client, err := oss.New(config.GetOssEndpoint(), config.GetOssAccessKeyId(), config.GetOssAccessKeySecret())
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(config.GetOssBucket())
	if err != nil {
		return err
	}
	return bucket.PutObjectFromFile(path, name)
}

/*oss 上传字节流，path: 云端路径（含文件名），data: 文件内容*/
func putOssObjectFromBytes(path string, data []byte) error {
	client, err := oss.New(config.GetOssEndpoint(), config.GetOssAccessKeyId(), config.GetOssAccessKeySecret())
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(config.GetOssBucket())
	if err != nil {
		return err
	}
	return bucket.PutObject(path, bytes.NewReader(data))
}

/*oss 删除文件*/
func deleteOssObject(path string) error {
	client, err := oss.New(config.GetOssEndpoint(), config.GetOssAccessKeyId(), config.GetOssAccessKeySecret())
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(config.GetOssBucket())
	if err != nil {
		return err
	}
	return bucket.DeleteObject(path)
}

func signOssGetURL(objectKey string, expireSec int64) (string, error) {
	client, err := oss.New(config.GetOssEndpoint(), config.GetOssAccessKeyId(), config.GetOssAccessKeySecret())
	if err != nil {
		return "", err
	}
	bucket, err := client.Bucket(config.GetOssBucket())
	if err != nil {
		return "", err
	}
	return bucket.SignURL(objectKey, oss.HTTPGet, expireSec)
}

func signAvatarURL(objectKey string) (string, error) {
	return signOssGetURL(objectKey, avatarSignedExpireSec)
}
