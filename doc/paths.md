
<a name="paths"></a>
## Paths

<a name="article-list"></a>
### list article
```
GET /article
```


#### Description
List


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Query**|**id**  <br>*optional*|Filter by ids|< integer > array|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK|[ArticleCollection](#articlecollection)|
|**400**|Bad Request|[error](#error)|
|**404**|Not Found|No Content|


#### Produces

* `application/vnd.goa.error`
* `application/vnd.article+json; type=collection`


#### Tags

* article


#### Example HTTP response

##### Response 200
```
json :
"[{\"created_at\":\"2016-03-20T09:39:48Z\",\"created_by\":1,\"id\":1,\"text\":\"Voluptatem ea.\"},{\"created_at\":\"2016-03-20T09:39:48Z\",\"created_by\":1,\"id\":1,\"text\":\"Voluptatem ea.\"},{\"created_at\":\"2016-03-20T09:39:48Z\",\"created_by\":1,\"id\":1,\"text\":\"Voluptatem ea.\"}]"
```


##### Response 400
```
json :
{
  "code" : "invalid_value",
  "detail" : "Value of ID must be an integer",
  "id" : "3F1FKVRR",
  "meta" : {
    "timestamp" : 1458609066
  },
  "status" : "400"
}
```


<a name="article-show"></a>
### show article
```
GET /article/{id}
```


#### Description
Get arcitle by id


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**id**  <br>*required*|id|integer|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK|[Article](#article)|
|**400**|Bad Request|[error](#error)|
|**404**|Not Found|No Content|


#### Produces

* `application/vnd.goa.error`
* `application/vnd.article+json`


#### Tags

* article


#### Example HTTP response

##### Response 200
```
json :
{
  "created_at" : "2016-03-20T09:39:48Z",
  "created_by" : 1,
  "id" : 1,
  "text" : "Voluptatem ea."
}
```


##### Response 400
```
json :
{
  "code" : "invalid_value",
  "detail" : "Value of ID must be an integer",
  "id" : "3F1FKVRR",
  "meta" : {
    "timestamp" : 1458609066
  },
  "status" : "400"
}
```



