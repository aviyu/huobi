package main

import (
	"github.com/aviyu/huobi/cmd/accountclientexample"
	"github.com/aviyu/huobi/cmd/accountwebsocketclientexample"
	"github.com/aviyu/huobi/cmd/algoorderclientexample"
	"github.com/aviyu/huobi/cmd/commonclientexample"
	"github.com/aviyu/huobi/cmd/crossmarginclientexample"
	"github.com/aviyu/huobi/cmd/etfclientexample"
	"github.com/aviyu/huobi/cmd/isolatedmarginclientexample"
	"github.com/aviyu/huobi/cmd/marketclientexample"
	"github.com/aviyu/huobi/cmd/marketwebsocketclientexample"
	"github.com/aviyu/huobi/cmd/orderclientexample"
	"github.com/aviyu/huobi/cmd/orderwebsocketclientexample"
	"github.com/aviyu/huobi/cmd/stablecoinclientexample"
	"github.com/aviyu/huobi/cmd/subuserclientexample"
	"github.com/aviyu/huobi/cmd/walletclientexample"
	"github.com/aviyu/huobi/logging/perflogger"
)

func main() {
	runAll()
}

// Run all examples
func runAll() {
	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	algoorderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	subuserclientexample.RunAllExamples()
	stablecoinclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
	marketwebsocketclientexample.RunAllExamples()
	accountwebsocketclientexample.RunAllExamples()
	orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)

	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}
