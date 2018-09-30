package commandApi

import (
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"quizes/broker"
	"quizes/command"
)

type I interface {}
type Promise map[string] I

func CreateStudentHandler(ctx echo.Context) error {
	f := new(struct{
		Username string `json:"username"`
		Password string `json:"password"`
		FullName string `json:"full_name"`
		Address string `json:"address"`
		PhoneNumber string `json:"phone_number"`
		Age int `json:"age"`
	})

	if err := ctx.Bind(f); err != nil {
		return err
	}

	cmd := command.NewCreateStudentCommand(
		f.Username,
		f.Password,
		f.FullName,
		f.Address,
		f.PhoneNumber,
		f.Age)

	broker.Dispatch("message", cmd)

	return ctx.JSON(http.StatusCreated, Promise{"id": cmd.ID, "message": "success!"})
}

func CreateAdminHandler(ctx echo.Context) error {
	f := new(struct{
		Username string `json:"username"`
		Password string `json:"password"`
		FullName string `json:"full_name"`
		Address string `json:"address"`
		Age int `json:"age"`
	})

	if err := ctx.Bind(f); err != nil {
		return err
	}

	cmd := command.NewCreateAdminCommand(
		f.Username,
		f.Password,
		f.FullName,
		f.Address,
		f.Age)

	broker.Dispatch("message", cmd)

	return ctx.JSON(http.StatusCreated, Promise{"id": cmd.ID, "message": "success!"})
}

func CreateExamHandler(ctx echo.Context) error {
	f := new(struct{
		Title string `json:"title"`
		WrongAnswerWeighted int `json:"wrong_answer_weighted"`
		RightAnswerWeighted int `json:"right_answer_weighted"`
		DurationLimit int `json:"duration_limit"`
	})

	if err := ctx.Bind(f); err != nil {
		return err
	}

	cmd := command.NewCreateExamCommand(
		f.Title,
		f.WrongAnswerWeighted,
		f.RightAnswerWeighted,
		f.DurationLimit)

	broker.Dispatch("message", cmd)

	return ctx.JSON(http.StatusCreated, Promise{"id": cmd.ID, "message": "success!"})
}

func CreateQuestionHandler(ctx echo.Context) error {
	f := new(struct{
		ExamID uuid.UUID `json:"exam_id"`
		Text string `json:"text"`
	})

	if err := ctx.Bind(f); err != nil {
		return err
	}

	cmd := command.NewCreateQuestionCommand(
		f.ExamID,
		f.Text)

	broker.Dispatch("message", cmd)

	return ctx.JSON(http.StatusCreated, Promise{"id": cmd.ID, "message": "success!"})
}

func CreateChoiceHandler(ctx echo.Context) error {
	f := new(struct{
		ExamID uuid.UUID `json:"exam_id"`
		QuestionID uuid.UUID `json:"question_id"`
		Text string `json:"text"`
		IsCorrect bool `json:"is_correct"`
	})

	if err := ctx.Bind(f); err != nil {
		return err
	}

	cmd := command.NewCreateChoiceCommand(
		f.ExamID,
		f.QuestionID,
		f.Text,
		f.IsCorrect)

	broker.Dispatch("message", cmd)

	return ctx.JSON(http.StatusCreated, Promise{"id": cmd.ID, "message": "success!"})
}

func CreateAttemptHandler(ctx echo.Context) error {
	f := new(struct{
		UserID uuid.UUID `json:"user_id"`
		ExamID uuid.UUID `json:"exam_id"`
	})

	if err := ctx.Bind(f); err != nil {
		return err
	}

	cmd := command.NewAttemptCommand(
		f.UserID,
		f.ExamID)

	broker.Dispatch("message", cmd)

	return ctx.JSON(http.StatusCreated, Promise{"id": cmd.ID, "message": "success!"})
}

func CreateUserAnswerHandler(ctx echo.Context) error {
	f := new(struct{
		UserID uuid.UUID `json:"user_id"`
		ExamID uuid.UUID `json:"exam_id"`
		QuestionID uuid.UUID `json:"question_id"`
		ChoiceID uuid.UUID `json:"choice_id"`
		AttemptID uuid.UUID `json:"attempt_id"`
	})

	if err := ctx.Bind(f); err != nil {
		return err
	}

	cmd := command.NewUserAnswerCommand(
		f.UserID,
		f.ExamID,
		f.QuestionID,
		f.ChoiceID,
		f.AttemptID)

	broker.Dispatch("message", cmd)

	return ctx.JSON(http.StatusCreated, Promise{"id": cmd.ID, "message": "success!"})
}

func CreateSubmissionHandler(ctx echo.Context) error {
	f := new(struct{
		UserID uuid.UUID `json:"user_id"`
		ExamID uuid.UUID `json:"exam_id"`
		AttemptID uuid.UUID `json:"attempt_id"`
	})

	if err := ctx.Bind(f); err != nil {
		return err
	}

	cmd := command.NewSubmissionCommand(
		f.UserID,
		f.ExamID,
		f.AttemptID)

	broker.Dispatch("message", cmd)

	return ctx.JSON(http.StatusCreated, Promise{"id": cmd.ID, "message": "success!"})
}

