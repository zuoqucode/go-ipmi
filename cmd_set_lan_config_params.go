package ipmi

// 23.1 Set LAN Configuration Parameters Command
type SetLanConfigParamsRequest struct {
	ChannelNumber uint8
	ParamSelector LanParamSelector
	ConfigData    []byte
}

type SetLanConfigParamsResponse struct {
	// emtpy
}

func (req *SetLanConfigParamsRequest) Pack() []byte {
	out := make([]byte, 2+len(req.ConfigData))
	packUint8(req.ChannelNumber, out, 0)
	packUint8(uint8(req.ParamSelector), out, 1)
	packBytes(req.ConfigData, out, 2)
	return out
}

func (req *SetLanConfigParamsRequest) Command() Command {
	return CommandSetLanConfigParams
}

func (res *SetLanConfigParamsResponse) CompletionCodes() map[uint8]string {
	return map[uint8]string{
		0x80: "parameter not supported.",
		0x81: "attempt to set the 'set in progress' value (in parameter #0) when not in the 'set complete' state.",
		0x82: "attempt to write read-only parameter",
		0x83: "attempt to read write-only parameter",
	}
}

func (res *SetLanConfigParamsResponse) Unpack(msg []byte) error {
	return nil
}

func (res *SetLanConfigParamsResponse) Format() string {
	return ""
}

// Todo
func (c *Client) SetLanConfigParams(channum uint8, paramData LanParamSelector, data []byte) (response *SetLanConfigParamsResponse, err error) {
	request := &SetLanConfigParamsRequest{
		ChannelNumber: channum,
		ParamSelector: paramData,
		ConfigData:    data,
	}
	response = &SetLanConfigParamsResponse{}
	err = c.Exchange(request, response)
	return
}
