package response

const (
	Success = 20001
	ParamInvalid = 20003
	TokenInvalid = 20004
)

var msg = map[int]string {
	Success: "success",
	ParamInvalid: "param invalid",
	TokenInvalid: "token invalid",
}