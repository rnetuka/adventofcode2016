package day10

type Factory struct {
	bots	 map[int]*Bot
	outputs  map[int]*Output
	listener *ComparisonListener
}

func NewFactory() *Factory {
	factory := new(Factory)
	factory.bots = make(map[int]*Bot)
	factory.outputs = make(map[int]*Output)
	return factory
}

func (factory *Factory) Bot(id int) *Bot {
	if factory.bots[id] == nil {
		bot := new(Bot)
		bot.id = id
		bot.listener = factory.listener
		factory.bots[id] = bot
	}
	return factory.bots[id]
}

func (factory *Factory) Output(id int) *Output {
	if factory.outputs[id] == nil {
		factory.outputs[id] = new(Output)
	}
	return factory.outputs[id]
}

func (factory *Factory) WorkFinished() bool {
	for _, bot := range factory.bots {
		if len(bot.instructions) > 0 {
			return false
		}
	}
	return true
}