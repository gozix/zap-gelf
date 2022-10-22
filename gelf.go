// Copyright 2018 Sergey Novichkov. All rights reserved.
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package gelf

import (
	"github.com/gozix/di"
	gzGlue "github.com/gozix/glue/v3"
	gzViper "github.com/gozix/viper/v3"
	gzZap "github.com/gozix/zap/v3"
)

// Bundle implements the glue.Bundle interface.
type Bundle struct{}

// BundleName is default definition name.
const BundleName = "zap-gelf"

// Bundle implements glue.Bundle.
var _ gzGlue.Bundle = (*Bundle)(nil)

// NewBundle create bundle instance.
func NewBundle() *Bundle {
	return new(Bundle)
}

// Name implements the glue.Bundle interface.
func (*Bundle) Name() string {
	return BundleName
}

func (b *Bundle) DependsOn() []string {
	return []string{
		gzZap.BundleName,
		gzViper.BundleName,
	}
}

// Build implements the glue.Bundle interface.
func (b *Bundle) Build(builder di.Builder) error {
	return builder.Provide(b.provideGelfCore, gzZap.AsCoreFactory())
}

func (b *Bundle) provideGelfCore() gzZap.CoreFactory {
	return &gelfCore{}
}
