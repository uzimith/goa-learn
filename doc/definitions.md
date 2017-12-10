
<a name="definitions"></a>
## Definitions

<a name="article"></a>
### Article
Article (default view)


|Name|Description|Schema|
|---|---|---|
|**created_by**  <br>*optional*|作成者  <br>**Example** : `1`|integer (int64)|
|**id**  <br>*required*|ID  <br>**Example** : `1`|integer (int64)|
|**text**  <br>*required*|text  <br>**Example** : `"Iure voluptatem."`|string|


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



