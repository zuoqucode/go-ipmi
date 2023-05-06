package ipmi

import "sync"

var (
	cmd  Command
	muID sync.Mutex
)

// 20.2 Raw cmd
type RawInfoRequest struct {
	ConfigData []byte
}

type RawInfoResponse struct {
	Data []byte
}

func (req *RawInfoRequest) Command() Command {
	return cmd
}

func (req *RawInfoRequest) Pack() []byte {
	out := make([]byte, len(req.ConfigData))
	packBytes(req.ConfigData, out, 0)
	return out
}

func (res *RawInfoResponse) CompletionCodes() map[uint8]string {
	return map[uint8]string{}
}

func (res *RawInfoResponse) Unpack(msg []byte) error {
	res.Data = msg
	return nil
}

func (res *RawInfoResponse) Format() string {
	return ""
}

func (c *Client) SendRawInfo(mycmd Command, configData []byte) (response *RawInfoResponse, err error) {
	muID.Lock()
	cmd = mycmd
	muID.Unlock()
	request := &RawInfoRequest{
		ConfigData: configData,
	}
	response = &RawInfoResponse{}
	err = c.Exchange(request, response)
	return
}
