package util

import (
. "github.com/onsi/ginkgo"
. "github.com/onsi/gomega"
)
var _ = Describe("Main", func() {
	BeforeSuite(func() {
		DB.LogMode(true)
		Migrate()
	})

	Describe("Database", func() {
		Context("when i call Migrate", func() {
			BeforeEach(func() {
				Migrate();
			})

			It("should migrating tables", func() {
				Expect(DB.HasTable("users")).To(BeTrue())
				Expect(DB.HasTable("students")).To(BeTrue())
				Expect(DB.HasTable("admins")).To(BeTrue())
				Expect(DB.HasTable("exams")).To(BeTrue())
				Expect(DB.HasTable("questions")).To(BeTrue())
				Expect(DB.HasTable("choices")).To(BeTrue())
				Expect(DB.HasTable("attempts")).To(BeTrue())
				Expect(DB.HasTable("user_answers")).To(BeTrue())
				Expect(DB.HasTable("submissions")).To(BeTrue())
			})
		})
	})
})

