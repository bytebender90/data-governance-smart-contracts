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
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type JPyzqDxq struct{}

func (kuuXUPaZ *JPyzqDxq) Init(vhZhSyYv shim.ChaincodeStubInterface) pb.Response {
	return  /*line NUfqfwTa.go:1*/ shim.Success(nil)
}

                   
func (kuuXUPaZ *JPyzqDxq) Invoke(vhZhSyYv shim.ChaincodeStubInterface) pb.Response {

	mtyaWx76, y5yvGZQA :=  /*line QEomIeVa.go:1*/ vhZhSyYv.GetFunctionAndParameters()

	switch mtyaWx76 {
	case  /*line GshNt_1P.go:1*/ func() string {
		data :=  /*line GshNt_1P.go:1*/ []byte("u\xd1\xc3\u007f\xc0dMe\xdf\xcc\xc3\xc6\xd3\xdc")
		positions := [...]byte{11, 10, 12, 12, 3, 13, 12, 8, 9, 13, 4, 1, 2, 8}
		for i := 0; i < 14; i += 2 {
			localKey :=  /*line GshNt_1P.go:1*/ byte(i) +  /*line GshNt_1P.go:1*/ byte(positions[i]^positions[i+1]) + 161
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line GshNt_1P.go:1*/ string(data)
	}():
		return  /*line XN2oz9yn.go:1*/ kuuXUPaZ.aolkF3KY(vhZhSyYv, y5yvGZQA)
	case  /*line Jjm2Azxb.go:1*/ func() string {
		key :=  /*line Jjm2Azxb.go:1*/ []byte("ӛ\u0083\xc8\x1c\xabtW_W;><")
		data :=  /*line Jjm2Azxb.go:1*/ []byte("H\v&\xe4<\x81\xf8\xd9\xcb\xc0\xbb\x9c\xb2\x9d")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line Jjm2Azxb.go:1*/ string(data)
	}():
		return  /*line UYEGgoAI.go:1*/ kuuXUPaZ.sgFZSQF7(vhZhSyYv, y5yvGZQA)
	case  /*line DnwL6NRB.go:1*/ func() string {
		seed :=  /*line DnwL6NRB.go:1*/ byte(147)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line DnwL6NRB.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/  /*line DnwL6NRB.go:1*/ fnc(246)(251)(233)(206)(175)(79)(125)(23)(65)(111)(208)(188)(126)(249)(227)(200)(120)(8)(31)(43)(89)(175)(113)(207)
		return  /*line DnwL6NRB.go:1*/ string(data)
	}():
		return  /*line IT4_QgYQ.go:1*/ kuuXUPaZ.uDyCIlIB(vhZhSyYv, y5yvGZQA)
	case  /*line yP75Y0qD.go:1*/ func() string {
		data :=  /*line yP75Y0qD.go:1*/ []byte("\x10-KctLJ\xf5\xfb<SP>=ce\xf8\tta<aP?")
		positions := [...]byte{9, 1, 8, 12, 12, 16, 13, 2, 12, 7, 6, 17, 12, 6, 0, 5, 12, 11, 7, 5, 23, 20, 3, 22}
		for i := 0; i < 24; i += 2 {
			localKey :=  /*line yP75Y0qD.go:1*/ byte(i) +  /*line yP75Y0qD.go:1*/ byte(positions[i]^positions[i+1]) + 196
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line yP75Y0qD.go:1*/ string(data)
	}():
		return  /*line DRkMvuwR.go:1*/ kuuXUPaZ.efFVT2Sf(vhZhSyYv, y5yvGZQA)
	case  /*line f798Fvjk.go:1*/ func() string {
		fullData :=  /*line f798Fvjk.go:1*/ []byte("\x81\xd2\r\x15\x82\x1a\xbb\x14;O\x80\v%\xa1\xf9Q\x91\x9f%\xd2eD\x908\x97\xdc\xde.<\x1a@\xba\x02\xa8Z\xa8G\xe3")
		var data []byte
		data =  /*line f798Fvjk.go:1*/ append(data, fullData[20]^fullData[7], fullData[24]+fullData[26], fullData[37]+fullData[4], fullData[5]-fullData[33], fullData[14]-fullData[10], fullData[2]+fullData[30], fullData[34]+fullData[11], fullData[28]+fullData[23], fullData[36]+fullData[29], fullData[17]-fullData[8], fullData[21]^fullData[12], fullData[25]^fullData[35], fullData[27]^fullData[9], fullData[0]+fullData[19], fullData[18]^fullData[15], fullData[32]-fullData[22], fullData[6]+fullData[31], fullData[16]+fullData[1], fullData[3]-fullData[13])
		return  /*line f798Fvjk.go:1*/ string(data)
	}():
		return  /*line kM8nWpsh.go:1*/ kuuXUPaZ.leLg2iBJ(vhZhSyYv)
	case  /*line y0OK5sRL.go:1*/ func() string {
		key :=  /*line y0OK5sRL.go:1*/ []byte("\xde\fs\x11ڨm\x83\xb5\xe7\xe1\xd6\x11a\x01\xfaGCf")
		data :=  /*line y0OK5sRL.go:1*/ []byte("O\x81\u0603S\xf5\xd2\xf7\x16KBJr\xa3zH\xa8\xb0\xcb")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line y0OK5sRL.go:1*/ string(data)
	}():
		return  /*line ukvGEzmz.go:1*/ kuuXUPaZ.jddDqpYa(vhZhSyYv, y5yvGZQA)
	case  /*line xzZx4zxw.go:1*/ func() string {
		seed :=  /*line xzZx4zxw.go:1*/ byte(112)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line xzZx4zxw.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/  /*line xzZx4zxw.go:1*/ fnc(1)(4)(16)(247)(5)(204)(40)(1)(23)(233)(23)(249)(231)(47)(229)(197)(39)(25)(231)(62)(196)(26)(251)(231)(14)(48)(237)
		return  /*line xzZx4zxw.go:1*/ string(data)
	}():
		return  /*line PJnnPy_p.go:1*/ kuuXUPaZ.fKDeNQZA(vhZhSyYv, y5yvGZQA)
	case  /*line ZSSVbl13.go:1*/ func() string {
		data :=  /*line ZSSVbl13.go:1*/ []byte("p\xddYryManC\xfd\xd2t\xc8\xd1y\xe6\xd8\xddkr")
		positions := [...]byte{9, 8, 9, 7, 10, 18, 9, 10, 9, 2, 17, 2, 6, 16, 15, 15, 16, 16, 12, 0, 12, 1, 9, 13}
		for i := 0; i < 24; i += 2 {
			localKey :=  /*line ZSSVbl13.go:1*/ byte(i) +  /*line ZSSVbl13.go:1*/ byte(positions[i]^positions[i+1]) + 155
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line ZSSVbl13.go:1*/ string(data)
	}():
		return  /*line CXBALqcS.go:1*/ kuuXUPaZ.pxhwNbVy(vhZhSyYv)
	case  /*line mAurtpuz.go:1*/ func() string {
		seed :=  /*line mAurtpuz.go:1*/ byte(210)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line mAurtpuz.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/  /*line mAurtpuz.go:1*/ fnc(159)(4)(240)(13)(7)(212)(24)(15)(237)(3)(253)(19)(237)(225)(55)(214)(40)(247)(247)(13)(207)(45)(246)(224)(29)(19)(237)(18)(242)(15)(218)(19)(12)(248)
		return  /*line mAurtpuz.go:1*/ string(data)
	}():
		return  /*line zzcOp0SA.go:1*/ kuuXUPaZ.yxMJzmq_(vhZhSyYv, y5yvGZQA)
	case  /*line rxCUWtNB.go:1*/ func() string {
		fullData :=  /*line rxCUWtNB.go:1*/ []byte("\x92J\x80>\xcfҋŮ\x98\xefna\xdb\xe9;\x9d\t\x9a=7L\xaa\x1d\x1e\xa6<\xd7\xef\xa4~Q4'\x00\xf7y\xe3P|\x14>̪4\xb3")
		var data []byte
		data =  /*line rxCUWtNB.go:1*/ append(data, fullData[32]+fullData[19], fullData[40]^fullData[12], fullData[11]-fullData[17], fullData[21]^fullData[3], fullData[36]-fullData[34], fullData[14]^fullData[29], fullData[35]-fullData[0], fullData[24]-fullData[43], fullData[38]-fullData[10], fullData[8]-fullData[1], fullData[9]-fullData[20], fullData[7]-fullData[31], fullData[39]^fullData[23], fullData[30]-fullData[26], fullData[18]^fullData[37], fullData[13]-fullData[6], fullData[41]-fullData[42], fullData[28]-fullData[2], fullData[16]-fullData[33], fullData[15]-fullData[5], fullData[45]^fullData[27], fullData[4]^fullData[22], fullData[25]-fullData[44])
		return  /*line rxCUWtNB.go:1*/ string(data)
	}():
		return  /*line l6CZGHoW.go:1*/ kuuXUPaZ.wYYPiY8F(vhZhSyYv)
	case  /*line togK3R8K.go:1*/ func() string {
		fullData :=  /*line togK3R8K.go:1*/ []byte("\x8e\x17V\xef>l\xe5\x8f\xd9\xe6V\xb3\x00\xb8#\xf8\xfe\xcb\xeb\xfa\x9d\xad\xb6\x9d\x96\xc3J?\xdaI؆a\x8d\xb4\xdd\x17\x99V\xac\x84h\xf2\x8f$ƿ\x10\r(2\xa0\xd3\x1f\xba\fr\x88]=|e\xaa\xf9\xb4\x8cp\xcd\xe5\x1cL\xb9\xd9^")
		var data []byte
		data =  /*line togK3R8K.go:1*/ append(data, fullData[60]^fullData[48], fullData[15]^fullData[33], fullData[1]^fullData[56], fullData[20]^fullData[3], fullData[58]^fullData[44], fullData[47]+fullData[59], fullData[73]-fullData[63], fullData[21]^fullData[8], fullData[52]+fullData[0], fullData[2]^fullData[50], fullData[51]-fullData[27], fullData[19]-fullData[31], fullData[62]^fullData[17], fullData[61]-fullData[14], fullData[22]+fullData[25], fullData[11]+fullData[23], fullData[4]^fullData[70], fullData[13]-fullData[29], fullData[9]-fullData[66], fullData[6]+fullData[40], fullData[10]-fullData[42], fullData[53]-fullData[54], fullData[45]^fullData[64], fullData[36]^fullData[38], fullData[34]^fullData[28], fullData[18]^fullData[43], fullData[49]^fullData[5], fullData[67]^fullData[39], fullData[12]-fullData[65], fullData[71]^fullData[30], fullData[35]+fullData[24], fullData[16]-fullData[37], fullData[69]^fullData[41], fullData[7]+fullData[46], fullData[57]+fullData[72], fullData[55]^fullData[32], fullData[26]-fullData[68])
		return  /*line togK3R8K.go:1*/ string(data)
	}():
		return  /*line brdynTFU.go:1*/ kuuXUPaZ.z0P02F2H(vhZhSyYv, y5yvGZQA)
	case  /*line EkdjOJK7.go:1*/ func() string {
		var data []byte
		i := 11
		decryptKey := 156
		for counter := 0; i != 4; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 10:
				i = 7
				data =  /*line EkdjOJK7.go:1*/ append(data, 30)
			case 2:
				i = 1
				data =  /*line EkdjOJK7.go:1*/ append(data, "\x1bU"...,
				)
			case 12:
				data =  /*line EkdjOJK7.go:1*/ append(data, "9G?"...,
				)
				i = 9
			case 0:
				for y := range data {
					data[y] = data[y] +  /*line EkdjOJK7.go:1*/ byte(decryptKey^y)
				}
				i = 4
			case 3:
				i = 5
				data =  /*line EkdjOJK7.go:1*/ append(data, ")5"...,
				)
			case 14:
				i = 10
				data =  /*line EkdjOJK7.go:1*/ append(data, "7<>"...,
				)
			case 8:
				data =  /*line EkdjOJK7.go:1*/ append(data, 44)
				i = 3
			case 7:
				i = 0
				data =  /*line EkdjOJK7.go:1*/ append(data, "A71="...,
				)
			case 9:
				i = 2
				data =  /*line EkdjOJK7.go:1*/ append(data, "AAS;"...,
				)
			case 11:
				data =  /*line EkdjOJK7.go:1*/ append(data, 71)
				i = 13
			case 1:
				i = 8
				data =  /*line EkdjOJK7.go:1*/ append(data, "*8"...,
				)
			case 13:
				data =  /*line EkdjOJK7.go:1*/ append(data, "J="...,
				)
				i = 6
			case 5:
				i = 14
				data =  /*line EkdjOJK7.go:1*/ append(data, "+;%7"...,
				)
			case 6:
				i = 12
				data =  /*line EkdjOJK7.go:1*/ append(data, "IK\x1e"...,
				)
			}
		}
		return  /*line EkdjOJK7.go:1*/ string(data)
	}():
		return  /*line Mv0Ob2Xu.go:1*/ kuuXUPaZ.dwIdWa90(vhZhSyYv)
	case  /*line GM6o4owF.go:1*/ func() string {
		data :=  /*line GM6o4owF.go:1*/ []byte("quAr+';\x19GkF>a!yOrg2W\x01W$y&\r\\q\x05\x04e\x1c#nd~ati@\x00\x1cN1'>")
		positions := [...]byte{27, 20, 4, 8, 8, 6, 32, 6, 25, 31, 26, 44, 25, 35, 7, 7, 7, 43, 45, 39, 9, 11, 39, 41, 9, 22, 44, 27, 24, 6, 21, 19, 2, 38, 13, 44, 19, 38, 23, 28, 29, 5, 4, 23, 38, 28, 25, 23, 24, 25, 8, 2, 43, 41, 40, 10, 32, 18, 38, 13, 40, 39, 31, 43}
		for i := 0; i < 64; i += 2 {
			localKey :=  /*line GM6o4owF.go:1*/ byte(i) +  /*line GM6o4owF.go:1*/ byte(positions[i]^positions[i+1]) + 9
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line GM6o4owF.go:1*/ string(data)
	}():
		return  /*line Zi21uSAz.go:1*/ kuuXUPaZ.cuM598CU(vhZhSyYv, y5yvGZQA)
	case  /*line UukjTo_X.go:1*/ func() string {
		seed :=  /*line UukjTo_X.go:1*/ byte(204)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line UukjTo_X.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/  /*line UukjTo_X.go:1*/ fnc(165)(4)(240)(13)(7)(212)(24)(15)(237)(3)(253)(19)(237)(225)(55)(216)(36)(240)(13)(7)(218)(33)(254)(247)(5)(249)
		return  /*line UukjTo_X.go:1*/ string(data)
	}():
		return  /*line lNIvr3ej.go:1*/ kuuXUPaZ.kbgtOUnn(vhZhSyYv, y5yvGZQA)
	case  /*line H9reJwTQ.go:1*/ func() string {
		data :=  /*line H9reJwTQ.go:1*/ []byte("\x17\xbet'e\xcd\x06\xe9\xd9\xe0\x1dPr\x94\xfa+d\\\xd8")
		positions := [...]byte{7, 13, 7, 9, 17, 8, 5, 17, 18, 3, 1, 17, 8, 17, 10, 0, 15, 13, 9, 7, 3, 14, 6, 17}
		for i := 0; i < 24; i += 2 {
			localKey :=  /*line H9reJwTQ.go:1*/ byte(i) +  /*line H9reJwTQ.go:1*/ byte(positions[i]^positions[i+1]) + 50
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line H9reJwTQ.go:1*/ string(data)
	}():
		return  /*line sc9VbZD3.go:1*/ kuuXUPaZ.fnIzwh3h(vhZhSyYv, y5yvGZQA)
	case  /*line KF1CzAV7.go:1*/ func() string {
		key :=  /*line KF1CzAV7.go:1*/ []byte("/\xc5\xdat\xcb.dIC\t\x16\xf7\xa3\xf5d/t\xf3\x99#\xfa")
		data :=  /*line KF1CzAV7.go:1*/ []byte("H\xa0\xae9\xaeZ\x05-\"}w\xbcƌ\x13@\x06\x97\xcdB\x9d")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line KF1CzAV7.go:1*/ string(data)
	}():
		return  /*line HaFf0oj9.go:1*/ kuuXUPaZ.gwgznRFf(vhZhSyYv, y5yvGZQA)
	case  /*line ARAACpZG.go:1*/ func() string {
		var data []byte
		i := 1
		decryptKey := 149
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 7:
				data =  /*line ARAACpZG.go:1*/ append(data, "\x02\x06"...,
				)
				i = 2
			case 4:
				for y := range data {
					data[y] = data[y] -  /*line ARAACpZG.go:1*/ byte(decryptKey^y)
				}
				i = 0
			case 1:
				data =  /*line ARAACpZG.go:1*/ append(data, "\xf0\xed\xff\xd7"...,
				)
				i = 6
			case 5:
				data =  /*line ARAACpZG.go:1*/ append(data, "\xf3\xe8\xf7"...,
				)
				i = 8
			case 3:
				i = 5
				data =  /*line ARAACpZG.go:1*/ append(data, "\xe2\xf4\xe4\xc7"...,
				)
			case 8:
				i = 7
				data =  /*line ARAACpZG.go:1*/ append(data, 245)
			case 6:
				i = 3
				data =  /*line ARAACpZG.go:1*/ append(data, "\xf2\x00\xf0\xf2"...,
				)
			case 2:
				data =  /*line ARAACpZG.go:1*/ append(data, 15)
				i = 4
			}
		}
		return  /*line ARAACpZG.go:1*/ string(data)
	}():
		return  /*line Emk0LcOw.go:1*/ kuuXUPaZ.knYHC9Gp(vhZhSyYv, y5yvGZQA)
	case  /*line jtgaNzOf.go:1*/ func() string {
		data :=  /*line jtgaNzOf.go:1*/ []byte("\br8.wePp6v\x10N\x1f5\x02na/#")
		positions := [...]byte{11, 7, 1, 3, 17, 18, 11, 3, 15, 0, 8, 2, 0, 4, 13, 1, 0, 14, 12, 18, 15, 13, 12, 7, 13, 10}
		for i := 0; i < 26; i += 2 {
			localKey :=  /*line jtgaNzOf.go:1*/ byte(i) +  /*line jtgaNzOf.go:1*/ byte(positions[i]^positions[i+1]) + 67
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line jtgaNzOf.go:1*/ string(data)
	}():
		return  /*line B6zHHFuz.go:1*/ kuuXUPaZ.mUd4z6cw(vhZhSyYv, y5yvGZQA)
	case  /*line Ytyks_No.go:1*/ func() string {
		key :=  /*line Ytyks_No.go:1*/ []byte("d-\xc3\xf5\x86\x9c\x9b@\x9c\xd8hD\x05I\xafdn\xcco\xe4\xde6U\r\x1b")
		data :=  /*line Ytyks_No.go:1*/ []byte("\rH\xa2}\xf3\xb1\xca4Ō\xf90\\\xf9\xca\xe5\xf6\x99\xff\x90\x8b0\x14XW")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line Ytyks_No.go:1*/ string(data)
	}():
		return  /*line JDOsZ46X.go:1*/ kuuXUPaZ.gaj_brT8(vhZhSyYv, y5yvGZQA)
	}

	 /*line HjUgsUEW.go:1*/ fmt.Println( /*line QdOE9RzG.go:1*/ func() string {
		data :=  /*line QdOE9RzG.go:1*/ []byte("\x8e9\xfa\xff\x88r\xf1\x87\xa3\fޮo\xac\xebf\xdb\x13e|fiݛ:\xd0")
		positions := [...]byte{11, 19, 9, 8, 4, 18, 3, 0, 3, 18, 21, 2, 9, 1, 23, 10, 13, 5, 7, 5, 13, 17, 0, 11, 14, 9, 22, 11, 7, 25, 13, 6, 2, 5, 16, 4}
		for i := 0; i < 36; i += 2 {
			localKey :=  /*line QdOE9RzG.go:1*/ byte(i) +  /*line QdOE9RzG.go:1*/ byte(positions[i]^positions[i+1]) + 90
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line QdOE9RzG.go:1*/ string(data)
	}() + mtyaWx76)	       
	return  /*line Dze1uQ2W.go:1*/ shim.Error( /*line sFgV2bDr.go:1*/ func() string {
		var data []byte
		i := 17
		decryptKey := 42
		for counter := 0; i != 11; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 6:
				i = 13
				data =  /*line sFgV2bDr.go:1*/ append(data, "\xfb\xfb\x06\xfc"...,
				)
			case 17:
				i = 0
				data =  /*line sFgV2bDr.go:1*/ append(data, "\xd3\xe5"...,
				)
			case 13:
				i = 3
				data =  /*line sFgV2bDr.go:1*/ append(data, "\xb1\xf6"...,
				)
			case 14:
				i = 9
				data =  /*line sFgV2bDr.go:1*/ append(data, 4)
			case 5:
				data =  /*line sFgV2bDr.go:1*/ append(data, "\v\x02\xff"...,
				)
				i = 4
			case 10:
				data =  /*line sFgV2bDr.go:1*/ append(data, "\x12\x10"...,
				)
				i = 15
			case 2:
				data =  /*line sFgV2bDr.go:1*/ append(data, "\xa9\xfd\xf9\xf5"...,
				)
				i = 6
			case 0:
				data =  /*line sFgV2bDr.go:1*/ append(data, 230)
				i = 8
			case 7:
				i = 14
				data =  /*line sFgV2bDr.go:1*/ append(data, "\x05\a\xb8"...,
				)
			case 12:
				data =  /*line sFgV2bDr.go:1*/ append(data, 248)
				i = 1
			case 8:
				i = 16
				data =  /*line sFgV2bDr.go:1*/ append(data, "\xe7\xee"...,
				)
			case 16:
				i = 2
				data =  /*line sFgV2bDr.go:1*/ append(data, "\xfa\xec\xea"...,
				)
			case 3:
				i = 12
				data =  /*line sFgV2bDr.go:1*/ append(data, "\b\x00"...,
				)
			case 4:
				i = 10
				data =  /*line sFgV2bDr.go:1*/ append(data, "\x15\t"...,
				)
			case 15:
				i = 11
				for y := range data {
					data[y] = data[y] -  /*line sFgV2bDr.go:1*/ byte(decryptKey^y)
				}
			case 1:
				data =  /*line sFgV2bDr.go:1*/ append(data, "\b\x00"...,
				)
				i = 7
			case 9:
				data =  /*line sFgV2bDr.go:1*/ append(data, "\b\x13"...,
				)
				i = 5
			}
		}
		return  /*line sFgV2bDr.go:1*/ string(data)
	}())
}

func main() {
	xonUbr_n :=  /*line k3TjStHd.go:1*/ shim.Start( /*line EKNDeEy5.go:1*/ new(JPyzqDxq))
	if xonUbr_n != nil {
		 /*line VI5Dfxc8.go:1*/ fmt.Printf( /*line nku0uqud.go:1*/ func() string {
			var data []byte
			i := 8
			decryptKey := 36
			for counter := 0; i != 1; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 0:
					i = 5
					data =  /*line nku0uqud.go:1*/ append(data, "'y"...,
					)
				case 12:
					data =  /*line nku0uqud.go:1*/ append(data, ",'!+"...,
					)
					i = 6
				case 5:
					i = 4
					data =  /*line nku0uqud.go:1*/ append(data, "\\X\r"...,
					)
				case 13:
					i = 0
					data =  /*line nku0uqud.go:1*/ append(data, "0-"...,
					)
				case 3:
					i = 12
					data =  /*line nku0uqud.go:1*/ append(data, "r0$"...,
					)
				case 4:
					for y := range data {
						data[y] = data[y] ^  /*line nku0uqud.go:1*/ byte(decryptKey^y)
					}
					i = 1
				case 8:
					i = 10
					data =  /*line nku0uqud.go:1*/ append(data, "\x19/,0"...,
					)
				case 6:
					i = 11
					data =  /*line nku0uqud.go:1*/ append(data, "&..d"...,
					)
				case 10:
					data =  /*line nku0uqud.go:1*/ append(data, "*y"...,
					)
					i = 2
				case 2:
					i = 7
					data =  /*line nku0uqud.go:1*/ append(data, 41)
				case 11:
					data =  /*line nku0uqud.go:1*/ append(data, "6'*"...,
					)
					i = 13
				case 7:
					i = 9
					data =  /*line nku0uqud.go:1*/ append(data, "/5'\""...,
					)
				case 9:
					data =  /*line nku0uqud.go:1*/ append(data, ">>6"...,
					)
					i = 3
				}
			}
			return  /*line nku0uqud.go:1*/ string(data)
		}(), xonUbr_n)
	}
}
