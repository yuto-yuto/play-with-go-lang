package main

import (
	channelcontext "play-with-go-lang/channel_context"
	"play-with-go-lang/enum"
	packagestructure "play-with-go-lang/package_structure"
	"play-with-go-lang/scanner"
	"play-with-go-lang/utils"
)

func unused(x ...interface{}) {}

func main() {
	unused(scanner.ErrScannerNotFound,
		utils.ConvertByteToString,
		channelcontext.RunChannelContext,
		enum.RunEnum,
		packagestructure.PrintGreeting)

	// scanner.Run()
	// utils.ConvertByteToString()
	// utils.StringHandling()
	// utils.HandleJson()
	// utils.DoStructTest()
	// channelcontext.RunChannelContext()
	// channelcontext.RunChannelDirection()
	// utils.RunOptionalParams()
	// enum.RunEnum()
	// utils.FormatTime()
	// utils.SliceTest()
	// utils.CreateFile()
	// packagestructure.PrintGreeting()
	// utils.RunArbitraryNumberOfArgs()
	// utils.GetValueFromInterface()
	// utils.SplitString()
	// utils.ConvertBetweenStringAndNumber()
	// utils.RunMap()
	// utils.ReplaceString()
	// utils.RunRegex()
	// utils.RunRange()
	// utils.RunDefer()
	// utils.RunServer()
	// utils.RunTilda()
	// utils.ToFromJSON()
	// utils.RunCustomError()
	// utils.RunAtomic()
	// utils.RunMultiAssign()
	// utils.RunTimer()
	// utils.RunMemory()
	// utils.RunDynamicAccess()
	// utils.RunDeleteFunc()
	// utils.RunCond()
	// utils.RunMethodTest()
	// utils.RunOnce()
	utils.SetDefaultValue()

	// utils.RunSliceArrayDeclaration()
}
