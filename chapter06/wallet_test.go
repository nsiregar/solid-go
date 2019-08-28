package wallet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nsiregar/solid-go/chapter06"
)

var _ = Describe("Bitcoin Wallet", func() {
	var (
		ewallet wallet.Wallet
	)

	BeforeEach(func() {
		ewallet = wallet.Wallet{}
	})

	Describe("#Deposit", func() {
		It("increase total balance", func() {
			ewallet.Deposit(wallet.Bitcoin(10))

			Expect(ewallet.Balance()).To(Equal(wallet.Bitcoin(10)))
		})
	})

	Describe("#Withdraw", func() {
		BeforeEach(func() {
			ewallet.Deposit(wallet.Bitcoin(10))
		})

		Context("Sufficient funds", func() {
			It("decrease total balance", func() {
				ewallet.Withdraw(wallet.Bitcoin(5))

				Expect(ewallet.Balance()).To(Equal(wallet.Bitcoin(5)))
			})
		})

		Context("Insufficient funds", func() {
			It("raise error", func() {
				err := ewallet.Withdraw(wallet.Bitcoin(15))

				Expect(err).To(HaveOccurred())
			})
		})
	})
})
