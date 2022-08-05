package etherscan

import "context"

type Client interface {
	GetSourceCode(ctx context.Context, address string) (*GetSourceCodeResponse, error)
}
