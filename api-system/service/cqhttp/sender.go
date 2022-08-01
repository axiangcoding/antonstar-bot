package cqhttp

type SendGroupMsgForm struct {
	MessagePrefix string `json:"message_prefix,omitempty"`
	GroupId       int64  `json:"group_id,omitempty"`
	Message       string `json:"message,omitempty"`
}

type CommonResponse struct {
	Status  string `json:"status,omitempty"`
	Retcode int    `json:"retcode,omitempty"`
}
