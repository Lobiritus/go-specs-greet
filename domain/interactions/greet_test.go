package interactions

import (
	"testing"

	go_specs_greet "github.com/lobiritus/go-specs-greet"
	"github.com/lobiritus/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(go_specs_greet.Greet))

}
