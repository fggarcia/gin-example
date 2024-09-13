package memory

import (
	"log/slog"
	"testing"

	"github.com/KimMachineGun/automemlimit/memlimit"
)

func TestAutoMemLimit(T *testing.T) {
	memlimit.SetGoMemLimitWithOpts(
		memlimit.WithRatio(0.9),
		memlimit.WithProvider(memlimit.ApplyFallback(memlimit.FromCgroup, memlimit.FromSystem)),
		memlimit.WithLogger(slog.Default()),
	)
}
