package decorator

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type MetricsClient interface {
	Inc(key string, value int)
}

type commnadMetricsClient[C any] struct {
	base CommandHandler[C]
	client MetricsClient
}

func (d commnadMetricsClient[C]) Handle(ctx context.Context, cmd C) (err error) {
	start := time.Now()

	actionName := strings.ToLower(generateActionName(cmd))

	defer func() {
		end := time.Since(start)

		d.client.Inc(fmt.Sprintf("command.%s.duration", actionName), int(end.Seconds()))

		if err == nil {
			d.client.Inc(fmt.Sprint("command.%s.success", actionName), 1)
		} else {
			d.client.Inc(fmt.Sprint("command.%s.failure", actionName), 1)
		}
	}()
	
	return d.base.Handle(ctx, cmd)
}