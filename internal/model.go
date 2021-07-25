package internal

type MinerConfigResponseJson struct {
	JSONPath string `json:"jsonPath"`
}

type MinerConfigRequestJson struct {
	Method string `json:"method"`
	Body   string `json:"body"`
	URL    string `json:"url"`
}

type MinerConfigJson struct {
	Name     string                  `json:"name"`
	Request  MinerConfigRequestJson  `json:"request"`
	Response MinerConfigResponseJson `json:"response"`
}

type MinersConfigJson struct {
	TelegramToken    string            `json:"telegramToken"`
	TelegramUsername string            `json:"telegramUsername"`
	Miners           []MinerConfigJson `json:"miners"`
}

type MinerResult struct {
	Name    string
	Success bool
	Result  string
}

type MinersResult struct {
	Miners []MinerResult
}
