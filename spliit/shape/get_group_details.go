package shape

// GetGroupDetailsRequest is the input for the get group details endpoint.
type GetGroupDetailsRequest = GetGroupRequest

// GetGroupDetailsResponse extends GetGroupResponse with additional detail fields.
type GetGroupDetailsResponse struct {
	*GetGroupResponse

	ParticipantsWithExpenses []string `json:"participantsWithExpenses"`
}
