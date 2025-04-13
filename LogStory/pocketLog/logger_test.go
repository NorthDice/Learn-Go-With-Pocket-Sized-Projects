package pocketLog_test

import "github.com/NorthDice/LogStory/pocketLog"

func ExampleLogger_Debugf() {
	debugLogger := pocketLog.New(pocketLog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	//Output: Hello, world
}
