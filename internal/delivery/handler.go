package delivery

type handler struct {
	telebotHandler *telebotHandler
}

func SetupHandler(container *Container) {
	h := &handler{
		telebotHandler: NewTelebotHandler(container.Bot, container.TelebotService),
	}

	h.telebotHandler.TelebotHandler()
}
