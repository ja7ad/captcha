package captcha

type Challenge struct {
	Success        bool     `json:"success"`
	ChallengeTs    string   `json:"challenge_ts"`
	Hostname       string   `json:"hostname,omitempty"`
	ErrorCodes     []string `json:"error-codes,omitempty"`
	ApkPackageName string   `json:"apk_package_name,omitempty"`
	Credit         bool     `json:"credit,omitempty"`
	Score          float64  `json:"score,omitempty"`
	ScoreReason    []any    `json:"score_reason,omitempty"`
}

type Provider string

const (
	GoogleProvider     Provider = "google"
	HCaptchaProvider   Provider = "hcaptcha"
	CloudflareProvider Provider = "cloudflare"
)

func (p Provider) String() string {
	return string(p)
}
