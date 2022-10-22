// Copyright 2018 Sergey Novichkov. All rights reserved.
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package gelf

import (
	"os"

	gzZap "github.com/gozix/zap/v3"

	gelf "github.com/snovichkov/zap-gelf"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type gelfCore struct{}

// gelfCore implements zap.CoreFactory.
var _ gzZap.CoreFactory = (*gelfCore)(nil)

func (g *gelfCore) Name() string {
	return "gelf"
}

func (g *gelfCore) New(cfg *viper.Viper, path string) (_ zapcore.Core, err error) {
	var (
		key     = path + ".level"
		options = make([]gelf.Option, 0, 3)
	)

	if cfg.IsSet(key) {
		options = append(options, gelf.LevelString(cfg.GetString(key)))
	}

	key = path + ".addr"
	if cfg.IsSet(key) {
		options = append(options, gelf.Addr(cfg.GetString(key)))
	}

	var host = cfg.GetString(path + ".host")
	if len(host) == 0 {
		if host, err = os.Hostname(); err != nil {
			return nil, err
		}
	}

	options = append(options, gelf.Host(host))

	return gelf.NewCore(options...)
}
