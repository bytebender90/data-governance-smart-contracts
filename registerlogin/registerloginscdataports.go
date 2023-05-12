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

type ZohimWzR struct{}

                
func (u3xzONWW *ZohimWzR) Init(ipEgxdqE shim.ChaincodeStubInterface) pb.Response {
	return  /*line IeZwh80u.go:1*/ shim.Success(nil)
}

                   
func (u3xzONWW *ZohimWzR) Invoke(ipEgxdqE shim.ChaincodeStubInterface) pb.Response {

	gTiIOKPl, lsLFSQCl :=  /*line yeMUOgHD.go:1*/ ipEgxdqE.GetFunctionAndParameters()

	switch gTiIOKPl {
	case  /*line MtnHe7Jg.go:1*/ func() string {
		fullData :=  /*line MtnHe7Jg.go:1*/ []byte("3dbw\xa5\x13_\xd1j\x9a\x91\xfd\xd6i\x9b\x00\xb9\x1e\xd7\xd8\b\xd7")
		var data []byte
		data =  /*line MtnHe7Jg.go:1*/ append(data, fullData[6]+fullData[5], fullData[2]-fullData[11], fullData[12]+fullData[10], fullData[13]+fullData[15], fullData[21]-fullData[1], fullData[4]^fullData[7], fullData[17]-fullData[16], fullData[18]+fullData[14], fullData[3]+fullData[19], fullData[20]+fullData[8], fullData[9]-fullData[0])
		return  /*line MtnHe7Jg.go:1*/ string(data)
	}():
		return  /*line C1YFYJx4.go:1*/ u3xzONWW.ovHz35Wm(ipEgxdqE, lsLFSQCl)
	case  /*line LpPZxno2.go:1*/ func() string {
		fullData :=  /*line LpPZxno2.go:1*/ []byte("\b\x88-\xff\x06~\x8d\xe34\xa5\xbaH\n!\x05\xc0l\xe1k8\x85`\xee\x84")
		var data []byte
		data =  /*line LpPZxno2.go:1*/ append(data, fullData[10]-fullData[11], fullData[15]^fullData[9], fullData[7]^fullData[23], fullData[17]^fullData[1], fullData[20]+fullData[22], fullData[12]^fullData[5], fullData[2]+fullData[19], fullData[6]^fullData[3], fullData[13]+fullData[8], fullData[0]+fullData[18], fullData[14]^fullData[21], fullData[4]+fullData[16])
		return  /*line LpPZxno2.go:1*/ string(data)
	}():
		return  /*line Rwub6tXz.go:1*/ u3xzONWW.ukTSJP_V(ipEgxdqE, lsLFSQCl)
	case  /*line kMyRXxe4.go:1*/ func() string {
		data :=  /*line kMyRXxe4.go:1*/ []byte("upd\xaa\xb0\xc7O\x9e\xae\xb0*\xa5\xb5")
		positions := [...]byte{3, 5, 12, 4, 8, 8, 9, 10, 9, 3, 7, 3, 11, 11}
		for i := 0; i < 14; i += 2 {
			localKey :=  /*line kMyRXxe4.go:1*/ byte(i) +  /*line kMyRXxe4.go:1*/ byte(positions[i]^positions[i+1]) + 181
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line kMyRXxe4.go:1*/ string(data)
	}():
		return  /*line iIxEIZnA.go:1*/ u3xzONWW.l171la8N(ipEgxdqE, lsLFSQCl)
	case  /*line Fn62eY2a.go:1*/ func() string {
		key :=  /*line Fn62eY2a.go:1*/ []byte("\xf3\xbe%A\x00@`\xb5n\xd0:F\xe0\x00\xb3H\xf9\x00G\xf7`ְt\xe38^/")
		data :=  /*line Fn62eY2a.go:1*/ []byte("\x82\xb2? t%\xf5\xbe\xf7\xa2\x16\x1b\x93s\xc4'yd\bk\f\x93\xb7\xed\x917\x14J")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line Fn62eY2a.go:1*/ string(data)
	}():
		return  /*line DNqRjsq2.go:1*/ u3xzONWW.ayGq0Qam(ipEgxdqE, lsLFSQCl)
	case  /*line tyL8qAbn.go:1*/ func() string {
		key :=  /*line tyL8qAbn.go:1*/ []byte("\b\xb5ٖc\xe1\xf9\xa6\xbe\x10f\xd97J")
		data :=  /*line tyL8qAbn.go:1*/ []byte("o\x1aM\xe5\xd5Hl\xe87b\xc7G\x9e\xaf")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line tyL8qAbn.go:1*/ string(data)
	}():
		return  /*line YZoGPCa2.go:1*/ u3xzONWW.spza01Hy(ipEgxdqE)
	case  /*line Fwwz7XzX.go:1*/ func() string {
		var data []byte
		i := 4
		decryptKey := 141
		for counter := 0; i != 11; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 10:
				data =  /*line Fwwz7XzX.go:1*/ append(data, "\x8b\xa9\x93"...,
				)
				i = 8
			case 7:
				i = 9
				data =  /*line Fwwz7XzX.go:1*/ append(data, 155)
			case 0:
				i = 7
				data =  /*line Fwwz7XzX.go:1*/ append(data, 146)
			case 6:
				i = 10
				data =  /*line Fwwz7XzX.go:1*/ append(data, 133)
			case 4:
				i = 1
				data =  /*line Fwwz7XzX.go:1*/ append(data, "\x8a\x89\x9b"...,
				)
			case 5:
				data =  /*line Fwwz7XzX.go:1*/ append(data, "\x9c\x8e\x98\x8b"...,
				)
				i = 6
			case 2:
				data =  /*line Fwwz7XzX.go:1*/ append(data, "\xb3\x9d"...,
				)
				i = 0
			case 8:
				i = 3
				data =  /*line Fwwz7XzX.go:1*/ append(data, 135)
			case 3:
				i = 2
				data =  /*line Fwwz7XzX.go:1*/ append(data, "\xa1\x9b"...,
				)
			case 9:
				i = 11
				for y := range data {
					data[y] = data[y] ^  /*line Fwwz7XzX.go:1*/ byte(decryptKey^y)
				}
			case 1:
				data =  /*line Fwwz7XzX.go:1*/ append(data, "\xa7\x87"...,
				)
				i = 5
			}
		}
		return  /*line Fwwz7XzX.go:1*/ string(data)
	}():
		return  /*line w8lNWgMz.go:1*/ u3xzONWW.zxDvS5er(ipEgxdqE)
	case  /*line LdsOPbw1.go:1*/ func() string {
		fullData :=  /*line LdsOPbw1.go:1*/ []byte("Ⴋ\xfa\xe0~\xbf\xfalj\xe0\x8a\x85\xfc\x83\x92\xdf\x1c\x0f\xfd")
		var data []byte
		data =  /*line LdsOPbw1.go:1*/ append(data, fullData[14]-fullData[17], fullData[12]^fullData[4], fullData[10]-fullData[8], fullData[6]^fullData[3], fullData[7]+fullData[5], fullData[2]^fullData[16], fullData[0]-fullData[15], fullData[13]-fullData[11], fullData[19]+fullData[9], fullData[1]-fullData[18])
		return  /*line LdsOPbw1.go:1*/ string(data)
	}():
		return  /*line kw3amxrs.go:1*/ u3xzONWW.oaQpob1F(ipEgxdqE)
	case  /*line i6zVcady.go:1*/ func() string {
		var data []byte
		i := 7
		decryptKey := 251
		for counter := 0; i != 6; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				i = 4
				data =  /*line i6zVcady.go:1*/ append(data, 130)
			case 4:
				for y := range data {
					data[y] = data[y] ^  /*line i6zVcady.go:1*/ byte(decryptKey^y)
				}
				i = 6
			case 5:
				data =  /*line i6zVcady.go:1*/ append(data, "\x9e\x89"...,
				)
				i = 0
			case 7:
				data =  /*line i6zVcady.go:1*/ append(data, "\x8e\x8d\x9f\xbf"...,
				)
				i = 5
			case 0:
				data =  /*line i6zVcady.go:1*/ append(data, 157)
				i = 2
			case 2:
				i = 3
				data =  /*line i6zVcady.go:1*/ append(data, "\x9d\xa3\x99"...,
				)
			case 3:
				i = 1
				data =  /*line i6zVcady.go:1*/ append(data, "\xb1\x83\x8b\x83"...,
				)
			}
		}
		return  /*line i6zVcady.go:1*/ string(data)
	}():
		return  /*line rjr0XRqe.go:1*/ u3xzONWW.jNh2zf1L(ipEgxdqE)
	case  /*line BqC8Ce1C.go:1*/ func() string {
		fullData :=  /*line BqC8Ce1C.go:1*/ []byte("\xeb\xad\xd9۷\x8eb\x16}\r")
		var data []byte
		data =  /*line BqC8Ce1C.go:1*/ append(data, fullData[3]^fullData[4], fullData[6]+fullData[9], fullData[5]+fullData[2], fullData[7]-fullData[1], fullData[0]-fullData[8])
		return  /*line BqC8Ce1C.go:1*/ string(data)
	}():
		return  /*line H8lxHS7q.go:1*/ u3xzONWW.dnDmY7x_(ipEgxdqE, lsLFSQCl)
	case  /*line qyst9LQk.go:1*/ func() string {
		fullData :=  /*line qyst9LQk.go:1*/ []byte("\x88\xe1G^Z\x18Ǵ\x92\xb4(\x90\xc5\xfdN\x9c\x11\xb9\x93;n] \xe6ݷ(\x8b\xf7\xdc<\xad")
		var data []byte
		data =  /*line qyst9LQk.go:1*/ append(data, fullData[26]-fullData[25], fullData[21]^fullData[10], fullData[24]+fullData[0], fullData[12]+fullData[31], fullData[18]+fullData[23], fullData[11]-fullData[19], fullData[1]^fullData[8], fullData[28]+fullData[20], fullData[14]^fullData[30], fullData[9]^fullData[6], fullData[15]+fullData[7], fullData[17]^fullData[29], fullData[13]-fullData[27], fullData[16]^fullData[3], fullData[4]+fullData[5], fullData[2]^fullData[22])
		return  /*line qyst9LQk.go:1*/ string(data)
	}():
		return  /*line lISaBG5j.go:1*/ u3xzONWW.x7xK1un1(ipEgxdqE)
	case  /*line mbQTXtuh.go:1*/ func() string {
		fullData :=  /*line mbQTXtuh.go:1*/ []byte("\x8e\xc5,\x91P\x8eji{Z̟\x8c\xc4c\xab!\x05\xd7(4\xd1\x12\xfd+\xa9|\xe4\v\xe7\x8c\xeaDYP\f:`\xbaA\xadu!\xa6\xd4\u007f\xef\xd5")
		var data []byte
		data =  /*line mbQTXtuh.go:1*/ append(data, fullData[34]^fullData[16], fullData[33]^fullData[2], fullData[38]+fullData[15], fullData[10]-fullData[9], fullData[17]-fullData[12], fullData[20]+fullData[42], fullData[45]^fullData[35], fullData[25]-fullData[32], fullData[1]+fullData[40], fullData[44]+fullData[11], fullData[7]+fullData[29], fullData[0]+fullData[18], fullData[47]-fullData[14], fullData[21]-fullData[30], fullData[19]+fullData[4], fullData[46]-fullData[8], fullData[43]-fullData[39], fullData[27]+fullData[5], fullData[31]-fullData[26], fullData[28]^fullData[6], fullData[3]^fullData[23], fullData[36]^fullData[41], fullData[37]^fullData[22], fullData[24]-fullData[13])
		return  /*line mbQTXtuh.go:1*/ string(data)
	}():
		return  /*line IWucKBTo.go:1*/ u3xzONWW.oOf32tmm(ipEgxdqE)
	case  /*line kBE1onis.go:1*/ func() string {
		var data []byte
		i := 1
		decryptKey := 16
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 7:
				data =  /*line kBE1onis.go:1*/ append(data, "\x91\x9cj\xa0"...,
				)
				i = 8
			case 6:
				i = 9
				data =  /*line kBE1onis.go:1*/ append(data, "\x80\xa2"...,
				)
			case 4:
				for y := range data {
					data[y] = data[y] -  /*line kBE1onis.go:1*/ byte(decryptKey^y)
				}
				i = 0
			case 5:
				i = 6
				data =  /*line kBE1onis.go:1*/ append(data, "\x89\x86\x92"...,
				)
			case 2:
				data =  /*line kBE1onis.go:1*/ append(data, 132)
				i = 5
			case 3:
				data =  /*line kBE1onis.go:1*/ append(data, 157)
				i = 7
			case 8:
				data =  /*line kBE1onis.go:1*/ append(data, "s\x86\x92"...,
				)
				i = 2
			case 9:
				i = 4
				data =  /*line kBE1onis.go:1*/ append(data, "\xaa\xa5\xa9"...,
				)
			case 1:
				data =  /*line kBE1onis.go:1*/ append(data, "\x96\x93\xa1{"...,
				)
				i = 3
			}
		}
		return  /*line kBE1onis.go:1*/ string(data)
	}():
		return  /*line SGEmQTAz.go:1*/ u3xzONWW.pWAPy15K(ipEgxdqE)
	case  /*line kul_rU3H.go:1*/ func() string {
		fullData :=  /*line kul_rU3H.go:1*/ []byte("\xf0\xc4Ǚ\x9cP\x03b\xdc^\x96\r\xbb\xb9\xd6\xe2\x15\xa8\xa0\xf1")
		var data []byte
		data =  /*line kul_rU3H.go:1*/ append(data, fullData[18]^fullData[2], fullData[8]^fullData[13], fullData[10]^fullData[15], fullData[19]+fullData[9], fullData[7]-fullData[0], fullData[6]-fullData[4], fullData[11]-fullData[12], fullData[14]+fullData[3], fullData[17]^fullData[1], fullData[5]+fullData[16])
		return  /*line kul_rU3H.go:1*/ string(data)
	}():
		return  /*line wTbSxBl0.go:1*/ u3xzONWW.ipny_KJh(ipEgxdqE, lsLFSQCl)
	case  /*line J2F2sZ6n.go:1*/ func() string {
		fullData :=  /*line J2F2sZ6n.go:1*/ []byte("\"\xd9VBaR!\x02J\xff\xe8Dr}\xfcx\x02\x05\vot7p\xd3")
		var data []byte
		data =  /*line J2F2sZ6n.go:1*/ append(data, fullData[1]-fullData[12], fullData[5]^fullData[21], fullData[0]^fullData[2], fullData[8]+fullData[18], fullData[15]-fullData[17], fullData[6]+fullData[11], fullData[22]+fullData[7], fullData[9]+fullData[20], fullData[23]+fullData[19], fullData[14]+fullData[13], fullData[4]+fullData[10], fullData[3]+fullData[16])
		return  /*line J2F2sZ6n.go:1*/ string(data)
	}():
		return  /*line wa2qsUOz.go:1*/ u3xzONWW.jzlHeYA_(ipEgxdqE, lsLFSQCl)
	case  /*line jmhuVelO.go:1*/ func() string {
		var data []byte
		i := 1
		decryptKey := 156
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 3:
				data =  /*line jmhuVelO.go:1*/ append(data, 192)
				i = 9
			case 7:
				for y := range data {
					data[y] = data[y] ^  /*line jmhuVelO.go:1*/ byte(decryptKey^y)
				}
				i = 0
			case 6:
				i = 8
				data =  /*line jmhuVelO.go:1*/ append(data, "\xd9\xd8\xd9"...,
				)
			case 9:
				i = 6
				data =  /*line jmhuVelO.go:1*/ append(data, "\xc0\xdd\xfc\xde"...,
				)
			case 4:
				data =  /*line jmhuVelO.go:1*/ append(data, "\xc4\xcf"...,
				)
				i = 5
			case 2:
				data =  /*line jmhuVelO.go:1*/ append(data, "\xca\xc3\xd5\xca"...,
				)
				i = 4
			case 8:
				i = 2
				data =  /*line jmhuVelO.go:1*/ append(data, "\xd9\xf3\xd5\xed"...,
				)
			case 1:
				i = 3
				data =  /*line jmhuVelO.go:1*/ append(data, "\xc4\xd2\xc0"...,
				)
			case 5:
				i = 7
				data =  /*line jmhuVelO.go:1*/ append(data, 198)
			}
		}
		return  /*line jmhuVelO.go:1*/ string(data)
	}():
		return  /*line idA1QgGT.go:1*/ u3xzONWW.efR9N_Jt(ipEgxdqE)
	case  /*line D8bz4RDK.go:1*/ func() string {
		fullData :=  /*line D8bz4RDK.go:1*/ []byte("\xc4\xd9l\x9be\x16\xa2\xf9\x8e.\x0f/1W\x1bd+\xbe0o\xe9\n\xe9]\xc7\vY1A\x9e*TCICt")
		var data []byte
		data =  /*line D8bz4RDK.go:1*/ append(data, fullData[15]^fullData[5], fullData[31]^fullData[12], fullData[10]+fullData[4], fullData[20]-fullData[35], fullData[23]^fullData[11], fullData[34]+fullData[16], fullData[14]+fullData[27], fullData[28]+fullData[9], fullData[17]-fullData[13], fullData[1]+fullData[8], fullData[6]^fullData[24], fullData[25]+fullData[26], fullData[32]^fullData[21], fullData[29]-fullData[18], fullData[0]-fullData[19], fullData[33]+fullData[30], fullData[2]+fullData[7], fullData[3]^fullData[22])
		return  /*line D8bz4RDK.go:1*/ string(data)
	}():
		return  /*line ZMh1IzOE.go:1*/ u3xzONWW.eewx5evd(ipEgxdqE)
	case  /*line u8BPc0kw.go:1*/ func() string {
		seed :=  /*line u8BPc0kw.go:1*/ byte(54)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line u8BPc0kw.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line u8BPc0kw.go:1*/  /*line u8BPc0kw.go:1*/  /*line u8BPc0kw.go:1*/  /*line u8BPc0kw.go:1*/  /*line u8BPc0kw.go:1*/  /*line u8BPc0kw.go:1*/  /*line u8BPc0kw.go:1*/  /*line u8BPc0kw.go:1*/  /*line u8BPc0kw.go:1*/  /*line u8BPc0kw.go:1*/  /*line u8BPc0kw.go:1*/  /*line u8BPc0kw.go:1*/  /*line u8BPc0kw.go:1*/  /*line u8BPc0kw.go:1*/ fnc(63)(251)(244)(253)(19)(241)(240)(30)(242)(13)(215)(37)(248)(9)
		return  /*line u8BPc0kw.go:1*/ string(data)
	}():
		return  /*line ruRQYnhM.go:1*/ u3xzONWW.cTp0CzYJ(ipEgxdqE, lsLFSQCl)
	case  /*line Hc7ZyxZS.go:1*/ func() string {
		seed :=  /*line Hc7ZyxZS.go:1*/ byte(129)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line Hc7ZyxZS.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line Hc7ZyxZS.go:1*/  /*line Hc7ZyxZS.go:1*/  /*line Hc7ZyxZS.go:1*/  /*line Hc7ZyxZS.go:1*/  /*line Hc7ZyxZS.go:1*/  /*line Hc7ZyxZS.go:1*/  /*line Hc7ZyxZS.go:1*/  /*line Hc7ZyxZS.go:1*/  /*line Hc7ZyxZS.go:1*/  /*line Hc7ZyxZS.go:1*/  /*line Hc7ZyxZS.go:1*/  /*line Hc7ZyxZS.go:1*/  /*line Hc7ZyxZS.go:1*/  /*line Hc7ZyxZS.go:1*/ fnc(230)(2)(29)(211)(42)(230)(27)(198)(51)(56)(216)(236)(16)(229)
		return  /*line Hc7ZyxZS.go:1*/ string(data)
	}():
		return  /*line VMwAdDH9.go:1*/ u3xzONWW.mdfCz3FT(ipEgxdqE, lsLFSQCl)
	case  /*line vWGKKTpn.go:1*/ func() string {
		key :=  /*line vWGKKTpn.go:1*/ []byte("\x02\x14\xc6\x0e\xaa{\xd3uS\x9c\xf3t\x1e-\x96\f<")
		data :=  /*line vWGKKTpn.go:1*/ []byte("fy'q\x1e\xe4I\xd6\xc7\x014ׁ\x9c\vz\xb0")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line vWGKKTpn.go:1*/ string(data)
	}():
		return  /*line Jx73jzYU.go:1*/ u3xzONWW.iYOlV2HM(ipEgxdqE, lsLFSQCl)
	case  /*line M6uluDBb.go:1*/ func() string {
		seed :=  /*line M6uluDBb.go:1*/ byte(48)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line M6uluDBb.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line M6uluDBb.go:1*/  /*line M6uluDBb.go:1*/  /*line M6uluDBb.go:1*/  /*line M6uluDBb.go:1*/  /*line M6uluDBb.go:1*/  /*line M6uluDBb.go:1*/  /*line M6uluDBb.go:1*/  /*line M6uluDBb.go:1*/  /*line M6uluDBb.go:1*/  /*line M6uluDBb.go:1*/  /*line M6uluDBb.go:1*/  /*line M6uluDBb.go:1*/  /*line M6uluDBb.go:1*/  /*line M6uluDBb.go:1*/  /*line M6uluDBb.go:1*/ fnc(145)(36)(89)(167)(91)(161)(85)(155)(18)(70)(140)(36)(78)(149)(48)
		return  /*line M6uluDBb.go:1*/ string(data)
	}():
		return  /*line _ocxUWZu.go:1*/ u3xzONWW._IMfdl2d(ipEgxdqE, lsLFSQCl)
	case  /*line fPFLADif.go:1*/ func() string {
		key :=  /*line fPFLADif.go:1*/ []byte("%6\xcd\xc6*[{\x85!l\r7\x1e\xe8\xc2\xecr.$?v\xfd\xe2")
		data :=  /*line fPFLADif.go:1*/ []byte("TC\xa8\xb4S\x0e\b\xe0S\x1fOVm\x8d\xa6\x83\x1c}P^\x02\x88\x91")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line fPFLADif.go:1*/ string(data)
	}():
		return  /*line JxRonBpL.go:1*/ u3xzONWW.uhEjaEas(ipEgxdqE, lsLFSQCl)
	case  /*line ZxQWKNE5.go:1*/ func() string {
		key :=  /*line ZxQWKNE5.go:1*/ []byte("X\x85kgK°>e\xb6\b\x93\x05XU\"\x91\xec\xfd)E\\\xc5pa\x11\u007f\xf4\r")
		data :=  /*line ZxQWKNE5.go:1*/ []byte("\x19\xf0\xfa\v.\x93\xc3'\r\xbd:\xcen\r\x0fM\xddgw8/\x19\xae\xd1\rS\xd0~Z")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line ZxQWKNE5.go:1*/ string(data)
	}():
		return  /*line ai0XXaB9.go:1*/ u3xzONWW.dlTruzqG(ipEgxdqE, lsLFSQCl)
	case  /*line Sq0fHGUy.go:1*/ func() string {
		data :=  /*line Sq0fHGUy.go:1*/ []byte("L\xd1;\xc8\xdc\xdd\xdbs=\xd9o\xeaN")
		positions := [...]byte{2, 11, 5, 3, 1, 4, 12, 9, 4, 4, 6, 8, 4, 0, 4, 9, 6, 11}
		for i := 0; i < 18; i += 2 {
			localKey :=  /*line Sq0fHGUy.go:1*/ byte(i) +  /*line Sq0fHGUy.go:1*/ byte(positions[i]^positions[i+1]) + 128
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line Sq0fHGUy.go:1*/ string(data)
	}():
		return  /*line Fb895goI.go:1*/ u3xzONWW.iiQph8om(ipEgxdqE, lsLFSQCl)
	case  /*line XUghwCgk.go:1*/ func() string {
		data :=  /*line XUghwCgk.go:1*/ []byte("@)\x1bU\x1bB&G\xee#A")
		positions := [...]byte{7, 8, 9, 1, 2, 4, 2, 9, 5, 0, 7, 6, 4, 10, 6, 1}
		for i := 0; i < 16; i += 2 {
			localKey :=  /*line XUghwCgk.go:1*/ byte(i) +  /*line XUghwCgk.go:1*/ byte(positions[i]^positions[i+1]) + 24
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line XUghwCgk.go:1*/ string(data)
	}():
		return  /*line fMhrZksF.go:1*/ u3xzONWW.zHAosy7V(ipEgxdqE)
	case  /*line CU3jPKVH.go:1*/ func() string {
		seed :=  /*line CU3jPKVH.go:1*/ byte(0)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line CU3jPKVH.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/  /*line CU3jPKVH.go:1*/ fnc(99)(203)(147)(36)(80)(118)(33)(45)(98)(199)(131)(7)(24)(41)(66)(162)(54)(121)(238)(207)(170)(76)
		return  /*line CU3jPKVH.go:1*/ string(data)
	}():
		return  /*line iIAZCzzQ.go:1*/ u3xzONWW.qxnshQm_(ipEgxdqE, lsLFSQCl)
	}

	 /*line t5Dt24GZ.go:1*/ fmt.Println( /*line WM2UNAuj.go:1*/ func() string {
		seed :=  /*line WM2UNAuj.go:1*/ byte(206)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line WM2UNAuj.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/  /*line WM2UNAuj.go:1*/ fnc(55)(115)(238)(213)(166)(70)(71)(210)(169)(77)(86)(250)(245)(239)(138)(90)(183)(115)(220)(116)(46)(107)(207)(147)(253)(224)
		return  /*line WM2UNAuj.go:1*/ string(data)
	}() + gTiIOKPl)	       
	return  /*line Af2iuBeb.go:1*/ shim.Error( /*line Rzl4K3FU.go:1*/ func() string {
		seed :=  /*line Rzl4K3FU.go:1*/ byte(37)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line Rzl4K3FU.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/  /*line Rzl4K3FU.go:1*/ fnc(119)(249)(246)(238)(16)(255)(237)(17)(166)(89)(235)(27)(229)(31)(248)(233)(80)(166)(19)(23)(243)(247)(19)(226)(1)(80)(169)(7)(6)(25)(236)(26)(225)(31)(250)(225)
		return  /*line Rzl4K3FU.go:1*/ string(data)
	}())
}

func main() {
	dTakwHTF :=  /*line QitceHqX.go:1*/ shim.Start( /*line CD3Z6Mau.go:1*/ new(ZohimWzR))
	if dTakwHTF != nil {
		 /*line XkeQDvnZ.go:1*/ fmt.Printf( /*line TDHDyAbS.go:1*/ func() string {
			seed :=  /*line TDHDyAbS.go:1*/ byte(155)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line TDHDyAbS.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/  /*line TDHDyAbS.go:1*/ fnc(222)(11)(246)(21)(253)(172)(75)(247)(27)(231)(8)(237)(31)(247)(167)(77)(19)(239)(20)(255)(243)(236)(11)(31)(185)(33)(18)(232)(29)(230)(21)(191)(100)(141)(70)
			return  /*line TDHDyAbS.go:1*/ string(data)
		}(), dTakwHTF)
	}
}

var garbleActionID = "zhK629o2d41qPIujwrmR"
