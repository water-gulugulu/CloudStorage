// @File  : Local.go 
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
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` // 本地文件路径
}
type LocalUpload struct {
	Config Local
}


// @author: [wulala](https://github.com/water-gulugulu)
// @object: *LocalUpload
// @function: UploadFile
// @description: 上传文件
// @param: file *multipart.FileHeader 
// @return: Path string
// @return: FileName string
// @return: err error
func (l *LocalUpload) UploadFile(file *multipart.FileHeader) (Path, FileName string, err error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(l.Config.Path, os.ModePerm)
	if mkdirErr != nil {
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := l.Config.Path + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, filename, nil
}

// @author: [wulala](https://github.com/water-gulugulu)
// @object: *LocalUpload
// @function: DeleteFile
// @description: 删除文件
// @param: key string
// @return: error

func (l *LocalUpload) DeleteFile(key string) error {
	p := l.Config.Path + "/" + key
	if strings.Contains(p, l.Config.Path) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
