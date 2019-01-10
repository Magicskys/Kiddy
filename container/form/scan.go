package form

type Action_Sqlmap struct {
	Action string `form:"action" binding:"required"`
	Id []string `form:"uid[]" binding:"required"`
}