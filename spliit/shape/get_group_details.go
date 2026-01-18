package shape

type GetGroupDetailsRequest = GetGroupRequest
type GetGroupDetailsResponse struct {
	*GetGroupResponse

	ParticipantsWithExpenses []string `json:"participantsWithExpenses"`
}
