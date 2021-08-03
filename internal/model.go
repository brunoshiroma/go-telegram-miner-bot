package internal

import (
	"time"

	_ "gorm.io/gorm"
)

type MinerConfigResponseJson struct {
	JSONPath string `json:"jsonPath"`
}

type MinerConfigRequestJson struct {
	Method string `json:"method"`
	Body   string `json:"body"`
	URL    string `json:"url"`
}

type MinerConfigDBJson struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type AutoCheckConfigJson struct {
	Enabled  bool              `json:"enabled"`
	Interval uint              `json:"interval"`
	DBConfig MinerConfigDBJson `json:"db"`
}

type MinerConfigJson struct {
	Name     string                  `json:"name"`
	Request  MinerConfigRequestJson  `json:"request"`
	Response MinerConfigResponseJson `json:"response"`
}

type MinersConfigJson struct {
	TelegramToken    string              `json:"telegramToken"`
	TelegramUsername string              `json:"telegramUsername"`
	Miners           []MinerConfigJson   `json:"miners"`
	AutoCheck        AutoCheckConfigJson `json:"autocheck"`
}

type MinerResult struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	Success   bool
	Result    string
	CreatedAt time.Time `gorm:"autoCreateTime;index"`
}

type MinersResult struct {
	Miners []MinerResult
}
