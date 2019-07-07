package go_kit_sample

import (
	"context"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

func Handler() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		return strings.ToUpper(request.(string)), nil
	}
}
