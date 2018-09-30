package queryApi

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"quizes/util"
)

func ListStudentHandler(ctx echo.Context) error {
	var listStudent []struct {
		ID uuid.UUID `json:"id"`
		Username string `json:"username"`
		FullName string `json:"full_name"`
		Address string `json:"address"`
		PhoneNumber string `json:"phone_number"`
		Age int `json:"age"`
	}

	util.DB.Raw(`
		select 
			u.id as id, 
			u.username as username, 
			s.full_name as full_name, 
			s.address as address, 
			s.phone_number as phone_number, 
			s.age as age 
		from users u, students s 
		where u.model_id = s.id
	`).Scan(&listStudent)

	return ctx.JSON(http.StatusOK, listStudent)
}

func ListAdminHandler(ctx echo.Context) error {
	var listAdmin []struct {
		ID uuid.UUID `json:"id"`
		Username string `json:"username"`
		FullName string `json:"full_name"`
		Address string `json:"address"`
		Age int `json:"age"`
	}

	util.DB.Raw(`
		select 
			u.id as id, 
			u.username as username, 
			a.full_name as full_name, 
			a.address as address, 
			a.age as age 
		from users u, admins a 
		where u.model_id = a.id
	`).Scan(&listAdmin)

	return ctx.JSON(http.StatusOK, listAdmin)
}

func ListExamHandler(ctx echo.Context) error {
	var listExam []struct {
		ID uuid.UUID `json:"id"`
		Title string `json:"title"`
		WrongAnswerWeighted int `json:"wrong_answer_weighted"`
		RightAnswerWeighted int `json:"right_answer_weighted"`
		DurationLimit int `json:"duration_limit"`
	}

	util.DB.Raw(`
		select 
			id, 
			title, 
			wrong_answer_weighted, 
			right_answer_weighted, 
			duration_limit 
		from exams
	`).Scan(&listExam)

	return ctx.JSON(http.StatusOK, listExam)
}

func DetailExamHandler(ctx echo.Context) error {
	var (
		exam util.Exam
		detailExam struct {
			ID uuid.UUID `json:"id"`
			Title string `json:"title"`
			WrongAnswerWeighted int `json:"wrong_answer_weighted"`
			RightAnswerWeighted int `json:"right_answer_weighted"`
			DurationLimit int `json:"duration_limit"`
			Questions []struct {
				ID uuid.UUID `json:"id"`
				Text string `json:"text"`
				Choices []struct {
					ID uuid.UUID `json:"id"`
					Text string `json:"text"`
				} `json:"choices"`
			} `json:"questions"`
		}
	)

	examID := ctx.Param("exam_id")

	util.DB.Model(&exam).
		Preload("Questions.Choices").
		Where("id = ?", examID).
		Find(&exam)

	copier.Copy(&detailExam, &exam)

	return ctx.JSON(http.StatusOK, detailExam)
}

func DetailSubmissionHandler(ctx echo.Context) error {
	var (
		detailSubmission struct {
			ID uuid.UUID `json:"id"`
			UserID uuid.UUID `json:"user_id"`
			ExamID uuid.UUID `json:"exam_id"`
			AttemptID uuid.UUID `json:"attempt_id"`
			WrongAnswerTotal int `json:"wrong_answer_total"`
			RightAnswerTotal int `json:"right_answer_total"`
			NotAnswerTotal int `json:"not_answer_total"`
			Percentage float32 `json:"percentage"`
			TotalScore int `json:"total_score"`
			DurationComplete string `json:"duration_complete"`
		}
	)

	userID := ctx.Param("user_id")
	examID := ctx.Param("exam_id")
	attemptID := ctx.Param("attempt_id")

	util.DB.Raw(`
		select * from submissions 
			where user_id = ? and 
			exam_id = ? and 
			attempt_id = ?
	`, userID, examID, attemptID).Scan(&detailSubmission)

	return ctx.JSON(http.StatusOK, detailSubmission)
}

func ListRankExam(ctx echo.Context) error {
	var (
		submissions []util.Submission
		listSubmission []struct {
			ID uuid.UUID `json:"id"`
			UserID uuid.UUID `json:"user_id"`
			ExamID uuid.UUID `json:"exam_id"`
			AttemptID uuid.UUID `json:"attempt_id"`
			WrongAnswerTotal int `json:"wrong_answer_total"`
			RightAnswerTotal int `json:"right_answer_total"`
			NotAnswerTotal int `json:"not_answer_total"`
			Percentage float32 `json:"percentage"`
			TotalScore int `json:"total_score"`
			DurationComplete string `json:"duration_complete"`
			User struct {
				ID uuid.UUID `json:"id"`
				Username string `json:"username"`
			}
		}
	)

	examID := ctx.Param("exam_id")

	util.DB.Where(`exam_id = ?`, examID).Order(`total_score asc`).Preload("User").Find(&submissions)

	copier.Copy(&listSubmission, &submissions)

	return ctx.JSON(http.StatusOK, listSubmission)
}