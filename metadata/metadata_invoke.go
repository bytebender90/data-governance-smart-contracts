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
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"reflect"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (kuuXUPaZ *JPyzqDxq) aolkF3KY(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) pb.Response {
	var xonUbr_n error
	var tvV5RxQz []string
	var yfxmHRDX string

	aTwOY3WH :=  /*line sIHFAXpA.go:1*/ func() string {
		fullData :=  /*line sIHFAXpA.go:1*/ []byte("3\xfe\xa0O\x00\xc1M\xd2F\x9b.&\x02v\x14B")
		var data []byte
		data =  /*line sIHFAXpA.go:1*/ append(data, fullData[3]+fullData[1], fullData[4]-fullData[9], fullData[8]-fullData[7], fullData[5]^fullData[2], fullData[15]^fullData[11], fullData[10]+fullData[0], fullData[13]^fullData[12], fullData[6]+fullData[14])
		return  /*line sIHFAXpA.go:1*/ string(data)
	}()
	dAMBvbjT := y5yvGZQA[0]

	jVqGRYFt, xonUbr_n :=  /*line FUZJ3MQm.go:1*/ kuuXUPaZ.iiTHmzLy(vhZhSyYv)
	if xonUbr_n != nil {
		return  /*line KhrARkGE.go:1*/ shim.Error( /*line J6jxm_cs.go:1*/ xonUbr_n.Error())
	}
	if jVqGRYFt ==  /*line BAVli4K1.go:1*/ func() string {
		seed :=  /*line BAVli4K1.go:1*/ byte(40)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line BAVli4K1.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line BAVli4K1.go:1*/  /*line BAVli4K1.go:1*/  /*line BAVli4K1.go:1*/  /*line BAVli4K1.go:1*/  /*line BAVli4K1.go:1*/ fnc(62)(251)(11)(7)(242)
		return  /*line BAVli4K1.go:1*/ string(data)
	}() {
		yfxmHRDX = y5yvGZQA[1]

	} else if jVqGRYFt ==  /*line DAh_sGuX.go:1*/ func() string {
		key :=  /*line DAh_sGuX.go:1*/ []byte("\x19d W")
		data :=  /*line DAh_sGuX.go:1*/ []byte("[\x0eU\x0e")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line DAh_sGuX.go:1*/ string(data)
	}() {
		yfxmHRDX, xonUbr_n =  /*line U13w9YnD.go:1*/ kuuXUPaZ.a8t6_82N(vhZhSyYv)
		if xonUbr_n != nil {
			return  /*line mTlXrbp7.go:1*/ shim.Error( /*line Oe8iCO3m.go:1*/ xonUbr_n.Error())
		}
	} else {
		return  /*line upz8iTYQ.go:1*/ shim.Error( /*line FwdUztP5.go:1*/ func() string {
			var data []byte
			i := 3
			decryptKey := 33
			for counter := 0; i != 6; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 0:
					i = 8
					data =  /*line FwdUztP5.go:1*/ append(data, "xufr"...,
					)
				case 5:
					data =  /*line FwdUztP5.go:1*/ append(data, 128)
					i = 2
				case 8:
					for y := range data {
						data[y] = data[y] -  /*line FwdUztP5.go:1*/ byte(decryptKey^y)
					}
					i = 6
				case 4:
					i = 0
					data =  /*line FwdUztP5.go:1*/ append(data, "\x91<"...,
					)
				case 7:
					i = 4
					data =  /*line FwdUztP5.go:1*/ append(data, "\x8dG\x84\x96"...,
					)
				case 3:
					data =  /*line FwdUztP5.go:1*/ append(data, 104)
					i = 5
				case 2:
					i = 1
					data =  /*line FwdUztP5.go:1*/ append(data, "|~\x86\x8d"...,
					)
				case 1:
					data =  /*line FwdUztP5.go:1*/ append(data, "\x834\x84\x88"...,
					)
					i = 7
				}
			}
			return  /*line FwdUztP5.go:1*/ string(data)
		}())
	}

	bXQWJgv9 := y5yvGZQA[2]
	lohinFr6 := y5yvGZQA[3]
	m5mWDJgx := y5yvGZQA[4]
	pzKghUuG := y5yvGZQA[5]
	xwCUJUQo := y5yvGZQA[6]
	gEutPS3G := y5yvGZQA[7]
	lMODNH3n := y5yvGZQA[8]
	u7imAyzE := y5yvGZQA[9]
	kA4KsH6J := y5yvGZQA[10]
	bg2bcVdk := y5yvGZQA[11]
	gOh1F7Le := y5yvGZQA[12]
	q2zUyIw9 := y5yvGZQA[13]
	zJrRGUxe, xonUbr_n :=  /*line xuCnnUY3.go:1*/ strconv.Atoi(y5yvGZQA[14])
	if xonUbr_n != nil {
		 /*line H7k6tm01.go:1*/ fmt.Println(xonUbr_n)
		 /*line AiCdnSC8.go:1*/ fmt.Printf( /*line OhHZBoTu.go:1*/ func() string {
			var data []byte
			i := 3
			decryptKey := 180
			for counter := 0; i != 1; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 6:
					i = 5
					data =  /*line OhHZBoTu.go:1*/ append(data, 150)
				case 8:
					data =  /*line OhHZBoTu.go:1*/ append(data, 202)
					i = 4
				case 7:
					data =  /*line OhHZBoTu.go:1*/ append(data, 191)
					i = 0
				case 5:
					data =  /*line OhHZBoTu.go:1*/ append(data, 187)
					i = 7
				case 3:
					data =  /*line OhHZBoTu.go:1*/ append(data, 186)
					i = 8
				case 0:
					data =  /*line OhHZBoTu.go:1*/ append(data, 239)
					i = 2
				case 4:
					i = 6
					data =  /*line OhHZBoTu.go:1*/ append(data, 189)
				case 2:
					i = 1
					for y := range data {
						data[y] = data[y] ^  /*line OhHZBoTu.go:1*/ byte(decryptKey^y)
					}
				}
			}
			return  /*line OhHZBoTu.go:1*/ string(data)
		}(), zJrRGUxe, zJrRGUxe)
	}
	gFdLvPNs := y5yvGZQA[15]
	cLeWlEtC := y5yvGZQA[16]
	n9CWGG_9 := y5yvGZQA[17]
	d2IBoq9l := y5yvGZQA[18]
	h_2IS0Jn, xonUbr_n :=  /*line p5IdHg_J.go:1*/ strconv.Atoi(y5yvGZQA[19])
	if xonUbr_n != nil {
		 /*line KG6K7y3d.go:1*/ fmt.Println(xonUbr_n)
		 /*line Dld6sA96.go:1*/ fmt.Printf( /*line S7An85Xj.go:1*/ func() string {
			fullData :=  /*line S7An85Xj.go:1*/ []byte("^\xfc\x9c\xe2->\bzʞB\x89\xc28")
			var data []byte
			data =  /*line S7An85Xj.go:1*/ append(data, fullData[11]+fullData[2], fullData[9]^fullData[8], fullData[5]+fullData[3], fullData[10]-fullData[13], fullData[12]+fullData[0], fullData[6]^fullData[4], fullData[1]+fullData[7])
			return  /*line S7An85Xj.go:1*/ string(data)
		}(), h_2IS0Jn, h_2IS0Jn)
	}
	zkXVyJDc := y5yvGZQA[20]
	nxzmIpzM :=  /*line BZCZ58Cf.go:1*/ strings.Split(y5yvGZQA[21],  /*line pFFB0rO5.go:1*/ func() string {
		var data []byte
		i := 2
		decryptKey := 193
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				for y := range data {
					data[y] = data[y] +  /*line pFFB0rO5.go:1*/ byte(decryptKey^y)
				}
				i = 1
			case 2:
				data =  /*line pFFB0rO5.go:1*/ append(data, 107)
				i = 0
			}
		}
		return  /*line pFFB0rO5.go:1*/ string(data)
	}())
	aWCTWkGJ := y5yvGZQA[22]
	jRnJ4yRc := y5yvGZQA[23]
	caxsmtLS := y5yvGZQA[24]
	_4svNVmw := y5yvGZQA[25]
	aoWFJM4q := y5yvGZQA[26]
	c66JfjVV := y5yvGZQA[27]
	ffzZHiCx :=  /*line ygR7Fh6K.go:1*/ strings.Split(y5yvGZQA[28],  /*line eBWzK8V1.go:1*/ func() string {
		fullData :=  /*line eBWzK8V1.go:1*/ []byte("\xd3\xff")
		var data []byte
		data =  /*line eBWzK8V1.go:1*/ append(data, fullData[1]^fullData[0])
		return  /*line eBWzK8V1.go:1*/ string(data)
	}())
	ji1qnmZT := y5yvGZQA[29]
	cpiU8vtM := y5yvGZQA[30]
	eEOQot0t := y5yvGZQA[31]
	zrdHgEga := 1

	                    
	uySB84CB := y5yvGZQA[32]
	oEFPNejh := y5yvGZQA[33]
	ec7lyFOw := y5yvGZQA[34]
	rcHo4AXH := y5yvGZQA[35]
	ctIHY4K4 := y5yvGZQA[36]
	cmqQYjmi := y5yvGZQA[37]
	h4hiesOE := y5yvGZQA[38]
	igm8gjBT := y5yvGZQA[39]
	qWPLffM8 := y5yvGZQA[40]
	hF9UDa5l := y5yvGZQA[41]
	bH3Pa427 := y5yvGZQA[42]
	mn_A4VXe := y5yvGZQA[43]
	hsYNgxdA, xonUbr_n :=  /*line QzL9a_E7.go:1*/ strconv.ParseBool(y5yvGZQA[44])
	if xonUbr_n != nil {
		 /*line tSk6dFBM.go:1*/ fmt.Printf( /*line jBNYdsTg.go:1*/ func() string {
			var data []byte
			i := 8
			decryptKey := 159
			for counter := 0; i != 4; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 1:
					data =  /*line jBNYdsTg.go:1*/ append(data, 151)
					i = 3
				case 3:
					i = 0
					data =  /*line jBNYdsTg.go:1*/ append(data, 149)
				case 7:
					data =  /*line jBNYdsTg.go:1*/ append(data, 198)
					i = 6
				case 0:
					data =  /*line jBNYdsTg.go:1*/ append(data, 197)
					i = 2
				case 5:
					i = 4
					for y := range data {
						data[y] = data[y] ^  /*line jBNYdsTg.go:1*/ byte(decryptKey^y)
					}
				case 2:
					i = 5
					data =  /*line jBNYdsTg.go:1*/ append(data, 184)
				case 6:
					i = 1
					data =  /*line jBNYdsTg.go:1*/ append(data, 140)
				case 8:
					data =  /*line jBNYdsTg.go:1*/ append(data, 145)
					i = 7
				}
			}
			return  /*line jBNYdsTg.go:1*/ string(data)
		}(), y5yvGZQA[44], hsYNgxdA)
	}

	                                                 
	fBolCUH_, xonUbr_n :=  /*line fHC7EwR9.go:1*/ vhZhSyYv.GetState(pzKghUuG)
	if xonUbr_n != nil {
		return  /*line nyTTp6Vf.go:1*/ shim.Error( /*line zg2cuTKe.go:1*/ func() string {
			fullData :=  /*line zg2cuTKe.go:1*/ []byte("H\xabeFx\xca\x18\x13se\xd8S\xe9h\xa1\xf6\x8c@5\xc6`=\xbb\xcch#\x1c\xf1\xe4\xd4\xe5\xafڹ\xdf\x1e\xfd\xb9\t\xf7\xeen\b\xc0\u007f\xfaJ\xb0\xb7\x18\xb7\xb6`6\x04l\x13-\x03t\xa5\xdbz|\xd6o\xf1\xd7@g\xbfSM;\xf5\xbb\xb3«\xb2t\x1c\xc1\x13\\$\xb8x\x16v\xb5\x04\xa1R\x00\xbc")
			var data []byte
			data =  /*line zg2cuTKe.go:1*/ append(data, fullData[68]-fullData[45], fullData[5]^fullData[1], fullData[21]-fullData[29], fullData[38]^fullData[9], fullData[47]+fullData[90], fullData[48]-fullData[71], fullData[49]+fullData[42], fullData[82]+fullData[76], fullData[18]-fullData[19], fullData[13]-fullData[0], fullData[20]-fullData[36], fullData[50]^fullData[34], fullData[3]^fullData[25], fullData[52]+fullData[58], fullData[65]^fullData[54], fullData[91]+fullData[81], fullData[64]^fullData[70], fullData[40]+fullData[87], fullData[22]+fullData[2], fullData[92]-fullData[57], fullData[33]+fullData[31], fullData[4]-fullData[83], fullData[39]^fullData[67], fullData[85]+fullData[46], fullData[93]-fullData[27], fullData[63]+fullData[66], fullData[89]^fullData[56], fullData[69]+fullData[37], fullData[84]+fullData[7], fullData[32]-fullData[80], fullData[11]^fullData[8], fullData[75]+fullData[79], fullData[28]-fullData[44], fullData[59]+fullData[94], fullData[61]-fullData[62], fullData[95]^fullData[10], fullData[43]+fullData[14], fullData[53]-fullData[77], fullData[74]+fullData[55], fullData[73]+fullData[30], fullData[16]^fullData[12], fullData[86]-fullData[17], fullData[51]-fullData[72], fullData[6]-fullData[60], fullData[24]^fullData[26], fullData[35]-fullData[78], fullData[23]+fullData[41], fullData[88]-fullData[15])
			return  /*line zg2cuTKe.go:1*/ string(data)
		}() +  /*line QzeqpAL_.go:1*/ xonUbr_n.Error())
	} else if fBolCUH_ != nil {
		return  /*line HuupFDAh.go:1*/ shim.Error( /*line MtvdsGKX.go:1*/ func() string {
			data :=  /*line MtvdsGKX.go:1*/ []byte("\v}\x01\xaf \xfea\xecK\a\xc2l\re\x01\x06\xf9 \xe7\xfa\xba\xfbss\x1a\x86")
			positions := [...]byte{21, 7, 19, 16, 14, 18, 14, 25, 8, 21, 1, 22, 10, 14, 3, 9, 9, 21, 0, 20, 2, 15, 5, 14, 25, 25, 12, 5, 22, 24, 0, 12, 9, 8, 25, 1}
			for i := 0; i < 36; i += 2 {
				localKey :=  /*line MtvdsGKX.go:1*/ byte(i) +  /*line MtvdsGKX.go:1*/ byte(positions[i]^positions[i+1]) + 124
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line MtvdsGKX.go:1*/ string(data)
		}() + pzKghUuG)
	}

	                     
	var w82K8EZR = &DataModelsStruct{
		ec7lyFOw,
		rcHo4AXH,
		ctIHY4K4,
	}
	var zlBaOzr4 = &DataProvidedStruct{
		cmqQYjmi,
		h4hiesOE,
		igm8gjBT,
	}

	var xhED8KEJ = &DataSourceMetadata{
		uySB84CB,
		oEFPNejh,
		w82K8EZR,
		zlBaOzr4,
		qWPLffM8,
		hF9UDa5l,
		bH3Pa427,
		mn_A4VXe,
		hsYNgxdA,
	}

	                          
	var ls5l3RtF = &DatasetMetadata{aTwOY3WH,
		dAMBvbjT,
		yfxmHRDX,
		bXQWJgv9,
		lohinFr6,
		m5mWDJgx,
		pzKghUuG,
		xwCUJUQo,
		gEutPS3G,
		lMODNH3n,
		u7imAyzE,
		kA4KsH6J,
		bg2bcVdk,
		gOh1F7Le,
		q2zUyIw9,
		zJrRGUxe,
		gFdLvPNs,
		cLeWlEtC,
		n9CWGG_9,
		d2IBoq9l,
		h_2IS0Jn,
		zkXVyJDc,
		nxzmIpzM,
		aWCTWkGJ,
		jRnJ4yRc,
		caxsmtLS,
		_4svNVmw,
		aoWFJM4q,
		c66JfjVV,
		ffzZHiCx,
		ji1qnmZT,
		cpiU8vtM,
		zrdHgEga,
		eEOQot0t,
		xhED8KEJ,
	}

	Qk2Ir6Pj, _ :=  /*line HCwywvsW.go:1*/ json.Marshal(ls5l3RtF)
	xonUbr_n =  /*line JkA5PeKd.go:1*/ vhZhSyYv.PutState(pzKghUuG, Qk2Ir6Pj)
	if xonUbr_n != nil {
		return  /*line kQ59gPaL.go:1*/ shim.Error( /*line _GRbnVWq.go:1*/ xonUbr_n.Error())
	}

	tvV5RxQz =  /*line EDgSxzIr.go:1*/ append(tvV5RxQz,  /*line N0CId8o3.go:1*/ func() string {
		data :=  /*line N0CId8o3.go:1*/ []byte("\xe4x\xf4\xe2\xf3}d")
		positions := [...]byte{5, 1, 5, 2, 4, 0, 3, 1}
		for i := 0; i < 8; i += 2 {
			localKey :=  /*line N0CId8o3.go:1*/ byte(i) +  /*line N0CId8o3.go:1*/ byte(positions[i]^positions[i+1]) + 136
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line N0CId8o3.go:1*/ string(data)
	}())
	              
	iRZGm1gc, xonUbr_n :=  /*line cBkN9_1_.go:1*/ woUuOwcA(vhZhSyYv, []string{pzKghUuG, dAMBvbjT, yfxmHRDX,  /*line aWKghXhE.go:1*/ strings.Join(tvV5RxQz,  /*line uTt2fOSs.go:1*/ func() string {
		seed :=  /*line uTt2fOSs.go:1*/ byte(84)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line uTt2fOSs.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line uTt2fOSs.go:1*/ fnc(216)
		return  /*line uTt2fOSs.go:1*/ string(data)
	}()),  /*line IQMBaPd2.go:1*/ strconv.Itoa(zrdHgEga), zkXVyJDc, lMODNH3n})
	if xonUbr_n != nil {
		return  /*line vroGy9r4.go:1*/ shim.Error( /*line Ibc0ox7L.go:1*/ xonUbr_n.Error())
	}
	 /*line GJw0REFH.go:1*/ fmt.Println(iRZGm1gc)

	return  /*line zZwcZc3V.go:1*/ shim.Success(nil)

}

func (kuuXUPaZ *JPyzqDxq) sgFZSQF7(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) pb.Response {
	var xonUbr_n error
	var tvV5RxQz []string
	var yfxmHRDX string

	dAMBvbjT := y5yvGZQA[0]

	jVqGRYFt, xonUbr_n :=  /*line xKyYHFpl.go:1*/ kuuXUPaZ.iiTHmzLy(vhZhSyYv)
	if xonUbr_n != nil {
		return  /*line hTHsxMN0.go:1*/ shim.Error( /*line y1IBVwbu.go:1*/ xonUbr_n.Error())
	}

	if jVqGRYFt ==  /*line j9Ut9ilo.go:1*/ func() string {
		key :=  /*line j9Ut9ilo.go:1*/ []byte("*W1\x13\x04")
		data :=  /*line j9Ut9ilo.go:1*/ []byte("<\n;`a")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line j9Ut9ilo.go:1*/ string(data)
	}() {
		yfxmHRDX = y5yvGZQA[1]

	} else if jVqGRYFt ==  /*line DBqbt0qi.go:1*/ func() string {
		fullData :=  /*line DBqbt0qi.go:1*/ []byte("\xeak\x06Ƈ\xf3\x98\xb3")
		var data []byte
		data =  /*line DBqbt0qi.go:1*/ append(data, fullData[5]^fullData[4], fullData[0]^fullData[6], fullData[7]^fullData[3], fullData[1]-fullData[2])
		return  /*line DBqbt0qi.go:1*/ string(data)
	}() {
		yfxmHRDX, xonUbr_n =  /*line vBcgz5wy.go:1*/ kuuXUPaZ.a8t6_82N(vhZhSyYv)
		if xonUbr_n != nil {
			return  /*line mHjh7Nzx.go:1*/ shim.Error( /*line v07op_Y4.go:1*/ xonUbr_n.Error())
		}

	} else {
		return  /*line Um7Oi3fI.go:1*/ shim.Error( /*line ekHNWK34.go:1*/ func() string {
			var data []byte
			i := 7
			decryptKey := 228
			for counter := 0; i != 4; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 6:
					data =  /*line ekHNWK34.go:1*/ append(data, "\xc4\x14\x18"...,
					)
					i = 1
				case 8:
					data =  /*line ekHNWK34.go:1*/ append(data, "\x1d\x13"...,
					)
					i = 6
				case 0:
					i = 10
					data =  /*line ekHNWK34.go:1*/ append(data, 20)
				case 10:
					data =  /*line ekHNWK34.go:1*/ append(data, "&!"...,
					)
					i = 11
				case 2:
					i = 3
					data =  /*line ekHNWK34.go:1*/ append(data, 22)
				case 1:
					i = 0
					data =  /*line ekHNWK34.go:1*/ append(data, "\x1d\xd7"...,
					)
				case 5:
					i = 4
					for y := range data {
						data[y] = data[y] -  /*line ekHNWK34.go:1*/ byte(decryptKey^y)
					}
				case 7:
					data =  /*line ekHNWK34.go:1*/ append(data, "\xf8\x10"...,
					)
					i = 9
				case 11:
					i = 2
					data =  /*line ekHNWK34.go:1*/ append(data, "\xcc(%"...,
					)
				case 9:
					i = 8
					data =  /*line ekHNWK34.go:1*/ append(data, "\f\x0e\x16"...,
					)
				case 3:
					i = 5
					data =  /*line ekHNWK34.go:1*/ append(data, 34)
				}
			}
			return  /*line ekHNWK34.go:1*/ string(data)
		}())
	}
	bXQWJgv9 := y5yvGZQA[2]
	lohinFr6 := y5yvGZQA[3]
	m5mWDJgx := y5yvGZQA[4]
	pzKghUuG := y5yvGZQA[5]
	gEutPS3G := y5yvGZQA[6]
	u7imAyzE := y5yvGZQA[7]
	kA4KsH6J := y5yvGZQA[8]
	bg2bcVdk := y5yvGZQA[9]
	gOh1F7Le := y5yvGZQA[10]
	q2zUyIw9 := y5yvGZQA[11]
	zJrRGUxe, xonUbr_n :=  /*line w_bXty0z.go:1*/ strconv.Atoi(y5yvGZQA[12])
	if xonUbr_n != nil {
		 /*line GvSA8TpY.go:1*/ fmt.Println(xonUbr_n)
		 /*line EzvhV5Td.go:1*/ fmt.Printf( /*line cbBKFmXA.go:1*/ func() string {
			var data []byte
			i := 0
			decryptKey := 206
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 4:
					i = 1
					data =  /*line cbBKFmXA.go:1*/ append(data, 249)
				case 3:
					i = 4
					data =  /*line cbBKFmXA.go:1*/ append(data, 142)
				case 6:
					i = 7
					data =  /*line cbBKFmXA.go:1*/ append(data, 255)
				case 8:
					i = 2
					for y := range data {
						data[y] = data[y] ^  /*line cbBKFmXA.go:1*/ byte(decryptKey^y)
					}
				case 7:
					data =  /*line cbBKFmXA.go:1*/ append(data, 251)
					i = 5
				case 1:
					i = 6
					data =  /*line cbBKFmXA.go:1*/ append(data, 210)
				case 5:
					data =  /*line cbBKFmXA.go:1*/ append(data, 171)
					i = 8
				case 0:
					data =  /*line cbBKFmXA.go:1*/ append(data, 254)
					i = 3
				}
			}
			return  /*line cbBKFmXA.go:1*/ string(data)
		}(), zJrRGUxe, zJrRGUxe)
	}
	gFdLvPNs := y5yvGZQA[13]
	cLeWlEtC := y5yvGZQA[14]
	n9CWGG_9 := y5yvGZQA[15]
	d2IBoq9l := y5yvGZQA[16]
	h_2IS0Jn, xonUbr_n :=  /*line HrrzDd2H.go:1*/ strconv.Atoi(y5yvGZQA[17])
	if xonUbr_n != nil {
		 /*line AK1lblR3.go:1*/ fmt.Println(xonUbr_n)
		 /*line m6euWD1I.go:1*/ fmt.Printf( /*line H56x0szq.go:1*/ func() string {
			key :=  /*line H56x0szq.go:1*/ []byte("Y\xac^\x9e\xa8\xbe]")
			data :=  /*line H56x0szq.go:1*/ []byte("̨\xc2lxg\x19")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line H56x0szq.go:1*/ string(data)
		}(), h_2IS0Jn, h_2IS0Jn)
	}
	zkXVyJDc := y5yvGZQA[18]
	nxzmIpzM :=  /*line YzoumVbA.go:1*/ strings.Split(y5yvGZQA[19],  /*line dnsXAqYH.go:1*/ func() string {
		data :=  /*line dnsXAqYH.go:1*/ []byte("\xbc")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line dnsXAqYH.go:1*/ byte(i) +  /*line dnsXAqYH.go:1*/ byte(positions[i]^positions[i+1]) + 144
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line dnsXAqYH.go:1*/ string(data)
	}())
	aWCTWkGJ := y5yvGZQA[20]
	jRnJ4yRc := y5yvGZQA[21]
	caxsmtLS := y5yvGZQA[22]
	_4svNVmw := y5yvGZQA[23]
	aoWFJM4q := y5yvGZQA[24]
	c66JfjVV := y5yvGZQA[25]
	ffzZHiCx :=  /*line xwvPz_jP.go:1*/ strings.Split(y5yvGZQA[26],  /*line AkNHLmLF.go:1*/ func() string {
		data :=  /*line AkNHLmLF.go:1*/ []byte("\x91")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line AkNHLmLF.go:1*/ byte(i) +  /*line AkNHLmLF.go:1*/ byte(positions[i]^positions[i+1]) + 101
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line AkNHLmLF.go:1*/ string(data)
	}())
	ji1qnmZT := y5yvGZQA[27]
	cpiU8vtM := y5yvGZQA[28]
	eEOQot0t := y5yvGZQA[29]

	aifJoMiz := y5yvGZQA[30]

	                                         
	j9GLBFOm, xonUbr_n :=  /*line B34A_ACs.go:1*/ vhZhSyYv.GetState(pzKghUuG)
	if xonUbr_n != nil {
		return  /*line I26YT7QW.go:1*/ shim.Error( /*line kFeaIp6j.go:1*/ func() string {
			fullData :=  /*line kFeaIp6j.go:1*/ []byte("\x9a\xa4\xa3N'I\xe6\xe1\xdf\xf5ͨ\x02\x83\xd2\xc0$[\xf2\x9d%\x93M\xc9ӄp\xb0cF\xa7y\xdde\xad\xed\xe8Z\x96\xb0J\\\xba\xbd\x9d\xd8\xc5\x10\x82\xd9\xff\xf8:\xea\x17\xb5S,\xfel\xd1@7\x97\xa9b\xb0\x1f\x04\x8e\x8dz\xc1\xec\xa6E|\x8fs\x9c825\x84\x9c\xfe\x1cH\bq\xe2\x85y-\x03\r")
			var data []byte
			data =  /*line kFeaIp6j.go:1*/ append(data, fullData[20]-fullData[8], fullData[82]+fullData[57], fullData[72]+fullData[11], fullData[17]^fullData[62], fullData[87]^fullData[93], fullData[53]+fullData[71], fullData[42]-fullData[0], fullData[48]+fullData[18], fullData[38]-fullData[4], fullData[13]+fullData[19], fullData[24]^fullData[27], fullData[2]+fullData[46], fullData[37]-fullData[9], fullData[34]-fullData[40], fullData[22]-fullData[90], fullData[92]+fullData[30], fullData[39]^fullData[49], fullData[15]^fullData[74], fullData[52]+fullData[6], fullData[75]-fullData[60], fullData[68]-fullData[79], fullData[67]+fullData[29], fullData[43]-fullData[44], fullData[85]+fullData[26], fullData[21]-fullData[81], fullData[33]-fullData[51], fullData[86]^fullData[31], fullData[76]+fullData[1], fullData[7]+fullData[69], fullData[3]-fullData[36], fullData[56]^fullData[78], fullData[5]^fullData[16], fullData[70]+fullData[45], fullData[88]+fullData[59], fullData[14]+fullData[77], fullData[10]^fullData[64], fullData[28]+fullData[58], fullData[25]-fullData[47], fullData[65]+fullData[50], fullData[95]-fullData[35], fullData[55]+fullData[66], fullData[61]^fullData[80], fullData[91]^fullData[73], fullData[89]+fullData[12], fullData[63]+fullData[32], fullData[41]+fullData[54], fullData[94]-fullData[23], fullData[83]+fullData[84])
			return  /*line kFeaIp6j.go:1*/ string(data)
		}() +  /*line BPO9puB8.go:1*/ xonUbr_n.Error())
	} else if j9GLBFOm == nil {
		return  /*line BwEBQxcu.go:1*/ shim.Error( /*line VjdLZA4R.go:1*/ func() string {
			seed :=  /*line VjdLZA4R.go:1*/ byte(244)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line VjdLZA4R.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/  /*line VjdLZA4R.go:1*/ fnc(56)(155)(44)(102)(121)(64)(129)(7)(186)(185)(133)(251)(0)(1)(1)(175)(171)(78)(171)(67)(137)(15)(49)(79)(93)(17)(20)(51)(90)(108)(38)(63)(138)(12)(237)(192)
			return  /*line VjdLZA4R.go:1*/ string(data)
		}() + pzKghUuG)
	}
	ls5l3RtF := DatasetMetadata{}
	xonUbr_n =  /*line Jchs7Yb4.go:1*/ json.Unmarshal(j9GLBFOm, &ls5l3RtF)	                               
	if xonUbr_n != nil {
		return  /*line WTIJdyNS.go:1*/ shim.Error( /*line jVNpWqhz.go:1*/ xonUbr_n.Error())
	}

	qKQytQy6, xonUbr_n :=  /*line zBSfxAXJ.go:1*/ kuuXUPaZ.glBEweME(vhZhSyYv)
	if xonUbr_n != nil {
		return  /*line DjO2yxuE.go:1*/ shim.Error( /*line EfcS6tat.go:1*/ xonUbr_n.Error())
	}

	if ls5l3RtF.UsernameOfProvider != qKQytQy6 {
		return  /*line UvLLz7Dz.go:1*/ shim.Error( /*line Y7MS7IaI.go:1*/ func() string {
			key :=  /*line Y7MS7IaI.go:1*/ []byte("=\xa9\x00\xcb\xfa.v\xf7H\xd8\x1b\x9b\xb7\x8c \xa28\xd7}7j\xe8")
			data :=  /*line Y7MS7IaI.go:1*/ []byte("r\xc7l\xb2\xdaZ\x1e\x92h\xbcz\xef֬P\xd0W\xa1\x14S\x0f\x9a")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line Y7MS7IaI.go:1*/ string(data)
		}() + ls5l3RtF.UsernameOfProvider +  /*line Vn_lWl0R.go:1*/ func() string {
			fullData :=  /*line Vn_lWl0R.go:1*/ []byte("\xce\x1e\xe3r@\xc5NE\xb38\xf4B=\xa5\xaa\u007f\xde\xe7\x91ïcY=ҁ\x9c-u\xbf#rC\xb3\xa3_}\x01%\xe0\x9b\xec\x1f\x12ז\x92)\x01\xdb`H\xc0InP\x82\xadb\xc0\xfa\x86\xdf\xf1\rVN\xbd\xb2\x8e\xfd\xbfǤ\xba\xc7\xe0EE\x9e\x1c\t\v\x17\xe4 K)\xe3 \xb5\x00`\xf6")
			var data []byte
			data =  /*line Vn_lWl0R.go:1*/ append(data, fullData[63]+fullData[3], fullData[84]+fullData[36], fullData[27]-fullData[71], fullData[22]+fullData[72], fullData[93]+fullData[15], fullData[56]-fullData[43], fullData[79]^fullData[60], fullData[78]+fullData[80], fullData[87]-fullData[90], fullData[7]+fullData[85], fullData[58]^fullData[11], fullData[44]^fullData[34], fullData[61]-fullData[1], fullData[32]-fullData[16], fullData[48]+fullData[42], fullData[20]-fullData[86], fullData[24]^fullData[8], fullData[2]+fullData[18], fullData[70]-fullData[26], fullData[0]^fullData[67], fullData[62]^fullData[74], fullData[23]^fullData[53], fullData[45]-fullData[28], fullData[39]-fullData[52], fullData[81]+fullData[77], fullData[54]^fullData[37], fullData[25]-fullData[64], fullData[88]+fullData[12], fullData[33]^fullData[75], fullData[13]+fullData[19], fullData[83]+fullData[6], fullData[59]^fullData[76], fullData[29]+fullData[14], fullData[66]+fullData[89], fullData[57]^fullData[49], fullData[40]^fullData[10], fullData[30]^fullData[51], fullData[17]^fullData[69], fullData[9]^fullData[65], fullData[82]-fullData[73], fullData[92]-fullData[4], fullData[38]+fullData[55], fullData[35]-fullData[41], fullData[5]-fullData[50], fullData[91]+fullData[31], fullData[21]-fullData[47], fullData[68]-fullData[46])
			return  /*line Vn_lWl0R.go:1*/ string(data)
		}() + qKQytQy6)
	}

	if ls5l3RtF.UsernameOfOwner != lohinFr6 {
		tvV5RxQz =  /*line qMXhvrBp.go:1*/ append(tvV5RxQz,  /*line HMMEbLT_.go:1*/ func() string {
			key :=  /*line HMMEbLT_.go:1*/ []byte("\x1c\x9c9\nV\xe7\x96$_\xdeH\a,Su(*")
			data :=  /*line HMMEbLT_.go:1*/ []byte("Y\xd7,h\x18z\xd7A\xc1\x91\x1e\x19C$\xf9=H")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line HMMEbLT_.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.OrgOfOwner != m5mWDJgx {
		tvV5RxQz =  /*line neIv6ZUL.go:1*/ append(tvV5RxQz,  /*line Szmat8a8.go:1*/ func() string {
			key :=  /*line Szmat8a8.go:1*/ []byte(" \xd3ї\x1f\xc4(Z\xe7\x18\xdc4d\xa5r\x0e\xc0\xb6\x86-\xc8")
			data :=  /*line Szmat8a8.go:1*/ []byte("O\xa1\xb6\xf6q\xadR;\x93q\xb3ZD\xca\x14.\xaf\xc1\xe8H\xba")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line Szmat8a8.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.DatasetDescription != gEutPS3G {
		tvV5RxQz =  /*line yhGkg0Cl.go:1*/ append(tvV5RxQz,  /*line Wyfzk3QR.go:1*/ func() string {
			data :=  /*line Wyfzk3QR.go:1*/ []byte("vwws\u007fijtrwn")
			positions := [...]byte{8, 4, 2, 4, 2, 9, 2, 4, 0, 1, 4, 3, 6, 4}
			for i := 0; i < 14; i += 2 {
				localKey :=  /*line Wyfzk3QR.go:1*/ byte(i) +  /*line Wyfzk3QR.go:1*/ byte(positions[i]^positions[i+1]) + 10
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line Wyfzk3QR.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.ThemeCategory != kA4KsH6J {
		tvV5RxQz =  /*line mS7LBgip.go:1*/ append(tvV5RxQz,  /*line luFFg2AO.go:1*/ func() string {
			seed :=  /*line luFFg2AO.go:1*/ byte(37)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line luFFg2AO.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line luFFg2AO.go:1*/  /*line luFFg2AO.go:1*/  /*line luFFg2AO.go:1*/  /*line luFFg2AO.go:1*/  /*line luFFg2AO.go:1*/  /*line luFFg2AO.go:1*/  /*line luFFg2AO.go:1*/  /*line luFFg2AO.go:1*/  /*line luFFg2AO.go:1*/  /*line luFFg2AO.go:1*/  /*line luFFg2AO.go:1*/  /*line luFFg2AO.go:1*/  /*line luFFg2AO.go:1*/ fnc(81)(30)(241)(232)(8)(54)(202)(1)(19)(238)(24)(253)(245)
			return  /*line luFFg2AO.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.KeywordTag != bg2bcVdk {
		tvV5RxQz =  /*line F8MrFyDz.go:1*/ append(tvV5RxQz,  /*line Aonuhpfp.go:1*/ func() string {
			fullData :=  /*line Aonuhpfp.go:1*/ []byte("Ѫ\x96\x8c\xc8I\xfe;\x19\xc8\r\x8b\x95\xaf\x18*\xa8Q\xc55")
			var data []byte
			data =  /*line Aonuhpfp.go:1*/ append(data, fullData[12]^fullData[6], fullData[7]+fullData[15], fullData[0]+fullData[16], fullData[10]-fullData[2], fullData[18]+fullData[1], fullData[11]-fullData[8], fullData[17]^fullData[19], fullData[3]+fullData[9], fullData[14]+fullData[5], fullData[13]^fullData[4])
			return  /*line Aonuhpfp.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.Language != gOh1F7Le {
		tvV5RxQz =  /*line etznjPAY.go:1*/ append(tvV5RxQz,  /*line CmnPQ8z9.go:1*/ func() string {
			seed :=  /*line CmnPQ8z9.go:1*/ byte(131)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line CmnPQ8z9.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line CmnPQ8z9.go:1*/  /*line CmnPQ8z9.go:1*/  /*line CmnPQ8z9.go:1*/  /*line CmnPQ8z9.go:1*/  /*line CmnPQ8z9.go:1*/  /*line CmnPQ8z9.go:1*/  /*line CmnPQ8z9.go:1*/  /*line CmnPQ8z9.go:1*/ fnc(233)(245)(13)(249)(14)(236)(6)(254)
			return  /*line CmnPQ8z9.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.Distribution != q2zUyIw9 {
		tvV5RxQz =  /*line plqxmQzF.go:1*/ append(tvV5RxQz,  /*line btPt2mlX.go:1*/ func() string {
			seed :=  /*line btPt2mlX.go:1*/ byte(33)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line btPt2mlX.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line btPt2mlX.go:1*/  /*line btPt2mlX.go:1*/  /*line btPt2mlX.go:1*/  /*line btPt2mlX.go:1*/  /*line btPt2mlX.go:1*/  /*line btPt2mlX.go:1*/  /*line btPt2mlX.go:1*/  /*line btPt2mlX.go:1*/  /*line btPt2mlX.go:1*/  /*line btPt2mlX.go:1*/  /*line btPt2mlX.go:1*/  /*line btPt2mlX.go:1*/ fnc(67)(5)(10)(1)(254)(247)(249)(19)(255)(245)(6)(255)
			return  /*line btPt2mlX.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.DataVelocity != zJrRGUxe {
		tvV5RxQz =  /*line j74eG9Da.go:1*/ append(tvV5RxQz,  /*line GNz3_s_w.go:1*/ func() string {
			seed :=  /*line GNz3_s_w.go:1*/ byte(60)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line GNz3_s_w.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line GNz3_s_w.go:1*/  /*line GNz3_s_w.go:1*/  /*line GNz3_s_w.go:1*/  /*line GNz3_s_w.go:1*/  /*line GNz3_s_w.go:1*/  /*line GNz3_s_w.go:1*/  /*line GNz3_s_w.go:1*/  /*line GNz3_s_w.go:1*/  /*line GNz3_s_w.go:1*/  /*line GNz3_s_w.go:1*/  /*line GNz3_s_w.go:1*/  /*line GNz3_s_w.go:1*/  /*line GNz3_s_w.go:1*/ fnc(160)(61)(141)(7)(205)(240)(207)(165)(77)(142)(34)(79)(163)
			return  /*line GNz3_s_w.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.SpatialGeographicCoverage != gFdLvPNs {
		tvV5RxQz =  /*line Uz217SMR.go:1*/ append(tvV5RxQz,  /*line lIJb28QO.go:1*/ func() string {
			key :=  /*line lIJb28QO.go:1*/ []byte("Y\x9a#^4!\v2;\x8f\x99\x1c\xc7Mi\x8eː\xb7ۥ\x85\n\"z")
			data :=  /*line lIJb28QO.go:1*/ []byte("\x1a\xd6>\x165@a\x15*\xe0\xceV\x9a#\xffۘ\xb3\xb8\x9b\xc0\xedWE\xeb")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line lIJb28QO.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.TemporalCoverageStart != cLeWlEtC {
		tvV5RxQz =  /*line wpAKp6MS.go:1*/ append(tvV5RxQz,  /*line QF7LLwhA.go:1*/ func() string {
			seed :=  /*line QF7LLwhA.go:1*/ byte(209)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line QF7LLwhA.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/  /*line QF7LLwhA.go:1*/ fnc(165)(19)(228)(29)(229)(29)(237)(21)(205)(52)(249)(237)(7)(29)(254)(242)(218)(23)(27)(231)(8)
			return  /*line QF7LLwhA.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.TemporalCoverageEnd != n9CWGG_9 {
		tvV5RxQz =  /*line es8EgmMi.go:1*/ append(tvV5RxQz,  /*line YimwHc8I.go:1*/ func() string {
			var data []byte
			i := 4
			decryptKey := 114
			for counter := 0; i != 6; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 4:
					i = 7
					data =  /*line YimwHc8I.go:1*/ append(data, "\xf3\xe3\xea"...,
					)
				case 2:
					for y := range data {
						data[y] = data[y] -  /*line YimwHc8I.go:1*/ byte(decryptKey^y)
					}
					i = 6
				case 8:
					data =  /*line YimwHc8I.go:1*/ append(data, "\xda\xe4\xba\xe5"...,
					)
					i = 3
				case 5:
					data =  /*line YimwHc8I.go:1*/ append(data, 220)
					i = 0
				case 7:
					i = 8
					data =  /*line YimwHc8I.go:1*/ append(data, "\xec\xea\xec"...,
					)
				case 3:
					i = 1
					data =  /*line YimwHc8I.go:1*/ append(data, "\xeb\xd9\xe5\xd3"...,
					)
				case 0:
					i = 2
					data =  /*line YimwHc8I.go:1*/ append(data, 209)
				case 1:
					i = 5
					data =  /*line YimwHc8I.go:1*/ append(data, "\xd8մ"...,
					)
				}
			}
			return  /*line YimwHc8I.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.IndustryDomain != d2IBoq9l {
		tvV5RxQz =  /*line PxI5dzkB.go:1*/ append(tvV5RxQz,  /*line YWz8APET.go:1*/ func() string {
			seed :=  /*line YWz8APET.go:1*/ byte(39)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line YWz8APET.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line YWz8APET.go:1*/  /*line YWz8APET.go:1*/  /*line YWz8APET.go:1*/  /*line YWz8APET.go:1*/  /*line YWz8APET.go:1*/  /*line YWz8APET.go:1*/  /*line YWz8APET.go:1*/  /*line YWz8APET.go:1*/  /*line YWz8APET.go:1*/  /*line YWz8APET.go:1*/  /*line YWz8APET.go:1*/  /*line YWz8APET.go:1*/  /*line YWz8APET.go:1*/  /*line YWz8APET.go:1*/ fnc(78)(27)(244)(241)(6)(15)(248)(251)(57)(217)(226)(16)(232)(7)
			return  /*line YWz8APET.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.DataVolume != h_2IS0Jn {
		tvV5RxQz =  /*line JVgc5JOL.go:1*/ append(tvV5RxQz,  /*line W3xwLgcs.go:1*/ func() string {
			var data []byte
			i := 2
			decryptKey := 194
			for counter := 0; i != 11; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 7:
					i = 0
					data =  /*line W3xwLgcs.go:1*/ append(data, 203)
				case 1:
					i = 9
					data =  /*line W3xwLgcs.go:1*/ append(data, 196)
				case 6:
					i = 8
					data =  /*line W3xwLgcs.go:1*/ append(data, 209)
				case 4:
					i = 11
					for y := range data {
						data[y] = data[y] +  /*line W3xwLgcs.go:1*/ byte(decryptKey^y)
					}
				case 5:
					i = 1
					data =  /*line W3xwLgcs.go:1*/ append(data, 207)
				case 9:
					i = 4
					data =  /*line W3xwLgcs.go:1*/ append(data, 189)
				case 0:
					i = 5
					data =  /*line W3xwLgcs.go:1*/ append(data, 197)
				case 3:
					i = 7
					data =  /*line W3xwLgcs.go:1*/ append(data, 177)
				case 8:
					i = 3
					data =  /*line W3xwLgcs.go:1*/ append(data, 191)
				case 2:
					data =  /*line W3xwLgcs.go:1*/ append(data, 195)
					i = 10
				case 10:
					data =  /*line W3xwLgcs.go:1*/ append(data, 193)
					i = 6
				}
			}
			return  /*line W3xwLgcs.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.Comments != zkXVyJDc {
		tvV5RxQz =  /*line _lVKIouC.go:1*/ append(tvV5RxQz,  /*line VBdn0Yzp.go:1*/ func() string {
			var data []byte
			i := 5
			decryptKey := 170
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 7:
					data =  /*line VBdn0Yzp.go:1*/ append(data, 200)
					i = 8
				case 8:
					data =  /*line VBdn0Yzp.go:1*/ append(data, 210)
					i = 9
				case 9:
					i = 6
					data =  /*line VBdn0Yzp.go:1*/ append(data, 213)
				case 5:
					i = 4
					data =  /*line VBdn0Yzp.go:1*/ append(data, 202)
				case 6:
					i = 3
					data =  /*line VBdn0Yzp.go:1*/ append(data, 213)
				case 3:
					for y := range data {
						data[y] = data[y] +  /*line VBdn0Yzp.go:1*/ byte(decryptKey^y)
					}
					i = 0
				case 2:
					data =  /*line VBdn0Yzp.go:1*/ append(data, 210)
					i = 1
				case 4:
					i = 2
					data =  /*line VBdn0Yzp.go:1*/ append(data, 215)
				case 1:
					i = 7
					data =  /*line VBdn0Yzp.go:1*/ append(data, 211)
				}
			}
			return  /*line VBdn0Yzp.go:1*/ string(data)
		}())
	}

	if  /*line NueWaag5.go:1*/ len(ls5l3RtF.AccessRights) !=  /*line EO6Em__z.go:1*/ len(nxzmIpzM) {
		tvV5RxQz =  /*line R8Yzz_ZC.go:1*/ append(tvV5RxQz,  /*line ia7Yp2i5.go:1*/ func() string {
			fullData :=  /*line ia7Yp2i5.go:1*/ []byte("\x92\x1dr\xbfqY\xa7[\xa01J\xcc\xc0\xa3~\xd53\x02\xfd\xb7 \xe7\x1d\x01\xd4\n")
			var data []byte
			data =  /*line ia7Yp2i5.go:1*/ append(data, fullData[0]-fullData[9], fullData[5]+fullData[25], fullData[12]^fullData[13], fullData[21]+fullData[14], fullData[17]^fullData[4], fullData[3]^fullData[11], fullData[22]-fullData[18], fullData[6]^fullData[15], fullData[20]-fullData[19], fullData[10]+fullData[1], fullData[16]^fullData[7], fullData[8]^fullData[24], fullData[23]^fullData[2])
			return  /*line ia7Yp2i5.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.ContractAgreementHash != aWCTWkGJ {
		tvV5RxQz =  /*line p9M1262c.go:1*/ append(tvV5RxQz,  /*line gYrSDeaq.go:1*/ func() string {
			seed :=  /*line gYrSDeaq.go:1*/ byte(35)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line gYrSDeaq.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/  /*line gYrSDeaq.go:1*/ fnc(134)(24)(47)(100)(198)(123)(248)(1)(174)(157)(64)(139)(9)(18)(44)(80)(169)(88)(92)(0)(249)(4)(253)
			return  /*line gYrSDeaq.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.TermsConditionsHash != caxsmtLS {
		tvV5RxQz =  /*line CXfyweO0.go:1*/ append(tvV5RxQz,  /*line CkgThH2L.go:1*/ func() string {
			fullData :=  /*line CkgThH2L.go:1*/ []byte("\xcaX\xe3,lr\xb9\x80+\xefx\xe6\xf9ދ\x06\xba;\xaf*\x96\xf4D5\xfc{\xf5\x8e\x9c\b\x82BH\xd5\xe8K")
			var data []byte
			data =  /*line CkgThH2L.go:1*/ append(data, fullData[11]+fullData[27], fullData[30]+fullData[2], fullData[10]-fullData[15], fullData[23]^fullData[1], fullData[4]-fullData[12], fullData[29]-fullData[34], fullData[28]^fullData[16], fullData[33]+fullData[35], fullData[26]^fullData[20], fullData[7]+fullData[9], fullData[22]+fullData[19], fullData[32]^fullData[3], fullData[8]^fullData[31], fullData[18]-fullData[17], fullData[14]+fullData[13], fullData[25]+fullData[21], fullData[24]+fullData[5], fullData[0]^fullData[6])
			return  /*line CkgThH2L.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.ContractAgreementURL != jRnJ4yRc {
		tvV5RxQz =  /*line g5irz_Vi.go:1*/ append(tvV5RxQz,  /*line FCYXzkCd.go:1*/ func() string {
			key :=  /*line FCYXzkCd.go:1*/ []byte("S\x8d\x11\xf70R\xd0lJ\xd0s\xbb}:\x8f\xe0\x18\x19O\xcf\"\x85\xb3")
			data :=  /*line FCYXzkCd.go:1*/ []byte("\x10\xe2]}B\x0f\x93\b֑\xf4\xb7\xe8+ޅV[$Q3͙")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line FCYXzkCd.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.TermsConditionsURL != _4svNVmw {
		tvV5RxQz =  /*line faePmV2u.go:1*/ append(tvV5RxQz,  /*line ptOdeSDj.go:1*/ func() string {
			key :=  /*line ptOdeSDj.go:1*/ []byte("\xc8gcu+B\xbd\xc5\xfe\xf6\xe8\xc1n\x01K\x06M\x90\x04\xdck$")
			data :=  /*line ptOdeSDj.go:1*/ []byte("<\xcc\xd5\xe2\x9eb\xe3\xe5aeV%\xd7u\xb4u\xbb\x03$1\xbdp")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line ptOdeSDj.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.CopyrightHash != aoWFJM4q {
		tvV5RxQz =  /*line o87ujiGY.go:1*/ append(tvV5RxQz,  /*line dPjxV1E8.go:1*/ func() string {
			fullData :=  /*line dPjxV1E8.go:1*/ []byte("\x86ݲ\x02\x97\x91\xaf\x8f#I\xd9Y\x96\xd0eV\xe6F`\x19\x0f\xa0\x01\xb1\x1e\xd3")
			var data []byte
			data =  /*line dPjxV1E8.go:1*/ append(data, fullData[23]+fullData[2], fullData[20]^fullData[18], fullData[16]^fullData[12], fullData[15]+fullData[8], fullData[11]+fullData[19], fullData[6]-fullData[17], fullData[14]^fullData[3], fullData[7]+fullData[10], fullData[4]+fullData[1], fullData[9]-fullData[22], fullData[5]+fullData[13], fullData[21]+fullData[25], fullData[0]-fullData[24])
			return  /*line dPjxV1E8.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.CopyrightURL != c66JfjVV {
		tvV5RxQz =  /*line DGBw9hPW.go:1*/ append(tvV5RxQz,  /*line ParZVqWc.go:1*/ func() string {
			data :=  /*line ParZVqWc.go:1*/ []byte("x\xe6p\xf2\xe2\xe3g\xf1\xf6U\xdeL")
			positions := [...]byte{7, 1, 5, 5, 10, 10, 0, 3, 1, 8, 8, 8, 4, 3}
			for i := 0; i < 14; i += 2 {
				localKey :=  /*line ParZVqWc.go:1*/ byte(i) +  /*line ParZVqWc.go:1*/ byte(positions[i]^positions[i+1]) + 136
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line ParZVqWc.go:1*/ string(data)
		}())
	}
	v2y3tYFs :=  /*line soSALhu2.go:1*/ reflect.DeepEqual(ls5l3RtF.CustomAccessRights, ffzZHiCx)
	if v2y3tYFs == false {
		tvV5RxQz =  /*line YFeS3nsi.go:1*/ append(tvV5RxQz,  /*line K7OuTr1T.go:1*/ func() string {
			seed :=  /*line K7OuTr1T.go:1*/ byte(47)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line K7OuTr1T.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/  /*line K7OuTr1T.go:1*/ fnc(146)(54)(106)(213)(165)(72)(67)(199)(144)(32)(66)(146)(36)(245)(60)(111)(220)(185)(126)(251)
			return  /*line K7OuTr1T.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.Endpoint != ji1qnmZT {
		tvV5RxQz =  /*line efnaPKzf.go:1*/ append(tvV5RxQz,  /*line eI7eXMwC.go:1*/ func() string {
			var data []byte
			i := 9
			decryptKey := 42
			for counter := 0; i != 5; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 8:
					data =  /*line eI7eXMwC.go:1*/ append(data, 121)
					i = 6
				case 2:
					data =  /*line eI7eXMwC.go:1*/ append(data, 115)
					i = 0
				case 4:
					i = 2
					data =  /*line eI7eXMwC.go:1*/ append(data, 122)
				case 3:
					data =  /*line eI7eXMwC.go:1*/ append(data, 126)
					i = 8
				case 0:
					data =  /*line eI7eXMwC.go:1*/ append(data, 102)
					i = 3
				case 6:
					data =  /*line eI7eXMwC.go:1*/ append(data, 125)
					i = 1
				case 7:
					for y := range data {
						data[y] = data[y] ^  /*line eI7eXMwC.go:1*/ byte(decryptKey^y)
					}
					i = 5
				case 1:
					data =  /*line eI7eXMwC.go:1*/ append(data, 102)
					i = 7
				case 9:
					data =  /*line eI7eXMwC.go:1*/ append(data, 112)
					i = 4
				}
			}
			return  /*line eI7eXMwC.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.Blockchain != cpiU8vtM {
		tvV5RxQz =  /*line pa8sNZJ3.go:1*/ append(tvV5RxQz,  /*line CZt05e7y.go:1*/ func() string {
			fullData :=  /*line CZt05e7y.go:1*/ []byte("\x89\x92\xdd\xed\xaeu{P\xfdR\x18\xbd\xbde\xe7,\x84\xfb\xde\xf1")
			var data []byte
			data =  /*line CZt05e7y.go:1*/ append(data, fullData[13]+fullData[8], fullData[16]-fullData[10], fullData[15]-fullData[12], fullData[7]-fullData[3], fullData[4]+fullData[11], fullData[18]-fullData[6], fullData[2]-fullData[5], fullData[9]-fullData[19], fullData[17]^fullData[1], fullData[14]^fullData[0])
			return  /*line CZt05e7y.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.HasAgent != eEOQot0t {
		tvV5RxQz =  /*line eUQco1_n.go:1*/ append(tvV5RxQz,  /*line GDq3VG21.go:1*/ func() string {
			fullData :=  /*line GDq3VG21.go:1*/ []byte("o\xd2Q^\xaf\xc1\xc7\xf0\b\xa2")
			var data []byte
			data =  /*line GDq3VG21.go:1*/ append(data, fullData[2]-fullData[7], fullData[0]^fullData[8], fullData[6]^fullData[9], fullData[5]^fullData[4], fullData[1]-fullData[3])
			return  /*line GDq3VG21.go:1*/ string(data)
		}())
	}

	if ls5l3RtF.DataSourceMetadata.DataSourceID != aifJoMiz {
		tvV5RxQz =  /*line H2f3JgYc.go:1*/ append(tvV5RxQz,  /*line CqxWeG5m.go:1*/ func() string {
			data :=  /*line CqxWeG5m.go:1*/ []byte("W\n\xf6\x12\xe6\x1c\x04\x18\x01+I\x1c")
			positions := [...]byte{2, 8, 1, 0, 9, 7, 7, 8, 1, 11, 7, 6, 5, 3, 8, 4}
			for i := 0; i < 16; i += 2 {
				localKey :=  /*line CqxWeG5m.go:1*/ byte(i) +  /*line CqxWeG5m.go:1*/ byte(positions[i]^positions[i+1]) + 107
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line CqxWeG5m.go:1*/ string(data)
		}())
	}

	ls5l3RtF.UsernameOfProvider = dAMBvbjT
	ls5l3RtF.OrgOfProvider = yfxmHRDX
	ls5l3RtF.EmailOfProvider = bXQWJgv9
	ls5l3RtF.UsernameOfOwner = lohinFr6
	ls5l3RtF.OrgOfOwner = m5mWDJgx
	ls5l3RtF.DatasetName = pzKghUuG
	ls5l3RtF.DatasetDescription = gEutPS3G
	ls5l3RtF.ModificationDateTime = u7imAyzE

	ls5l3RtF.ThemeCategory = kA4KsH6J
	ls5l3RtF.KeywordTag = bg2bcVdk
	ls5l3RtF.Language = gOh1F7Le
	ls5l3RtF.Distribution = q2zUyIw9
	ls5l3RtF.DataVelocity = zJrRGUxe
	ls5l3RtF.SpatialGeographicCoverage = gFdLvPNs
	ls5l3RtF.TemporalCoverageStart = cLeWlEtC
	ls5l3RtF.TemporalCoverageEnd = n9CWGG_9
	ls5l3RtF.IndustryDomain = d2IBoq9l
	ls5l3RtF.DataVolume = h_2IS0Jn
	ls5l3RtF.Comments = zkXVyJDc

	ls5l3RtF.AccessRights = nxzmIpzM
	ls5l3RtF.ContractAgreementHash = aWCTWkGJ
	ls5l3RtF.ContractAgreementURL = jRnJ4yRc
	ls5l3RtF.TermsConditionsHash = caxsmtLS
	ls5l3RtF.TermsConditionsURL = _4svNVmw
	ls5l3RtF.CopyrightHash = aoWFJM4q
	ls5l3RtF.CopyrightURL = c66JfjVV
	ls5l3RtF.CustomAccessRights = ffzZHiCx

	ls5l3RtF.Endpoint = ji1qnmZT
	ls5l3RtF.Blockchain = cpiU8vtM
	ls5l3RtF.Version = ls5l3RtF.Version + 1
	ls5l3RtF.DataSourceMetadata.DataSourceID = aifJoMiz
	ls5l3RtF.HasAgent = eEOQot0t

	Qk2Ir6Pj, _ :=  /*line HgoL6b1K.go:1*/ json.Marshal(ls5l3RtF)
	xonUbr_n =  /*line NzS4GzkC.go:1*/ vhZhSyYv.PutState(pzKghUuG, Qk2Ir6Pj)
	if xonUbr_n != nil {
		return  /*line y0K4faIi.go:1*/ shim.Error( /*line u9b4ZSp1.go:1*/ xonUbr_n.Error())
	}

	              
	iRZGm1gc, xonUbr_n :=  /*line NI8gmeJ7.go:1*/ woUuOwcA(vhZhSyYv, []string{pzKghUuG, dAMBvbjT, yfxmHRDX,  /*line VRLAlmJ9.go:1*/ strings.Join(tvV5RxQz,  /*line Ys57Zo36.go:1*/ func() string {
		var data []byte
		i := 0
		decryptKey := 151
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				i = 2
				data =  /*line Ys57Zo36.go:1*/ append(data, 185)
			case 2:
				i = 1
				for y := range data {
					data[y] = data[y] ^  /*line Ys57Zo36.go:1*/ byte(decryptKey^y)
				}
			}
		}
		return  /*line Ys57Zo36.go:1*/ string(data)
	}()),  /*line IVFTmxj1.go:1*/ strconv.Itoa(ls5l3RtF.Version), ls5l3RtF.Comments, u7imAyzE})
	if xonUbr_n != nil {
		return  /*line lmIFXXIk.go:1*/ shim.Error( /*line jPboP63r.go:1*/ xonUbr_n.Error())
	}
	 /*line UpoeX0TJ.go:1*/ fmt.Println(iRZGm1gc)
	return  /*line DMr03Uaz.go:1*/ shim.Success(nil)

}

func (kuuXUPaZ *JPyzqDxq) uDyCIlIB(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) pb.Response {
	var tvV5RxQz []string
	mP0iLz9a := y5yvGZQA[0]
	pzKghUuG := y5yvGZQA[1]
	uySB84CB := y5yvGZQA[2]
	oEFPNejh := y5yvGZQA[3]
	ec7lyFOw := y5yvGZQA[4]
	rcHo4AXH := y5yvGZQA[5]
	ctIHY4K4 := y5yvGZQA[6]
	cmqQYjmi := y5yvGZQA[7]
	h4hiesOE := y5yvGZQA[8]
	igm8gjBT := y5yvGZQA[9]
	qWPLffM8 := y5yvGZQA[10]
	hF9UDa5l := y5yvGZQA[11]
	bH3Pa427 := y5yvGZQA[12]
	mn_A4VXe := y5yvGZQA[13]
	hsYNgxdA, xonUbr_n :=  /*line lArAJPjk.go:1*/ strconv.ParseBool(y5yvGZQA[14])
	if xonUbr_n != nil {
		 /*line zdUv6SFS.go:1*/ fmt.Printf( /*line z2N2fjEm.go:1*/ func() string {
			key :=  /*line z2N2fjEm.go:1*/ []byte("\xc9p\xd4\xe0N\x8b\x1a")
			data :=  /*line z2N2fjEm.go:1*/ []byte("\xec\x03\xee\xc0k\xff\x10")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line z2N2fjEm.go:1*/ string(data)
		}(), y5yvGZQA[14], hsYNgxdA)
	}

	 /*line LLZ1aqpa.go:1*/ fmt.Println( /*line Lhl2Pxdd.go:1*/ func() string {
		key :=  /*line Lhl2Pxdd.go:1*/ []byte("٣u\xd2\x12\xec\xe1\xad5\b\x88\x88\x0f\xc9\x00d\xd4\xff\xfdE$\xae*\x02\xa3)\xd7\xe0\xfb\xeb\x0f\x99\xfdt[\xaa\x15L\xdf;\x98\x9a\x03\x88\xa4DM\xdf\xd5")
		data :=  /*line Lhl2Pxdd.go:1*/ []byte("T}\xfe\xa2O\x86\x93s,\\\xdc\xe1_\x9e \x00\x8dud\xdbO\xc1Kp\xc0<I\x8dj\x89R\xcbd\x00\x06vQ#\x93\xe5\xcc\xc7q\xd9\xcf!'[K")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line Lhl2Pxdd.go:1*/ string(data)
	}(), pzKghUuG)
	                                                 
	cJaXNN6u, xonUbr_n :=  /*line g64dBWPt.go:1*/ vhZhSyYv.GetState(pzKghUuG)
	if xonUbr_n != nil {
		return  /*line klESwcrG.go:1*/ shim.Error( /*line _IsygTzu.go:1*/ func() string {
			data :=  /*line _IsygTzu.go:1*/ []byte("F\x97\xdbl\xddd \xe0o \xa0\x9c\xdaD\xbe \x82\x84[\xf8h\xe0\xa4n\xd8\xd8\u007f h\x94{l\\\x1f\xe7VaZa\xc4\xddji\x8bO\xcc:`")
			positions := [...]byte{43, 30, 39, 37, 13, 17, 16, 26, 28, 22, 25, 44, 11, 1, 47, 1, 45, 12, 41, 4, 29, 34, 24, 26, 29, 16, 41, 4, 11, 40, 32, 7, 2, 35, 13, 14, 29, 2, 7, 31, 40, 44, 47, 37, 17, 26, 4, 28, 34, 19, 10, 18, 4, 21, 10, 14, 2, 40, 43, 17, 26, 25, 24, 1, 25, 40, 16, 7, 25, 33}
			for i := 0; i < 70; i += 2 {
				localKey :=  /*line _IsygTzu.go:1*/ byte(i) +  /*line _IsygTzu.go:1*/ byte(positions[i]^positions[i+1]) + 54
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line _IsygTzu.go:1*/ string(data)
		}() +  /*line eI2XSpbl.go:1*/ xonUbr_n.Error())
	} else if cJaXNN6u == nil {
		return  /*line Dh_hqz71.go:1*/ shim.Error( /*line iPOXDCsw.go:1*/ func() string {
			var data []byte
			i := 1
			decryptKey := 208
			for counter := 0; i != 5; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					i = 2
					data =  /*line iPOXDCsw.go:1*/ append(data, "1\xed70"...,
					)
				case 8:
					i = 6
					data =  /*line iPOXDCsw.go:1*/ append(data, "K7"...,
					)
				case 9:
					data =  /*line iPOXDCsw.go:1*/ append(data, 60)
					i = 10
				case 15:
					i = 0
					data =  /*line iPOXDCsw.go:1*/ append(data, "(8&\xe2"...,
					)
				case 10:
					data =  /*line iPOXDCsw.go:1*/ append(data, "L:\xf6>"...,
					)
					i = 4
				case 2:
					data =  /*line iPOXDCsw.go:1*/ append(data, "<**"...,
					)
					i = 15
				case 12:
					i = 9
					data =  /*line iPOXDCsw.go:1*/ append(data, 62)
				case 1:
					i = 13
					data =  /*line iPOXDCsw.go:1*/ append(data, "2G"...,
					)
				case 7:
					i = 5
					for y := range data {
						data[y] = data[y] -  /*line iPOXDCsw.go:1*/ byte(decryptKey^y)
					}
				case 11:
					i = 3
					data =  /*line iPOXDCsw.go:1*/ append(data, "2<2"...,
					)
				case 4:
					i = 8
					data =  /*line iPOXDCsw.go:1*/ append(data, 67)
				case 13:
					data =  /*line iPOXDCsw.go:1*/ append(data, "A\xfd"...,
					)
					i = 12
				case 6:
					data =  /*line iPOXDCsw.go:1*/ append(data, "E>"...,
					)
					i = 11
				case 0:
					data =  /*line iPOXDCsw.go:1*/ append(data, ")/38"...,
					)
					i = 14
				case 14:
					i = 7
					data =  /*line iPOXDCsw.go:1*/ append(data, 31)
				}
			}
			return  /*line iPOXDCsw.go:1*/ string(data)
		}() + pzKghUuG +  /*line QEeOyzdj.go:1*/ func() string {
			seed :=  /*line QEeOyzdj.go:1*/ byte(14)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line QEeOyzdj.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line QEeOyzdj.go:1*/  /*line QEeOyzdj.go:1*/  /*line QEeOyzdj.go:1*/  /*line QEeOyzdj.go:1*/  /*line QEeOyzdj.go:1*/  /*line QEeOyzdj.go:1*/  /*line QEeOyzdj.go:1*/  /*line QEeOyzdj.go:1*/  /*line QEeOyzdj.go:1*/  /*line QEeOyzdj.go:1*/  /*line QEeOyzdj.go:1*/  /*line QEeOyzdj.go:1*/  /*line QEeOyzdj.go:1*/  /*line QEeOyzdj.go:1*/  /*line QEeOyzdj.go:1*/ fnc(46)(160)(75)(140)(38)(249)(64)(129)(7)(186)(185)(133)(251)(0)(1)
			return  /*line QEeOyzdj.go:1*/ string(data)
		}())
	}

	xhED8KEJ := DatasetMetadata{}
	xonUbr_n =  /*line SshQ_I81.go:1*/ json.Unmarshal(cJaXNN6u, &xhED8KEJ)	                               
	if xonUbr_n != nil {
		return  /*line vgCTwCNg.go:1*/ shim.Error( /*line r1dNz673.go:1*/ xonUbr_n.Error())
	}
	xhED8KEJ.DataSourceMetadata.DataSourceID = uySB84CB
	xhED8KEJ.DataSourceMetadata.DataSourceType = oEFPNejh
	xhED8KEJ.DataSourceMetadata.DataModels.Type = ec7lyFOw
	xhED8KEJ.DataSourceMetadata.DataModels.Value = rcHo4AXH
	xhED8KEJ.DataSourceMetadata.DataModels.Metadata = ctIHY4K4
	xhED8KEJ.DataSourceMetadata.DataProvided.Type = cmqQYjmi
	xhED8KEJ.DataSourceMetadata.DataProvided.Value = h4hiesOE
	xhED8KEJ.DataSourceMetadata.DataProvided.Metadata = igm8gjBT
	xhED8KEJ.DataSourceMetadata.Attributes = qWPLffM8
	xhED8KEJ.DataSourceMetadata.Service = hF9UDa5l
	xhED8KEJ.DataSourceMetadata.ServicePath = bH3Pa427
	xhED8KEJ.DataSourceMetadata.Mapping = mn_A4VXe
	xhED8KEJ.DataSourceMetadata.DataportsDataModelandFormat = hsYNgxdA

	w5g0TK0q :=  /*line s9_XzGSx.go:1*/ func() string {
		seed :=  /*line s9_XzGSx.go:1*/ byte(59)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line s9_XzGSx.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line s9_XzGSx.go:1*/  /*line s9_XzGSx.go:1*/  /*line s9_XzGSx.go:1*/  /*line s9_XzGSx.go:1*/  /*line s9_XzGSx.go:1*/  /*line s9_XzGSx.go:1*/  /*line s9_XzGSx.go:1*/ fnc(158)(75)(137)(14)(47)(79)(157)
		return  /*line s9_XzGSx.go:1*/ string(data)
	}()
	tvV5RxQz =  /*line rj8RTEHz.go:1*/ append(tvV5RxQz,  /*line AupaG7ue.go:1*/ func() string {
		var data []byte
		i := 13
		decryptKey := 181
		for counter := 0; i != 5; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				i = 2
				data =  /*line AupaG7ue.go:1*/ append(data, "\xceʺ"...,
				)
			case 2:
				data =  /*line AupaG7ue.go:1*/ append(data, "\xbbu\xc1\xb8"...,
				)
				i = 12
			case 9:
				i = 1
				data =  /*line AupaG7ue.go:1*/ append(data, 201)
			case 0:
				data =  /*line AupaG7ue.go:1*/ append(data, "\xc0\xd2"...,
				)
				i = 6
			case 3:
				data =  /*line AupaG7ue.go:1*/ append(data, 171)
				i = 10
			case 12:
				i = 11
				data =  /*line AupaG7ue.go:1*/ append(data, 198)
			case 13:
				data =  /*line AupaG7ue.go:1*/ append(data, 196)
				i = 0
			case 6:
				data =  /*line AupaG7ue.go:1*/ append(data, "\xbe|"...,
				)
				i = 7
			case 4:
				i = 8
				data =  /*line AupaG7ue.go:1*/ append(data, "\xb0®l"...,
				)
			case 11:
				i = 4
				data =  /*line AupaG7ue.go:1*/ append(data, "\xb2\xb4"...,
				)
			case 7:
				data =  /*line AupaG7ue.go:1*/ append(data, 206)
				i = 9
			case 8:
				i = 3
				data =  /*line AupaG7ue.go:1*/ append(data, "\xac\xae\xad\xad"...,
				)
			case 10:
				i = 5
				for y := range data {
					data[y] = data[y] +  /*line AupaG7ue.go:1*/ byte(decryptKey^y)
				}
			}
		}
		return  /*line AupaG7ue.go:1*/ string(data)
	}())

	 /*line DZMtfT0C.go:1*/ fmt.Println( /*line CgfspRd0.go:1*/ func() string {
		data :=  /*line CgfspRd0.go:1*/ []byte("1\xf0,+L\x02\x0e\xdao\xcb\xfd~r\xf3\xe0.a\xe106\xf2d$")
		positions := [...]byte{19, 11, 19, 15, 7, 9, 4, 7, 20, 13, 10, 6, 17, 13, 9, 1, 22, 14, 22, 9, 18, 1, 3, 20, 5, 0, 0, 4, 10, 2, 15, 6}
		for i := 0; i < 32; i += 2 {
			localKey :=  /*line CgfspRd0.go:1*/ byte(i) +  /*line CgfspRd0.go:1*/ byte(positions[i]^positions[i+1]) + 23
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line CgfspRd0.go:1*/ string(data)
	}())
	              
	iRZGm1gc, xonUbr_n :=  /*line I3WXuR3S.go:1*/ qnwYAENh(vhZhSyYv, []string{pzKghUuG, w5g0TK0q,  /*line T2ZZj7cN.go:1*/ strings.Join(tvV5RxQz,  /*line byZpF_Ko.go:1*/ func() string {
		seed :=  /*line byZpF_Ko.go:1*/ byte(115)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line byZpF_Ko.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line byZpF_Ko.go:1*/ fnc(185)
		return  /*line byZpF_Ko.go:1*/ string(data)
	}()), mP0iLz9a})
	if xonUbr_n != nil {
		return  /*line X3CXQOHf.go:1*/ shim.Error( /*line HVszhrjv.go:1*/ xonUbr_n.Error())
	}
	 /*line GqtltiOO.go:1*/ fmt.Println(iRZGm1gc)

	gBsMzs0O, _ :=  /*line Me2jF6I_.go:1*/ json.Marshal(xhED8KEJ)
	xonUbr_n =  /*line SxVn2mRK.go:1*/ vhZhSyYv.PutState(pzKghUuG, gBsMzs0O)
	if xonUbr_n != nil {
		return  /*line KovouMDm.go:1*/ shim.Error( /*line LBzsVbDa.go:1*/ xonUbr_n.Error())
	}

	return  /*line J4GUYmSe.go:1*/ shim.Success(nil)

}

func (kuuXUPaZ *JPyzqDxq) efFVT2Sf(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) pb.Response {
	var tvV5RxQz []string
	pzKghUuG := y5yvGZQA[0]
	ec7lyFOw := y5yvGZQA[1]
	rcHo4AXH := y5yvGZQA[2]
	ctIHY4K4 := y5yvGZQA[3]
	cmqQYjmi := y5yvGZQA[4]
	h4hiesOE := y5yvGZQA[5]
	igm8gjBT := y5yvGZQA[6]
	qWPLffM8 := y5yvGZQA[7]
	hF9UDa5l := y5yvGZQA[8]
	bH3Pa427 := y5yvGZQA[9]
	mn_A4VXe := y5yvGZQA[10]
	hsYNgxdA, xonUbr_n :=  /*line IgGZkyz7.go:1*/ strconv.ParseBool(y5yvGZQA[11])
	if xonUbr_n != nil {
		 /*line iTVU4hrr.go:1*/ fmt.Printf( /*line wIBAurkP.go:1*/ func() string {
			var data []byte
			i := 3
			decryptKey := 152
			for counter := 0; i != 1; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					i = 2
					data =  /*line wIBAurkP.go:1*/ append(data, 184)
				case 4:
					data =  /*line wIBAurkP.go:1*/ append(data, 176)
					i = 6
				case 8:
					i = 5
					data =  /*line wIBAurkP.go:1*/ append(data, 159)
				case 2:
					data =  /*line wIBAurkP.go:1*/ append(data, 5)
					i = 7
				case 5:
					for y := range data {
						data[y] = data[y] -  /*line wIBAurkP.go:1*/ byte(decryptKey^y)
					}
					i = 1
				case 7:
					i = 4
					data =  /*line wIBAurkP.go:1*/ append(data, 203)
				case 0:
					i = 8
					data =  /*line wIBAurkP.go:1*/ append(data, 10)
				case 6:
					i = 0
					data =  /*line wIBAurkP.go:1*/ append(data, 188)
				}
			}
			return  /*line wIBAurkP.go:1*/ string(data)
		}(), y5yvGZQA[11], hsYNgxdA)
	}
	mP0iLz9a := y5yvGZQA[12]

	 /*line EYA6k_IJ.go:1*/ fmt.Println( /*line Bab4nJJE.go:1*/ func() string {
		data :=  /*line Bab4nJJE.go:1*/ []byte("aIBtaG, \x9f\x8cIaqi{eKVat\xcd\x04z:>\bY\rSm$\u007fYR)t\x01hf6r\xa1d\xf2z9ne\x1dD ")
		positions := [...]byte{41, 44, 26, 43, 34, 0, 23, 23, 31, 37, 15, 26, 26, 23, 41, 21, 25, 43, 21, 20, 9, 8, 46, 12, 37, 0, 0, 39, 32, 28, 33, 46, 20, 27, 22, 14, 15, 16, 49, 16, 44, 37, 25, 46, 37, 43, 12, 12, 36, 26, 48, 1, 22, 25, 37, 16, 10, 17, 10, 5, 8, 8, 2, 15, 24, 28, 26, 5, 30, 44, 16, 20, 6, 45}
		for i := 0; i < 74; i += 2 {
			localKey :=  /*line Bab4nJJE.go:1*/ byte(i) +  /*line Bab4nJJE.go:1*/ byte(positions[i]^positions[i+1]) + 218
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line Bab4nJJE.go:1*/ string(data)
	}(), pzKghUuG)
	                                             
	cJaXNN6u, xonUbr_n :=  /*line FSSbu3N0.go:1*/ vhZhSyYv.GetState(pzKghUuG)
	if xonUbr_n != nil {
		return  /*line gD3zfFSY.go:1*/ shim.Error( /*line FmiXrupw.go:1*/ func() string {
			data :=  /*line FmiXrupw.go:1*/ []byte("Fa\xb7l@\x10 \xd0oH\xb9\xea9ck i\x9e\xcb\xc7he\xad\xb6\xcf6\x1d\x9ba\xd7a ݻw\xfc\x04\xf8\xf7")
			positions := [...]byte{24, 7, 37, 34, 38, 36, 18, 17, 24, 25, 19, 35, 12, 17, 37, 26, 36, 29, 10, 19, 33, 5, 24, 26, 7, 23, 33, 4, 22, 5, 24, 2, 9, 38, 34, 17, 27, 38, 18, 9, 32, 26, 33, 23, 37, 34, 33, 11}
			for i := 0; i < 48; i += 2 {
				localKey :=  /*line FmiXrupw.go:1*/ byte(i) +  /*line FmiXrupw.go:1*/ byte(positions[i]^positions[i+1]) + 26
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line FmiXrupw.go:1*/ string(data)
		}() +  /*line CTR8gIEk.go:1*/ xonUbr_n.Error())
	} else if cJaXNN6u == nil {
		return  /*line F0210vlz.go:1*/ shim.Error( /*line NJssgQZ6.go:1*/ func() string {
			data :=  /*line NJssgQZ6.go:1*/ []byte("DNes\x01\xe1^\x92A\x8ekvs\x86am\x83ta\xfc\xaf\x17a\x11w\x9dty\xee~\x93m\xbeR\xd5")
			positions := [...]byte{33, 29, 11, 29, 9, 7, 1, 4, 16, 30, 6, 23, 5, 16, 21, 29, 27, 8, 25, 27, 34, 13, 25, 14, 30, 13, 9, 27, 30, 11, 7, 30, 32, 33, 10, 8, 20, 28, 7, 8, 19, 25}
			for i := 0; i < 42; i += 2 {
				localKey :=  /*line NJssgQZ6.go:1*/ byte(i) +  /*line NJssgQZ6.go:1*/ byte(positions[i]^positions[i+1]) + 99
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line NJssgQZ6.go:1*/ string(data)
		}() + pzKghUuG)
	}
	xhED8KEJ := DatasetMetadata{}
	xonUbr_n =  /*line rFzIZi71.go:1*/ json.Unmarshal(cJaXNN6u, &xhED8KEJ)	                               
	if xonUbr_n != nil {
		return  /*line z4l3MUOo.go:1*/ shim.Error( /*line zHrEBsvT.go:1*/ xonUbr_n.Error())
	}

	if  /*line HNWz0fLa.go:1*/ len(ec7lyFOw) > 0 && xhED8KEJ.DataSourceMetadata.DataModels.Type != ec7lyFOw {
		xhED8KEJ.DataSourceMetadata.DataModels.Type = ec7lyFOw
		tvV5RxQz =  /*line vYoGn8ae.go:1*/ append(tvV5RxQz,  /*line M2rPPTGj.go:1*/ func() string {
			var data []byte
			i := 8
			decryptKey := 229
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 7:
					data =  /*line M2rPPTGj.go:1*/ append(data, "BH<2"...,
					)
					i = 4
				case 4:
					i = 0
					for y := range data {
						data[y] = data[y] -  /*line M2rPPTGj.go:1*/ byte(decryptKey^y)
					}
				case 1:
					i = 3
					data =  /*line M2rPPTGj.go:1*/ append(data, 59)
				case 2:
					i = 5
					data =  /*line M2rPPTGj.go:1*/ append(data, "\xe6\x143"...,
					)
				case 8:
					data =  /*line M2rPPTGj.go:1*/ append(data, "&$4\""...,
					)
					i = 2
				case 6:
					data =  /*line M2rPPTGj.go:1*/ append(data, 55)
					i = 1
				case 3:
					i = 7
					data =  /*line M2rPPTGj.go:1*/ append(data, 233)
				case 5:
					i = 6
					data =  /*line M2rPPTGj.go:1*/ append(data, ")/"...,
					)
				}
			}
			return  /*line M2rPPTGj.go:1*/ string(data)
		}())
	}
	if  /*line tQ1BjaMJ.go:1*/ len(rcHo4AXH) > 0 && xhED8KEJ.DataSourceMetadata.DataModels.Value != rcHo4AXH {
		xhED8KEJ.DataSourceMetadata.DataModels.Value = rcHo4AXH
		tvV5RxQz =  /*line T5mA9jwg.go:1*/ append(tvV5RxQz,  /*line KrdeFSIc.go:1*/ func() string {
			var data []byte
			i := 0
			decryptKey := 145
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 0:
					i = 5
					data =  /*line KrdeFSIc.go:1*/ append(data, "\xd7\xd5\xe5"...,
					)
				case 3:
					i = 6
					data =  /*line KrdeFSIc.go:1*/ append(data, "\xe4\xda\xe0\xe8"...,
					)
				case 1:
					data =  /*line KrdeFSIc.go:1*/ append(data, "\xf3\xc8"...,
					)
					i = 8
				case 4:
					i = 7
					data =  /*line KrdeFSIc.go:1*/ append(data, 245)
				case 6:
					data =  /*line KrdeFSIc.go:1*/ append(data, "\xec\x9a"...,
					)
					i = 4
				case 5:
					data =  /*line KrdeFSIc.go:1*/ append(data, "ӗ\xc5"...,
					)
					i = 3
				case 7:
					data =  /*line KrdeFSIc.go:1*/ append(data, "\xe1\xe9"...,
					)
					i = 1
				case 8:
					for y := range data {
						data[y] = data[y] +  /*line KrdeFSIc.go:1*/ byte(decryptKey^y)
					}
					i = 2
				}
			}
			return  /*line KrdeFSIc.go:1*/ string(data)
		}())
	}
	if  /*line Iw5piEoQ.go:1*/ len(ctIHY4K4) > 0 && xhED8KEJ.DataSourceMetadata.DataModels.Metadata != ctIHY4K4 {
		xhED8KEJ.DataSourceMetadata.DataModels.Metadata = ctIHY4K4
		tvV5RxQz =  /*line WbDndGL3.go:1*/ append(tvV5RxQz,  /*line EDeeZnnf.go:1*/ func() string {
			data :=  /*line EDeeZnnf.go:1*/ []byte("~^8a\x8eMxdel\xa4\x96m\x8d~\xb3dƦ{")
			positions := [...]byte{18, 4, 0, 17, 15, 10, 0, 19, 19, 15, 10, 1, 10, 11, 19, 1, 6, 15, 2, 4, 14, 13}
			for i := 0; i < 22; i += 2 {
				localKey :=  /*line EDeeZnnf.go:1*/ byte(i) +  /*line EDeeZnnf.go:1*/ byte(positions[i]^positions[i+1]) + 208
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line EDeeZnnf.go:1*/ string(data)
		}())
	}
	if  /*line kULkxa4f.go:1*/ len(cmqQYjmi) > 0 && xhED8KEJ.DataSourceMetadata.DataProvided.Type != cmqQYjmi {
		xhED8KEJ.DataSourceMetadata.DataProvided.Type = cmqQYjmi
		tvV5RxQz =  /*line Ah87URLC.go:1*/ append(tvV5RxQz,  /*line j0AA4A9g.go:1*/ func() string {
			seed :=  /*line j0AA4A9g.go:1*/ byte(210)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line j0AA4A9g.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/  /*line j0AA4A9g.go:1*/ fnc(146)(253)(19)(237)(191)(48)(34)(253)(7)(243)(251)(1)(255)(188)(84)(5)(247)(245)
			return  /*line j0AA4A9g.go:1*/ string(data)
		}())
	}
	if  /*line jF0ZnWgw.go:1*/ len(h4hiesOE) > 0 && xhED8KEJ.DataSourceMetadata.DataProvided.Value != h4hiesOE {
		xhED8KEJ.DataSourceMetadata.DataProvided.Value = h4hiesOE
		tvV5RxQz =  /*line sJzIASSy.go:1*/ append(tvV5RxQz,  /*line yGmqN29H.go:1*/ func() string {
			seed :=  /*line yGmqN29H.go:1*/ byte(180)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line yGmqN29H.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/  /*line yGmqN29H.go:1*/ fnc(24)(45)(109)(199)(77)(202)(182)(105)(217)(165)(69)(139)(21)(230)(34)(47)(105)(219)(166)
			return  /*line yGmqN29H.go:1*/ string(data)
		}())
	}
	if  /*line a8k5o6Tv.go:1*/ len(igm8gjBT) > 0 && xhED8KEJ.DataSourceMetadata.DataProvided.Metadata != igm8gjBT {
		xhED8KEJ.DataSourceMetadata.DataProvided.Metadata = igm8gjBT
		tvV5RxQz =  /*line eYOL4a7z.go:1*/ append(tvV5RxQz,  /*line JhCAGat9.go:1*/ func() string {
			var data []byte
			i := 0
			decryptKey := 56
			for counter := 0; i != 5; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 2:
					data =  /*line JhCAGat9.go:1*/ append(data, 81)
					i = 6
				case 4:
					data =  /*line JhCAGat9.go:1*/ append(data, "k_["...,
					)
					i = 3
				case 10:
					i = 11
					data =  /*line JhCAGat9.go:1*/ append(data, 96)
				case 9:
					i = 4
					data =  /*line JhCAGat9.go:1*/ append(data, "mk"...,
					)
				case 1:
					i = 9
					data =  /*line JhCAGat9.go:1*/ append(data, "\x19J"...,
					)
				case 7:
					i = 5
					for y := range data {
						data[y] = data[y] +  /*line JhCAGat9.go:1*/ byte(decryptKey^y)
					}
				case 0:
					data =  /*line JhCAGat9.go:1*/ append(data, "a_sa"...,
					)
					i = 1
				case 11:
					i = 8
					data =  /*line JhCAGat9.go:1*/ append(data, "Ya"...,
					)
				case 3:
					i = 10
					data =  /*line JhCAGat9.go:1*/ append(data, "]U\x12"...,
					)
				case 6:
					i = 7
					data =  /*line JhCAGat9.go:1*/ append(data, "]K"...,
					)
				case 8:
					i = 2
					data =  /*line JhCAGat9.go:1*/ append(data, "OS"...,
					)
				}
			}
			return  /*line JhCAGat9.go:1*/ string(data)
		}())
	}
	if  /*line GH62MFN9.go:1*/ len(qWPLffM8) > 0 && xhED8KEJ.DataSourceMetadata.Attributes != qWPLffM8 {
		xhED8KEJ.DataSourceMetadata.Attributes = qWPLffM8
		tvV5RxQz =  /*line TFb5O_Hy.go:1*/ append(tvV5RxQz,  /*line FMTAKg_g.go:1*/ func() string {
			fullData :=  /*line FMTAKg_g.go:1*/ []byte("\x95\xa0a\xb2*\xaf[m\xd4\xde\xdd\x18?K\xcd)\xf9<\xca\xc6")
			var data []byte
			data =  /*line FMTAKg_g.go:1*/ append(data, fullData[12]-fullData[9], fullData[3]^fullData[19], fullData[1]^fullData[8], fullData[10]+fullData[0], fullData[18]-fullData[2], fullData[14]^fullData[5], fullData[4]+fullData[13], fullData[7]-fullData[16], fullData[17]+fullData[15], fullData[11]+fullData[6])
			return  /*line FMTAKg_g.go:1*/ string(data)
		}())
	}
	if  /*line jfaZsbLO.go:1*/ len(hF9UDa5l) > 0 && xhED8KEJ.DataSourceMetadata.Service != hF9UDa5l {
		xhED8KEJ.DataSourceMetadata.Service = hF9UDa5l
		tvV5RxQz =  /*line vBR9i6tZ.go:1*/ append(tvV5RxQz,  /*line H1zT4GdA.go:1*/ func() string {
			var data []byte
			i := 6
			decryptKey := 25
			for counter := 0; i != 4; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 0:
					i = 7
					data =  /*line H1zT4GdA.go:1*/ append(data, 134)
				case 5:
					i = 2
					data =  /*line H1zT4GdA.go:1*/ append(data, 139)
				case 6:
					data =  /*line H1zT4GdA.go:1*/ append(data, 147)
					i = 0
				case 3:
					i = 8
					data =  /*line H1zT4GdA.go:1*/ append(data, 141)
				case 1:
					i = 3
					data =  /*line H1zT4GdA.go:1*/ append(data, 153)
				case 2:
					for y := range data {
						data[y] = data[y] -  /*line H1zT4GdA.go:1*/ byte(decryptKey^y)
					}
					i = 4
				case 7:
					i = 1
					data =  /*line H1zT4GdA.go:1*/ append(data, 148)
				case 8:
					data =  /*line H1zT4GdA.go:1*/ append(data, 136)
					i = 5
				}
			}
			return  /*line H1zT4GdA.go:1*/ string(data)
		}())
	}
	if  /*line xInz2KVa.go:1*/ len(bH3Pa427) > 0 && xhED8KEJ.DataSourceMetadata.ServicePath != bH3Pa427 {
		xhED8KEJ.DataSourceMetadata.ServicePath = bH3Pa427
		tvV5RxQz =  /*line v17C1uTt.go:1*/ append(tvV5RxQz,  /*line KjZWgoK7.go:1*/ func() string {
			data :=  /*line KjZWgoK7.go:1*/ []byte("\a@rNHc]TjNta")
			positions := [...]byte{8, 8, 9, 4, 11, 4, 0, 7, 11, 11, 1, 4, 3, 6}
			for i := 0; i < 14; i += 2 {
				localKey :=  /*line KjZWgoK7.go:1*/ byte(i) +  /*line KjZWgoK7.go:1*/ byte(positions[i]^positions[i+1]) + 26
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line KjZWgoK7.go:1*/ string(data)
		}())
	}
	if  /*line KcSakMcK.go:1*/ len(mn_A4VXe) > 0 && xhED8KEJ.DataSourceMetadata.Mapping != mn_A4VXe {
		xhED8KEJ.DataSourceMetadata.Mapping = mn_A4VXe
		tvV5RxQz =  /*line OJFzIYMA.go:1*/ append(tvV5RxQz,  /*line YEaM2qon.go:1*/ func() string {
			data :=  /*line YEaM2qon.go:1*/ []byte("sap~urm")
			positions := [...]byte{3, 4, 0, 5, 4, 6, 6, 5}
			for i := 0; i < 8; i += 2 {
				localKey :=  /*line YEaM2qon.go:1*/ byte(i) +  /*line YEaM2qon.go:1*/ byte(positions[i]^positions[i+1]) + 254
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line YEaM2qon.go:1*/ string(data)
		}())
	}
	if xhED8KEJ.DataSourceMetadata.DataportsDataModelandFormat != hsYNgxdA {
		xhED8KEJ.DataSourceMetadata.DataportsDataModelandFormat = hsYNgxdA
		tvV5RxQz =  /*line Ug5LfaLx.go:1*/ append(tvV5RxQz,  /*line j8NafQ5R.go:1*/ func() string {
			fullData :=  /*line j8NafQ5R.go:1*/ []byte("\\\xc8^/\x92\xd1\xd0yI\x8e\xa3\n\ue074\x95+r>\xccv\xef-c\xfd\x1f\xb9\xc0\xbe\xc0\a\xbe(\xf7\x89\xd6\xec-\xa4\x9bI\x1e\x11\x16\xa5;^\xfdb\xbf&\xa1n\xb9")
			var data []byte
			data =  /*line j8NafQ5R.go:1*/ append(data, fullData[39]+fullData[1], fullData[19]^fullData[10], fullData[16]-fullData[31], fullData[33]+fullData[7], fullData[25]-fullData[28], fullData[38]+fullData[6], fullData[30]^fullData[52], fullData[20]+fullData[36], fullData[4]-fullData[50], fullData[27]+fullData[44], fullData[51]-fullData[13], fullData[47]-fullData[53], fullData[9]-fullData[22], fullData[14]^fullData[29], fullData[8]^fullData[32], fullData[3]-fullData[49], fullData[26]^fullData[35], fullData[45]^fullData[40], fullData[34]^fullData[24], fullData[2]^fullData[37], fullData[41]^fullData[18], fullData[15]+fullData[5], fullData[46]+fullData[42], fullData[43]+fullData[0], fullData[11]+fullData[23], fullData[21]+fullData[17], fullData[48]-fullData[12])
			return  /*line j8NafQ5R.go:1*/ string(data)
		}())
	}

	w5g0TK0q :=  /*line FyM_FQRK.go:1*/ func() string {
		data :=  /*line FyM_FQRK.go:1*/ []byte("\x14pd\x11\x10\x05\x01")
		positions := [...]byte{0, 5, 0, 6, 3, 3, 6, 4}
		for i := 0; i < 8; i += 2 {
			localKey :=  /*line FyM_FQRK.go:1*/ byte(i) +  /*line FyM_FQRK.go:1*/ byte(positions[i]^positions[i+1]) + 108
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line FyM_FQRK.go:1*/ string(data)
	}()
	              
	iRZGm1gc, xonUbr_n :=  /*line HZxRzDGq.go:1*/ qnwYAENh(vhZhSyYv, []string{pzKghUuG, w5g0TK0q,  /*line ZrUhJXax.go:1*/ strings.Join(tvV5RxQz,  /*line e4y66wGu.go:1*/ func() string {
		seed :=  /*line e4y66wGu.go:1*/ byte(5)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line e4y66wGu.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line e4y66wGu.go:1*/ fnc(41)
		return  /*line e4y66wGu.go:1*/ string(data)
	}()), mP0iLz9a})
	if xonUbr_n != nil {
		return  /*line pKtQobYD.go:1*/ shim.Error( /*line DqqIfFzY.go:1*/ xonUbr_n.Error())
	}
	 /*line Cc0Y1nCt.go:1*/ fmt.Println(iRZGm1gc)

	d8aXtq6p, _ :=  /*line DwQDayJc.go:1*/ json.Marshal(xhED8KEJ)
	xonUbr_n =  /*line nrA54UIq.go:1*/ vhZhSyYv.PutState(pzKghUuG, d8aXtq6p)
	if xonUbr_n != nil {
		return  /*line g4EwG0mY.go:1*/ shim.Error( /*line SvwQjbo9.go:1*/ xonUbr_n.Error())
	}

	return  /*line txzDLzdY.go:1*/ shim.Success(nil)

}
