# goa learn


<a name="overview"></a>
## Overview
learning goa


### URI scheme
*Host* : localhost:8080  
*Schemes* : HTTP


### Consumes

* `application/json`
* `application/xml`
* `application/gob`
* `application/x-gob`


### Produces

* `application/json`
* `application/xml`
* `application/gob`
* `application/x-gob`




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




<a name="definitions"></a>
## Definitions

<a name="article"></a>
### Article
Article (default view)


|Name|Description|Schema|
|---|---|---|
|**created_at**  <br>*required*|作成日時  <br>**Example** : `"2016-03-20T09:39:48Z"`|string (date-time)|
|**created_by**  <br>*optional*|作成者  <br>**Example** : `1`|integer (int64)|
|**id**  <br>*required*|ID  <br>**Example** : `1`|integer (int64)|
|**text**  <br>*required*|text  <br>**Example** : `"Voluptatem ea."`|string|


<a name="articlecollection"></a>
### ArticleCollection
ArticleCollection is the media type for an array of Article (default view)

*Type* : < [Article](#article) > array


<a name="error"></a>
### error
Error response media type (default view)


|Name|Description|Schema|
|---|---|---|
|**code**  <br>*optional*|an application-specific error code, expressed as a string value.  <br>**Example** : `"invalid_value"`|string|
|**detail**  <br>*optional*|a human-readable explanation specific to this occurrence of the problem.  <br>**Example** : `"Value of ID must be an integer"`|string|
|**id**  <br>*optional*|a unique identifier for this particular occurrence of the problem.  <br>**Example** : `"3F1FKVRR"`|string|
|**meta**  <br>*optional*|a meta object containing non-standard meta-information about the error.  <br>**Example** : `{<br>  "timestamp" : 1458609066<br>}`|object|
|**status**  <br>*optional*|the HTTP status code applicable to this problem, expressed as a string value.  <br>**Example** : `"400"`|string|





