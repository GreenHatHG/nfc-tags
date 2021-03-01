package library

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"io/ioutil"
	"net/http"
	"strings"
)

func UploadToXiaoMi(imgPath string) (string, error) {
	img, err := ioutil.ReadFile(imgPath)
	if err != nil {
		return "", gerror.New(err.Error())
	}

	body, err := formPost(img, gtime.Now().String()+".png",
		"http://likeyunba.com/upload/xiaomi.php")
	if err != nil {
		return "", err
	}

	resp := gconv.Map(body)
	if gconv.Int(resp["code"]) == 200 {
		return gconv.String(resp["path"]), nil
	}
	return "", gerror.New(body)
}

func formPost(imgContent []byte, imgName, url string) (string, error) {
	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"file\"; " +
		"filename=\"" + imgName + "\"\r\nContent-Type: image/png\r\n\r\n" + string(imgContent) + "\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", gerror.New(err.Error())
	}
	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", gerror.New(err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", gerror.New(err.Error())
	}

	return string(body), nil
}
