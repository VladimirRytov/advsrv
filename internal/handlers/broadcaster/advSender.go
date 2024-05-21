package broadcaster

import (
	"encoding/json"

	"github.com/VladimirRytov/advsrv/internal/logging"
)

func (sh *Broadcaster) SendData(data []byte, kind string, action int) error {

	raw, err := json.Marshal(&SendWrapper{Type: kind, Action: action, Entry: sh.b64.ToBase64(data)})
	if err != nil {
		return err
	}
	logging.Logger.Debug("broadcaster.SendData: sending data", "data", string(raw))
	sh.buf <- raw
	return nil
}
