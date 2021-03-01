package service

import (
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gproc"
	"nfc-tags/library"
	"os"
	"path"
)

func (a *serviceAnalysis) StandingRestImageUrl() (string, error) {
	dir, _ := os.Getwd()
	htmlPath := path.Join(dir, "snapshot", "calendar.html")
	r, err := gproc.ShellExec(fmt.Sprintf(`snapshot %s png 0`, htmlPath))
	if err != nil {
		return "", gerror.New(err.Error())
	}
	g.Log().Println(r)

	url, err := library.UploadToXiaoMi(path.Join(dir, "output.png"))
	if err != nil {
		return "", err
	}
	return url, nil
}
