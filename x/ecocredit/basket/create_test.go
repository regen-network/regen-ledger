package basket_test

import (
	"context"

	"github.com/cucumber/godog"
)

func aBasket(ctx context.Context, arg1 string) error {
	return godog.ErrPending
}

func aBasketWithAutoretireDisables(arg1 string) error {
	return godog.ErrPending
}

func something() error {
	return godog.ErrPending
}

func theOtherThing() error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	})
	ctx.Step(`^a basket "([^"]*)"$`, aBasket)
	ctx.Step(`^a basket "([^"]*)", with auto-retire disables$`, aBasketWithAutoretireDisables)
	ctx.Step(`^something$`, something)
	ctx.Step(`^the other thing$`, theOtherThing)
}
