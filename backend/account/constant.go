package main

import "time"

var ChinaTime *time.Location

func Init() error {
	var err error

	ChinaTime, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}

	return nil
}
