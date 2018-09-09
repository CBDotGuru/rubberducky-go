package rubberducky

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestBuildPackage(t *testing.T) {
	os.Setenv("TEST", "true")
	var got string
	var err error
	got, err = buildPackage("Test")
	if err != nil {
		fmt.Println(err)
	}
	testmanifest := "../../../test/testdata/testmanifest.json"
	file, err := os.Open(filepath.FromSlash(testmanifest))
	if err != nil {
		return
	}
	info, _ := file.Stat()
	manifest := make([]byte, info.Size())
	_, err = file.Read(manifest)
	if err != nil {
		return
	}
	manifest = manifest[:len(manifest)-1]
	want := string(manifest)

	if got != want {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}
}

func TestWritePackage(t *testing.T) {
	os.Setenv("TEST", "true")
	var got error
	json := `{"contract_types":{"BasicMathLib":{"abi":[{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"times","outputs":[{"name":"err","type":"bool"},{"name":"res","type":"uint256"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"dividedBy","outputs":[{"name":"err","type":"bool"},{"name":"i","type":"uint256"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"plus","outputs":[{"name":"err","type":"bool"},{"name":"res","type":"uint256"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}],"name":"minus","outputs":[{"name":"err","type":"bool"},{"name":"res","type":"uint256"}],"payable":false,"stateMutability":"pure","type":"function"}],"compiler":{"name":"solc","version":"0.4.24+commit.e67f0147.Emscripten.clang"},"deployment_bytecode":{"bytecode":"0x6102a1610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f3007300000000000000000000000000000000000000003014608060405260043610610079576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631d3b9edf1461007e57806366098d4f146100c7578063e39bbf6814610110578063f4f3bdc114610159575b600080fd5b6100a660048036038101908080359060200190929190803590602001909291905050506101a2565b60405180831515151581526020018281526020019250505060405180910390f35b6100ef60048036038101908080359060200190929190803590602001909291905050506101d0565b60405180831515151581526020018281526020019250505060405180910390f35b6101386004803603810190808035906020019092919080359060200190929190505050610203565b60405180831515151581526020018281526020019250505060405180910390f35b610181600480360381019080803590602001909291908035906020019092919050505061023f565b60405180831515151581526020018281526020019250505060405180910390f35b60008082840290508383820414831517600081146101bf576101c8565b60019250600091505b509250929050565b600080828401905082811483821117848483031416600081146101f2576101fb565b60019250600091505b509250929050565b600080600083156000811461021f576001935060009250610236565b848604915060405182602082015260208101519350505b50509250929050565b600080828403905060018482148583101785858401141614600081146102645761026d565b60019250600091505b5092509290505600a165627a7a723058200190864cc40da4490c7e07413b820a1a2d9dbef9a3526bf4bf485f3b11a06ef10029"},"runtime_bytecode":{"bytecode":"0x7300000000000000000000000000000000000000003014608060405260043610610079576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631d3b9edf1461007e57806366098d4f146100c7578063e39bbf6814610110578063f4f3bdc114610159575b600080fd5b6100a660048036038101908080359060200190929190803590602001909291905050506101a2565b60405180831515151581526020018281526020019250505060405180910390f35b6100ef60048036038101908080359060200190929190803590602001909291905050506101d0565b60405180831515151581526020018281526020019250505060405180910390f35b6101386004803603810190808035906020019092919080359060200190929190505050610203565b60405180831515151581526020018281526020019250505060405180910390f35b610181600480360381019080803590602001909291908035906020019092919050505061023f565b60405180831515151581526020018281526020019250505060405180910390f35b60008082840290508383820414831517600081146101bf576101c8565b60019250600091505b509250929050565b600080828401905082811483821117848483031416600081146101f2576101fb565b60019250600091505b509250929050565b600080600083156000811461021f576001935060009250610236565b848604915060405182602082015260208101519350505b50509250929050565b600080828403905060018482148583101785858401141614600081146102645761026d565b60019250600091505b5092509290505600a165627a7a723058200190864cc40da4490c7e07413b820a1a2d9dbef9a3526bf4bf485f3b11a06ef10029"}}},"manifest_version":"2","meta":{"authors":["Modular Inc"],"description":"Basic math operations on Ethereum with overflow protection","license":"MIT","links":{"documentation":"https://github.com/Modular-Network/ethereum-libraries-basic-math"}},"package_name":"ethereum-libraries-basic-math","sources":{"./contracts/BasicMathLib.sol":"https://github.com/Modular-Network/ethereum-libraries-basic-math.git@618d432d34cb294ff0ec8c9825348e592c0a9cad"},"version":"1.2.7"}
`
	got = writePackage("test", json)
	if got != nil {
		t.Fatalf("Got '%v', expected '<nil>'", got)
	}
}
