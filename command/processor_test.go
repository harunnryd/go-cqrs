package command

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"quizes/util"
)

var _ = Describe("Processor", func() {
	var (
		student CreateStudentCommand
		admin CreateAdminCommand
		exam CreateExamCommand
		question CreateQuestionCommand
		choice1 CreateChoiceCommand
		choice2 CreateChoiceCommand
		choice3 CreateChoiceCommand
		attempt AttemptCommand
		userAnswer UserAnswerCommand
		submission SubmissionCommand
		userTotal int
		studentTotal int
		adminTotal int
		examTotal int
		questionTotal int
		choiceTotal int
		attemptTotal int
		userAnswerTotal int
		userAnswerResult util.UserAnswer
		submissionResult util.Submission
	)

	BeforeSuite(func() {
		util.Migrate()
		util.DB.LogMode(true)
	})

	Describe("CreateStudentCommand Process", func() {
		Context("when i call process", func() {
			BeforeEach(func() {
				student = NewCreateStudentCommand(
					"juminten",
					"12345678",
					"juminten kasian pak",
					"jl. kebeneran!",
					"081234567890",
					21)
				student.Process()
			})

			Context("the table users & students", func() {
				BeforeEach(func() {
					util.DB.Raw(`
						select count(*) from users
						where username = ?
					`, student.Username).Count(&userTotal)
					util.DB.Raw(`
						select count(*) from students
							where full_name = ? and
							address = ? and
							phone_number = ? and
							age = ?
					`, student.FullName, student.Address, student.PhoneNumber, student.Age).Count(&studentTotal)
				})

				It("has one user and one student", func() {
					Expect(userTotal).To(Equal(1))
					Expect(studentTotal).To(Equal(1))
				})
			})
		})

		Describe("CreateAdminCommand Process", func() {
			Context("when i call process", func() {
				BeforeEach(func() {
					admin = NewCreateAdminCommand(
						"alioncom",
						"12345678",
						"ali oncom sayur",
						"jl. jalan mulu, abis dah tuh duit :(",
						25)
					admin.Process()
				})
				Context("the table users & admins", func() {
					BeforeEach(func() {
						util.DB.Raw(`
								select count(*) from users
									where username = ?
							`, admin.Username).Count(&userTotal)
						util.DB.Raw(`
								select count(*) from admins
									where full_name = ? and
									address = ? and
									age = ?
							`, admin.FullName, admin.Address, admin.Age).Count(&adminTotal)
					})

					It("has one user and admins", func() {
						Expect(userTotal).To(Equal(1))
						Expect(adminTotal).To(Equal(1))
					})
				})
			})

			Describe("CreateExamCommand Process", func() {
				Context("when i call process", func() {
					BeforeEach(func() {
						exam = NewCreateExamCommand("math exams", -2, 4, 10)
						exam.Process()
					})

					Context("the table exams", func() {
						BeforeEach(func() {
							util.DB.Raw(`
								select count(*) from exams
									where title = ? and
									wrong_answer_weighted = ? and
									right_answer_weighted = ? and
									duration_limit = ?
							`, exam.Title, exam.WrongAnswerWeighted, exam.RightAnswerWeighted, exam.DurationLimit).Count(&examTotal)
						})

						It("has one exam", func() {
							Expect(examTotal).To(Equal(1))
						})
					})
				})

				Describe("CreateQuestionCommand Process", func() {
					Context("when i call process", func() {
						BeforeEach(func() {
							question = NewCreateQuestionCommand(exam.ID, "1 + 1 = ?")
							question.Process()
						})

						Context("the table questions", func() {
							BeforeEach(func() {
								util.DB.Raw(`
									select count(*) from questions
										where exam_id = ? and
										text = ?
								`, question.ExamID, question.Text).Count(&questionTotal)
							})

							It("has one question", func() {
								Expect(questionTotal).To(Equal(1))
							})
						})

					})
					Describe("CreateChoiceCommand Process", func() {
						Context("when i call process", func() {
							BeforeEach(func() {
								choice1 = NewCreateChoiceCommand(exam.ID, question.ID, "2", true)
								choice2 = NewCreateChoiceCommand(exam.ID, question.ID, "3", false)
								choice3 = NewCreateChoiceCommand(exam.ID, question.ID, "20", false)
								choice1.Process()
								choice2.Process()
								choice3.Process()
							})

							Context("the table choices", func() {
								BeforeEach(func() {
									util.DB.Raw(`
											select count(*) from choices 
												where exam_id = ? and
												question_id = ?
										`, exam.ID, question.ID).Count(&choiceTotal)
								})

								It("has three choices", func() {
									Expect(choiceTotal).To(Equal(3))
								})
							})
						})

						Describe("AttempCommand Process", func() {
							Context("when i call process", func() {
								BeforeEach(func() {
									attempt = NewAttemptCommand(student.ID, exam.ID)
									attempt.Process()
								})

								Context("the table attempts", func() {
									BeforeEach(func() {
										util.DB.Raw(`
												select count(*) from attempts
													where user_id = ? and
													exam_id = ?
											`, attempt.UserID, attempt.ExamID).Count(&attemptTotal)
									})

									It("has one attempt", func() {
										Expect(attemptTotal).To(Equal(1))
									})
								})
							})

							Describe("UserAnswerCommand Process", func() {
								Context("when i call process and fulfill the answer with false choice", func() {
									BeforeEach(func() {
										userAnswer = NewUserAnswerCommand(student.ID, exam.ID, question.ID, choice2.ID, attempt.ID)
										userAnswer.Process()
									})

									Context("the table user_answers", func() {
										BeforeEach(func() {
											util.DB.Raw(`
												select count(*) from user_answers 
													where user_id = ? and 
													exam_id = ? and 
													question_id = ? and 
													choice_id = ? and 
													attempt_id = ?
											`, userAnswer.UserID, userAnswer.ExamID, userAnswer.QuestionID, userAnswer.ChoiceID, userAnswer.AttemptID).Count(&userAnswerTotal)

											util.DB.Raw(`
												select * from user_answers 
													where user_id = ? and 
													exam_id = ? and 
													question_id = ? and 
													choice_id = ? and 
													attempt_id = ?
											`, userAnswer.UserID, userAnswer.ExamID, userAnswer.QuestionID, userAnswer.ChoiceID, userAnswer.AttemptID).Scan(&userAnswerResult)
										})

										It("has one user answer", func() {
											Expect(userAnswerTotal).To(Equal(1))
										})

										It("has fulfill user answer", func() {
											Expect(userAnswerResult.IsCorrect).To(BeFalse())
											Expect(userAnswerResult.AnswerWeighted).To(Equal(-2))
										})
									})
								})

								Describe("SubmissionCommand Process", func() {
									Context("when i call process", func() {
										BeforeEach(func() {
											submission = NewSubmissionCommand(student.ID, exam.ID, attempt.ID)
											submission.Process()
										})

										Context("the table submissions", func() {
											BeforeEach(func() {
												util.DB.Raw(`
													select * from submissions
														where user_id = ? and
														exam_id = ? and
														attempt_id = ?
													limit 1
												`, student.ID, exam.ID, attempt.ID).Scan(&submissionResult)
											})

											It("have one submission", func() {
												Expect(submissionResult.UserID).To(Equal(student.ID))
												Expect(submissionResult.ExamID).To(Equal(exam.ID))
												Expect(submissionResult.AttemptID).To(Equal(attempt.ID))
												Expect(submissionResult.WrongAnswerTotal).To(Equal(1))
												Expect(submissionResult.RightAnswerTotal).To(Equal(0))
												Expect(submissionResult.NotAnswerTotal).To(Equal(0))
												Expect(submissionResult.Percentage).To(Equal(float32(0)))
												Expect(submissionResult.TotalScore).To(Equal(-2))
												Expect(submissionResult.DurationComplete).NotTo(BeNil())
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
	})
})