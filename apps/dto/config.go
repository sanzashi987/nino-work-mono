package dto

type GeneralConfig struct {
	Avatar            string `json:"avatar"`
	SendKey           string `json:"send_key"`
	Theme             string `json:"theme"`
	Language          string `json:"language"`
	FontSize          string `json:"font_size"`
	SendPreviewBubble bool   `json:"send_preview_bubble"`
}

type MaskConfig struct {
	StartWithMask bool `json:"start_with_mask"`
	HidePresets   bool `json:"hide_presets"`
}

type ServiceConfig struct {
	APIKey string `json:"api_key"`
}

type UserPreference struct {
	GeneralConfig GeneralConfig `json:"general_config"`
	MaskConfig    MaskConfig    `json:"mask_config"`
	ChatConfig    ChatConfig    `json:"chat_config"`
	ServiceConfig ServiceConfig `json:"service_config"`
}

type DialogPreference struct {
	GPTAvatar  string     `json:"gpt_avatar"`
	Title      string     `json:"title"`
	UseGlobal  bool       `json:"use_global"`
	ChatConfig ChatConfig `json:"chat_config"`
}

type ChatConfig struct {
	Model             string  `json:"model"`
	Temperature       float32 `json:"temperature"`
	MaxTokens         int     `json:"max_tokens"`
	TopP              float32 `json:"top_p,omitempty"`
	PresencePenalty   float32 `json:"presence_penalty"`
	FrequencyPenalty  float32 `json:"frequency_penalty,omitempty"`
	HistoryCount      int     `json:"history_count"`
	CompressThreshold int     `json:"compress_threshold"`
	Memory            bool    `json:"memory"`
}
