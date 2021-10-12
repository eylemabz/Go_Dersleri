package error_handling

import (
	"fmt"
)

type BorderException struct {
	message   string
	parameter int
}

func (b *BorderException) Error() string {
	return fmt.Sprintf("%d----%s", b.parameter, b.message)
}

func TahminEt1(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &BorderException{"Sırınların dışında kaldı", tahmin}
	}
	return "bldiniz", nil
}
