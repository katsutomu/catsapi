package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SearchCatAction", func() {
	Context("type = 1 でリクエストする", func() {
		cat := SearchCatAction(1)
		It("三毛猫が返却される", func() {
			Expect(cat).To(Equal("三毛猫"))
		})
	})
	Context("type != 1 でリクエストする", func() {
		cat := SearchCatAction(0)
		It("いないよ！が返却される", func() {
			Expect(cat).To(Equal("いないよ！"))
		})
	})
})
