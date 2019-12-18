package yandexTTS

import (
	"os"
)

const (
	YANDEX_KEY           = "YANDEX_KEY"
	AUDIO_PATH           = "AUDIO_PATH"
	DEFAULT_LANG         = "DEFAULT_LANG"
	DEFAULT_SPEED        = "DEFAULT_SPEED"
	DEFAULT_EMOTION      = "DEFAULT_EMOTION"
	DEFAULT_SPEAKER      = "DEFAULT_SPEAKER"
	DEFAULT_QUALITY      = "DEFAULT_QUALITY"
	DEFAULT_AUDIO_FORMAT = "DEFAULT_AUDIO_FORMAT"
)

type Env struct{}

// global var for env api namespace, eg. env.GetYandexKey()
var env Env

func (e Env) GetYandexKey() string {
	return os.Getenv(YANDEX_KEY)
}

func (e Env) GetDefaultLang() string {
	return os.Getenv(DEFAULT_LANG)
}

func (e Env) GetDefaultSpeed() string {
	return os.Getenv(DEFAULT_SPEED)
}

func (e Env) GetDefaultEmotion() string {
	return os.Getenv(DEFAULT_EMOTION)
}

func (e Env) GetDefaultSpeaker() string {
	return os.Getenv(DEFAULT_SPEAKER)
}

func (e Env) GetDefaultQuality() string {
	return os.Getenv(DEFAULT_QUALITY)
}

func (e Env) GetDefaultAudioFormat() string {
	return os.Getenv(DEFAULT_AUDIO_FORMAT)
}

func (e Env) GetPathForAudio() string {
	return os.Getenv(AUDIO_PATH)
}
