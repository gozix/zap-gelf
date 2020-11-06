package gelf

import (
	"os"

	gzViper "github.com/gozix/viper/v2"
	gzZap "github.com/gozix/zap/v2"

	"github.com/sarulabs/di/v2"
	gelf "github.com/snovichkov/zap-gelf"
	"go.uber.org/zap/zapcore"
)

// Bundle implements the glue.Bundle interface.
type Bundle struct{}

// BundleName is default definition name.
const BundleName = "zap-gelf"

// NewBundle create bundle instance.
func NewBundle() *Bundle {
	return new(Bundle)
}

// Name implements the glue.Bundle interface.
func (*Bundle) Name() string {
	return BundleName
}

// Name implements the glue.BundleDependsOn interface.
func (b *Bundle) DependsOn() []string {
	return []string{
		gzZap.BundleName,
		gzViper.BundleName,
	}
}

// Build implements the glue.Bundle interface.
func (*Bundle) Build(builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: "zap.core.gelf",
		Tags: []di.Tag{{
			Name: gzZap.TagCoreFactory,
			Args: map[string]string{
				gzZap.ArgCoreType: "gelf",
			},
		}},
		Build: func(ctn di.Container) (interface{}, error) {
			return func(path string) (_ zapcore.Core, err error) {
				var cfg *gzViper.Viper
				if err = ctn.Fill(gzViper.BundleName, &cfg); err != nil {
					return nil, err
				}

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
			}, nil
		},
	})
}
