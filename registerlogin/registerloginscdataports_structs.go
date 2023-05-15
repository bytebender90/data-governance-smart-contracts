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

                              
type Org struct {
	ObjectType	string		`json:"docType"`
	Name	string		`json:"name"`
	Role	string		`json:"role"`
	Vat_Number	string		`json:"vat_number"`
	Country	string		`json:"country"`
	City	string		`json:"city"`
	Postal_code	string		`json:"Postal_code"`
	Address	string		`json:"address"`
	Phone	string		`json:"phone"`
	DateTime	string		`json:"dateTime"`
	IsExternalOrg	bool		`json:"isExternalOrg"`
	ExternalOrgName	string		`json:"externalOrgName"`
	Created_by	string		`json:"created_by"`
}

                      
type User struct {
	ObjectType	string		`json:"docType"`
	Name	string		`json:"name"`
	Surname	string		`json:"surname"`
	Username	string		`json:"username"`
	Email	string		`json:"email"`
	Organization	string		`json:"organization"`
	Type	string		`json:"type"`
	Status	string		`json:"status"`
	DateTime	string		`json:"dateTime"`
	ExternalOrg	string		`json:"externalOrg"`
	IsExternalUser	bool		`json:"isExternalUser"`
}

                             
type UserCredentials struct {
	ObjectType	string		`json:"docType"`
	UserId	string		`json:"userid"`
	Username	string		`json:"username"`
	Password	string		`json:"password"`
	Name	string		`json:"name"`
	Surname	string		`json:"surname"`
	Email	string		`json:"email"`
	Organization	string		`json:"organization"`
	Type	string		`json:"type"`
	Created_by	string		`json:"created_by"`
}
