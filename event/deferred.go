package event

import internalevent "github.com/kintar/ebitenui/v2/internal/event"

// ExecuteDeferred processes the queue of deferred actions and executes them. This should only be called by UI.
// Additionally, it can be called in unit tests to process events programmatically.
func ExecuteDeferred() {
	internalevent.ExecuteDeferred()
}
