package labapi

import "github.com/ethereum/go-ethereum/params"

type Backend interface {
	ChainConfig() *params.ChainConfig
}

type API struct {
	backend Backend
}

func NewAPI(backend Backend) *API {
	return &API{backend: backend}
}

func (api *API) ChainInfo() map[string]interface{} {
	cfg := api.backend.ChainConfig()
	chainID := ""
	var lab *params.LabConfig
	if cfg != nil {
		if cfg.ChainID != nil {
			chainID = cfg.ChainID.String()
		}
		lab = cfg.Lab
	}

	if lab == nil {
		lab = new(params.LabConfig)
	}

	return map[string]interface{}{
		"chainName":   lab.ChainName,
		"chainId":     chainID,
		"nativeToken": lab.NativeToken,
		"fees":        lab.Fees,
		"consensus":   lab.Consensus,
	}
}
