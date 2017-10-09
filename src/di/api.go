package di

import (
	"github.com/morikuni/chat/src/interface/api"
)

func InjectAPI() api.API {
	return api.NewAPI(
		InjectPostingService(),
		InjectChatReader(),
		InjectLogger(),
	)
}
