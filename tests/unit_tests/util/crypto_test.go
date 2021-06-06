package util_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/wallet-tracky/Golang-backend/model"
	"github.com/wallet-tracky/Golang-backend/util"
)

var _=Describe("Test password encryption",func ()  {
	
	savedPassword:="password"

	Context("Positive cases :",func(){

		It("Should not return error on encrypting password",func(){
			_,err:=util.EncryptPassword(savedPassword)

			Expect(err).To(BeNil())
			
		})

		It("Should match on comparsion of correct password",func ()  {
			encryptedPassword,_:=util.EncryptPassword(savedPassword)

			user:=model.User{Password: encryptedPassword}

			inputPassword:="password" //same as variable savedPassword
			isMatching:=user.ComparePassword(inputPassword)

			Expect(isMatching).To(Equal(true))
		})
	})

	Context("Negative cases: ",func() {
		It("Should not give match on comparison of incorrect password",func(){

			encryptedPassword,_:=util.EncryptPassword(savedPassword)

			user:=model.User{Password: encryptedPassword}

			inputPassword:="notpassword" //different from variable savedPassword
			isMatching:=user.ComparePassword(inputPassword)

			Expect(isMatching).To(Equal(false))
		})
	})
})


var _= Describe("Test token generation ",func() {

	Context("Positive scenario :",func ()  {
		
		It("Should not return error for proper arguments for token generation",func() {

			_,err:=util.GenerateToken(1,"access token")

			Expect(err).To(BeNil())
		})

		It("Shoud return proper claim or resolving token",func ()  {
			userId:=1

			token,_:=util.GenerateToken(userId,"access token")

			claim,_:=util.VerifyToken(token)

			Expect(claim.UserId).To(Equal(userId))
		})
	})

})



