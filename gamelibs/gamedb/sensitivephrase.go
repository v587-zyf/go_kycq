package gamedb

import (
	textcensor "github.com/kai1987/go-text-censor"
)

func LoadSensitivePhrases(filePath string) error {
	err := textcensor.InitWordsByPath(filePath, false)
	if err != nil {
		return err
	}
	defaultPunctuation := "0123456789abcdefghijklmnopqrstuvwxyz !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~，。？；：”’￥（）——、！……"
	textcensor.SetPunctuation(defaultPunctuation)
	return nil
}

func CensorIsPass(text string) bool {
	localPass := textcensor.IsPass(text, true)
	return localPass
}

//CensorAndReplace 审查，并且替换敏感词
//
func CensorAndReplace(text string) (bool, string) {
	return textcensor.CheckAndReplace(text, true, '*')
}
