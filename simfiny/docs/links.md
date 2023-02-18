# Links

```http title="HTTP"
BASE URI - /api/v1/financial-service/links
```

## Get All Links 
Lists all of the links for the currently authenticated user.

```http title="HTTP"
GET - /api/v1/financial-service/links/user/{userID}
```

## Get A Specific Link
Retrieve a single specific link using the link's unique Id.

```http title="HTTP"
GET - /api/v1/financial-service/links/link/{linkId}
```

## Create A Manual Link 
Creates a manual link

```http title="HTTP"
POST - /api/v1/financial-service/links
```

### Request
```json 
"institutionName": ""
"customInstitutionName": ""
"description": ""
```

## Update A Link
Update an existing link.

```http title="HTTP"
PUT - /api/v1/financial-service/links/link/{linkId}
```

## Convert A Link To A Manual One 
Convert an existing link into a manual one.

```http title="HTTP"
PUT - /api/v1/financial-service/links/link/{linkId}/convert
```


## Delete A link 
Remove a link from your account. This will remove
- All bank accounts associated with this link.
- All spending objects associated with each of those bank accounts.
- All transactions for the those bank accounts.
This cannot be undone and data cannot be recovered.
If the link specified is a Plaid link, then the access_token associated with that link will also be
revoked. Link data is deleted in the background, so if you need to "wait" for all of the link's data to
be properly deleted. Then you should poll the `/link/wait` endpoint.

```http title="HTTP"
DELETE - /api/v1/financial-service/links/link/{linkId}
```

## Wait For A Deleted Link 
This endpoint is used to "wait" for all of the data associated with a link to be deleted. If the link is
is already deleted then a simple **200** is returned to the caller. If the link is not deleted then this
endpoint will block for up to 30 seconds at a time while it waits for the link to be removed. If it is
removed while the endpoint is blocking then it will return 200 at that time.

```http title="HTTP"
DELETE - /api/v1/financial-service/links/link/{linkId}/wait
```
