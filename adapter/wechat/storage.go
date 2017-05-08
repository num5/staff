package wechat

import (
	"encoding/json"
	"github.com/num5/ider"
	"github.com/num5/webot"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type message struct {
	ID          int64              `json:"id"`
	Msg         webot.EventMsgData `json:"msg"`
	ReceiveTime time.Time          `json:"receive_time"`
}

func Stor(storpath, file string, data webot.EventMsgData) error {
	msg := new(message)
	id := ider.NewID(1).Next()
	msg.ID = id
	msg.Msg = data
	msg.ReceiveTime = time.Now()

	return write(storpath, file, msg)
}

func write(storpath, file string, value interface{}) error {
	// 检测文件夹是否存在   若不存在  创建文件夹
	if _, err := os.Stat(storpath); err != nil {

		if os.IsNotExist(err) {

			err = os.MkdirAll(storpath, os.ModePerm)

			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	content, err := json.Marshal(value)

	if err != nil {
		return err
	}
	storfile := path.Join(storpath, file)
	return ioutil.WriteFile(storfile, content, os.ModePerm)
}
