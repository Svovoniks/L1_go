package main

import (
	"fmt"
	"strconv"
)

type DataProducer interface {
	getDataString() string
}

type Printer struct{}

func (p *Printer) Print(dp DataProducer) {
	fmt.Println(dp.getDataString())
}

type StringDataProvider struct{}

func (sp *StringDataProvider) getDataString() string {
	return "string"
}

type IntDataProvider struct{}

func (ip *IntDataProvider) getDataInt() int {
	return 1
}

type IntAdapter struct {
	ip *IntDataProvider
}

func (ia *IntAdapter) getDataString() string {
	return strconv.Itoa(ia.ip.getDataInt())
}

func main() {
	printer := &Printer{}
	strProvider := &StringDataProvider{}

	printer.Print(strProvider)

	intProvider := &IntDataProvider{}
	intProviderAdapter := &IntAdapter{
		ip: intProvider,
	}

	printer.Print(intProviderAdapter)
}
