package util

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Time", func() {
	var (
		then time.Time
		now time.Time
		year int
		month int
		day int
		hour int
		min int
		sec int
	)

	Describe("Diff", func() {
		Context("when i call Diff", func() {
			BeforeEach(func() {
				then = time.Date(
					2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
				now = time.Now()
				year, month, day, hour, min, sec = Diff(then, now)
			})

			It("should be fulfill year", func() {
				Expect(year).NotTo(BeNil())
			})

			It("should fulfill month", func() {
				Expect(month).NotTo(BeNil())
			})

			It("should fulfill day", func() {
				Expect(day).NotTo(BeNil())
			})

			It("should fulfill hour", func() {
				Expect(hour).NotTo(BeNil())
			})

			It("should fulfill min", func() {
				Expect(hour).NotTo(BeNil())
			})

			It("should fulfill min", func() {
				Expect(min).NotTo(BeNil())
			})

			It("should fulfill sec", func() {
				Expect(sec).NotTo(BeNil())
			})
		})
	})
})
