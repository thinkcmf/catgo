package catgo

import "testing"

func TestDump(t *testing.T) {
	Dump(map[string]string{
		"ddd":   "dd",
		"ddddd": "ssss",
	})

	Dump(struct {
		Ddd string
		sss string
	}{
		Ddd: "DDD",
		sss:"sss",
	})
}
