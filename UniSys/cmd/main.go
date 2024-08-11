package main

import "UniSys/uni"

const InputPath = "input/dane.csv"
const OutputPath = "output/"
const Format = "json"

func main() {
	uni.Export(InputPath, OutputPath, Format)
}
