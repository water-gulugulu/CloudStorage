// @File  : Tencent.go 
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

import (
	"mime/multipart"
)

type TencentCos struct {
	Config TencentCOS
}
type TencentCOS struct {
	Bucket     string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Region     string `mapstructure:"region" json:"region" yaml:"region"`
	SecretID   string `mapstructure:"secret-id" json:"secretID" yaml:"secret-id"`
	SecretKey  string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	BaseURL    string `mapstructure:"base-url" json:"baseURL" yaml:"base-url"`
	PathPrefix string `mapstructure:"path-prefix" json:"pathPrefix" yaml:"path-prefix"`
}

func (t TencentCos) UploadFile(file *multipart.FileHeader) (path, filename string, err error) {
	panic("implement me")
}

func (t TencentCos) DeleteFile(key string) error {
	panic("implement me")
}
