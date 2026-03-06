package response

import "hab/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
