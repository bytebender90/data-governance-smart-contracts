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

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (u3xzONWW *ZohimWzR) zHAosy7V(ipEgxdqE shim.ChaincodeStubInterface) pb.Response {
	iztzjRXX, dTakwHTF :=  /*line Z3_oBJHv.go:1*/ u3xzONWW.rjGOlCPK(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line GJF7CpL7.go:1*/ shim.Error( /*line MQTUhhPC.go:1*/ dTakwHTF.Error())
	}
	_br98sk5, dTakwHTF :=  /*line ByDzVUPk.go:1*/ u3xzONWW.qbaWnNO0(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line AQDJYzA6.go:1*/ shim.Error( /*line HVRkpfyw.go:1*/ dTakwHTF.Error())
	}
	d9zLne_w :=  /*line FEfRLji2.go:1*/ fmt.Sprintf( /*line ZP9iRb4B.go:1*/ func() string {
		data :=  /*line ZP9iRb4B.go:1*/ []byte("\xf6\xcb҃\"\xa9\x19to\x12\x8d\xab\x80Zdb\xddTo\xb1\x15\xf7I\xd7\x19\u007f.r\x92\xb1ln\xf5\xbf#ux٦n\xddm\xb0thús\xf9\x83\x18o\xaega\xf1P\x9ca\xcc\xe7t?\x99G\"wr\"\x16}")
		positions := [...]byte{30, 20, 15, 9, 9, 24, 22, 10, 61, 59, 62, 23, 6, 56, 29, 50, 45, 25, 2, 52, 19, 9, 37, 40, 48, 20, 20, 11, 67, 1, 52, 16, 29, 25, 38, 55, 33, 60, 1, 59, 34, 46, 3, 57, 33, 34, 10, 11, 46, 45, 42, 33, 29, 15, 16, 63, 5, 0, 5, 50, 30, 6, 23, 21, 32, 34, 46, 26, 28, 12, 64, 29, 36, 29, 66, 44, 56, 49, 43, 64, 67, 38, 69, 62, 4, 33, 2, 51, 36, 62, 52, 13, 15, 31, 18, 31}
		for i := 0; i < 96; i += 2 {
			localKey :=  /*line ZP9iRb4B.go:1*/ byte(i) +  /*line ZP9iRb4B.go:1*/ byte(positions[i]^positions[i+1]) + 149
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line ZP9iRb4B.go:1*/ string(data)
	}(), iztzjRXX, _br98sk5)
	 /*line Hx4cXnck.go:1*/ fmt.Println(d9zLne_w)
	_hnXa_XS, dTakwHTF :=  /*line vNHTUrnP.go:1*/ wphqE9Co(ipEgxdqE, d9zLne_w)
	if dTakwHTF != nil {
		return  /*line GQbacz7x.go:1*/ shim.Error( /*line Uf3PI6oJ.go:1*/ dTakwHTF.Error())
	}
	return  /*line v_DziOlq.go:1*/ shim.Success(_hnXa_XS)
}

func (u3xzONWW *ZohimWzR) mdfCz3FT(ipEgxdqE shim.ChaincodeStubInterface, lsLFSQCl []string) pb.Response {
	f3MEWGdr := lsLFSQCl[0]

	d9zLne_w :=  /*line AKxT293P.go:1*/ fmt.Sprintf( /*line OO2zR3f1.go:1*/ func() string {
		fullData :=  /*line OO2zR3f1.go:1*/ []byte("\xa9\xb4\xe5#\xd2\xfb\xfe\xf0\xa0:\xa2JG\x04T\x15\x890 \xbb:\xdcTqyY\xc4t\x97x\x84\x8cぶ\x130/a\x99\x15\xe0\xf1\xd9\xe8\xbe _\x181\x9eJ9\xdb:B\x9f\xca\r\xa6\xb8\xdbؘY\xb0\x8e\x18S\x91$C\xb5\xfb\x97\x05\"\x9bX);F\xee\xcfda5\xf4\xf7\a\x16<\x0e\xd5\xf3\xd0")
		var data []byte
		data =  /*line OO2zR3f1.go:1*/ append(data, fullData[76]^fullData[64], fullData[49]^fullData[35], fullData[21]+fullData[74], fullData[50]-fullData[52], fullData[24]+fullData[94], fullData[42]+fullData[27], fullData[36]^fullData[68], fullData[45]+fullData[34], fullData[11]-fullData[53], fullData[10]^fullData[95], fullData[79]-fullData[89], fullData[82]-fullData[1], fullData[56]-fullData[70], fullData[59]^fullData[30], fullData[60]-fullData[14], fullData[83]^fullData[8], fullData[57]+fullData[39], fullData[86]^fullData[38], fullData[0]-fullData[17], fullData[67]+fullData[78], fullData[25]-fullData[87], fullData[72]^fullData[28], fullData[37]^fullData[15], fullData[40]+fullData[58], fullData[69]-fullData[91], fullData[33]-fullData[92], fullData[3]+fullData[55], fullData[12]-fullData[93], fullData[19]^fullData[4], fullData[22]^fullData[9], fullData[18]+fullData[81], fullData[51]-fullData[61], fullData[43]^fullData[5], fullData[13]-fullData[62], fullData[44]+fullData[54], fullData[47]^fullData[20], fullData[84]-fullData[88], fullData[23]+fullData[7], fullData[16]-fullData[46], fullData[41]+fullData[31], fullData[65]-fullData[66], fullData[6]-fullData[26], fullData[71]^fullData[85], fullData[80]-fullData[90], fullData[29]+fullData[73], fullData[75]-fullData[32], fullData[48]-fullData[77], fullData[63]+fullData[2])
		return  /*line OO2zR3f1.go:1*/ string(data)
	}(), f3MEWGdr)
	 /*line nrykYzNv.go:1*/ fmt.Println(d9zLne_w)
	_hnXa_XS, dTakwHTF :=  /*line EhQBpURA.go:1*/ wphqE9Co(ipEgxdqE, d9zLne_w)
	if dTakwHTF != nil {
		return  /*line zjscE1uz.go:1*/ shim.Error( /*line Irjq8dl5.go:1*/ dTakwHTF.Error())
	}
	return  /*line DzbHKeig.go:1*/ shim.Success(_hnXa_XS)
}

func (u3xzONWW *ZohimWzR) qxnshQm_(ipEgxdqE shim.ChaincodeStubInterface, lsLFSQCl []string) pb.Response {
	iztzjRXX := lsLFSQCl[0]

	d9zLne_w :=  /*line Edj4syzK.go:1*/ fmt.Sprintf( /*line pP4MfFSH.go:1*/ func() string {
		seed :=  /*line pP4MfFSH.go:1*/ byte(155)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line pP4MfFSH.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/  /*line pP4MfFSH.go:1*/ fnc(224)(167)(81)(242)(7)(249)(254)(17)(251)(3)(176)(24)(65)(167)(66)(11)(244)(241)(37)(247)(245)(189)(24)(232)(51)(30)(242)(13)(247)(5)(248)(9)(179)(10)(246)(83)(254)(242)(13)(252)(243)(12)(248)(189)(24)(232)(3)(78)(175)(91)(0)
		return  /*line pP4MfFSH.go:1*/ string(data)
	}(), iztzjRXX)
	 /*line hMfJLKVr.go:1*/ fmt.Println(d9zLne_w)
	_hnXa_XS, dTakwHTF :=  /*line fARgncXt.go:1*/ wphqE9Co(ipEgxdqE, d9zLne_w)
	if dTakwHTF != nil {
		return  /*line _D3K_RkY.go:1*/ shim.Error( /*line xTUjzrWm.go:1*/ dTakwHTF.Error())
	}
	return  /*line oKIkf0Fk.go:1*/ shim.Success(_hnXa_XS)
}

func (u3xzONWW *ZohimWzR) spza01Hy(ipEgxdqE shim.ChaincodeStubInterface) pb.Response {
	d9zLne_w :=  /*line IxX6bIhN.go:1*/ func() string {
		data :=  /*line IxX6bIhN.go:1*/ []byte("\x18a\vel&4D#r\x11pM\x10BH/tq'aV\x1c\"=DYNn*'p1}")
		positions := [...]byte{5, 14, 27, 10, 22, 26, 18, 20, 10, 10, 14, 19, 7, 18, 25, 12, 20, 30, 2, 15, 22, 15, 20, 21, 0, 10, 8, 32, 24, 31, 10, 26, 16, 29, 10, 26, 6, 27, 13, 19, 22, 1, 26, 15, 31, 17, 11, 11}
		for i := 0; i < 48; i += 2 {
			localKey :=  /*line IxX6bIhN.go:1*/ byte(i) +  /*line IxX6bIhN.go:1*/ byte(positions[i]^positions[i+1]) + 28
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line IxX6bIhN.go:1*/ string(data)
	}()
	 /*line ltlr3rKT.go:1*/ fmt.Println(d9zLne_w)
	_hnXa_XS, dTakwHTF :=  /*line eAnL5XwL.go:1*/ wphqE9Co(ipEgxdqE, d9zLne_w)
	if dTakwHTF != nil {
		return  /*line KOrDb0CA.go:1*/ shim.Error( /*line UH0ViV75.go:1*/ dTakwHTF.Error())
	}
	return  /*line hZOwEb3l.go:1*/ shim.Success(_hnXa_XS)
}

func (u3xzONWW *ZohimWzR) zxDvS5er(ipEgxdqE shim.ChaincodeStubInterface) pb.Response {
	yzWCGFWr, dTakwHTF :=  /*line IFHgwjLA.go:1*/ u3xzONWW.vz2OhzsG(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line ClCvBhGc.go:1*/ shim.Error( /*line F6Rdhpqo.go:1*/ dTakwHTF.Error())
	}

	if yzWCGFWr ==  /*line EVxSH6Wn.go:1*/ func() string {
		var data []byte
		i := 1
		decryptKey := 49
		for counter := 0; i != 4; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 2:
				i = 4
				for y := range data {
					data[y] = data[y] +  /*line EVxSH6Wn.go:1*/ byte(decryptKey^y)
				}
			case 1:
				data =  /*line EVxSH6Wn.go:1*/ append(data, 69)
				i = 3
			case 0:
				i = 2
				data =  /*line EVxSH6Wn.go:1*/ append(data, 63)
			case 3:
				i = 5
				data =  /*line EVxSH6Wn.go:1*/ append(data, 66)
			case 5:
				data =  /*line EVxSH6Wn.go:1*/ append(data, 51)
				i = 0
			}
		}
		return  /*line EVxSH6Wn.go:1*/ string(data)
	}() {

		return  /*line BkaDwt3N.go:1*/ shim.Error( /*line FmQqvYJl.go:1*/ func() string {
			fullData :=  /*line FmQqvYJl.go:1*/ []byte("!\x8be\f\xe6\x1a\xc8\n\xa5&7\f,l\xaa-:s\t璓\x06\x1aL\xa8v$\x19b_\x17\xf6,IC\x19\x86\xfd\xa8\xb9u\x8d՞\xd7\xff\xc7\xc1\xaa\xe4ҿ\xb5\xbak\x8e\xf3\xe653H\xc4W\xa9(\xab\xa5\xb8y\xe4\xd8n\x13\x141\xd5|\x8f\xbat&\xf4\xe5K\xf6\xec\xb4]\xbe\x04\x1b\f\x05}\aK\x95wi\x81\x8d\a\xcd\xf5\x91y\xc9\xed\xb8\xb2Mp\x16U~\x98ow\xdf\x1c\x06\a\n\xaa\xc7h\xaa\xf2N\x14\xe3>\xc98\xb3hk|Z=\xf5b.>/\xc4\xe4")
			var data []byte
			data =  /*line FmQqvYJl.go:1*/ append(data, fullData[124]+fullData[67], fullData[104]+fullData[106], fullData[14]-fullData[132], fullData[117]+fullData[123], fullData[118]+fullData[64], fullData[113]+fullData[84], fullData[112]^fullData[130], fullData[98]^fullData[23], fullData[4]^fullData[78], fullData[131]^fullData[42], fullData[80]^fullData[102], fullData[61]-fullData[65], fullData[13]+fullData[54], fullData[91]+fullData[93], fullData[74]+fullData[139], fullData[73]^fullData[26], fullData[20]^fullData[58], fullData[145]-fullData[68], fullData[62]-fullData[114], fullData[101]^fullData[46], fullData[3]+fullData[30], fullData[147]-fullData[146], fullData[18]^fullData[126], fullData[15]-fullData[133], fullData[22]^fullData[137], fullData[2]+fullData[90], fullData[136]+fullData[121], fullData[71]^fullData[66], fullData[1]+fullData[97], fullData[144]^fullData[88], fullData[44]-fullData[140], fullData[8]-fullData[10], fullData[110]+fullData[72], fullData[27]-fullData[135], fullData[57]^fullData[37], fullData[100]^fullData[70], fullData[81]+fullData[24], fullData[79]+fullData[52], fullData[128]+fullData[143], fullData[129]+fullData[9], fullData[111]-fullData[83], fullData[12]-fullData[47], fullData[5]^fullData[16], fullData[19]+fullData[115], fullData[29]^fullData[92], fullData[21]-fullData[0], fullData[28]-fullData[127], fullData[51]^fullData[89], fullData[59]-fullData[107], fullData[103]+fullData[116], fullData[35]-fullData[119], fullData[60]+fullData[108], fullData[138]-fullData[122], fullData[32]+fullData[94], fullData[34]-fullData[50], fullData[6]+fullData[49], fullData[69]^fullData[7], fullData[76]^fullData[141], fullData[36]^fullData[99], fullData[77]-fullData[31], fullData[41]+fullData[38], fullData[11]^fullData[33], fullData[120]^fullData[17], fullData[55]+fullData[95], fullData[105]^fullData[85], fullData[87]^fullData[43], fullData[142]-fullData[82], fullData[48]^fullData[25], fullData[86]+fullData[56], fullData[109]-fullData[63], fullData[96]-fullData[45], fullData[134]+fullData[75], fullData[39]^fullData[125], fullData[53]+fullData[40])
			return  /*line FmQqvYJl.go:1*/ string(data)
		}())

	}

	_br98sk5, dTakwHTF :=  /*line MKkegRzx.go:1*/ u3xzONWW.qbaWnNO0(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line DwNNM_zc.go:1*/ shim.Error( /*line VYXedsk7.go:1*/ dTakwHTF.Error())
	}
	d9zLne_w :=  /*line Vqwi46yl.go:1*/ fmt.Sprintf( /*line PKFwxSXR.go:1*/ func() string {
		key :=  /*line PKFwxSXR.go:1*/ []byte("C_\xa3e\xa8ʼ\xb8\x9a\x9c\x94^\x8f\xbd\xa6\x95\x0eR<\x9cBCY\xcd\xdbp\xc9B\xb2z\xa5@\xd5\xc7\xd0\xc6\xd6\xd1\xfd\x83\xc4\x16\xd7v\x94c")
		data :=  /*line PKFwxSXR.go:1*/ []byte("\xbe\x81\x16\xca\x14/\x1f,\t\x0e\xb6\x98\n\xdf\n\x04q\xa6\xb5\f\xa7e\x93\xef*\xe20\xab \xe0\x14b\x01\xe9>'C6\x1f\xbd\xe6;J\x98\x11\xe0")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line PKFwxSXR.go:1*/ string(data)
	}(), _br98sk5)
	 /*line hoJRRRty.go:1*/ fmt.Println(d9zLne_w)
	_hnXa_XS, dTakwHTF :=  /*line DG3mwmmn.go:1*/ wphqE9Co(ipEgxdqE, d9zLne_w)
	if dTakwHTF != nil {
		return  /*line jUX55_qw.go:1*/ shim.Error( /*line zCo5ZMt5.go:1*/ dTakwHTF.Error())
	}
	return  /*line BNzN8cPS.go:1*/ shim.Success(_hnXa_XS)
}

func (u3xzONWW *ZohimWzR) oaQpob1F(ipEgxdqE shim.ChaincodeStubInterface) pb.Response {
	yzWCGFWr, dTakwHTF :=  /*line CciVMZTV.go:1*/ u3xzONWW.vz2OhzsG(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line jMBcnytK.go:1*/ shim.Error( /*line UyUWHb8z.go:1*/ dTakwHTF.Error())
	}

	if yzWCGFWr ==  /*line k6VM4E0m.go:1*/ func() string {
		var data []byte
		i := 2
		decryptKey := 43
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 3:
				i = 1
				for y := range data {
					data[y] = data[y] -  /*line k6VM4E0m.go:1*/ byte(decryptKey^y)
				}
			case 0:
				data =  /*line k6VM4E0m.go:1*/ append(data, 156)
				i = 3
			case 2:
				data =  /*line k6VM4E0m.go:1*/ append(data, 158)
				i = 4
			case 5:
				data =  /*line k6VM4E0m.go:1*/ append(data, 144)
				i = 0
			case 4:
				data =  /*line k6VM4E0m.go:1*/ append(data, 155)
				i = 5
			}
		}
		return  /*line k6VM4E0m.go:1*/ string(data)
	}() {

		return  /*line SIXq85ZG.go:1*/ shim.Error( /*line CN7_zc_I.go:1*/ func() string {
			data :=  /*line CN7_zc_I.go:1*/ []byte("O\xd0\xe5\x14 \xd6d{i3s.\xdc ʥ\xaf\xcao\xdek\xfb\xefd\xb6i\xef\xe9\xb1\xc6\xd0n \xcdus%\x82^?UT\xf4\xd1n\b=\xa2ԃd\xd2j\xb6\x1cBG \x03\xa6< oj\xf9+yiza\xa7i\x9dn")
			positions := [...]byte{56, 16, 12, 72, 40, 7, 15, 46, 40, 26, 43, 2, 39, 2, 5, 66, 28, 65, 70, 7, 38, 22, 41, 45, 42, 43, 9, 35, 36, 28, 58, 59, 54, 5, 11, 46, 14, 64, 43, 29, 30, 47, 14, 11, 30, 29, 55, 14, 9, 27, 14, 7, 60, 64, 51, 27, 58, 17, 53, 14, 3, 22, 1, 33, 39, 19, 15, 54, 24, 45, 42, 7, 38, 64, 70, 49, 70, 37, 2, 35, 28, 63, 52, 21, 48, 56, 29, 5, 22, 39, 48, 17, 52, 16, 46, 55, 40, 72}
			for i := 0; i < 98; i += 2 {
				localKey :=  /*line CN7_zc_I.go:1*/ byte(i) +  /*line CN7_zc_I.go:1*/ byte(positions[i]^positions[i+1]) + 67
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line CN7_zc_I.go:1*/ string(data)
		}())

	}

	d9zLne_w :=  /*line lJ77h9d6.go:1*/ fmt.Sprintf( /*line BObD0nyA.go:1*/ func() string {
		fullData :=  /*line BObD0nyA.go:1*/ []byte("J\x8bV4#$\x99\x13<7\xb7\x90\xdcd\xa9\xe2\x99H\xe4\xab\xce7\xdaE\u007fH\x88\x87\xf6Gx1$Sl\xa4\xc4\xdfwCi\x90kA\xfb\x8a߰\x90\x8d\xb2\xd62\x93\x92)\x84\xdc(\x9bGv\x1d\xeb\x902a\x90\xc8\xe4\xe9=<\xb6\x80\v\x95\xab\x90\xc8\x11\xc5\xc5\x15\x98\xa1\x98:\x8a\xd9\xdb\xfdʼ\x11K\xbc\u0088\x0ez\x8d\xf5T\x00\x90\xf1\x02\xb3H")
		var data []byte
		data =  /*line BObD0nyA.go:1*/ append(data, fullData[101]^fullData[28], fullData[50]^fullData[11], fullData[56]-fullData[94], fullData[30]^fullData[62], fullData[26]+fullData[69], fullData[58]+fullData[71], fullData[92]+fullData[16], fullData[18]+fullData[48], fullData[95]+fullData[32], fullData[82]-fullData[33], fullData[39]^fullData[66], fullData[25]-fullData[99], fullData[3]+fullData[29], fullData[9]^fullData[83], fullData[65]+fullData[52], fullData[102]+fullData[100], fullData[97]+fullData[85], fullData[37]^fullData[1], fullData[70]^fullData[41], fullData[104]-fullData[78], fullData[2]-fullData[106], fullData[27]+fullData[59], fullData[107]-fullData[79], fullData[7]^fullData[31], fullData[53]^fullData[57], fullData[73]+fullData[96], fullData[14]^fullData[20], fullData[0]^fullData[4], fullData[24]-fullData[80], fullData[67]+fullData[51], fullData[21]-fullData[68], fullData[44]-fullData[89], fullData[13]^fullData[109], fullData[105]+fullData[54], fullData[12]+fullData[49], fullData[6]+fullData[22], fullData[74]^fullData[81], fullData[8]-fullData[36], fullData[46]-fullData[42], fullData[43]^fullData[5], fullData[17]^fullData[87], fullData[108]-fullData[23], fullData[72]-fullData[90], fullData[40]-fullData[91], fullData[35]+fullData[19], fullData[64]+fullData[15], fullData[93]+fullData[77], fullData[76]^fullData[10], fullData[47]^fullData[88], fullData[86]-fullData[61], fullData[34]-fullData[60], fullData[63]-fullData[38], fullData[84]+fullData[45], fullData[98]-fullData[75], fullData[55]^fullData[103])
		return  /*line BObD0nyA.go:1*/ string(data)
	}(), true)
	 /*line UJ2wSKqk.go:1*/ fmt.Println(d9zLne_w)
	_hnXa_XS, dTakwHTF :=  /*line MUHGohZP.go:1*/ wphqE9Co(ipEgxdqE, d9zLne_w)
	if dTakwHTF != nil {
		return  /*line UtUls6RE.go:1*/ shim.Error( /*line UWRvaxVc.go:1*/ dTakwHTF.Error())
	}
	return  /*line tgyRvFy9.go:1*/ shim.Success(_hnXa_XS)
}

func (u3xzONWW *ZohimWzR) pWAPy15K(ipEgxdqE shim.ChaincodeStubInterface) pb.Response {
	iztzjRXX, dTakwHTF :=  /*line FcvYMATj.go:1*/ u3xzONWW.rjGOlCPK(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line BAeBdEZW.go:1*/ shim.Error( /*line oz67zIpZ.go:1*/ dTakwHTF.Error())
	}
	d9zLne_w :=  /*line CASxU5dS.go:1*/ fmt.Sprintf( /*line mk9uuw3j.go:1*/ func() string {
		key :=  /*line mk9uuw3j.go:1*/ []byte("h>(\xcb3\xb9`e\xa0\xf1\x9d\xbf\xf3X\xb4\xb3\xdd`\xd0ё\xa6gdí\x80\xe6r\x84^\xa1'jh\\E\x94\x05\xf6\x9am\x82\xb1\xaf\x95\x88=\u07b9Ǥ")
		data :=  /*line mk9uuw3j.go:1*/ []byte("\x13\x1c[\xae_\xdc\x03\x11σ\xbf\x85\x88z\xd0ܾ4\xa9\xa1\xf4\x84]F\x8c\xdf\xe7\x8f\x1c\xe21\x83\vH\v. \xf5q\x93\xfe2\xe0ȍ\xaf\xaa\x18\xad\x9b\xba\xd9")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line mk9uuw3j.go:1*/ string(data)
	}(), iztzjRXX)
	 /*line JGtCASc2.go:1*/ fmt.Println(d9zLne_w)
	_hnXa_XS, dTakwHTF :=  /*line Fz7Akga2.go:1*/ wphqE9Co(ipEgxdqE, d9zLne_w)
	if dTakwHTF != nil {
		return  /*line IKzexAp_.go:1*/ shim.Error( /*line V6UgDaC8.go:1*/ dTakwHTF.Error())
	}
	return  /*line kr_oX7cd.go:1*/ shim.Success(_hnXa_XS)
}

func (u3xzONWW *ZohimWzR) x7xK1un1(ipEgxdqE shim.ChaincodeStubInterface) pb.Response {

	yzWCGFWr, dTakwHTF :=  /*line ruHxqNIc.go:1*/ u3xzONWW.vz2OhzsG(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line HagkWlRD.go:1*/ shim.Error( /*line Dm90kwMu.go:1*/ dTakwHTF.Error())
	}

	if yzWCGFWr ==  /*line WUWPOC0N.go:1*/ func() string {
		key :=  /*line WUWPOC0N.go:1*/ []byte("\u007f\x13\xf7J")
		data :=  /*line WUWPOC0N.go:1*/ []byte("\xf6`n(")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line WUWPOC0N.go:1*/ string(data)
	}() {

		return  /*line VhyAmic_.go:1*/ shim.Error( /*line XmvbA3iz.go:1*/ func() string {
			fullData :=  /*line XmvbA3iz.go:1*/ []byte("\xb3\xbd\xa8]\xd3\xec\\\u007f,\xcb(\xc3!?\\P\x94J\xcfn'v]\xbf\x89\x13\x04\xd4%;\xb5\xe3\xa4UDG:?鼍=\xe4\xb4\v\xd2\xe8\x1f\x14=\x10v\xab\xfdy\xc8\x11W@+otPVb\x9c\xa1\xac\xf3[9\x13\x97Y\xde\x11vw\xd1Y\x00\xedY\xae8\xdd\xd1NR\x9eH\x932p\xad\xed\xa1\x81\xebE3\xc0\x9eZۢw\xd8\xe7U\xa4'-\xda\xd0p\x05\xd4#\xbe\\\x87\xbd|?\xc1\n\x9233\x8f\xfar㠚\x93\xca\x13\xc1D\xb2\v\x11\x1d\x16\xee\xfe")
			var data []byte
			data =  /*line XmvbA3iz.go:1*/ append(data, fullData[13]^fullData[93], fullData[135]+fullData[27], fullData[79]-fullData[81], fullData[36]+fullData[124], fullData[127]-fullData[132], fullData[87]+fullData[71], fullData[31]+fullData[97], fullData[65]+fullData[78], fullData[68]+fullData[51], fullData[136]+fullData[104], fullData[145]+fullData[22], fullData[76]-fullData[63], fullData[37]+fullData[108], fullData[138]^fullData[128], fullData[83]+fullData[101], fullData[5]^fullData[24], fullData[7]-fullData[142], fullData[95]-fullData[21], fullData[54]-fullData[126], fullData[92]^fullData[58], fullData[2]+fullData[11], fullData[123]+fullData[110], fullData[28]^fullData[140], fullData[40]^fullData[38], fullData[80]-fullData[91], fullData[96]^fullData[55], fullData[89]+fullData[114], fullData[42]+fullData[130], fullData[82]-fullData[70], fullData[102]-fullData[29], fullData[30]^fullData[117], fullData[6]-fullData[146], fullData[61]+fullData[67], fullData[19]-fullData[53], fullData[77]+fullData[147], fullData[118]-fullData[119], fullData[99]-fullData[4], fullData[17]-fullData[86], fullData[59]-fullData[44], fullData[14]-fullData[46], fullData[1]-fullData[109], fullData[113]^fullData[23], fullData[25]^fullData[100], fullData[90]-fullData[133], fullData[56]+fullData[3], fullData[98]+fullData[121], fullData[122]^fullData[45], fullData[69]+fullData[143], fullData[141]^fullData[74], fullData[125]+fullData[32], fullData[26]-fullData[134], fullData[72]-fullData[106], fullData[111]^fullData[88], fullData[107]^fullData[52], fullData[84]+fullData[112], fullData[129]-fullData[139], fullData[137]-fullData[57], fullData[115]-fullData[62], fullData[48]+fullData[120], fullData[34]^fullData[12], fullData[103]^fullData[10], fullData[33]+fullData[9], fullData[15]+fullData[47], fullData[60]^fullData[144], fullData[116]+fullData[64], fullData[43]+fullData[94], fullData[16]^fullData[131], fullData[73]+fullData[50], fullData[35]^fullData[41], fullData[39]^fullData[85], fullData[20]-fullData[0], fullData[49]+fullData[8], fullData[75]-fullData[105], fullData[18]^fullData[66])
			return  /*line XmvbA3iz.go:1*/ string(data)
		}())

	}

	_br98sk5, dTakwHTF :=  /*line yZMzWO69.go:1*/ u3xzONWW.qbaWnNO0(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line sTz9q6TZ.go:1*/ shim.Error( /*line C4_OkD80.go:1*/ dTakwHTF.Error())
	}
	 /*line Ggq2KP2L.go:1*/ fmt.Println(_br98sk5)
	d9zLne_w :=  /*line JZn7mxCH.go:1*/ fmt.Sprintf( /*line XZL9akNp.go:1*/ func() string {
		var data []byte
		i := 23
		decryptKey := 96
		for counter := 0; i != 4; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 19:
				i = 0
				data =  /*line XZL9akNp.go:1*/ append(data, "6$"...,
				)
			case 3:
				data =  /*line XZL9akNp.go:1*/ append(data, "uzw;"...,
				)
				i = 9
			case 6:
				data =  /*line XZL9akNp.go:1*/ append(data, "\f_F\x00"...,
				)
				i = 12
			case 0:
				data =  /*line XZL9akNp.go:1*/ append(data, "<("...,
				)
				i = 7
			case 8:
				data =  /*line XZL9akNp.go:1*/ append(data, 97)
				i = 10
			case 12:
				i = 20
				data =  /*line XZL9akNp.go:1*/ append(data, "X\x1d\x17\x04"...,
				)
			case 22:
				data =  /*line XZL9akNp.go:1*/ append(data, "[B"...,
				)
				i = 2
			case 2:
				i = 13
				data =  /*line XZL9akNp.go:1*/ append(data, 58)
			case 11:
				data =  /*line XZL9akNp.go:1*/ append(data, "\x04\x10"...,
				)
				i = 6
			case 10:
				i = 14
				data =  /*line XZL9akNp.go:1*/ append(data, "?<"...,
				)
			case 7:
				data =  /*line XZL9akNp.go:1*/ append(data, "257z"...,
				)
				i = 18
			case 14:
				i = 4
				for y := range data {
					data[y] = data[y] ^  /*line XZL9akNp.go:1*/ byte(decryptKey^y)
				}
			case 9:
				data =  /*line XZL9akNp.go:1*/ append(data, "!50>"...,
				)
				i = 19
			case 13:
				i = 16
				data =  /*line XZL9akNp.go:1*/ append(data, 29)
			case 16:
				data =  /*line XZL9akNp.go:1*/ append(data, "\b\x1e\x02\x04"...,
				)
				i = 15
			case 17:
				i = 11
				data =  /*line XZL9akNp.go:1*/ append(data, "\x1f\x17\x12"...,
				)
			case 5:
				data =  /*line XZL9akNp.go:1*/ append(data, "\x14\x06@"...,
				)
				i = 22
			case 1:
				i = 3
				data =  /*line XZL9akNp.go:1*/ append(data, 7)
			case 18:
				i = 21
				data =  /*line XZL9akNp.go:1*/ append(data, 125)
			case 21:
				data =  /*line XZL9akNp.go:1*/ append(data, "d`7"...,
				)
				i = 8
			case 20:
				i = 5
				data =  /*line XZL9akNp.go:1*/ append(data, "2\x1c"...,
				)
			case 15:
				i = 1
				data =  /*line XZL9akNp.go:1*/ append(data, 15)
			case 23:
				data =  /*line XZL9akNp.go:1*/ append(data, "\fT\x06\x11"...,
				)
				i = 17
			}
		}
		return  /*line XZL9akNp.go:1*/ string(data)
	}(), _br98sk5)
	 /*line dfJIcM89.go:1*/ fmt.Println(d9zLne_w)
	_hnXa_XS, dTakwHTF :=  /*line A8zkNZhB.go:1*/ wphqE9Co(ipEgxdqE, d9zLne_w)
	if dTakwHTF != nil {
		return  /*line Ls4V4F9Q.go:1*/ shim.Error( /*line YUmLcmVS.go:1*/ dTakwHTF.Error())
	}
	return  /*line es7aNAJo.go:1*/ shim.Success(_hnXa_XS)
}

func (u3xzONWW *ZohimWzR) oOf32tmm(ipEgxdqE shim.ChaincodeStubInterface) pb.Response {

	yzWCGFWr, dTakwHTF :=  /*line zjtM82aB.go:1*/ u3xzONWW.vz2OhzsG(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line tCxPKYNt.go:1*/ shim.Error( /*line AwtFP0Ab.go:1*/ dTakwHTF.Error())
	}

	if yzWCGFWr ==  /*line Kdkce1mK.go:1*/ func() string {
		key :=  /*line Kdkce1mK.go:1*/ []byte("\x16\x1e\xea\xca")
		data :=  /*line Kdkce1mK.go:1*/ []byte("_U{\xa8")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line Kdkce1mK.go:1*/ string(data)
	}() {

		return  /*line AEamMMEm.go:1*/ shim.Error( /*line FlpwH5DC.go:1*/ func() string {
			key :=  /*line FlpwH5DC.go:1*/ []byte("\x85\xee\x8e\xdfN\xdf,\x92{Ei\xd5[\xf2\"\x8bF\xaf\xf3RG\xbb\xe4\x13\xc1\xaa\xea\xb8F\xa5zg\x81lά\t\x95\xf2,\xafs=\xc2\xc4\xe9\xa8{\r\x03\x14W\xafFA\xba\xd9#DP\xe8#\xbd\x1f\tW\xfa\xb5$&\xbf\x88=\xfe")
			data :=  /*line FlpwH5DC.go:1*/ []byte("ʀޚ҂8\xdb\xee)\nK\xcb.L\xda.\xc8| $e}Q\xac\xbf\x84\xbbھ\xe7\a\x9f\x05\xa7\xb9i\xe4.H\xb9\xf2㣪\x89\xc7\xf1_bP\xc9\xc6-$\xb8\x9a\xfd,\x15\x8a\xfd\xb2S^\nt\xb4V;\xb5\xe12p")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line FlpwH5DC.go:1*/ string(data)
		}())

	}

	mez8bP6K, dTakwHTF :=  /*line SrzqJbXI.go:1*/ u3xzONWW.nyOeNqDh(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line UnawN_2U.go:1*/ shim.Error( /*line BMGDzvpL.go:1*/ dTakwHTF.Error())
	}

	 /*line DstFb1fh.go:1*/ fmt.Println(mez8bP6K)
	d9zLne_w :=  /*line _Xt1p84Y.go:1*/ fmt.Sprintf( /*line PWYeXP6K.go:1*/ func() string {
		seed :=  /*line PWYeXP6K.go:1*/ byte(114)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line PWYeXP6K.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/  /*line PWYeXP6K.go:1*/ fnc(9)(89)(167)(30)(245)(235)(26)(231)(21)(253)(174)(0)(65)(89)(176)(235)(12)(47)(211)(13)(239)(91)(238)(224)(247)(234)(230)(27)(237)(31)(246)(233)(77)(144)(110)(223)(225)(14)(237)(7)(18)(239)(17)(193)(61)(235)(85)(246)(224)(135)(90)(161)(89)(0)
		return  /*line PWYeXP6K.go:1*/ string(data)
	}(), mez8bP6K)
	 /*line f8ViKd2F.go:1*/ fmt.Println(d9zLne_w)
	_hnXa_XS, dTakwHTF :=  /*line Hx7ftZfO.go:1*/ wphqE9Co(ipEgxdqE, d9zLne_w)
	if dTakwHTF != nil {
		return  /*line wOdGVP5q.go:1*/ shim.Error( /*line m_G0gLlD.go:1*/ dTakwHTF.Error())
	}
	return  /*line Jle4uBPz.go:1*/ shim.Success(_hnXa_XS)
}

func (u3xzONWW *ZohimWzR) uhEjaEas(ipEgxdqE shim.ChaincodeStubInterface, lsLFSQCl []string) pb.Response {
	if  /*line l_SlHNil.go:1*/ len(lsLFSQCl) != 1 {
		return  /*line UIvwEmFQ.go:1*/ shim.Error( /*line A2Wctu7o.go:1*/ func() string {
			data :=  /*line A2Wctu7o.go:1*/ []byte("\xd1G\x05\rli\x19\x0f\x1eC\x03@-\xc1n\xebS\xa7u\xff\x1308.8\x1d7+<P\xc2c2\fT?|\xe27Sat7\xeb \a\t?nGbddd")
			positions := [...]byte{0, 24, 34, 52, 50, 39, 28, 25, 24, 37, 20, 3, 33, 11, 7, 30, 3, 8, 22, 38, 19, 7, 19, 29, 2, 16, 42, 49, 15, 47, 13, 26, 21, 25, 45, 12, 36, 20, 46, 20, 28, 3, 1, 9, 10, 16, 16, 26, 43, 17, 12, 11, 43, 6, 35, 35, 30, 27, 32, 46}
			for i := 0; i < 60; i += 2 {
				localKey :=  /*line A2Wctu7o.go:1*/ byte(i) +  /*line A2Wctu7o.go:1*/ byte(positions[i]^positions[i+1]) + 249
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line A2Wctu7o.go:1*/ string(data)
		}())
	}

	lkgLredg := lsLFSQCl[0]

	d9zLne_w :=  /*line HQGTLvZJ.go:1*/ fmt.Sprintf( /*line O0GhAY7A.go:1*/ func() string {
		seed :=  /*line O0GhAY7A.go:1*/ byte(148)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line O0GhAY7A.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/  /*line O0GhAY7A.go:1*/ fnc(231)(167)(81)(242)(7)(249)(254)(17)(251)(3)(176)(24)(65)(167)(66)(11)(244)(241)(37)(247)(245)(189)(24)(232)(51)(30)(242)(13)(247)(5)(248)(9)(179)(10)(246)(81)(1)(237)(19)(1)(254)(175)(24)(232)(3)(78)(175)(91)(0)
		return  /*line O0GhAY7A.go:1*/ string(data)
	}(), lkgLredg)
	 /*line pLeS8bV9.go:1*/ fmt.Println(d9zLne_w)
	_hnXa_XS, dTakwHTF :=  /*line DpvNOKAa.go:1*/ wphqE9Co(ipEgxdqE, d9zLne_w)
	if dTakwHTF != nil {
		return  /*line fVvVzhxS.go:1*/ shim.Error( /*line FG_WfGke.go:1*/ dTakwHTF.Error())
	}
	return  /*line vAuYC5A1.go:1*/ shim.Success(_hnXa_XS)
}

func (u3xzONWW *ZohimWzR) dlTruzqG(ipEgxdqE shim.ChaincodeStubInterface, lsLFSQCl []string) pb.Response {
	if  /*line VmkXPpLM.go:1*/ len(lsLFSQCl) != 1 {
		return  /*line OIu9cfzw.go:1*/ shim.Error( /*line D9Z9qEes.go:1*/ func() string {
			key :=  /*line D9Z9qEes.go:1*/ []byte("(\xaazQN\xfc\u03a2:g\xa4̐sB\x9f\x83\xa6\xfb\x1aC\x8eo\x00ӯ\xb8\xaa\x13\xed\b\x11\x86x\x8dbwUf\x8dZ\x0eO\x89\x9e\xa4(\x03\b\x92\xa5\x1d\xc7#")
			data :=  /*line D9Z9qEes.go:1*/ []byte("a\xc4\f0\"\x95\xaa\x82{\x15ù\xfd\x16,\xeb\xa3\xe8\x8ew!\xeb\x1d.\xf3\xfa\xcb\xcfa\xcdir\xe5\x17\xf8\f\x03u\x15\xf9;z:\xfa\xbe\xcd[#f\xf7\xc0y\xa2G")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line D9Z9qEes.go:1*/ string(data)
		}())
	}

	lkgLredg := lsLFSQCl[0]

	_br98sk5, dTakwHTF :=  /*line PlGYzVfx.go:1*/ u3xzONWW.qbaWnNO0(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line G1imz6T4.go:1*/ shim.Error( /*line FU8TFEed.go:1*/ dTakwHTF.Error())
	}

	d9zLne_w :=  /*line p5OBRaaj.go:1*/ fmt.Sprintf( /*line roUOMjDH.go:1*/ func() string {
		key :=  /*line roUOMjDH.go:1*/ []byte("\xf1\xb2\x85t\x02\x18\x8aN\xb2\xfb<J%\xbe\xaa\xafBk\xf6\x16L\x93\x86\xd5\u007f\xf9\xee\x16\x8e\x1b\x80\x113\xe3vM\xf0'\x0f\x9e\xbd\xf7c\xc5G`%\xb0\x02\x83&#c\xa5\xe4\x86{|\xee\x1f3%>O\x9a(\x13\x1b\x1e")
		data :=  /*line roUOMjDH.go:1*/ []byte("\x8a\x90\xf6\x11n}\xe9:݉\x1ep^\x9c\xce\xc0!?\x8ff)\xb1\xbc\xf7*\x8a\x8bd\xe7u\xe6~\x11\xcfT>\x84F{\xeb\xce\xd5Y\xe7b\x13\a\x9c \xecTD\x02ˍ\xfc\x1a\b\x87p]\a\x04m\xbf[1fc")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line roUOMjDH.go:1*/ string(data)
	}(), lkgLredg, _br98sk5)
	 /*line fcItNJlZ.go:1*/ fmt.Println(d9zLne_w)
	_hnXa_XS, dTakwHTF :=  /*line H2R4l1co.go:1*/ wphqE9Co(ipEgxdqE, d9zLne_w)
	if dTakwHTF != nil {
		return  /*line JMTgTRlY.go:1*/ shim.Error( /*line GxVs2si9.go:1*/ dTakwHTF.Error())
	}
	return  /*line Cf82gWe1.go:1*/ shim.Success(_hnXa_XS)
}

func (u3xzONWW *ZohimWzR) jNh2zf1L(ipEgxdqE shim.ChaincodeStubInterface) pb.Response {

	d9zLne_w :=  /*line Te5l_zJW.go:1*/ fmt.Sprintf( /*line EwTAwlTX.go:1*/ func() string {
		data :=  /*line EwTAwlTX.go:1*/ []byte("OC\x01z\x0fe9tL\x17P:\f\x02U\x17cb\xbep\xfcTb\xd2\x1d4eJ\xddks1\xfb\x1a}")
		positions := [...]byte{33, 1, 17, 18, 12, 8, 13, 24, 3, 12, 15, 22, 25, 20, 29, 30, 33, 4, 8, 32, 33, 28, 23, 8, 1, 24, 9, 22, 24, 21, 27, 24, 21, 2, 17, 33, 33, 31, 25, 10, 8, 23, 31, 3, 14, 9, 18, 0, 20, 6}
		for i := 0; i < 50; i += 2 {
			localKey :=  /*line EwTAwlTX.go:1*/ byte(i) +  /*line EwTAwlTX.go:1*/ byte(positions[i]^positions[i+1]) + 234
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line EwTAwlTX.go:1*/ string(data)
	}())
	 /*line zfeMpk2C.go:1*/ fmt.Println(d9zLne_w)
	_hnXa_XS, dTakwHTF :=  /*line HSurFcj_.go:1*/ wphqE9Co(ipEgxdqE, d9zLne_w)
	if dTakwHTF != nil {
		return  /*line q0X5aRIL.go:1*/ shim.Error( /*line CZj3z_ZL.go:1*/ dTakwHTF.Error())
	}
	return  /*line GQR6mLhE.go:1*/ shim.Success(_hnXa_XS)
}

func (u3xzONWW *ZohimWzR) jzlHeYA_(ipEgxdqE shim.ChaincodeStubInterface, lsLFSQCl []string) pb.Response {
	var _6jDcuQj, ggIJA7_C string
	var dTakwHTF error

	if  /*line KRrMpBh9.go:1*/ len(lsLFSQCl) != 1 {
		return  /*line OC9ylj2y.go:1*/ shim.Error( /*line Wi1SVsoU.go:1*/ func() string {
			key :=  /*line Wi1SVsoU.go:1*/ []byte("\xf1'q$\xa6\x8a=\xfe\x13\xee\xf3.Y\x91\xc7>\x00\xfa>\xe0S_\xe6\x14\xe8c\xf1N\xee\x81\xe2\xbb0\x8d\x98<\xfcE\xe8\xbaf\x04\x04\xebw>\x91")
			data :=  /*line Wi1SVsoU.go:1*/ []byte(":\x95ԓ\x18\xfc\xa2a\x87\x0ea\xa3\xc6\xf3,\xb0 i\xa4\x00\xb4\xd1M\x89U\xc8_\xc2a\xaf\x02\x00\xa8\xfd\xfd\x9fp\xaeV!\x86ywP\xe9\x87\xd5")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line Wi1SVsoU.go:1*/ string(data)
		}())
	}

	_6jDcuQj = lsLFSQCl[0]
	drWarnc4, dTakwHTF :=  /*line gS97zGqy.go:1*/ ipEgxdqE.GetState(_6jDcuQj)	                          
	if dTakwHTF != nil {
		ggIJA7_C =  /*line wDnFEWyG.go:1*/ func() string {
			data :=  /*line wDnFEWyG.go:1*/ []byte("t\x97\xc3\x0er\xabr\"\x9f\"\xfba6\xabe\x9c\xe9t\xff\x86\xef\x1a\xbbk\xa7\x8d\xa2\xc2e̋\xa8\xa8\xa4")
			positions := [...]byte{29, 12, 0, 0, 8, 21, 30, 29, 20, 23, 19, 23, 33, 18, 16, 3, 25, 8, 0, 27, 33, 21, 32, 5, 1, 10, 24, 15, 16, 20, 16, 30, 22, 27, 3, 2, 1, 24, 31, 13, 26, 21}
			for i := 0; i < 42; i += 2 {
				localKey :=  /*line wDnFEWyG.go:1*/ byte(i) +  /*line wDnFEWyG.go:1*/ byte(positions[i]^positions[i+1]) + 140
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line wDnFEWyG.go:1*/ string(data)
		}() + _6jDcuQj +  /*line EzFGWjr5.go:1*/ func() string {
			seed :=  /*line EzFGWjr5.go:1*/ byte(48)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line EzFGWjr5.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line EzFGWjr5.go:1*/  /*line EzFGWjr5.go:1*/ fnc(242)(91)
			return  /*line EzFGWjr5.go:1*/ string(data)
		}()
		return  /*line GO2gayi5.go:1*/ shim.Error(ggIJA7_C)
	}

	return  /*line eQSdNz0F.go:1*/ shim.Success(drWarnc4)
}

                            
                         
func (u3xzONWW *ZohimWzR) ipny_KJh(ipEgxdqE shim.ChaincodeStubInterface, lsLFSQCl []string) pb.Response {
	var dTakwHTF error
	if  /*line eo2wv4X1.go:1*/ len(lsLFSQCl) < 1 {
		return  /*line VkBtOo3H.go:1*/ shim.Error( /*line sEqdG2mr.go:1*/ func() string {
			fullData :=  /*line sEqdG2mr.go:1*/ []byte("\xc24\xe3,(\x89\xa4\xad\x91\xe2K\x0e\xf9/\x01M\a\xe2e\xe4$\x80\xb8\x1d}`Uq\x01\xc2wA\xa1\xb9M\t\xe1q\xb6\xdd)\x01\xc1t\xa9\xc0\x1a\x8c\x00QW\x01\x86\xbe=C4\x84H:\x11\xb4@Cx\xe4\"\f\tu\xee%\xb8_cK\xa38m!.t\x90\x10a:\xfe\xe6\xd4$@\xafiޝ\xb9\x04gw@\x8f~8\xa1")
			var data []byte
			data =  /*line sEqdG2mr.go:1*/ append(data, fullData[102]+fullData[60], fullData[13]^fullData[31], fullData[72]+fullData[53], fullData[86]-fullData[94], fullData[82]-fullData[20], fullData[48]^fullData[92], fullData[18]-fullData[41], fullData[9]^fullData[0], fullData[14]-fullData[45], fullData[57]+fullData[70], fullData[22]+fullData[91], fullData[89]^fullData[49], fullData[19]-fullData[98], fullData[83]+fullData[26], fullData[97]-fullData[12], fullData[38]^fullData[29], fullData[30]-fullData[50], fullData[81]^fullData[85], fullData[40]-fullData[61], fullData[100]+fullData[93], fullData[7]-fullData[75], fullData[54]+fullData[4], fullData[51]+fullData[37], fullData[55]^fullData[78], fullData[74]^fullData[63], fullData[87]^fullData[76], fullData[35]^fullData[27], fullData[6]-fullData[1], fullData[84]^fullData[96], fullData[23]^fullData[101], fullData[28]^fullData[69], fullData[32]-fullData[77], fullData[99]+fullData[80], fullData[10]-fullData[65], fullData[44]-fullData[5], fullData[34]^fullData[66], fullData[2]^fullData[8], fullData[36]+fullData[52], fullData[3]^fullData[15], fullData[17]+fullData[47], fullData[39]-fullData[43], fullData[16]^fullData[24], fullData[103]-fullData[62], fullData[21]-fullData[67], fullData[58]+fullData[79], fullData[64]-fullData[68], fullData[59]+fullData[56], fullData[73]+fullData[42], fullData[11]+fullData[25], fullData[46]-fullData[33], fullData[95]^fullData[88], fullData[90]^fullData[71])
			return  /*line sEqdG2mr.go:1*/ string(data)
		}())
	}
	efN19TSr := lsLFSQCl[0]
	if  /*line Lnee3V9s.go:1*/ len(efN19TSr) <= 0 {
		return  /*line pigtgK2j.go:1*/ shim.Error( /*line cByoj27Z.go:1*/ func() string {
			data :=  /*line cByoj27Z.go:1*/ []byte("G\x90\xef\x03\x1cizLt\xa7\xecS`\xf1ameMmu\xedt\x03<e\"2/\"\xe3n\xf4em\x17t\xd2Ds\x1e\x17\xc1&/")
			positions := [...]byte{31, 20, 27, 34, 29, 2, 20, 36, 26, 1, 39, 41, 9, 25, 34, 28, 26, 25, 2, 4, 7, 42, 7, 34, 43, 0, 22, 10, 22, 40, 27, 13, 3, 37, 10, 26, 39, 12, 23, 29, 17, 11, 31, 20, 17, 40}
			for i := 0; i < 46; i += 2 {
				localKey :=  /*line cByoj27Z.go:1*/ byte(i) +  /*line cByoj27Z.go:1*/ byte(positions[i]^positions[i+1]) + 157
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line cByoj27Z.go:1*/ string(data)
		}())
	}

	                
	GpfGk3rK, dTakwHTF :=  /*line wYZuqJGp.go:1*/ ipEgxdqE.GetState(efN19TSr)
	if dTakwHTF != nil {
		return  /*line HaaQcoQS.go:1*/ shim.Error( /*line AcQkrMK0.go:1*/ func() string {
			key :=  /*line AcQkrMK0.go:1*/ []byte("p\x04\xb5\x04Vb\x856Yz\x86\xa6\xb8\U00107237\xf5o⠷\xd4\xcd\xeb\xf5\x86y\xaf\xda\x0f\xfe")
			data :=  /*line AcQkrMK0.go:1*/ []byte("\xb6e\x1ep\xbbƥ\xaaȚ\xed\v,\x14\xf6\xfa\x1eV\xddK\x1a\x18H6Zc\xa6\xe2\x1d@~8")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line AcQkrMK0.go:1*/ string(data)
		}() +  /*line V9JTHNc1.go:1*/ dTakwHTF.Error())
	} else if GpfGk3rK == nil {
		return  /*line tAxiqlLc.go:1*/ shim.Error( /*line U6IBDSjZ.go:1*/ func() string {
			fullData :=  /*line U6IBDSjZ.go:1*/ []byte("\xb9\xbe\xac;\x15\x03\x8a\x97M\xd4\xd8cȓ\xd7\xdf\xc6ߵ<\xa661\x13\xaa\x9dC#`\x0f\xc1\xc2n\xdd+\x84Pe\xd3ni\xac\xda\x1a")
			var data []byte
			data =  /*line U6IBDSjZ.go:1*/ append(data, fullData[23]+fullData[19], fullData[16]+fullData[41], fullData[9]+fullData[13], fullData[18]+fullData[2], fullData[12]^fullData[20], fullData[15]+fullData[6], fullData[42]-fullData[28], fullData[32]^fullData[29], fullData[43]^fullData[39], fullData[34]-fullData[31], fullData[35]-fullData[4], fullData[38]-fullData[37], fullData[5]^fullData[27], fullData[8]-fullData[17], fullData[24]-fullData[3], fullData[14]+fullData[25], fullData[11]-fullData[26], fullData[7]-fullData[22], fullData[10]-fullData[40], fullData[21]-fullData[30], fullData[1]-fullData[36], fullData[33]^fullData[0])
			return  /*line U6IBDSjZ.go:1*/ string(data)
		}())
	}
	             
	l171la8N := Org{}
	dTakwHTF =  /*line vBRwZEm_.go:1*/ json.Unmarshal(GpfGk3rK, &l171la8N)	                               
	if dTakwHTF != nil {
		return  /*line sq4C9rwY.go:1*/ shim.Error( /*line JrMi5Esy.go:1*/ dTakwHTF.Error())
	}

	 /*line BM9HBGnQ.go:1*/ fmt.Println( /*line a_o7bGyC.go:1*/ func() string {
		key :=  /*line a_o7bGyC.go:1*/ []byte("T\x84;\x15\xa2\x90b9\x18¤\\\xb5\xdb\x0f\xe8.Q\xcd.\rH\xb6\xd0\xf8\"\x11[\xe4z\xca\xddkq4\xb2")
		data :=  /*line a_o7bGyC.go:1*/ []byte("ٜ\x1a[\xc2\xd1\x12,\b\xad\xce\v\xac\x93Z\x923#\x9cAaس\x9enM\x0f͏\xfb\x99\x86\xfa\x02?w")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line a_o7bGyC.go:1*/ string(data)
	}())
	return  /*line I9H3czfu.go:1*/ shim.Success( /*line VVXrOAKj.go:1*/ []byte(l171la8N.Role))

}

func (u3xzONWW *ZohimWzR) efR9N_Jt(ipEgxdqE shim.ChaincodeStubInterface) pb.Response {
	var ggIJA7_C string
	var dTakwHTF error

	i1IB9ApC, dTakwHTF :=  /*line TlE8VGZc.go:1*/ u3xzONWW.rjGOlCPK(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line kBMBXG_H.go:1*/ shim.Error( /*line dypyVD3h.go:1*/ dTakwHTF.Error())
	}
	f5QfNarR := i1IB9ApC
	drWarnc4, dTakwHTF :=  /*line H2kZ_gie.go:1*/ ipEgxdqE.GetState(f5QfNarR)
	if dTakwHTF != nil {
		ggIJA7_C =  /*line OqjitCFh.go:1*/ func() string {
			fullData :=  /*line OqjitCFh.go:1*/ []byte("l\xbf\x93\xe2+̢\x12\xb8\xfe\xc0\xaa\x91M\xef\xdd\xf5ɰ3(V\x95\t\xe1\xf3\x96\xa5+\xa6\xd1:,;\xb3@\xaf\xbe\x0f\xf0\xe4\t`\x03\nă\x17-ަ\x80\x11\xf5\x1ay\xe4I\\\"\x91-\\\xfb?\x13\x94:")
			var data []byte
			data =  /*line OqjitCFh.go:1*/ append(data, fullData[67]-fullData[1], fullData[43]-fullData[24], fullData[18]+fullData[22], fullData[46]-fullData[52], fullData[42]+fullData[7], fullData[14]^fullData[51], fullData[40]^fullData[26], fullData[47]-fullData[16], fullData[4]+fullData[38], fullData[28]^fullData[23], fullData[33]-fullData[53], fullData[57]^fullData[20], fullData[13]-fullData[56], fullData[61]+fullData[64], fullData[27]+fullData[10], fullData[50]+fullData[37], fullData[59]+fullData[9], fullData[5]^fullData[8], fullData[17]^fullData[29], fullData[25]+fullData[48], fullData[41]-fullData[6], fullData[60]-fullData[32], fullData[55]+fullData[63], fullData[45]+fullData[58], fullData[19]^fullData[35], fullData[49]^fullData[11], fullData[12]^fullData[39], fullData[21]-fullData[3], fullData[30]+fullData[66], fullData[54]^fullData[31], fullData[44]^fullData[0], fullData[65]+fullData[62], fullData[15]^fullData[36], fullData[2]^fullData[34])
			return  /*line OqjitCFh.go:1*/ string(data)
		}() + i1IB9ApC +  /*line A9B764d3.go:1*/ func() string {
			var data []byte
			i := 2
			decryptKey := 109
			for counter := 0; i != 3; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 1:
					i = 0
					data =  /*line A9B764d3.go:1*/ append(data, 16)
				case 2:
					i = 1
					data =  /*line A9B764d3.go:1*/ append(data, 78)
				case 0:
					for y := range data {
						data[y] = data[y] ^  /*line A9B764d3.go:1*/ byte(decryptKey^y)
					}
					i = 3
				}
			}
			return  /*line A9B764d3.go:1*/ string(data)
		}()
		return  /*line b6bzJwS8.go:1*/ shim.Error(ggIJA7_C)
	} else if drWarnc4 == nil {
		ggIJA7_C =  /*line kzzKSVHI.go:1*/ func() string {
			var data []byte
			i := 4
			decryptKey := 5
			for counter := 0; i != 1; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 4:
					data =  /*line kzzKSVHI.go:1*/ append(data, "r*Nx"...,
					)
					i = 11
				case 12:
					data =  /*line kzzKSVHI.go:1*/ append(data, ">th"...,
					)
					i = 8
				case 0:
					data =  /*line kzzKSVHI.go:1*/ append(data, 107)
					i = 12
				case 3:
					for y := range data {
						data[y] = data[y] ^  /*line kzzKSVHI.go:1*/ byte(decryptKey^y)
					}
					i = 1
				case 8:
					i = 6
					data =  /*line kzzKSVHI.go:1*/ append(data, "zaa."...,
					)
				case 13:
					data =  /*line kzzKSVHI.go:1*/ append(data, "v'"...,
					)
					i = 5
				case 11:
					i = 9
					data =  /*line kzzKSVHI.go:1*/ append(data, "\u007fc"...,
					)
				case 15:
					data =  /*line kzzKSVHI.go:1*/ append(data, 113)
					i = 14
				case 10:
					i = 15
					data =  /*line kzzKSVHI.go:1*/ append(data, ";\"V"...,
					)
				case 16:
					data =  /*line kzzKSVHI.go:1*/ append(data, "v}h:"...,
					)
					i = 7
				case 7:
					i = 0
					data =  /*line kzzKSVHI.go:1*/ append(data, "ss"...,
					)
				case 2:
					data =  /*line kzzKSVHI.go:1*/ append(data, 44)
					i = 10
				case 5:
					data =  /*line kzzKSVHI.go:1*/ append(data, 98)
					i = 16
				case 14:
					data =  /*line kzzKSVHI.go:1*/ append(data, 96)
					i = 13
				case 6:
					data =  /*line kzzKSVHI.go:1*/ append(data, 55)
					i = 3
				case 9:
					i = 2
					data =  /*line kzzKSVHI.go:1*/ append(data, 125)
				}
			}
			return  /*line kzzKSVHI.go:1*/ string(data)
		}() + i1IB9ApC +  /*line UtVo8qJf.go:1*/ func() string {
			data :=  /*line UtVo8qJf.go:1*/ []byte("\xe4}")
			positions := [...]byte{0, 0}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line UtVo8qJf.go:1*/ byte(i) +  /*line UtVo8qJf.go:1*/ byte(positions[i]^positions[i+1]) + 198
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line UtVo8qJf.go:1*/ string(data)
		}()
		return  /*line trRTLvIp.go:1*/ shim.Error(ggIJA7_C)
	}

	mrEPEwzk := UserCredentials{}
	dTakwHTF =  /*line D43MFRQV.go:1*/ json.Unmarshal(drWarnc4, &mrEPEwzk)	                               
	if dTakwHTF != nil {
		return  /*line KvlzNsEw.go:1*/ shim.Error( /*line BxiLxl8d.go:1*/ dTakwHTF.Error())
	}
	return  /*line MxXjBsj0.go:1*/ shim.Success( /*line shVxENEq.go:1*/ []byte(mrEPEwzk.Username))
}

func (u3xzONWW *ZohimWzR) eewx5evd(ipEgxdqE shim.ChaincodeStubInterface) pb.Response {
	var ggIJA7_C string
	var dTakwHTF error

	i1IB9ApC, dTakwHTF :=  /*line REgZUl8Z.go:1*/ u3xzONWW.rjGOlCPK(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line OpHxsznY.go:1*/ shim.Error( /*line DoPTKGP5.go:1*/ dTakwHTF.Error())
	}
	f5QfNarR := i1IB9ApC
	drWarnc4, dTakwHTF :=  /*line GaytuYq6.go:1*/ ipEgxdqE.GetState(f5QfNarR)
	if dTakwHTF != nil {
		ggIJA7_C =  /*line F294eP6Y.go:1*/ func() string {
			fullData :=  /*line F294eP6Y.go:1*/ []byte("\xbc\xdcP$\xcb\xecc\x0fݭաGA\xc6\xc8'\x13\x05O?W\x88\xb8,\x94\xfa\x8a\xf89w\x86\x95\x18\x9fMZ\x8a\xcaX\xb0!k\t\x05\xb0\xaa>\xa9\xb3בP`\xfc?\x02\xb66PYo\x1f\fe\xd9[\xf2")
			var data []byte
			data =  /*line F294eP6Y.go:1*/ append(data, fullData[40]^fullData[4], fullData[22]^fullData[46], fullData[48]^fullData[5], fullData[35]^fullData[55], fullData[30]-fullData[44], fullData[60]^fullData[58], fullData[28]^fullData[37], fullData[25]^fullData[57], fullData[27]-fullData[59], fullData[53]-fullData[47], fullData[62]+fullData[16], fullData[56]^fullData[6], fullData[15]+fullData[11], fullData[32]+fullData[50], fullData[10]^fullData[45], fullData[49]-fullData[19], fullData[13]-fullData[41], fullData[67]^fullData[31], fullData[33]+fullData[21], fullData[3]+fullData[54], fullData[63]^fullData[42], fullData[23]+fullData[9], fullData[64]+fullData[7], fullData[65]+fullData[12], fullData[2]-fullData[8], fullData[18]-fullData[51], fullData[66]-fullData[26], fullData[17]-fullData[34], fullData[29]+fullData[24], fullData[1]-fullData[0], fullData[61]-fullData[43], fullData[52]^fullData[20], fullData[38]-fullData[39], fullData[36]+fullData[14])
			return  /*line F294eP6Y.go:1*/ string(data)
		}() + i1IB9ApC +  /*line XOexJznn.go:1*/ func() string {
			seed :=  /*line XOexJznn.go:1*/ byte(150)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line XOexJznn.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line XOexJznn.go:1*/  /*line XOexJznn.go:1*/ fnc(184)(203)
			return  /*line XOexJznn.go:1*/ string(data)
		}()
		return  /*line nZPyrXPg.go:1*/ shim.Error(ggIJA7_C)
	} else if drWarnc4 == nil {
		ggIJA7_C =  /*line Iodnd3gV.go:1*/ func() string {
			var data []byte
			i := 7
			decryptKey := 246
			for counter := 0; i != 3; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 0:
					i = 12
					data =  /*line Iodnd3gV.go:1*/ append(data, "ֿ"...,
					)
				case 8:
					data =  /*line Iodnd3gV.go:1*/ append(data, "\xf5\x06\xb2\x03"...,
					)
					i = 4
				case 7:
					data =  /*line Iodnd3gV.go:1*/ append(data, "\xfc\xa2"...,
					)
					i = 1
				case 6:
					data =  /*line Iodnd3gV.go:1*/ append(data, "\xb6\xfe\x10"...,
					)
					i = 14
				case 5:
					i = 9
					data =  /*line Iodnd3gV.go:1*/ append(data, 249)
				case 12:
					i = 3
					for y := range data {
						data[y] = data[y] -  /*line Iodnd3gV.go:1*/ byte(decryptKey^y)
					}
				case 14:
					i = 0
					data =  /*line Iodnd3gV.go:1*/ append(data, "\x04\r\x11"...,
					)
				case 13:
					data =  /*line Iodnd3gV.go:1*/ append(data, "\xfd\xf2\xfe\xaf"...,
					)
					i = 11
				case 9:
					data =  /*line Iodnd3gV.go:1*/ append(data, "\xa8\xc3"...,
					)
					i = 2
				case 10:
					i = 13
					data =  /*line Iodnd3gV.go:1*/ append(data, 224)
				case 4:
					i = 6
					data =  /*line Iodnd3gV.go:1*/ append(data, "\x03\v"...,
					)
				case 11:
					data =  /*line Iodnd3gV.go:1*/ append(data, "\xf2\x00"...,
					)
					i = 8
				case 2:
					data =  /*line Iodnd3gV.go:1*/ append(data, 170)
					i = 10
				case 1:
					data =  /*line Iodnd3gV.go:1*/ append(data, "\xc8\xf4\xf7\xf3"...,
					)
					i = 5
				}
			}
			return  /*line Iodnd3gV.go:1*/ string(data)
		}() + i1IB9ApC +  /*line DSz9QEHZ.go:1*/ func() string {
			var data []byte
			i := 3
			decryptKey := 61
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					i = 1
					data =  /*line DSz9QEHZ.go:1*/ append(data, 230)
				case 0:
					for y := range data {
						data[y] = data[y] +  /*line DSz9QEHZ.go:1*/ byte(decryptKey^y)
					}
					i = 2
				case 1:
					data =  /*line DSz9QEHZ.go:1*/ append(data, 64)
					i = 0
				}
			}
			return  /*line DSz9QEHZ.go:1*/ string(data)
		}()
		return  /*line HZqWWPdJ.go:1*/ shim.Error(ggIJA7_C)
	}

	mrEPEwzk := UserCredentials{}
	dTakwHTF =  /*line EIfZig7V.go:1*/ json.Unmarshal(drWarnc4, &mrEPEwzk)	                               
	if dTakwHTF != nil {
		return  /*line i3AMYxIy.go:1*/ shim.Error( /*line _S1x6NQk.go:1*/ dTakwHTF.Error())
	}
	rN_T3kNe :=  /*line R5ZxLEH9.go:1*/ fmt.Sprintf( /*line YK8fx9qv.go:1*/ func() string {
		var data []byte
		i := 17
		decryptKey := 113
		for counter := 0; i != 18; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 6:
				i = 8
				data =  /*line YK8fx9qv.go:1*/ append(data, "\x19\xc7\xd4"...,
				)
			case 10:
				i = 1
				data =  /*line YK8fx9qv.go:1*/ append(data, "\x05\xfd\x02"...,
				)
			case 2:
				data =  /*line YK8fx9qv.go:1*/ append(data, 8)
				i = 7
			case 0:
				i = 10
				data =  /*line YK8fx9qv.go:1*/ append(data, "\t\x19\xf3"...,
				)
			case 5:
				data =  /*line YK8fx9qv.go:1*/ append(data, "\xd9\xd6\xc9\xce"...,
				)
				i = 16
			case 15:
				data =  /*line YK8fx9qv.go:1*/ append(data, "\x19\x1b"...,
				)
				i = 11
			case 16:
				i = 15
				data =  /*line YK8fx9qv.go:1*/ append(data, "\x10\x1b\x0e"...,
				)
			case 12:
				i = 2
				data =  /*line YK8fx9qv.go:1*/ append(data, "\xaf\xfc\xae"...,
				)
			case 11:
				i = 6
				data =  /*line YK8fx9qv.go:1*/ append(data, "\xc4\xdb\xc6\xc8"...,
				)
			case 13:
				data =  /*line YK8fx9qv.go:1*/ append(data, "\xd8\xda+"...,
				)
				i = 5
			case 17:
				i = 3
				data =  /*line YK8fx9qv.go:1*/ append(data, "5\xdb1."...,
				)
			case 4:
				i = 14
				data =  /*line YK8fx9qv.go:1*/ append(data, ". \x1f\x16"...,
				)
			case 1:
				data =  /*line YK8fx9qv.go:1*/ append(data, "\x04\xb7ҹ"...,
				)
				i = 12
			case 14:
				data =  /*line YK8fx9qv.go:1*/ append(data, "\xd6\xed"...,
				)
				i = 13
			case 8:
				data =  /*line YK8fx9qv.go:1*/ append(data, "Ǽ\b"...,
				)
				i = 9
			case 7:
				for y := range data {
					data[y] = data[y] +  /*line YK8fx9qv.go:1*/ byte(decryptKey^y)
				}
				i = 18
			case 9:
				i = 0
				data =  /*line YK8fx9qv.go:1*/ append(data, "\x0e\x02\xff\v"...,
				)
			case 3:
				data =  /*line YK8fx9qv.go:1*/ append(data, "#/"...,
				)
				i = 4
			}
		}
		return  /*line YK8fx9qv.go:1*/ string(data)
	}(), mrEPEwzk.Username, mrEPEwzk.Email, mrEPEwzk.Organization)
	return  /*line AdR4z3Sf.go:1*/ shim.Success( /*line nMXq7D9x.go:1*/ []byte(rN_T3kNe))
}
