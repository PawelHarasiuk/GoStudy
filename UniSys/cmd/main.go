package main

import "UniSys/uni"

const INPUT_PATH = "input/dane.csv"
const OUTPUT_PATH = "output/"
const FORMAT = "json"

func main() {
	uni.Export(INPUT_PATH, OUTPUT_PATH, FORMAT)
}
