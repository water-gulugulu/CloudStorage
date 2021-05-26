// @File  : Oss.go 
// @Author: JunLong.Liao&此处不应有BUG!
// @Date  : 2021/5/26
// @slogan: 又是不想写代码的一天，神兽保佑，代码无BUG！
//         ┏┓      ┏┓
//        ┏┛┻━━━━━━┛┻┓
//        ┃     ღ    ┃
//        ┃  ┳┛   ┗┳ ┃
//        ┃     ┻    ┃
//        ┗━┓      ┏━┛
//          ┃      ┗━━━┓
//          ┃ 神兽咆哮!  ┣┓
//          ┃         ┏┛
//          ┗┓┓┏━━━┳┓┏┛
//           ┃┫┫   ┃┫┫
//           ┗┻┛   ┗┻┛

package CloudStorage

import "mime/multipart"

type OSS interface {
	UploadFile(file *multipart.FileHeader) (path,filename string ,err error)
	DeleteFile(key string) error
}

type OssConfig struct {
	Type string // 类型 local-本地 tencent-腾讯云 aliyun-阿里云 qiniu-七牛云
	Config interface{} // 类型 根据type断言对应的类型
}


type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`                                // 存储区域
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`                          // 空间名称
	ImgPath       string `mapstructure:"img-path" json:"imgPath" yaml:"img-path"`                     // CDN加速域名
	UseHTTPS      bool   `mapstructure:"use-https" json:"useHttps" yaml:"use-https"`                  // 是否使用https
	AccessKey     string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`               // accessKey
	SecretKey     string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`               // secretKey
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"useCdnDomains" yaml:"use-cdn-domains"` // 上传是否使用CDN上传加速
}

type AliyunOSS struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `mapstructure:"access-key-id" json:"accessKeyId" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"accessKeySecret" yaml:"access-key-secret"`
	BucketName      string `mapstructure:"bucket-name" json:"bucketName" yaml:"bucket-name"`
	BucketUrl       string `mapstructure:"bucket-url" json:"bucketUrl" yaml:"bucket-url"`
}


func (c *OssConfig)NewOSS() OSS  {
	switch c.Type {
	case "local":
		return &LocalUpload{
			Config: c.Config.(Local),
		}
	case "tencent":
		return &TencentCos{
			Config: c.Config.(TencentCOS),
		}
	default:
		return &LocalUpload{
			Config: c.Config.(Local),
		}
	}
}