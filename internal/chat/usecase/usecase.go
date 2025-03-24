package chat

import repository "golang-clean-architecture/internal/chat/repository"

type ChatUsecase interface {
}

type chatUsecase struct {
	chatRepository repository.ChatRepository
}

func NewChatUsecase(repository repository.ChatRepository) ChatUsecase {
	return &chatUsecase{}
}
