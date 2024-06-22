package singbox

import (
	"github.com/sagernet/sing/common/wininet"
)

// 清除系统代理
func DisableProxy() error {
	err := wininet.ClearSystemProxy()
	if err != nil {
		return err
	}
	return nil
}
