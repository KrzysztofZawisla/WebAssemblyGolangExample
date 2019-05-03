package main

import (
	"fmt"
	"syscall/js"
	"encoding/hex"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/md5"
	"math/rand"
)

func main() {
	done := make(chan struct{}, 0)
	startUpText := "WebAssembly is ready to work"
	fmt.Println(startUpText)
	registerCallbacks()
	<-done
}

func registerCallbacks() {
	js.Global().Set("wasmAdd", js.Func(wasmAdd))
	js.Global().Set("wasmSubtract", js.Func(wasmSubtract))
	js.Global().Set("wasmMultiply", js.Func(wasmMultiply))
	js.Global().Set("wasmHash", js.Func(wasmHash))
	js.Global().Set("wasmMathRandom", js.Func(wasmMathRandom))
	js.Global().Set("wasmDivide", js.Func(wasmDivide))
}

var wasmAdd js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	var output int = args[0].Int()
	for i := 1; i<len(args); i++ {
		output += args[i].Int()
	}
	return output
})

var wasmSubtract js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	var output int = args[0].Int()
	for i := 1; i<len(args); i++ {
		output -= args[i].Int()
	}
	return output
})

var wasmMultiply js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	var output int = args[0].Int()
	for i := 1; i<len(args); i++ {
		output *= args[i].Int()
	}
	return output
})

var wasmDivide js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	var output int = args[0].Int()
	for i := 1; i<len(args); i++ {
		output /= args[i].Int()
	}
	return output
})

var wasmHash js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var argEncryptionMethod string
	if len(args) == 0 {
		return nil
	} else if len(args) == 1 {
		argEncryptionMethod = "SHA256"
	} else {
		argEncryptionMethod = args[1].String()
	}
	argString := args[0].String()
	if argEncryptionMethod == "SHA256" {
		crypto := sha256.New()
		crypto.Write([]byte(argString))
		hexValue := hex.EncodeToString(crypto.Sum(nil))
		return hexValue
	} else if argEncryptionMethod == "SHA512" {
		crypto := sha512.New()
		crypto.Write([]byte(argString))
		hexValue := hex.EncodeToString(crypto.Sum(nil))
		return hexValue
	} else if argEncryptionMethod == "MD5" {
		crypto := md5.New()
		crypto.Write([]byte(argString))
		hexValue := hex.EncodeToString(crypto.Sum(nil))
		return hexValue
	}
	return nil
})

var wasmMathRandom js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var min int
	if len(args) == 0 {
		return nil
	} else if len(args) == 1 {
		min = 0
	} else {
		min = args[0].Int()
	}
	max := args[1].Int()
	output := rand.Intn((max - min) + 1)
	output += min
	return output
})