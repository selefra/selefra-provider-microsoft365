package main

import (
	"github.com/selefra/selefra-provider-microsoft365/provider"
	"github.com/selefra/selefra-provider-sdk/grpc/serve"
)

func main() {
	myProvider := provider.GetProvider()
	serve.Serve(myProvider.Name, myProvider)
}
