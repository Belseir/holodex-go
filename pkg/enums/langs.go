package enums

import "fmt"

var Langs = newLangs()

type LangsType string

func newLangs() *langs {
	return &langs{
		ENGLISH:  "en",
		JAPANESE: "jp",
		CHINESE:  "zh",
	}
}

type langs struct {
	ENGLISH  LangsType
	JAPANESE LangsType
	CHINESE  LangsType
}

func (_ *langs) IsValid(l LangsType) (bool, error) {
	if l != Langs.ENGLISH && l != Langs.JAPANESE && l != Langs.CHINESE {
		return false, fmt.Errorf("Invalid Langs (%s) accessed. Expected 'en', 'jp' or 'zh'", l)
	}

	return true, nil
}
