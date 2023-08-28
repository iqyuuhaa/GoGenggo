package process

import "time"

type ChatbotHistoryProcess struct {
	Method                   string
	Identifier               string
	Datetime                 time.Time
	ProcessTime              float64
	DialogflowProcessTime    float64
	NonDialogflowProcessTime float64
	SendMessageProcessTime   float64
}
