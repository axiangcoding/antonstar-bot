package cardfight

type CardAction interface {
	displayName() string

	TakeStepWithCar(item *CardCarItem) string
	// StartFight 入场介绍
	StartFight() string
	IsDead() (bool, string)
	ModuleStatus() string
}

type CardCarAction interface {
	CardAction
	attack(enemy *CardCarItem) string
	repair() string
}

type matchAction interface {
	generateText(steps []string) string
	decideOrderBothCar() (*CardCarItem, *CardCarItem)

	Fight() string
}
