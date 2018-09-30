package command

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Events", func() {
	var (
		student CreateStudentCommand
		admin CreateAdminCommand
		exam CreateExamCommand
		question CreateQuestionCommand
		choice CreateChoiceCommand
		attempt AttemptCommand
		userAnswer UserAnswerCommand
		submission SubmissionCommand
	)

	Describe("CreateAdminCommand", func() {
		Context("when i call NewCreateAdminCommand", func() {
			BeforeEach(func() {
				admin = NewCreateAdminCommand("alioncom", "12345678", "ali oncom", "jl. makanya dijajanin biar jadian", 35)
			})

			It("should fulfill ID", func() {
				Expect(admin.ID).NotTo(BeNil())
			})

			It("should fulfill Type", func() {
				Expect(admin.Type).To(Equal("CreateAdminCommand"))
			})

			It("should fulfill Username", func() {
				Expect(admin.Username).To(Equal("alioncom"))
			})

			It("should fulfill Password", func() {
				Expect(admin.Password).NotTo(BeEmpty())
			})

			It("should fulfill FullName", func() {
				Expect(admin.FullName).To(Equal("ali oncom"))
			})

			It("should fulfill Address", func() {
				Expect(admin.Address).To(Equal("jl. makanya dijajanin biar jadian"))
			})

			It("should fulfill Age", func() {
				Expect(admin.Age).To(Equal(35))
			})
		})

		Describe("CreateStudentCommand", func() {
			Context("when i call NewCreateStudentCommand", func() {
				BeforeEach(func() {
					student = NewCreateStudentCommand("otoyunyu", "12345678", "otoy selalu ceria", "jl. jalan doank, jadian mah kagak", "081234567890", 21)
				})

				It("should fulfill ID", func() {
					Expect(student.ID).NotTo(BeNil())
				})

				It("should fulfill Type", func() {
					Expect(student.Type).To(Equal("CreateStudentCommand"))
				})

				It("should fulfill Username", func() {
					Expect(student.Username).To(Equal("otoyunyu"))
				})

				It("should fulfill Password", func() {
					Expect(student.Password).NotTo(BeEmpty())
				})

				It("should fulfill FullName", func() {
					Expect(student.FullName).To(Equal("otoy selalu ceria"))
				})

				It("should fulfill Address", func() {
					Expect(student.Address).To(Equal("jl. jalan doank, jadian mah kagak"))
				})

				It("should fulfill PhoneNumber", func() {
					Expect(student.PhoneNumber).To(Equal("081234567890"))
				})

				It("should fulfill Age", func() {
					Expect(student.Age).To(Equal(21))
				})
			})

			Describe("CreateExamCommand", func() {
				Context("when i call NewCreateExamCommand", func() {
					BeforeEach(func() {
						exam = NewCreateExamCommand("ujian akhir akhiran", -2, 4, 15)
					})

					It("should fulfill ID", func() {
						Expect(exam.ID).NotTo(BeEmpty())
					})

					It("should fulfill Type", func() {
						Expect(exam.Type).To(Equal("CreateExamCommand"))
					})

					It("should fulfill Title", func() {
						Expect(exam.Title).To(Equal("ujian akhir akhiran"))
					})

					It("should fulfill WrongAnswerWeighted", func() {
						Expect(exam.WrongAnswerWeighted).To(Equal(-2))
					})

					It("should fulfill RightAnswerWeighted", func() {
						Expect(exam.RightAnswerWeighted).To(Equal(4))
					})

					It("should fulfill DurationLimit", func() {
						Expect(exam.DurationLimit).To(Equal(15))
					})
				})

				Describe("CreateQuestionCommand", func() {
					Context("when i call NewCreateQuestionCommand", func() {
						BeforeEach(func() {
							question = NewCreateQuestionCommand(exam.ID, "siapa nama saya?")
						})

						It("should fulfill ID", func() {
							Expect(question.ID).NotTo(BeEmpty())
						})

						It("should fulfill Type", func() {
							Expect(question.Type).To(Equal("CreateQuestionCommand"))
						})

						It("should fulfill ExamID", func() {
							Expect(question.ExamID).NotTo(BeEmpty())
							Expect(question.ExamID).To(Equal(exam.ID))
						})

						It("should fulfill Text", func() {
							Expect(question.Text).To(Equal("siapa nama saya?"))
						})
					})

					Describe("CreateChoiceCommand", func() {
						Context("when i call NewCreateChoiceCommand", func() {
							BeforeEach(func() {
								choice = NewCreateChoiceCommand(exam.ID, question.ID, "otoy", false)
							})

							It("should fulfill ID", func() {
								Expect(choice.ID).NotTo(BeEmpty())
							})

							It("should fulfill Type", func() {
								Expect(choice.Type).To(Equal("CreateChoiceCommand"))
							})

							It("should fulfill ExamID", func() {
								Expect(choice.ExamID).NotTo(BeEmpty())
								Expect(choice.ExamID).To(Equal(exam.ID))
							})

							It("should fulfill QuestionID", func() {
								Expect(choice.QuestionID).NotTo(BeEmpty())
								Expect(choice.QuestionID).To(Equal(question.ID))
							})

							It("should fulfill Text", func() {
								Expect(choice.Text).To(Equal("otoy"))
							})

							It("should fulfill IsCorrect", func() {
								Expect(choice.IsCorrect).To(BeFalse())
							})
						})

						Describe("AttemptCommand", func() {
							Context("when i call NewAttemptCommand", func() {
								BeforeEach(func() {
									attempt = NewAttemptCommand(student.ID, exam.ID)
								})

								It("should fulfill ID", func() {
									Expect(attempt.ID).NotTo(BeEmpty())
								})

								It("should fulfill Type", func() {
									Expect(attempt.Type).To(Equal("AttemptCommand"))
								})

								It("should fulfill UserID", func() {
									Expect(attempt.UserID).NotTo(BeEmpty())
									Expect(attempt.UserID).To(Equal(student.ID))
								})

								It("should fulfill ExamID", func() {
									Expect(attempt.ExamID).NotTo(BeEmpty())
									Expect(attempt.ExamID).To(Equal(exam.ID))
								})
							})

							Describe("UserAnswerCommand", func() {
								Context("when i call NewUserAnswerCommand", func() {
									BeforeEach(func() {
										userAnswer = NewUserAnswerCommand(student.ID, exam.ID, question.ID, choice.ID, attempt.ID)
									})

									It("should fulfill ID", func() {
										Expect(userAnswer.ID).NotTo(BeEmpty())
									})

									It("should fulfill Type", func() {
										Expect(userAnswer.Type).To(Equal("UserAnswerCommand"))
									})

									It("should fulfill UserID", func() {
										Expect(userAnswer.UserID).NotTo(BeEmpty())
										Expect(userAnswer.UserID).To(Equal(student.ID))
									})

									It("should fulfill ExamID", func() {
										Expect(userAnswer.ExamID).NotTo(BeEmpty())
										Expect(userAnswer.ExamID).To(Equal(exam.ID))
									})

									It("should fulfill QuestionID", func() {
										Expect(userAnswer.QuestionID).NotTo(BeEmpty())
										Expect(userAnswer.QuestionID).To(Equal(question.ID))
									})

									It("should fulfill ChoiceID", func() {
										Expect(userAnswer.ChoiceID).NotTo(BeEmpty())
										Expect(userAnswer.ChoiceID).To(Equal(choice.ID))
									})

									It("should fulfill AttemptID", func() {
										Expect(userAnswer.AttemptID).NotTo(BeEmpty())
										Expect(userAnswer.AttemptID).To(Equal(attempt.ID))
									})
								})
							})

							Describe("SubmissionCommand", func() {
								Context("when i call NewSubmissionCommand", func() {
									BeforeEach(func() {
										submission = NewSubmissionCommand(student.ID, exam.ID, attempt.ID)
									})

									It("should fulfill ID", func() {
										Expect(submission.ID).NotTo(BeEmpty())
									})

									It("should fulfill Type", func() {
										Expect(submission.Type).To(Equal("SubmissionCommand"))
									})

									It("should fulfill UserID", func() {
										Expect(submission.UserID).NotTo(BeEmpty())
										Expect(submission.UserID).To(Equal(student.ID))
									})

									It("should fulfill ExamID", func() {
										Expect(submission.ExamID).NotTo(BeEmpty())
										Expect(submission.ExamID).To(Equal(exam.ID))
									})

									It("should fulfill AttemptID", func() {
										Expect(submission.AttemptID).NotTo(BeEmpty())
										Expect(submission.AttemptID).To(Equal(attempt.ID))
									})
								})
							})
						})
					})
				})
			})
		})
	})
})