package main

// CommandHandler обрабатывает команды и сохраняет события
type CommandHandler struct {
	EventStore *EventStore
}

// NewCommandHandler создает новый экземпляр CommandHandler
func NewCommandHandler(es *EventStore) *CommandHandler {
	return &CommandHandler{
		EventStore: es,
	}
}

// Handle обрабатывает команду и возвращает результирующие события
func (ch *CommandHandler) Handle(cmd Command) error {
	events, err := cmd.Execute(ch.EventStore)
	if err != nil {
		return err
	}
	ch.EventStore.AddEvents(events)
	return nil
}
