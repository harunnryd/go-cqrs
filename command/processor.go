package command

import (
	"fmt"
	"github.com/satori/go.uuid"
	u "quizes/util"
	"time"
)

type I interface {Process() error}

func (cmd *CreateStudentCommand) Process() error {
	var (
		student u.Student
		user u.User
	)

	student.ID = uuid.Must(uuid.NewV4())
	student.FullName = cmd.FullName
	student.Address = cmd.Address
	student.PhoneNumber = cmd.PhoneNumber
	student.Age = cmd.Age

	user.ID = cmd.ID
	user.Username = cmd.Username
	user.Password = cmd.Password
	user.ModelType = "students"
	user.ModelID = student.ID

	tx := u.DB.Begin()

	if err := tx.Create(&student).Error; err != nil {
		return tx.Rollback().Error
	}

	if err := tx.Create(&user).Error; err != nil {
		return tx.Rollback().Error
	}

	return tx.Commit().Error
}

func (cmd *CreateAdminCommand) Process() error {
	var (
		admin u.Admin
		user u.User
	)

	admin.ID = uuid.Must(uuid.NewV4())
	admin.FullName = cmd.FullName
	admin.Address = cmd.Address
	admin.Age = cmd.Age

	user.ID = cmd.ID
	user.Username = cmd.Username
	user.Password = cmd.Password
	user.ModelType = "admins"
	user.ModelID = admin.ID

	tx := u.DB.Begin()

	if err := tx.Create(&admin).Error; err != nil {
		return tx.Rollback().Error
	}

	if err := tx.Create(&user).Error; err != nil {
		return tx.Rollback().Error
	}

	return tx.Commit().Error
}

func (cmd *CreateExamCommand) Process() error {
	var (
		exam u.Exam
	)

	exam.ID = cmd.ID
	exam.Title = cmd.Title
	exam.WrongAnswerWeighted = cmd.WrongAnswerWeighted
	exam.RightAnswerWeighted = cmd.RightAnswerWeighted
	exam.DurationLimit = cmd.DurationLimit

	tx := u.DB.Begin()

	if err := tx.Create(&exam).Error; err != nil {
		return tx.Rollback().Error
	}

	return tx.Commit().Error
}

func (cmd *CreateQuestionCommand) Process() error {
	var (
		question u.Question
	)

	question.ID = cmd.ID
	question.ExamID = cmd.ExamID
	question.Text = cmd.Text

	tx := u.DB.Begin()

	if err := tx.Create(&question).Error; err != nil {
		return tx.Rollback().Error
	}

	return tx.Commit().Error
}

func (cmd *CreateChoiceCommand) Process() error {
	var (
		choice u.Choice
	)

	choice.ID = cmd.ID
	choice.ExamID = cmd.ExamID
	choice.QuestionID = cmd.QuestionID
	choice.Text = cmd.Text
	choice.IsCorrect = cmd.IsCorrect

	tx := u.DB.Begin()

	if err := tx.Create(&choice).Error; err != nil {
		return tx.Rollback().Error
	}

	return tx.Commit().Error
}

func (cmd *AttemptCommand) Process() error {
	var (
		attempt u.Attempt
		examResult struct { DurationLimit int }

	)

	u.DB.Raw("select duration_limit from exams where id = ? limit 1", cmd.ExamID).Scan(&examResult)


	attempt.ID = cmd.ID
	attempt.UserID = cmd.UserID
	attempt.ExamID = cmd.ExamID
	attempt.StartAt = time.Now().Local()
	attempt.EndAt = time.Now().Local().Add(time.Minute * time.Duration(examResult.DurationLimit))

	tx := u.DB.Begin()

	if err := tx.Create(&attempt).Error; err != nil {
		return tx.Rollback().Error
	}

	return tx.Commit().Error
}

func (cmd *UserAnswerCommand) Process() error {
	var (
		userAnswer u.UserAnswer
		attemptResult struct { EndAt time.Time }
		choiceResult struct { IsCorrect bool }
		examResult struct { WrongAnswerWeighted int; RightAnswerWeighted int}
		userAnswerTotal int
	)

	u.DB.Raw("select is_correct from choices where id = ? and question_id = ? and exam_id = ? limit 1", cmd.ChoiceID, cmd.QuestionID, cmd.ExamID).Scan(&choiceResult)
	u.DB.Raw("select wrong_answer_weighted, right_answer_weighted from exams where id = ? limit 1", cmd.ExamID).Scan(&examResult)
	u.DB.Raw("select count(*) from user_answers where user_id = ? and exam_id = ? and question_id = ?", cmd.UserID, cmd.ExamID, cmd.QuestionID).Count(&userAnswerTotal)
	u.DB.Raw("select end_at from attempts where user_id and exam_id = ? limit 1", cmd.UserID, cmd.AttemptID).Scan(&attemptResult)

	if userAnswerTotal != 0 {
		return nil
	}

	if attemptResult.EndAt.After(time.Now().Local()) {
		return nil
	}

	userAnswer.ID = cmd.ID
	userAnswer.UserID = cmd.UserID
	userAnswer.ExamID = cmd.ExamID
	userAnswer.QuestionID = cmd.QuestionID
	userAnswer.ChoiceID = cmd.ChoiceID
	userAnswer.AttemptID = cmd.AttemptID
	userAnswer.IsCorrect = choiceResult.IsCorrect
	if choiceResult.IsCorrect {
		userAnswer.AnswerWeighted = examResult.RightAnswerWeighted
	} else {
		userAnswer.AnswerWeighted = examResult.WrongAnswerWeighted
	}

	tx := u.DB.Begin()

	if err := tx.Create(&userAnswer).Error; err != nil {
		return tx.Rollback().Error
	}

	return tx.Commit().Error
}

func (cmd *SubmissionCommand) Process() error {
	var (
		submission u.Submission
		attemptResult struct { StartAt time.Time }
		userAnswerGroupResult struct {
			UserID uuid.UUID
			TotalScore int
			RightAnswerTotal int
			WrongAnswerTotal int
			NotAnswerTotal int
			TotalQuestion int
			Percentage float32
		}
	)

	u.DB.Raw(`select start_at from attempts where id = ? and user_id ? limit 1`, cmd.AttemptID, cmd.UserID).Scan(&attemptResult)

 	subQueryRightAnswerTotal := fmt.Sprintf(`
				select COUNT(*) from user_answers 
					where user_id = '%s' and 
					exam_id = '%s' and 
					attempt_id = '%s' and
					is_correct = true
			`, cmd.UserID, cmd.ExamID, cmd.AttemptID)

	subQueryWrongAnswerTotal := fmt.Sprintf(`
				select COUNT(*) from user_answers 
					where user_id = '%s' and 
					exam_id = '%s' and
					attempt_id = '%s' and
					is_correct = false
			`, cmd.UserID, cmd.ExamID, cmd.AttemptID)

	subQueryAnswerTotal := fmt.Sprintf(`
				select COUNT(*) from user_answers 
					where user_id = '%s' and 
					exam_id = '%s' and
					attempt_id = '%s'
			`, cmd.UserID, cmd.ExamID, cmd.AttemptID)

	subQueryTotalQuestion := fmt.Sprintf(`
				select COUNT(*) from questions 
					where exam_id = '%s'
			`, cmd.ExamID)

	subQueryTotalNotAnswer := fmt.Sprintf(`(%s) - (%s)`, subQueryAnswerTotal, subQueryTotalQuestion)
	subQueryPercentage := fmt.Sprintf(`(%s) * 100 / (%s)`, subQueryRightAnswerTotal, subQueryTotalQuestion)

	u.DB.Raw(fmt.Sprintf(`
		select 
			user_id, 
			SUM(answer_weighted) as total_score,
			COUNT(*) as total_answer,
			(%s) as right_answer_total,
			(%s) as wrong_answer_total,
			(%s) as total_question,
			(%s) as not_answer_total,
			(%s) as percentage
		from user_answers 
			where user_id = '%s' and
			exam_id = '%s' and
			attempt_id = '%s'
		group by user_id
	`,
	subQueryRightAnswerTotal,
	subQueryWrongAnswerTotal,
	subQueryTotalQuestion,
	subQueryTotalNotAnswer,
	subQueryPercentage,
	cmd.UserID,
	cmd.ExamID,
	cmd.AttemptID)).Scan(&userAnswerGroupResult)

	// diff time
	t := time.Now().Local()
	_, _, _, _, min, sec := u.Diff(attemptResult.StartAt, t)

	min = t.Minute() - min
	sec = t.Second() - sec

	submission.ID = cmd.ID
	submission.UserID = cmd.UserID
	submission.ExamID = cmd.ExamID
	submission.AttemptID = cmd.AttemptID
	submission.WrongAnswerTotal = userAnswerGroupResult.WrongAnswerTotal
	submission.RightAnswerTotal = userAnswerGroupResult.RightAnswerTotal
	submission.NotAnswerTotal = userAnswerGroupResult.NotAnswerTotal
	submission.Percentage = userAnswerGroupResult.Percentage
	submission.TotalScore = userAnswerGroupResult.TotalScore
	submission.DurationComplete = fmt.Sprintf("%d:%d", min, sec)

	tx := u.DB.Begin()

	if err := tx.Create(&submission).Error; err != nil {
		return tx.Rollback().Error
	}

	return tx.Commit().Error
}
