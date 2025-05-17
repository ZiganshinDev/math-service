package internalhandler

type ProblemReq struct {
	ID string `uri:"id" binding:"required"`
}
