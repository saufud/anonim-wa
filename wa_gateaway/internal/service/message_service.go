package service

import (
	"anon-wa/internal/util"
	"anon-wa/internal/sender"
)

type MessageService struct {
	Sender    sender.MessageSender
	Watermark string
}

func New(sender sender.MessageSender, watermark string) *MessageService {
	return &MessageService{Sender: sender, Watermark: watermark}
}

func (s *MessageService) Send(target, content string) error {
	finalMsg := util.Build(content, s.Watermark)
	return s.Sender.Send(target, finalMsg)
}
