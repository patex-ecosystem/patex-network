package compressor

import (
	"github.com/patex-ecosystem/patex-network/pt-node/rollup/derive"
)

type FactoryFunc func(Config) (derive.Compressor, error)

const RatioKind = "ratio"
const ShadowKind = "shadow"

var Kinds = map[string]FactoryFunc{
	RatioKind:  NewRatioCompressor,
	ShadowKind: NewShadowCompressor,
}

var KindKeys []string

func init() {
	for k := range Kinds {
		KindKeys = append(KindKeys, k)
	}
}
