package chat

type ChatUsecase interface {
}

type chatUsecase struct {
	chatRepository ChatRepository
}

func NewChatUsecase(repository ChatRepository) ChatUsecase {
	return &chatUsecase{}
}
