// Package spliit provides a typed client for the Spliit tRPC API.
//
// The package lets callers batch multiple endpoint calls into a single request
// while preserving typed request and response values.
//
// The usual flow is:
//   - create a Client,
//   - construct one or more TypedCall values with NewCall,
//   - send them with Client.SendRequests,
//   - read the populated Result or Err fields from the original call values.
//
// Example:
//
//	client := spliit.NewClient()
//	getGroupCall := spliit.NewCall(
//		&endpoint.GetGroup{},
//		&shape.GetGroupRequest{
//			GroupID: "trip-123",
//		},
//	)
//
//	_, err := client.SendRequests(context.Background(), getGroupCall)
//	if err != nil {
//		log.Fatal(err)
//	}
//	if getGroupCall.Err != nil {
//		log.Fatal(getGroupCall.Err)
//	}
//
//	log.Println(getGroupCall.Result.Group.Name)
package spliit
