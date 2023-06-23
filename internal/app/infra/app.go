package infra

import "time"

type (

	// @envconfig (prefix:"APP")
	App struct {
		Name     string `envconfig:"NAME" default:"Midgard Repayments" required:"true"`
		Key      string `envconfig:"KEY" default:"midgard-repayment" required:"true"`
		Debug    bool   `envconfig:"DEBUG" default:"false" required:"true"`
		Env      string `envconfig:"ENV" default:"local" required:"true"`
		Version  string `envconfig:"VERSION" default:"v1.0.0" required:"true"`
		TimeZone string `envconfig:"TIMEZONE" default:"Asia/Jakarta" required:"true"`

		Address      string        `envconfig:"ADDRESS" default:":8089" required:"true"`
		ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" default:"5s"`
		WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"10s"`
	}
)
