package upload

import (
	"context"
	"douyin/config"
	"mime/multipart"
	"strconv"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

//上传文件，本项目主要包含视频和图片
// 参考项目和文档
// https://gitee.com/pixelmax/gin-vue-admin/blob/main/server/utils/upload/aliyun_oss.go
// https://help.aliyun.com/product/31815.html?spm=5176.7933691.J_5253785160.6.272f4c59KogXWZ
// https://developer.qiniu.com/kodo

// ToQiNiu 上传文件到七牛云对象存储
func ToQiNiu(file multipart.File, fileSize, videoID int64) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: config.QiNiuBucket,
	}

	mac := qbox.NewMac(config.QiNiuAccessKey, config.QiNiuSecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	key := "video/" + strconv.FormatInt(videoID, 10) + ".mp4"
	err := formUploader.Put(context.Background(), &ret, upToken, key, file, fileSize, &putExtra)
	if err != nil {
		return "", err
	}
	url := "http://" + config.QiNiuServer + "/" + ret.Key
	return url, nil
}

// ToAliYun 上传文件到阿里云对象存储
func ToAliYun() {

}
