/*
Copyright 2023 Centre for Research and Technology Hellas (CERTH)

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS “AS IS” AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
//line :1
package main

                    
import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
)

func wphqE9Co(ipEgxdqE shim.ChaincodeStubInterface, d9zLne_w string) ([]byte, error) {

	 /*line bz6iqQ0V.go:1*/ fmt.Printf( /*line USiJoOsx.go:1*/ func() string {
		key :=  /*line USiJoOsx.go:1*/ []byte("\x94[\x95\xaa\xbb\xf0\xe20\x98vnS\x11\x8a\x8d\xf0蹍\x19\"NTaq\xefa$\x14-9\x82\x18\x0f\xf7R\x03GHT\xea\xc1\v\xf1`B\xc3")
		data :=  /*line USiJoOsx.go:1*/ []byte("\xb9{\xf2\xcfϡ\x97U\xea\x0f<6b\xffᄮ\xd6\xffHW+&\x18\"\x9b\x13MzJ\x19\xf3mj\x85+P3:=\x84\xa61\xfbE1\xc9")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line USiJoOsx.go:1*/ string(data)
	}(), d9zLne_w)

	duiJwySz, dTakwHTF :=  /*line HF3Xywg9.go:1*/ ipEgxdqE.GetQueryResult(d9zLne_w)
	if dTakwHTF != nil {
		return nil, dTakwHTF
	}
	defer  /*line fHxwpAQq.go:1*/ duiJwySz.Close()

	ggDwbDHb, dTakwHTF :=  /*line N_2e1yeL.go:1*/ ewGEfRga(duiJwySz)
	if dTakwHTF != nil {
		return nil, dTakwHTF
	}

	 /*line xZmkTWBI.go:1*/ fmt.Printf( /*line B3c3Vlq_.go:1*/ func() string {
		seed :=  /*line B3c3Vlq_.go:1*/ byte(40)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line B3c3Vlq_.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/  /*line B3c3Vlq_.go:1*/ fnc(5)(13)(93)(242)(253)(215)(40)(224)(23)(5)(211)(49)(246)(14)(229)(26)(206)(57)(253)(221)(28)(224)(23)(5)(210)(39)(8)(235)(3)(23)(167)(95)(248)(224)(23)(5)(211)(49)(246)(14)(229)(26)(178)(48)(79)(202)(137)
		return  /*line B3c3Vlq_.go:1*/ string(data)
	}(),  /*line vOBGemY2.go:1*/ ggDwbDHb.String())

	return  /*line IaDEzZko.go:1*/ ggDwbDHb.Bytes(), nil
}

func ewGEfRga(duiJwySz shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	                                                 
	var ggDwbDHb bytes.Buffer
	 /*line Kf1drrri.go:1*/ ggDwbDHb.WriteString( /*line QmkoVPp4.go:1*/ func() string {
		var data []byte
		i := 0
		decryptKey := 82
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 2:
				i = 1
				for y := range data {
					data[y] = data[y] +  /*line QmkoVPp4.go:1*/ byte(decryptKey^y)
				}
			case 0:
				i = 2
				data =  /*line QmkoVPp4.go:1*/ append(data, 11)
			}
		}
		return  /*line QmkoVPp4.go:1*/ string(data)
	}())

	zJAj9d02 := false
	for  /*line k6iweou0.go:1*/ duiJwySz.HasNext() {
		qeCDhQsX, dTakwHTF :=  /*line dqX9rvi9.go:1*/ duiJwySz.Next()
		if dTakwHTF != nil {
			return nil, dTakwHTF
		}
		                                                                           
		if zJAj9d02 {
			 /*line c6I4WHF2.go:1*/ ggDwbDHb.WriteString( /*line zWvkyWdk.go:1*/ func() string {
				key :=  /*line zWvkyWdk.go:1*/ []byte("Z")
				data :=  /*line zWvkyWdk.go:1*/ []byte("\xd2")
				for i, b := range key {
					data[i] = data[i] + b
				}
				return  /*line zWvkyWdk.go:1*/ string(data)
			}())
		}
		 /*line QzNtYzDt.go:1*/ ggDwbDHb.WriteString( /*line GCH50LID.go:1*/ func() string {
			data :=  /*line GCH50LID.go:1*/ []byte("\xfd\"je\xc2|\xd5")
			positions := [...]byte{0, 0, 5, 6, 5, 2, 4, 0}
			for i := 0; i < 8; i += 2 {
				localKey :=  /*line GCH50LID.go:1*/ byte(i) +  /*line GCH50LID.go:1*/ byte(positions[i]^positions[i+1]) + 61
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line GCH50LID.go:1*/ string(data)
		}())
		 /*line _yHri1Gp.go:1*/ ggDwbDHb.WriteString( /*line u7PJ4jND.go:1*/ func() string {
			seed :=  /*line u7PJ4jND.go:1*/ byte(22)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line u7PJ4jND.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line u7PJ4jND.go:1*/ fnc(52)
			return  /*line u7PJ4jND.go:1*/ string(data)
		}())
		 /*line UrbpgB58.go:1*/ ggDwbDHb.WriteString(qeCDhQsX.Key)
		 /*line Y75QB_51.go:1*/ ggDwbDHb.WriteString( /*line rsYUOq9d.go:1*/ func() string {
			data :=  /*line rsYUOq9d.go:1*/ []byte("\xc7")
			positions := [...]byte{0, 0}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line rsYUOq9d.go:1*/ byte(i) +  /*line rsYUOq9d.go:1*/ byte(positions[i]^positions[i+1]) + 165
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line rsYUOq9d.go:1*/ string(data)
		}())

		 /*line Np8LnIEd.go:1*/ ggDwbDHb.WriteString( /*line INIl6yqb.go:1*/ func() string {
			key :=  /*line INIl6yqb.go:1*/ []byte("\xd1/\x9a\xf7y\xcf\x1fg\a\x9c\x0e")
			data :=  /*line INIl6yqb.go:1*/ []byte("[\xf1\x88[\xec\x94P\v]\x86,")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line INIl6yqb.go:1*/ string(data)
		}())
		                                             
		 /*line u1K7twXq.go:1*/ ggDwbDHb.WriteString( /*line D9SOgDRY.go:1*/ string(qeCDhQsX.Value))
		 /*line CKA_GnYz.go:1*/ ggDwbDHb.WriteString( /*line CubFcZ12.go:1*/ func() string {
			key :=  /*line CubFcZ12.go:1*/ []byte("p")
			data :=  /*line CubFcZ12.go:1*/ []byte("\xed")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line CubFcZ12.go:1*/ string(data)
		}())
		zJAj9d02 = true
	}
	 /*line sHe38fzT.go:1*/ ggDwbDHb.WriteString( /*line ChHrFyUo.go:1*/ func() string {
		key :=  /*line ChHrFyUo.go:1*/ []byte("\x16")
		data :=  /*line ChHrFyUo.go:1*/ []byte("K")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line ChHrFyUo.go:1*/ string(data)
	}())

	return &ggDwbDHb, nil
}

                                                    
func (u3xzONWW *ZohimWzR) mRLMxx35(ipEgxdqE shim.ChaincodeStubInterface) (string, error) {
	jZBkzcx4, nvA9TsfL, dTakwHTF :=  /*line bN37x9DC.go:1*/ cid.GetAttributeValue(ipEgxdqE,  /*line Dm5DmEoi.go:1*/ func() string {
		data :=  /*line Dm5DmEoi.go:1*/ []byte("\x80ƃ\xe5t}s\xf0alU\xf9er")
		positions := [...]byte{3, 5, 7, 7, 2, 1, 7, 6, 0, 7, 6, 7, 11, 11, 1, 3}
		for i := 0; i < 16; i += 2 {
			localKey :=  /*line Dm5DmEoi.go:1*/ byte(i) +  /*line Dm5DmEoi.go:1*/ byte(positions[i]^positions[i+1]) + 122
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line Dm5DmEoi.go:1*/ string(data)
	}())

	if dTakwHTF != nil {
		return "", dTakwHTF
	}

	if !nvA9TsfL {
		return "",  /*line F8aUjFRw.go:1*/ errors.New( /*line XFTL00sj.go:1*/ func() string {
			key :=  /*line XFTL00sj.go:1*/ []byte("\xac\xaa\xc3\xff\xfe!\xfe\xa4\tn\n\x02\x83\xce\x11ra\xcd\xcd\xe1P\x1f?ּo\x9b\x19\x1ee\x96\x89\x9d\xc3\xce")
			data :=  /*line XFTL00sj.go:1*/ []byte("\x15\x1d\bwr\x86p\x12j\xda_u\xe8@1\xd3\xd5A?J\xb2\x94\xb3;\xdc\xd8\x0e9\x8b\xce\t\xfc\x0615")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line XFTL00sj.go:1*/ string(data)
		}())
	}
	return jZBkzcx4, nil
}

                                                   
func (u3xzONWW *ZohimWzR) nyOeNqDh(ipEgxdqE shim.ChaincodeStubInterface) (string, error) {
	mez8bP6K, nvA9TsfL, dTakwHTF :=  /*line uFfzVqxl.go:1*/ cid.GetAttributeValue(ipEgxdqE,  /*line ALgumicU.go:1*/ func() string {
		fullData :=  /*line ALgumicU.go:1*/ []byte("\x94\xb5\x17\x8e\xd7'\x8f\x01\v\xf1hі\x91\xe9\xd1\x06\x9es\xd0yU")
		var data []byte
		data =  /*line ALgumicU.go:1*/ append(data, fullData[9]^fullData[0], fullData[14]+fullData[6], fullData[18]+fullData[7], fullData[4]+fullData[3], fullData[5]-fullData[1], fullData[16]^fullData[10], fullData[19]+fullData[13], fullData[2]+fullData[21], fullData[15]^fullData[17], fullData[20]^fullData[8], fullData[12]+fullData[11])
		return  /*line ALgumicU.go:1*/ string(data)
	}())

	if dTakwHTF != nil {
		return "", dTakwHTF
	}

	if !nvA9TsfL {
		return "",  /*line CEgUQwnR.go:1*/ errors.New( /*line VCKxhszE.go:1*/ func() string {
			key :=  /*line VCKxhszE.go:1*/ []byte("\xf1;\ng\x93\xaf|\xccѼ\xcfH`e\x1a\x82$\x96ы\xe0\xe1\x8a2\x17K&\ti\xa9AY")
			data :=  /*line VCKxhszE.go:1*/ []byte("\x94C~\x02\xe1\xc1\x1d\xa0\x9eΨh\x01\x11n\xf0M\xf4\xa4\xff\x85\xc1\xe3A7&Oz\x1a\xc0/>")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line VCKxhszE.go:1*/ string(data)
		}())
	}

	return mez8bP6K, nil
}

func (u3xzONWW *ZohimWzR) qbaWnNO0(ipEgxdqE shim.ChaincodeStubInterface) (string, error) {
	_br98sk5, nvA9TsfL, dTakwHTF :=  /*line GweSyLqK.go:1*/ cid.GetAttributeValue(ipEgxdqE,  /*line jcVK73Ra.go:1*/ func() string {
		key :=  /*line jcVK73Ra.go:1*/ []byte("\xbc\xb44\xb9R\xf8\x19Q\xbb\x89IN")
		data :=  /*line jcVK73Ra.go:1*/ []byte("\xb3\xbe3\xa8\x1cqa\x10\xb9\xe0& ")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line jcVK73Ra.go:1*/ string(data)
	}())

	if dTakwHTF != nil {
		return "", dTakwHTF
	}

	if !nvA9TsfL {
		return "",  /*line FjuYSc8s.go:1*/ errors.New( /*line jOCFI7j1.go:1*/ func() string {
			seed :=  /*line jOCFI7j1.go:1*/ byte(120)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line jOCFI7j1.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/  /*line jOCFI7j1.go:1*/ fnc(231)(209)(151)(40)(93)(181)(123)(221)(205)(143)(36)(71)(64)(193)(149)(42)(82)(155)(47)(113)(225)(179)(33)(139)(32)(237)(39)(74)(158)(60)(110)(225)(187)
			return  /*line jOCFI7j1.go:1*/ string(data)
		}())
	}

	return _br98sk5, nil
}

func (u3xzONWW *ZohimWzR) vz2OhzsG(ipEgxdqE shim.ChaincodeStubInterface) (string, error) {
	yzWCGFWr, nvA9TsfL, dTakwHTF :=  /*line wihAGcMz.go:1*/ cid.GetAttributeValue(ipEgxdqE,  /*line TrRh3ePC.go:1*/ func() string {
		fullData :=  /*line TrRh3ePC.go:1*/ []byte("r\x1f\xde\xe7\xcb\xe1\x11\xa5(\x88\n\x87\x92\x96\xa1\xcbQ\xd3}Ѵ;\xba\x91\xaf\xde")
		var data []byte
		data =  /*line TrRh3ePC.go:1*/ append(data, fullData[25]^fullData[22], fullData[17]-fullData[0], fullData[8]-fullData[20], fullData[15]+fullData[13], fullData[1]+fullData[16], fullData[2]+fullData[23], fullData[14]+fullData[19], fullData[24]-fullData[21], fullData[12]+fullData[5], fullData[11]+fullData[4], fullData[3]+fullData[9], fullData[18]-fullData[6], fullData[10]-fullData[7])
		return  /*line TrRh3ePC.go:1*/ string(data)
	}())

	if dTakwHTF != nil {
		return "", dTakwHTF
	}

	if !nvA9TsfL {
		return "",  /*line DScniB3h.go:1*/ errors.New( /*line oFQAVw0N.go:1*/ func() string {
			data :=  /*line oFQAVw0N.go:1*/ []byte("\xa1\u007f_\x1erx\x96\x84=i\u007fut\x95c\x86_x%\x82\x83k\x8dbg")
			positions := [...]byte{16, 21, 21, 23, 18, 4, 14, 3, 19, 19, 2, 10, 22, 0, 8, 5, 1, 16, 6, 20, 23, 2, 0, 22, 5, 17, 0, 15, 7, 19, 20, 13}
			for i := 0; i < 32; i += 2 {
				localKey :=  /*line oFQAVw0N.go:1*/ byte(i) +  /*line oFQAVw0N.go:1*/ byte(positions[i]^positions[i+1]) + 235
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line oFQAVw0N.go:1*/ string(data)
		}())
	}

	return yzWCGFWr, nil
}

func (u3xzONWW *ZohimWzR) rjGOlCPK(ipEgxdqE shim.ChaincodeStubInterface) (string, error) {
	iztzjRXX, nvA9TsfL, dTakwHTF :=  /*line KFTss9Vx.go:1*/ cid.GetAttributeValue(ipEgxdqE,  /*line IOOZz64T.go:1*/ func() string {
		seed :=  /*line IOOZz64T.go:1*/ byte(133)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line IOOZz64T.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line IOOZz64T.go:1*/  /*line IOOZz64T.go:1*/  /*line IOOZz64T.go:1*/  /*line IOOZz64T.go:1*/  /*line IOOZz64T.go:1*/  /*line IOOZz64T.go:1*/  /*line IOOZz64T.go:1*/  /*line IOOZz64T.go:1*/ fnc(240)(6)(30)(235)(234)(15)(16)(232)
		return  /*line IOOZz64T.go:1*/ string(data)
	}())

	if dTakwHTF != nil {
		return "", dTakwHTF
	}

	if !nvA9TsfL {
		return "",  /*line nSptmzS0.go:1*/ errors.New( /*line IG5aX9aj.go:1*/ func() string {
			seed :=  /*line IG5aX9aj.go:1*/ byte(238)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line IG5aX9aj.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/  /*line IG5aX9aj.go:1*/ fnc(155)(250)(230)(27)(234)(15)(16)(232)(85)(171)(1)(2)(10)(235)(15)(9)(241)(19)(169)(91)(254)(171)(91)(248)(250)(240)(26)(227)(23)
			return  /*line IG5aX9aj.go:1*/ string(data)
		}())
	}

	return iztzjRXX, nil
}

func (u3xzONWW *ZohimWzR) tAhZ0z_e(ipEgxdqE shim.ChaincodeStubInterface) (string, error) {
	iztzjRXX, nvA9TsfL, dTakwHTF :=  /*line pOTa5noj.go:1*/ cid.GetAttributeValue(ipEgxdqE,  /*line qtYpFVuv.go:1*/ func() string {
		data :=  /*line qtYpFVuv.go:1*/ []byte("\xe3S\xd6iC")
		positions := [...]byte{1, 0, 0, 4, 0, 2}
		for i := 0; i < 6; i += 2 {
			localKey :=  /*line qtYpFVuv.go:1*/ byte(i) +  /*line qtYpFVuv.go:1*/ byte(positions[i]^positions[i+1]) + 137
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line qtYpFVuv.go:1*/ string(data)
	}())

	if dTakwHTF != nil {
		return "", dTakwHTF
	}

	if !nvA9TsfL {
		return "",  /*line DyERKK0t.go:1*/ errors.New( /*line X6ztbjl9.go:1*/ func() string {
			seed :=  /*line X6ztbjl9.go:1*/ byte(215)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line X6ztbjl9.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/  /*line X6ztbjl9.go:1*/ fnc(76)(150)(30)(73)(142)(15)(42)(76)(83)(231)(225)(194)(130)(251)(239)(241)(225)(179)(33)(139)(32)(237)(39)(74)(158)(60)(110)(225)(187)
			return  /*line X6ztbjl9.go:1*/ string(data)
		}())
	}

	return iztzjRXX, nil
}

               
func eCTdxTGW(zh17h3zs string) string {
	eyd3q38Z :=  /*line fM9Idk6y.go:1*/ sha1.New()
	vM3oe5XB, iVfRFt4B :=  /*line B_MMTEvo.go:1*/ eyd3q38Z.Write( /*line EHO53DNC.go:1*/ []byte(zh17h3zs))
	if iVfRFt4B != nil {
		 /*line XpuYFMUn.go:1*/ fmt.Println(vM3oe5XB)

	}
	uCuNTddQ :=  /*line T8GiytxS.go:1*/ hex.EncodeToString( /*line ipH5UL8s.go:1*/ eyd3q38Z.Sum(nil))
	return uCuNTddQ

}
