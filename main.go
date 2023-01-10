package main

import (
	pachdClient "github.com/pachyderm/pachyderm/src/client"
	pach2 "github.com/pachyderm/pachyderm/v2/src/client"
)

func main() {
	_, _ = pachdClient.NewFromAddress("")
	_, _ = pach2.NewFromURI("")
}
git init