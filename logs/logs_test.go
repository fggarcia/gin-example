package logs

import (
	"fmt"
	"github.com/proullon/ramsql/engine/log"
	"testing"
)

func TestLogs(t *testing.T) {
	log.Debug(func() string {
		fmt.Printf("selecting cheapest method for pack %+v\n", []string{"a", "b", "c"})
		// Esta línea solo se ejecutará si el nivel de logging es Debug o inferior
		return fmt.Sprintf("selecting cheapest from multiple methods for pack %+v", []string{"a", "b", "c"})
	}())

	log.Info("Este mensaje se mostrará siempre que el nivel de logging sea Info o inferior")
}
