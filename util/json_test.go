package util

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Json", func() {
	var (
		s interface{}
		b []byte
	)

	Describe("Encode", func() {
		Context("when i call Encode", func() {
			BeforeEach(func() {
				s = struct { Name string } {Name: "entong"}
				b = Encode(s)
			})

			It("should convert to []byte", func() {
				Expect(b).To(Equal([]byte(`{"Name":"entong"}`)))
			})
		})
	})

	Describe("Decode", func() {
		Context("when i call Decode", func() {
			BeforeEach(func() {
				Decode(b, &s)
			})

			It("should convert to map[string]interface {}", func() {
				Expect(s).To(Equal(map[string]interface{}{"Name": "entong"}))
			})
		})
	})
})