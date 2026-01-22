# Unofficial Spliit api

This module implements the Spliit api, allowing you to manage you Spliit group from Go, C or C++.

## Installation

`go get go.chrastecky.dev/spliit-api`

## Usage

First create a client:

```go
client := spliit.NewClient()
```

And then get sending some requests:

```go
client := spliit.NewClient()
calls, err := client.SendRequests(
    context.Background(),
    spliit.NewCall(
        &endpoint.GetGroup{},
        &shape.GetGroupDetailsRequest{
            GroupID: "abcde",
        },
    ),
)
```

And then work with the response:

```go
response := calls[0].(spliit.OutputCall[shape.GetGroupDetailsResponse]).Output()
log.Println(response.Group.Name)
```

If you don't like the type assertion, you can hold a reference to the initial call:

```go
getGroupCall := spliit.NewCall(
    &endpoint.GetGroup{},
    &shape.GetGroupDetailsRequest{
        GroupID: groupID,
    },
)

_, err := client.SendRequests(
    context.Background(),
    getGroupCall,
)

if err != nil {
    log.Fatal(err)
}

log.Println(getGroupCall.Result.Group.Name)
```

### Multiple calls

You can send as many calls as you want in a single request:

```go
getGroupCall := spliit.NewCall(
    &endpoint.GetGroup{},
    &shape.GetGroupDetailsRequest{
        GroupID: groupID,
    },
)
listExpensesCall := spliit.NewCall(
    &endpoint.ListExpenses{},
    &shape.ListExpensesRequest{
        GroupID: groupID,
    },
)

_, err := client.SendRequests(
    context.Background(),
    getGroupCall,
    listExpensesCall,
)

if err != nil {
    log.Fatal(err)
}

log.Println(getGroupCall.Result.Group.Name)
log.Println(listExpensesCall.Result.Expenses[0].Title)
```

## Error handling

In addition to the SendRequests() method returning an error, each call might have errored out as well:

```go
_, err := client.SendRequests(
    context.Background(),
    getGroupCall,
)

if err != nil {
    log.Fatal(err)
}
if getGroupCall.Err != nil {
    log.Fatal(getGroupCall.Err)
}
```

Or if using the calls return value:

```go
calls, err := client.SendRequests(
    context.Background(),
    getGroupCall,
)

if err != nil {
    log.Fatal(err)
}

if err = calls[0].(spliit.OutputCall[shape.GetGroupDetailsResponse]).ErrValue(); err != nil {
    log.Fatal(err)
}
```

## Inputs

The input request might be anything that can be correctly mapped to the request, so all these are equally valid:

```go
getGroupCall := spliit.NewCall(
    &endpoint.GetGroup{},
    &shape.GetGroupDetailsRequest{
        GroupID: "abcde",
    },
)

getGroupCall := spliit.NewCall(
    &endpoint.GetGroup{},
    map[string]string{
        "groupId": "abcde",
    },
)

getGroupCall := spliit.NewCall(
    &endpoint.GetGroup{},
    `{"groupId": "abcde"}`,
)
```

## Endpoints

The endpoints describe the expected inputs and outputs:

- [CreateExpense](spliit/endpoint/create_expense.go)
- [CreateGroup](spliit/endpoint/create_group.go)
- [DeleteExpense](spliit/endpoint/delete_expense.go)
- [GetExpense](spliit/endpoint/get_expense.go)
- [GetGroup](spliit/endpoint/get_group.go)
- [GetGroupDetails](spliit/endpoint/get_group_details.go)
- [GetStats](spliit/endpoint/get_stats.go)
- [ListActivities](spliit/endpoint/list_activities.go)
- [ListBalances](spliit/endpoint/list_balances.go)
- [ListCategories](spliit/endpoint/list_categories.go)
- [ListExpenses](spliit/endpoint/list_expenses.go)
- [ListGroups](spliit/endpoint/list_groups.go)
- [UpdateExpense](spliit/endpoint/update_expense.go)
- [UpdateGroup](spliit/endpoint/update_group.go)

## C bindings

In the [cbindings](cbindings) directory you can find exported C bindings that make it possible to use the library from
C/C++.

### Building

Run `make build-lib` and a new `build` directory will get created which will contain `libspliit.h` and `libspliit.so`.

The following C functions are exported:

```c
extern int Spliit_NewClient(uint64_t* outHandle);
extern int Spliit_SendRequests(uint64_t clientHandle, char* jsonCalls, SpliitResult** outResults, size_t* outCount);
extern void Spliit_FreeResults(SpliitResult* results, size_t count);
extern size_t Spliit_GetLastError(char* buf, size_t bufLen);
extern int Spliit_CloseHandle(uint64_t handle);
```

The usage is (in Qt, adapt to your favourite brand of strings etc.):

```c++
uint64_t spliit_handle;
auto result = Spliit_NewClient(&spliit_handle);
if (result != SPLIIT_SUCCESS) {
    std::size_t len = Spliit_GetLastError(nullptr, 0);
    QByteArray buf(static_cast<int>(len), Qt::Uninitialized);
    Spliit_GetLastError(buf.data(), static_cast<std::size_t>(buf.size()));
    auto error = QString::fromUtf8(buf.constData()); 
    // todo handle error
}

SpliitResult* results = nullptr;
size_t count = 0;

result = Spliit_SendRequests(
    spliit_handle,
    "{\"endpoint\": \"groups.getDetails\", \"input\": {\"groupId\": \"abcde\"}}",
    &results,
    &count
)
if (result != SPLIIT_SUCCESS) {
    // handle error
}

SpliitResult &res = results[0];
if (res.error != nullptr) {
    // handle error
}

const char* resultJson = res.result;
// do something with the result

Spliit_FreeResults(results, count);
```

The JSON that gets sent to the `Spliit_SendRequests` function is as follows:

```json5
{
  "endpoint": "<endpoint-name>",
  "input": { // the input will be deserialized directly into the relevant Go struct
    
  }
}
```

The endpoint names are taken directly from the endpoint structs:

- [groups.expenses.create](spliit/endpoint/create_expense.go)
- [groups.create](spliit/endpoint/create_group.go)
- [groups.expenses.delete](spliit/endpoint/delete_expense.go)
- [groups.expenses.get](spliit/endpoint/get_expense.go)
- [groups.get](spliit/endpoint/get_group.go)
- [groups.getDetails](spliit/endpoint/get_group_details.go)
- [groups.stats.get](spliit/endpoint/get_stats.go)
- [groups.activities.list](spliit/endpoint/list_activities.go)
- [groups.balances.list](spliit/endpoint/list_balances.go)
- [categories.list](spliit/endpoint/list_categories.go)
- [groups.expenses.list](spliit/endpoint/list_expenses.go)
- [groups.list](spliit/endpoint/list_groups.go)
- [groups.expenses.update](spliit/endpoint/update_expense.go)
- [groups.update](spliit/endpoint/update_group.go)
