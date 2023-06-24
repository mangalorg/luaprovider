# aes

 AES encryption and decryption.

## decrypt


```lua
function aes.decrypt(key: string, value: string)
  -> decrypted: string
```

 Decrypts a string using AES.

@*param* `key` — The key to use for decryption.

@*param* `value` — The string to decrypt.

@*return* `decrypted` — The decrypted string.

## encrypt


```lua
function aes.encrypt(key: string, value: string)
  -> encrypted: string
```

 Encrypts a string using AES.

@*param* `key` — The key to use for encryption.

@*param* `value` — The string to encrypt.

@*return* `encrypted` — The encrypted string.


---

# arg


Command-line arguments of Lua Standalone.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-arg)



```lua
string[]
```


---

# assert


Raises an error if the value of its argument v is false (i.e., `nil` or `false`); otherwise, returns all its arguments. In case of error, `message` is the error object; when absent, it defaults to `"assertion failed!"`

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-assert)


```lua
function assert(v?: <T>, message?: any, ...any)
  -> <T>
  2. ...any
```


---

# base64

 Base64 encoding and decoding.

## decode


```lua
function base64.decode(value: string, encoding?: userdata)
  -> decoded: string
```

 Decodes a base64 string.

@*param* `value` — The base64 string to decode.

@*param* `encoding` — The encoding to use. Defaults to `std_encoding`.

@*return* `decoded` — The decoded string.

## encode


```lua
function base64.encode(value: string, encoding?: userdata)
  -> encoded: string
```

 Encodes a string to base64.

@*param* `value` — The string to encode.

@*param* `encoding` — The encoding to use. Defaults to `std_encoding`.

@*return* `encoded` — The encoded string.

## raw_std_encoding


```lua
userdata
```

The standard raw, unpadded base64 encoding, as defined in RFC 4648.

## raw_url_encoding


```lua
userdata
```

The alternate raw, unpadded base64 encoding defined in RFC 4648. It is typically used in URLs and file names.

## std_encoding


```lua
userdata
```

The standard base64 encoding, as defined in RFC 4648.

## url_encoding


```lua
userdata
```

The alternate base64 encoding defined in RFC 4648. It is typically used in URLs and file names.


---

# base64_encoding


```lua
userdata
```


---

# collectgarbage


This function is a generic interface to the garbage collector. It performs different functions according to its first argument, `opt`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-collectgarbage)


```lua
opt:
   -> "collect" -- Performs a full garbage-collection cycle.
    | "stop" -- Stops automatic execution.
    | "restart" -- Restarts automatic execution.
    | "count" -- Returns the total memory in Kbytes.
    | "step" -- Performs a garbage-collection step.
    | "isrunning" -- Returns whether the collector is running.
    | "setpause" -- Set `pause`.
    | "setstepmul" -- Set `step multiplier`.
```


```lua
function collectgarbage(opt?: "collect"|"count"|"isrunning"|"restart"|"setpause"...(+3), arg?: integer)
  -> any
```


---

# coroutine




[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-coroutine)



```lua
coroutinelib
```


---

# coroutine.close


Closes coroutine `co` , closing all its pending to-be-closed variables and putting the coroutine in a dead state.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-coroutine.close)


```lua
function coroutine.close(co: thread)
  -> noerror: boolean
  2. errorobject: any
```


---

# coroutine.create


Creates a new coroutine, with body `f`. `f` must be a function. Returns this new coroutine, an object with type `"thread"`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-coroutine.create)


```lua
function coroutine.create(f: fun(...any):...unknown)
  -> thread
```


---

# coroutine.isyieldable


Returns true when the running coroutine can yield.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-coroutine.isyieldable)


```lua
function coroutine.isyieldable()
  -> boolean
```


---

# coroutine.resume


Starts or continues the execution of coroutine `co`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-coroutine.resume)


```lua
function coroutine.resume(co: thread, val1?: any, ...any)
  -> success: boolean
  2. ...any
```


---

# coroutine.running


Returns the running coroutine plus a boolean, true when the running coroutine is the main one.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-coroutine.running)


```lua
function coroutine.running()
  -> running: thread
  2. ismain: boolean
```


---

# coroutine.status


Returns the status of coroutine `co`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-coroutine.status)


```lua
return #1:
    | "running" -- Is running.
    | "suspended" -- Is suspended or not started.
    | "normal" -- Is active but not running.
    | "dead" -- Has finished or stopped with an error.
```


```lua
function coroutine.status(co: thread)
  -> "dead"|"normal"|"running"|"suspended"
```


---

# coroutine.wrap


Creates a new coroutine, with body `f`; `f` must be a function. Returns a function that resumes the coroutine each time it is called.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-coroutine.wrap)


```lua
function coroutine.wrap(f: fun(...any):...unknown)
  -> fun(...any):...unknown
```


---

# coroutine.yield


Suspends the execution of the calling coroutine.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-coroutine.yield)


```lua
(async) function coroutine.yield(...any)
  -> ...any
```


---

# crypto

 Various cryptographic functions.

## aes


```lua
aes
```

AES encryption and decryption.

## md5


```lua
md5
```

MD5 cryptographic hash function.

## sha1


```lua
sha1
```

SHA1 cryptographic hash function.

## sha256


```lua
sha256
```

SHA256 cryptographic hash function.

## sha512


```lua
sha512
```

SHA-512 cryptographic hash function.


---

# dofile


Opens the named file and executes its content as a Lua chunk. When called without arguments, `dofile` executes the content of the standard input (`stdin`). Returns all values returned by the chunk. In case of errors, `dofile` propagates the error to its caller. (That is, `dofile` does not run in protected mode.)

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-dofile)


```lua
function dofile(filename?: string)
  -> ...any
```


---

# encoding


## base64


```lua
base64
```

Base64 encoding and decoding.

## json


```lua
json
```

Provides functions for encoding and decoding JSON.


---

# error


Terminates the last protected function called and returns message as the error object.

Usually, `error` adds some information about the error position at the beginning of the message, if the message is a string.


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-error)


```lua
function error(message: any, level?: integer)
```


---

# getfenv


Returns the current environment in use by the function. `f` can be a Lua function or a number that specifies the function at that stack level.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-getfenv)


```lua
function getfenv(f?: integer|fun(...any):...unknown)
  -> table
```


---

# getmetatable


If object does not have a metatable, returns nil. Otherwise, if the object's metatable has a __metatable field, returns the associated value. Otherwise, returns the metatable of the given object.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-getmetatable)


```lua
function getmetatable(object: any)
  -> metatable: table
```


---

# headless

 Headless browser

## browser


```lua
function headless.browser()
  -> browser: headless_browser
```

 Creates a new headless browser

@*return* `browser` — headless browser instance


---

# headless_browser


## page


```lua
(method) headless_browser:page(url: string)
  -> page: headless_page
```

 Visit the page

@*param* `url` — url of the page to visit


---

# headless_element


## html


```lua
(method) headless_element:html()
  -> html: string
```


---

# headless_page


## click


```lua
(method) headless_page:click()
```

## html


```lua
(method) headless_page:html()
  -> html: string
```

## input


```lua
(method) headless_page:input(text: string)
```


---

# html

 This library provides functions for parsing HTML and querying it using CSS selectors. It is based on [goquery](https://github.com/PuerkitoBio/goquery).

## parse


```lua
function html.parse(html: string)
  -> document: html_document
```

 Parses the given HTML and returns a selection containing the root element.

@*param* `html` — The HTML to parse.

@*return* `document` — The document object.


---

# html_document

 Document represents an HTML document to be manipulated. Unlike jQuery, which is loaded as part of a DOM document, and thus acts upon its containing document, GoQuery doesn't know which HTML document to act upon. So it needs to be told, and that's what the document class is for. It holds the root document node to manipulate, and can make selections on this document.

## find


```lua
(method) html_document:find(selector: string)
  -> selection: html_selection
```

 Finds all elements that match the selector string. It returns a new selection object with the matched elements.

@*param* `selector` — The CSS selector to use to find the elements.

@*return* `selection` — A selection object.

## html


```lua
(method) html_document:html()
  -> html: string
```

 Gets the HTML contents of the first element in the set of matched elements. It includes text and comment nodes.

@*return* `html` — The HTML contents of the first element in the set of matched elements.

## markdown


```lua
(method) html_document:markdown()
  -> markdown: string
```

 Converts the document to Markdown. Can be used to show the contents of a document in info page

@*return* `markdown` — The Markdown representation of the document.

## selection


```lua
(method) html_document:selection()
  -> selection: html_selection
```

 Converts document to a selection object.

@*return* `selection` — A selection object.

## simplified


```lua
(method) html_document:simplified()
  -> html: html_document
```

 Gets the readable part of the document (simplified view). Similar to reader mode in browsers.

@*return* `html` — The simplified document


---

# html_selection


## add


```lua
(method) html_selection:add(selector: string)
  -> selection: html_selection
```

 Add adds the selector string's matching nodes to those in the current selection and returns a new selection object. The selector string is run in the context of the document of the current selection object.

@*param* `selector` — The CSS selector to use to find the elements.

@*return* `selection` — A selection object.

## add_back


```lua
(method) html_selection:add_back(selection: html_selection)
  -> selection: html_selection
```

 Adds the specified Selection object's nodes to those in the current selection and returns a new Selection object.

@*param* `selection` — A selection object.

@*return* `selection` — A selection object.

## add_class


```lua
(method) html_selection:add_class(class: string)
```

 Adds the specified class(es) to each of the set of matched elements.

@*param* `class` — One or more class names to be added to the class attribute of each matched element.

## add_selection


```lua
(method) html_selection:add_selection(selection: html_selection)
  -> selection: html_selection
```

 Adds the specified selection object's nodes to those in the current selection and returns a new selection object.

@*param* `selection` — A selection object.

@*return* `selection` — A selection object.

## attr


```lua
(method) html_selection:attr(name: string)
  -> value: string
  2. ok: boolean
```

 Returns the value of the given attribute of the first element in the selection.

@*param* `name` — The name of the attribute to get.

@*return* `value` — The value of the attribute, or nil if the attribute is not present.

@*return* `ok` — Whether the attribute was present.

## attr_or


```lua
(method) html_selection:attr_or(name: string, default: string)
  -> value: string
```

 Returns the value of the given attribute of the first element in the selection, or a default value if the attribute is not present.

@*param* `name` — The name of the attribute to get.

@*param* `default` — The default value to return if the attribute is not present.

@*return* `value` — The value of the attribute, or the default value if the attribute is absent.

## children


```lua
(method) html_selection:children()
  -> selection: html_selection
```

 Gets all direct child elements of each element in the Selection. It returns a new selection object containing the matched elements.

@*return* `selection` — A selection object.

## closest


```lua
(method) html_selection:closest(selector: string)
  -> selection: html_selection
```

 Gets the first element that matches the selector by testing the element itself and traversing up through its ancestors in the DOM tree.

@*param* `selector` — The CSS selector to use to find the elements.

@*return* `selection` — A selection object.

## contents


```lua
(method) html_selection:contents()
  -> selection: html_selection
```

 Contents gets the children of each element in the selection, including text and comment nodes. It returns a new selection object containing these elements

@*return* `selection` — A selection object.

## each


```lua
(method) html_selection:each(fn: fun(selection: html_selection, index: number))
```

 Iterates over all elements in the selection, calling the given function for each one.

@*param* `fn` — The function to call for each element.

## eq


```lua
(method) html_selection:eq(index: number)
  -> selection: html_selection
```

 Reduces the set of matched elements to the one at the specified index. If a negative index is given, it counts backwards starting at the end of the set. It returns a new Selection object, and an empty Selection object if the index is invalid.

@*param* `index` — The index of the element to select.

@*return* `selection` — A selection object.

## filter


```lua
(method) html_selection:filter(selector: string)
  -> selection: html_selection
```

 Filter reduces the set of matched elements to those that match the selector string. It returns a new Selection object for this subset of matching elements.

@*param* `selector` — The CSS selector to use to find the elements.

@*return* `selection` — A selection object.

## find


```lua
(method) html_selection:find(selector: string)
  -> selection: html_selection
```

 Finds all elements matching the given selector.

@*param* `selector` — The CSS selector to use to find the elements.

@*return* `selection` — A selection object.

## find_selection


```lua
(method) html_selection:find_selection(selection: html_selection)
  -> selection: html_selection
```

 gets the descendants of each element in the current selection, filtered by a selection. It returns a new selection object containing these matched elements.

@*param* `selection` — A selection object.

@*return* `selection` — A selection object.

## first


```lua
(method) html_selection:first()
  -> selection: html_selection
```

 Reduces the set of matched elements to the first in the set. It returns a new selection object, and an empty selection object if the the selection is empty.

@*return* `selection` — A selection object.

## has_class


```lua
(method) html_selection:has_class(class: string)
  -> ok: boolean
```

 determines whether any of the matched elements are assigned the given class.

@*param* `class` — The class to check for.

@*return* `ok` — Whether any of the matched elements have the given class.

## html


```lua
(method) html_selection:html()
  -> html: string
```

 Returns the HTML contents of the first element in the selection.

@*return* `html` — Gets the HTML contents of the first element in the set of matched elements. It includes text and comment nodes.

## is


```lua
(method) html_selection:is(selector: string)
  -> ok: boolean
```

 Checks the current matched set of elements against a selector and returns true if at least one of these elements matches.

@*param* `selector` — The CSS selector to use to find the elements.

@*return* `ok` — Whether any of the matched elements match the selector.

## last


```lua
(method) html_selection:last()
  -> selection: html_selection
```

 Reduces the set of matched elements to the last in the set. It returns a new selection object, and an empty selection object if the the selection is empty.

@*return* `selection` — A selection object.

## length


```lua
(method) html_selection:length()
  -> n: number
```

 Returns the number of elements in the selection.

@*return* `n` — The number of elements in the selection.

## map


```lua
(method) html_selection:map(fn: fun(selection: html_selection, index: number):any)
  -> results: any[]
```

 Iterates over a selection, executing a function for each matched element. The function's return value is added to the returned table.

@*param* `fn` — The function to execute for each element. It receives the index of the element in the selection and the element as arguments.

@*return* `results` — A table containing the return values of the function for each element.

## markdown


```lua
(method) html_selection:markdown()
  -> markdown: string
```

 Converts the selection to Markdown. Can be used to show the contents of a selection in info page

@*return* `markdown` — The Markdown representation of the selection.

## next


```lua
(method) html_selection:next()
  -> selection: html_selection
```

 Gets the immediately following sibling of each element in the set of matched elements, optionally filtered by a selector. It returns a new Selection object containing the matched elements.

@*return* `selection` — A selection object.

## next_all


```lua
(method) html_selection:next_all()
```

 Gets all the following siblings of each element in the Selection. It returns a new Selection object containing the matched elements.

## next_until


```lua
(method) html_selection:next_until(selector: string)
  -> selection: html_selection
```

 gets all following siblings of each element up to but not including the element matched by the selector. It returns a new Selection object containing the matched elements.

@*param* `selector` — The CSS selector to use to find the elements.

@*return* `selection` — A selection object.

## parent


```lua
(method) html_selection:parent()
  -> selection: html_selection
```

 Gets the parent of each element in the Selection. It returns a new Selection object containing the matched elements.

@*return* `selection` — A selection object.

## parents


```lua
(method) html_selection:parents()
  -> selection: html_selection
```

 Gets the ancestors of each element in the current Selection. It returns a new Selection object with the matched elements.

@*return* `selection` — A selection object.

## parents_until


```lua
(method) html_selection:parents_until(selector: string)
  -> selection: html_selection
```

 Gets the ancestors of each element in the current Selection, up to but not including the element matched by the selector. It returns a new Selection object with the matched elements.

@*param* `selector` — The CSS selector to use to find the elements.

@*return* `selection` — A selection object.

## prev


```lua
(method) html_selection:prev()
  -> selection: html_selection
```

 Gets the immediately preceding sibling of each element in the Selection. It returns a new selection object containing the matched elements.

@*return* `selection` — A selection object.

## prev_all


```lua
(method) html_selection:prev_all()
  -> selection: html_selection
```

 Gets all the preceding siblings of each element in the Selection. It returns a new selection object containing the matched elements.

@*return* `selection` — A selection object.

## prev_until


```lua
(method) html_selection:prev_until(selector: string)
  -> selection: html_selection
```

 Gets all preceding siblings of each element up to but not including the element matched by the selector. It returns a new selection object containing the matched elements.

@*param* `selector` — The CSS selector to use to find the elements.

@*return* `selection` — A selection object.

## remove


```lua
(method) html_selection:remove(selector: string)
  -> selection: html_selection
```

 Removes elements from the selection that match the selector string. It returns a new selection object with the matching elements removed.

@*param* `selector` — The CSS selector to use to find the elements.

@*return* `selection` — A selection object.

## remove_class


```lua
(method) html_selection:remove_class(class: string)
```

 Removes the given class(es) from each element in the set of matched elements. Multiple class names can be specified, separated by a space or via multiple arguments. If no class name is provided, all classes are removed.

@*param* `class` — The class to remove. If not provided, all classes are removed.

## siblings


```lua
(method) html_selection:siblings()
  -> selection: html_selection
```

 Gets all sibling elements of each element in the Selection. It returns a new selection object containing the matched elements.

@*return* `selection` — A selection object.

## simplified


```lua
(method) html_selection:simplified()
  -> selection: html_selection
```

 Gets the readable part of the selection (simplified view). Similar to reader mode in browsers.

@*return* `selection` — A selection object.

## slice


```lua
(method) html_selection:slice(start: number, finish: number)
  -> selection: html_selection
```

 Returns a selection containing a subset of the elements in the original selection. It returns a new selection object with the matched elements.

@*param* `start` — The index of the first element to include in the new selection.

@*param* `finish` — The index of the first element to exclude from the new selection.

@*return* `selection` — A selection object.

## terminate


```lua
(method) html_selection:terminate()
  -> selection: html_selection
```

 Ends the most recent filtering operation in the current chain and returns the set of matched elements to its previous state.

@*return* `selection` — A selection object.

## text


```lua
(method) html_selection:text()
  -> text: string
```

 Returns the text contents of the first element in the selection.

@*return* `text` — The text contents of the first element in the selection.

## toggle_class


```lua
(method) html_selection:toggle_class(class: string)
```

 Adds or removes the given class(es) for each element in the set of matched elements. Multiple class names can be specified, separated by a space or via multiple arguments.

@*param* `class` — The class to toggle.


---

# http

 Package http provides HTTP client implementations. Make HTTP (or HTTPS) requests

## METHOD_CONNECT


```lua
string
```

CONNECT HTTP Method

## METHOD_DELETE


```lua
string
```

DELETE HTTP Method

## METHOD_GET


```lua
string
```

GET HTTP Method

## METHOD_HEAD


```lua
string
```

HEAD HTTP Method

## METHOD_PATCH


```lua
string
```

PATCH HTTP Method

## METHOD_POST


```lua
string
```

POST HTTP Method

## METHOD_PUT


```lua
string
```

PUT HTTP Method

## STATUS_OK


```lua
number
```

Returned if the request was successful

## request


```lua
function http.request(method: 'CONNECT'|'DELETE'|'GET'|'HEAD'|'PATCH'...(+2), url: string, body?: string)
  -> request: http_request
```

 Create a new HTTP request

@*param* `method` — HTTP method

@*param* `url` — URL

@*param* `body` — Request body

@*return* `request` — Request

```lua
method:
    | 'GET'
    | 'HEAD'
    | 'POST'
    | 'PUT'
    | 'PATCH'
    | 'DELETE'
    | 'CONNECT'
```


---

# http_request

 HTTP Request

## content_length


```lua
(method) http_request:content_length(length?: number)
  -> length: number?
```

 Get or set request content length

@*param* `length` — Content length

@*return* `length` — Content length

## cookie


```lua
(method) http_request:cookie(key: string, value?: string)
  -> value: string?
```

 Get or set request cookie

@*param* `key` — Cookie key

@*param* `value` — Cookie value

@*return* `value` — Cookie value

## header


```lua
(method) http_request:header(key: string, value?: string)
  -> value: string?
```

 Get or set request header

@*param* `key` — Header key

@*param* `value` — Header value

@*return* `value` — Header value

## send


```lua
(method) http_request:send()
  -> response: http_response
```

 Perform request

@*return* `response` — Response


---

# http_response

 HTTP Response

## body


```lua
(method) http_response:body()
  -> body: string
```

 Get response body

@*return* `body` — Body

## content_length


```lua
(method) http_response:content_length()
  -> length: number
```

 Get response content length

@*return* `length` — Content length

## cookies


```lua
(method) http_response:cookies()
  -> cookies: { name: string, value: string }[]
```

 Get response cookies. Returns a list of cookies.

@*return* `cookies` — Cookies

## header


```lua
(method) http_response:header(key: string)
  -> value: string
```

 Get response header

@*param* `key` — Header key

@*return* `value` — Header value

## status


```lua
(method) http_response:status()
  -> status: number
```

 Get response status

@*return* `status` — Status


---

# ipairs


Returns three values (an iterator function, the table `t`, and `0`) so that the construction
```lua
    for i,v in ipairs(t) do body end
```
will iterate over the key–value pairs `(1,t[1]), (2,t[2]), ...`, up to the first absent index.


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-ipairs)


```lua
function ipairs(t: <T:table>)
  -> fun(table: <V>[], i?: integer):integer, <V>
  2. <T:table>
  3. i: integer
```


---

# js

 JavaScript execution.

## vm


```lua
function js.vm()
  -> vm: js_vm
```

 Creates a new JavaScript virtual machine.

@*return* `vm` — The new JavaScript virtual machine.


---

# js_vm

 A JavaScript virtual machine. This is used to execute JavaScript code.

## get


```lua
(method) js_vm:get(name: string)
  -> value: js_vm_value
```

 Gets the value of the given property on the global object.

@*param* `name` — The name of the property.

@*return* `value` — The value of the property.

## run


```lua
(method) js_vm:run(code: string)
  -> value: js_vm_value
```

 Runs the given JavaScript code.

@*param* `code` — The JavaScript code to run.

@*return* `value` — The value returned by the code.

## set


```lua
(method) js_vm:set(name: string, value: any)
```

 Sets the value of the given property on the global object. It will convert the given Lua value to a JavaScript value.

@*param* `name` — The name of the property.

@*param* `value` — The value to set.


---

# js_vm_value

 A value returned from a JavaScript VM.

## export


```lua
(method) js_vm_value:export()
  -> value: any
```

 Exports the value to a Lua value.

@*return* `value` — The exported value.

## is_boolean


```lua
(method) js_vm_value:is_boolean()
  -> ok: boolean
```

 Returns whether the value is a boolean.

@*return* `ok` — Whether the value is a boolean.

## is_function


```lua
(method) js_vm_value:is_function()
  -> ok: boolean
```

 Returns whether the value is a function.

@*return* `ok` — Whether the value is a function.

## is_nan


```lua
(method) js_vm_value:is_nan()
  -> ok: boolean
```

 Returns whether the value is NaN.

@*return* `ok` — Whether the value is NaN.

## is_null


```lua
(method) js_vm_value:is_null()
  -> ok: boolean
```

 Returns whether the value is null.

@*return* `ok` — Whether the value is null.

## is_number


```lua
(method) js_vm_value:is_number()
  -> ok: boolean
```

 Returns whether the value is a number.

@*return* `ok` — Whether the value is a number.

## is_object


```lua
(method) js_vm_value:is_object()
  -> ok: boolean
```

 Returns whether the value is an object.

@*return* `ok` — Whether the value is an object.

## is_string


```lua
(method) js_vm_value:is_string()
  -> ok: boolean
```

 Returns whether the value is a string.

@*return* `ok` — Whether the value is a string.

## is_undefined


```lua
(method) js_vm_value:is_undefined()
  -> ok: boolean
```

 Returns whether the value is undefined.

@*return* `ok` — Whether the value is undefined.

## to_string


```lua
(method) js_vm_value:to_string()
  -> string: string
```

 Converts the value to a string.

@*return* `string` — The string representation of the value.


---

# json

 Provides functions for encoding and decoding JSON.

## decode


```lua
function json.decode(json: string)
  -> value: any
```

 Decodes a JSON string into a Lua value.

@*param* `json` — The JSON string to decode.

@*return* `value` — The decoded value.

## encode


```lua
function json.encode(value: any)
  -> json: string
```

 Encodes a Lua value into a JSON string.

@*param* `value` — The value to encode.

@*return* `json` — The encoded JSON string.


---

# levenshtein

 Levenshtein distance algorithm

## distance


```lua
function levenshtein.distance(s1: string, s2: string)
  -> distance: number
```

 Compute Levenshtein distance between two strings

@*return* `distance` — Levenshtein distance between s1 and s2


---

# load


Loads a chunk.

If `chunk` is a string, the chunk is this string. If `chunk` is a function, `load` calls it repeatedly to get the chunk pieces. Each call to `chunk` must return a string that concatenates with previous results. A return of an empty string, `nil`, or no value signals the end of the chunk.


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-load)


```lua
mode:
    | "b" -- Only binary chunks.
    | "t" -- Only text chunks.
   -> "bt" -- Both binary and text.
```


```lua
function load(chunk: string|function, chunkname?: string, mode?: "b"|"bt"|"t", env?: table)
  -> function?
  2. error_message: string?
```


---

# loadfile


Loads a chunk from file `filename` or from the standard input, if no file name is given.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-loadfile)


```lua
mode:
    | "b" -- Only binary chunks.
    | "t" -- Only text chunks.
   -> "bt" -- Both binary and text.
```


```lua
function loadfile(filename?: string, mode?: "b"|"bt"|"t", env?: table)
  -> function?
  2. error_message: string?
```


---

# loadstring


Loads a chunk from the given string.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-loadstring)


```lua
function loadstring(text: string, chunkname?: string)
  -> function?
  2. error_message: string?
```


---

# math




[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math)



```lua
mathlib
```


---

# math.abs


Returns the absolute value of `x`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.abs)


```lua
function math.abs(x: <Number:number>)
  -> <Number:number>
```


---

# math.acos


Returns the arc cosine of `x` (in radians).

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.acos)


```lua
function math.acos(x: number)
  -> number
```


---

# math.asin


Returns the arc sine of `x` (in radians).

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.asin)


```lua
function math.asin(x: number)
  -> number
```


---

# math.atan


Returns the arc tangent of `x` (in radians).

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.atan)


```lua
function math.atan(y: number)
  -> number
```


---

# math.atan2


Returns the arc tangent of `y/x` (in radians).

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.atan2)


```lua
function math.atan2(y: number, x: number)
  -> number
```


---

# math.ceil


Returns the smallest integral value larger than or equal to `x`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.ceil)


```lua
function math.ceil(x: number)
  -> integer
```


---

# math.cos


Returns the cosine of `x` (assumed to be in radians).

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.cos)


```lua
function math.cos(x: number)
  -> number
```


---

# math.cosh


Returns the hyperbolic cosine of `x` (assumed to be in radians).

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.cosh)


```lua
function math.cosh(x: number)
  -> number
```


---

# math.deg


Converts the angle `x` from radians to degrees.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.deg)


```lua
function math.deg(x: number)
  -> number
```


---

# math.exp


Returns the value `e^x` (where `e` is the base of natural logarithms).

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.exp)


```lua
function math.exp(x: number)
  -> number
```


---

# math.floor


Returns the largest integral value smaller than or equal to `x`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.floor)


```lua
function math.floor(x: number)
  -> integer
```


---

# math.fmod


Returns the remainder of the division of `x` by `y` that rounds the quotient towards zero.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.fmod)


```lua
function math.fmod(x: number, y: number)
  -> number
```


---

# math.frexp


Decompose `x` into tails and exponents. Returns `m` and `e` such that `x = m * (2 ^ e)`, `e` is an integer and the absolute value of `m` is in the range [0.5, 1) (or zero when `x` is zero).

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.frexp)


```lua
function math.frexp(x: number)
  -> m: number
  2. e: number
```


---

# math.ldexp


Returns `m * (2 ^ e)` .

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.ldexp)


```lua
function math.ldexp(m: number, e: number)
  -> number
```


---

# math.log


Returns the logarithm of `x` in the given base.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.log)


```lua
function math.log(x: number, base?: integer)
  -> number
```


---

# math.log10


Returns the base-10 logarithm of x.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.log10)


```lua
function math.log10(x: number)
  -> number
```


---

# math.max


Returns the argument with the maximum value, according to the Lua operator `<`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.max)


```lua
function math.max(x: <Number:number>, ...<Number:number>)
  -> <Number:number>
```


---

# math.min


Returns the argument with the minimum value, according to the Lua operator `<`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.min)


```lua
function math.min(x: <Number:number>, ...<Number:number>)
  -> <Number:number>
```


---

# math.modf


Returns the integral part of `x` and the fractional part of `x`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.modf)


```lua
function math.modf(x: number)
  -> integer
  2. number
```


---

# math.pow


Returns `x ^ y` .

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.pow)


```lua
function math.pow(x: number, y: number)
  -> number
```


---

# math.rad


Converts the angle `x` from degrees to radians.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.rad)


```lua
function math.rad(x: number)
  -> number
```


---

# math.random


* `math.random()`: Returns a float in the range [0,1).
* `math.random(n)`: Returns a integer in the range [1, n].
* `math.random(m, n)`: Returns a integer in the range [m, n].


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.random)


```lua
function math.random(m: integer, n: integer)
  -> integer
```


---

# math.randomseed


Sets `x` as the "seed" for the pseudo-random generator.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.randomseed)


```lua
function math.randomseed(x: integer)
```


---

# math.sin


Returns the sine of `x` (assumed to be in radians).

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.sin)


```lua
function math.sin(x: number)
  -> number
```


---

# math.sinh


Returns the hyperbolic sine of `x` (assumed to be in radians).

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.sinh)


```lua
function math.sinh(x: number)
  -> number
```


---

# math.sqrt


Returns the square root of `x`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.sqrt)


```lua
function math.sqrt(x: number)
  -> number
```


---

# math.tan


Returns the tangent of `x` (assumed to be in radians).

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.tan)


```lua
function math.tan(x: number)
  -> number
```


---

# math.tanh


Returns the hyperbolic tangent of `x` (assumed to be in radians).

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.tanh)


```lua
function math.tanh(x: number)
  -> number
```


---

# math.tointeger


If the value `x` is convertible to an integer, returns that integer.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.tointeger)


```lua
function math.tointeger(x: any)
  -> integer?
```


---

# math.type


Returns `"integer"` if `x` is an integer, `"float"` if it is a float, or `nil` if `x` is not a number.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.type)


```lua
return #1:
    | "integer"
    | "float"
    | 'nil'
```


```lua
function math.type(x: any)
  -> "float"|"integer"|'nil'
```


---

# math.ult


Returns `true` if and only if `m` is below `n` when they are compared as unsigned integers.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-math.ult)


```lua
function math.ult(m: integer, n: integer)
  -> boolean
```


---

# md5

 MD5 cryptographic hash function.

## sum


```lua
function md5.sum(value: string)
  -> hash: string
```

 Returns the MD5 hash of the given string.

@*param* `value` — The string to hash.

@*return* `hash` — The MD5 hash of the given string.


---

# module


Creates a module.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-module)


```lua
function module(name: string, ...any)
```


---

# newproxy


```lua
function newproxy(proxy: boolean|table|userdata)
  -> userdata
```


---

# next


Allows a program to traverse all fields of a table. Its first argument is a table and its second argument is an index in this table. A call to `next` returns the next index of the table and its associated value. When called with `nil` as its second argument, `next` returns an initial index and its associated value. When called with the last index, or with `nil` in an empty table, `next` returns `nil`. If the second argument is absent, then it is interpreted as `nil`. In particular, you can use `next(t)` to check whether a table is empty.

The order in which the indices are enumerated is not specified, *even for numeric indices*. (To traverse a table in numerical order, use a numerical `for`.)

The behavior of `next` is undefined if, during the traversal, you assign any value to a non-existent field in the table. You may however modify existing fields. In particular, you may set existing fields to nil.


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-next)


```lua
function next(table: table<<K>, <V>>, index?: <K>)
  -> <K>?
  2. <V>?
```


---

# package




[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-package)



```lua
packagelib
```


---

# package.config


A string describing some compile-time configurations for packages.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-package.config)



```lua
string
```


---

# package.loaders


A table used by `require` to control how to load modules.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-package.loaders)



```lua
table
```


---

# package.loadlib


Dynamically links the host program with the C library `libname`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-package.loadlib)


```lua
function package.loadlib(libname: string, funcname: string)
  -> any
```


---

# package.searchers


A table used by `require` to control how to load modules.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-package.searchers)



```lua
table
```


---

# package.searchpath


Searches for the given `name` in the given `path`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-package.searchpath)


```lua
function package.searchpath(name: string, path: string, sep?: string, rep?: string)
  -> filename: string?
  2. errmsg: string?
```


---

# package.seeall


Sets a metatable for `module` with its `__index` field referring to the global environment, so that this module inherits values from the global environment. To be used as an option to function `module` .

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-package.seeall)


```lua
function package.seeall(module: table)
```


---

# pairs


If `t` has a metamethod `__pairs`, calls it with t as argument and returns the first three results from the call.

Otherwise, returns three values: the [next](http://www.lua.org/manual/5.2/manual.html#pdf-next) function, the table `t`, and `nil`, so that the construction
```lua
    for k,v in pairs(t) do body end
```
will iterate over all key–value pairs of table `t`.

See function [next](http://www.lua.org/manual/5.2/manual.html#pdf-next) for the caveats of modifying the table during its traversal.


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-pairs)


```lua
function pairs(t: <T:table>)
  -> fun(table: table<<K>, <V>>, index?: <K>):<K>, <V>
  2. <T:table>
```


---

# pcall


Calls the function `f` with the given arguments in *protected mode*. This means that any error inside `f` is not propagated; instead, `pcall` catches the error and returns a status code. Its first result is the status code (a boolean), which is true if the call succeeds without errors. In such case, `pcall` also returns all results from the call, after this first result. In case of any error, `pcall` returns `false` plus the error object.


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-pcall)


```lua
function pcall(f: fun(...any):...unknown, arg1?: any, ...any)
  -> success: boolean
  2. result: any
  3. ...any
```


---

# print


Receives any number of arguments and prints their values to `stdout`, converting each argument to a string following the same rules of [tostring](http://www.lua.org/manual/5.2/manual.html#pdf-tostring).
The function print is not intended for formatted output, but only as a quick way to show a value, for instance for debugging. For complete control over the output, use [string.format](http://www.lua.org/manual/5.2/manual.html#pdf-string.format) and [io.write](http://www.lua.org/manual/5.2/manual.html#pdf-io.write).


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-print)


```lua
function print(...any)
```


---

# rawequal


Checks whether v1 is equal to v2, without invoking the `__eq` metamethod.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-rawequal)


```lua
function rawequal(v1: any, v2: any)
  -> boolean
```


---

# rawget


Gets the real value of `table[index]`, without invoking the `__index` metamethod.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-rawget)


```lua
function rawget(table: table, index: any)
  -> any
```


---

# rawlen


Returns the length of the object `v`, without invoking the `__len` metamethod.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-rawlen)


```lua
function rawlen(v: string|table)
  -> len: integer
```


---

# rawset


Sets the real value of `table[index]` to `value`, without using the `__newindex` metavalue. `table` must be a table, `index` any value different from `nil` and `NaN`, and `value` any Lua value.
This function returns `table`.


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-rawset)


```lua
function rawset(table: table, index: any, value: any)
  -> table
```


---

# re

 Compiled regular expression

## find_submatch


```lua
(method) re:find_submatch(text: string)
  -> matches: table
```

 Returns a slice of strings holding the text of the leftmost match of the regular
 expression in s and the matches, if any, of its subexpressions. A return value of empty table indicates no match.

@*param* `text` — A string to search in

@*return* `matches` — Found matches

## groups


```lua
(method) re:groups(text: string)
  -> groups: table
```

 Returns a table of all capture groups.

@*param* `text` — The text to split

@*return* `groups` — A table of all capture groups

## match


```lua
(method) re:match(text: string)
  -> matched: table
```

 Reports whether the string s contains any match of the regular expression.

@*param* `text` — A string to search in

@*return* `matched` — Whether the string s contains any match of the regular expression.

## replace_all


```lua
(method) re:replace_all(text: string, replacement: string)
  -> replaced: string
```

 Returns a copy of text, replacing matches of the regexp with the replacement string.
 Inside replacement, $ signs are interpreted as in expand, so for instance $1 represents the text of the first submatch.

@*param* `text` — A string to replace matches in

@*param* `replacement` — A string to replace matches with

@*return* `replaced` — The result of the replacement

## replace_all_func


```lua
(method) re:replace_all_func(text: string, replacer: fun(match: string):string)
  -> replaced: string
```

 Returns a copy of text, replacing matches of the regexp with the replacement function.

@*param* `text` — A string to replace matches in

@*param* `replacer` — A function to replace matches with

@*return* `replaced` — The result of the replacement

## split


```lua
(method) re:split(text: string)
  -> parts: table
```

 Splits the given text into a table of strings.

@*param* `text` — The text to split

@*return* `parts` — The result of the split


---

# regexp

 A regular expression library.

## compile


```lua
function regexp.compile(pattern: string)
  -> regexp: re
```

 Compiles the given pattern into a regular expression.

@*param* `pattern` — The pattern to compile

@*return* `regexp` — The compiled regular expression

## match


```lua
function regexp.match(pattern: string, text: string)
  -> matched: table
```

 Matches the given pattern against the given text.

@*param* `pattern` — The pattern to match

@*param* `text` — The text to match against

@*return* `matched` — A table of all matches

## urls_relaxed


```lua
userdata
```

Matches all the urls it can find.

## urls_strict


```lua
userdata
```

Only matches urls with a scheme to avoid false positives.


---

# require


Loads the given module, returns any value returned by the given module(`true` when `nil`).

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-require)


```lua
function require(modname: string)
  -> unknown
```


---

# sdk

 Contains various utilities for making HTTP requests, working with JSON, HTML, and more.

## crypto


```lua
crypto
```

Various cryptographic functions.

## encoding


```lua
encoding
```


## headless


```lua
headless
```

Headless browser

## html


```lua
html
```

This library provides functions for parsing HTML and querying it using CSS selectors. It is based on [goquery](https://github.com/PuerkitoBio/goquery).

## http


```lua
http
```

Package http provides HTTP client implementations. Make HTTP (or HTTPS) requests

## js


```lua
js
```

JavaScript execution.

## levenshtein


```lua
levenshtein
```

Levenshtein distance algorithm

## regexp


```lua
regexp
```

A regular expression library.

## strings


```lua
strings
```

Simple functions to manipulate UTF-8 encoded strings.

## time


```lua
time
```

Time library

## urls


```lua
urls
```

URLs is a library for working with URLs.

## util


```lua
util
```

Functional helpers


---

# select


If `index` is a number, returns all arguments after argument number `index`; a negative number indexes from the end (`-1` is the last argument). Otherwise, `index` must be the string `"#"`, and `select` returns the total number of extra arguments it received.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-select)


```lua
index:
    | "#"
```


```lua
function select(index: integer|"#", ...any)
  -> any
```


---

# setfenv


Sets the environment to be used by the given function.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-setfenv)


```lua
function setfenv(f: fun(...any):...integer|unknown, table: table)
  -> function
```


---

# setmetatable


Sets the metatable for the given table. If `metatable` is `nil`, removes the metatable of the given table. If the original metatable has a `__metatable` field, raises an error.

This function returns `table`.

To change the metatable of other types from Lua code, you must use the debug library ([§6.10](http://www.lua.org/manual/5.2/manual.html#6.10)).


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-setmetatable)


```lua
function setmetatable(table: table, metatable?: table)
  -> table
```


---

# sha1

 SHA1 cryptographic hash function.

## sum


```lua
function sha1.sum(value: string)
  -> hash: string
```

 Returns the SHA1 hash of the given string.

@*param* `value` — The string to hash.

@*return* `hash` — The SHA1 hash of the given string.


---

# sha256

 SHA256 cryptographic hash function.

## sum


```lua
function sha256.sum(value: string)
  -> hash: string
```

 Returns the SHA256 hash of the given string.

@*param* `value` — The string to hash.

@*return* `hash` — The SHA256 hash of the given string.


---

# sha512

 SHA-512 cryptographic hash function.

## sum


```lua
function sha512.sum(value: string)
  -> hash: string
```

 Returns the SHA-512 hash of the given string.

@*param* `value` — The string to hash.

@*return* `hash` — The SHA-512 hash of the given string.


---

# string




[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string)



```lua
stringlib
```


---

# string.byte


Returns the internal numeric codes of the characters `s[i], s[i+1], ..., s[j]`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.byte)


```lua
function string.byte(s: string|number, i?: integer, j?: integer)
  -> ...integer
```


---

# string.char


Returns a string with length equal to the number of arguments, in which each character has the internal numeric code equal to its corresponding argument.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.char)


```lua
function string.char(byte: integer, ...integer)
  -> string
```


---

# string.dump


Returns a string containing a binary representation (a *binary chunk*) of the given function.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.dump)


```lua
function string.dump(f: fun(...any):...unknown, strip?: boolean)
  -> string
```


---

# string.find


Looks for the first match of `pattern` (see [§6.4.1](http://www.lua.org/manual/5.2/manual.html#6.4.1)) in the string.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.find)

@*return* `start`

@*return* `end`

@*return* `...` — captured


```lua
function string.find(s: string|number, pattern: string|number, init?: integer, plain?: boolean)
  -> start: integer
  2. end: integer
  3. ...any
```


---

# string.format


Returns a formatted version of its variable number of arguments following the description given in its first argument.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.format)


```lua
function string.format(s: string|number, ...any)
  -> string
```


---

# string.gmatch


Returns an iterator function that, each time it is called, returns the next captures from `pattern` (see [§6.4.1](http://www.lua.org/manual/5.2/manual.html#6.4.1)) over the string s.

As an example, the following loop will iterate over all the words from string s, printing one per line:
```lua
    s =
"hello world from Lua"
    for w in string.gmatch(s, "%a+") do
        print(w)
    end
```


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.gmatch)


```lua
function string.gmatch(s: string|number, pattern: string|number)
  -> fun():string, ...unknown
```


---

# string.gsub


Returns a copy of s in which all (or the first `n`, if given) occurrences of the `pattern` (see [§6.4.1](http://www.lua.org/manual/5.2/manual.html#6.4.1)) have been replaced by a replacement string specified by `repl`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.gsub)


```lua
function string.gsub(s: string|number, pattern: string|number, repl: string|number|function|table, n?: integer)
  -> string
  2. count: integer
```


---

# string.len


Returns its length.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.len)


```lua
function string.len(s: string|number)
  -> integer
```


---

# string.lower


Returns a copy of this string with all uppercase letters changed to lowercase.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.lower)


```lua
function string.lower(s: string|number)
  -> string
```


---

# string.match


Looks for the first match of `pattern` (see [§6.4.1](http://www.lua.org/manual/5.2/manual.html#6.4.1)) in the string.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.match)


```lua
function string.match(s: string|number, pattern: string|number, init?: integer)
  -> ...any
```


---

# string.pack


Returns a binary string containing the values `v1`, `v2`, etc. packed (that is, serialized in binary form) according to the format string `fmt` (see [§6.4.2](http://www.lua.org/manual/5.2/manual.html#6.4.2)) .

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.pack)


```lua
function string.pack(fmt: string, v1: string|number, v2: any, ...string|number)
  -> binary: string
```


---

# string.packsize


Returns the size of a string resulting from `string.pack` with the given format string `fmt` (see [§6.4.2](http://www.lua.org/manual/5.2/manual.html#6.4.2)) .

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.packsize)


```lua
function string.packsize(fmt: string)
  -> integer
```


---

# string.rep


Returns a string that is the concatenation of `n` copies of the string `s` separated by the string `sep`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.rep)


```lua
function string.rep(s: string|number, n: integer, sep?: string|number)
  -> string
```


---

# string.reverse


Returns a string that is the string `s` reversed.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.reverse)


```lua
function string.reverse(s: string|number)
  -> string
```


---

# string.sub


Returns the substring of the string that starts at `i` and continues until `j`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.sub)


```lua
function string.sub(s: string|number, i: integer, j?: integer)
  -> string
```


---

# string.unpack


Returns the values packed in string according to the format string `fmt` (see [§6.4.2](http://www.lua.org/manual/5.2/manual.html#6.4.2)) .

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.unpack)


```lua
function string.unpack(fmt: string, s: string, pos?: integer)
  -> ...any
  2. offset: integer
```


---

# string.upper


Returns a copy of this string with all lowercase letters changed to uppercase.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-string.upper)


```lua
function string.upper(s: string|number)
  -> string
```


---

# strings

 Simple functions to manipulate UTF-8 encoded strings.

## contains


```lua
function strings.contains(s: string, substr: string)
  -> ok: boolean
```

 Returns true if the string s contains substr.

@*param* `s` — The string to check.

@*param* `substr` — The substring to check for.

@*return* `ok` — True if the string contains the substring.

## contains_any


```lua
function strings.contains_any(s: string, chars: string)
  -> ok: boolean
```

 Returns true if the string s contains any of the runes in chars.
 If chars is the empty string, contains_any returns false.

@*param* `s` — The string to check.

@*param* `chars` — The characters to check for.

@*return* `ok` — True if the string contains any of the characters.

## count


```lua
function strings.count(s: string, substr: string)
  -> n: number
```

 Returns the number of non-overlapping instances of substr in s.

@*param* `s` — The string to check.

@*param* `substr` — The substring to check for.

@*return* `n` — The number of non-overlapping instances of substr in s.

## count_any


```lua
function strings.count_any(s: string, chars: string)
  -> n: number
```

 Returns the number of non-overlapping instances of any of the runes in chars in s.

@*param* `s` — The string to check.

@*param* `chars` — The characters to check for.

@*return* `n` — The number of non-overlapping instances of any of the runes in chars in s.

## duplicate


```lua
function strings.duplicate(s: string, count: number)
  -> s: string
```

 Returns a new string consisting of count copies of the string s. If count is zero or negative, returns the empty string.

@*param* `s` — The string to repeat.

@*param* `count` — The number of times to repeat the string.

@*return* `s` — A new string consisting of count copies of the string s.

## equal_fold


```lua
function strings.equal_fold(s: string, t: string)
  -> ok: boolean
```

 Returns true if the strings s and t, interpreted as UTF-8 strings, are equal under Unicode case-folding.

@*param* `s` — The first string to check.

@*param* `t` — The second string to check.

@*return* `ok` — True if the strings are equal under Unicode case-folding.

## has_prefix


```lua
function strings.has_prefix(s: string, prefix: string)
  -> ok: boolean
```

 Returns true if the string s begins with prefix.

@*param* `s` — The string to check.

@*param* `prefix` — The prefix to check for.

@*return* `ok` — True if the string has the prefix.

## has_suffix


```lua
function strings.has_suffix(s: string, suffix: string)
  -> ok: boolean
```

 Returns true if the string s ends with suffix.

@*param* `s` — The string to check.

@*param* `suffix` — The suffix to check for.

@*return* `ok` — True if the string has the suffix.

## index


```lua
function strings.index(s: string, substr: string)
  -> i: number
```

 Returns the index of the first instance of substr in s, or -1 if substr is not present in s.

@*param* `s` — The string to check.

@*param* `substr` — The substring to check for.

@*return* `i` — The index of the first instance of substr in s, or -1 if substr is not present in s.

## index_any


```lua
function strings.index_any(s: string, chars: string)
  -> i: number
```

 Returns the index of the first instance of any of the runes in chars in s, or -1 if none of the runes in chars are present in s.

@*param* `s` — The string to check.

@*param* `chars` — The characters to check for.

@*return* `i` — The index of the first instance of any of the runes in chars in s, or -1 if none of the runes in chars are present in s.

## last_index


```lua
function strings.last_index(s: string, substr: string)
  -> i: number
```

 Returns the index of the last instance of substr in s, or -1 if substr is not present in s.

@*param* `s` — The string to check.

@*param* `substr` — The substring to check for.

@*return* `i` — The index of the last instance of substr in s, or -1 if substr is not present in s.

## last_index_any


```lua
function strings.last_index_any(s: string, chars: string)
  -> i: number
```

 Returns the index of the last instance of any of the runes in chars in s, or -1 if none of the runes in chars are present in s.

@*param* `s` — The string to check.

@*param* `chars` — The characters to check for.

@*return* `i` — The index of the last instance of any of the runes in chars in s, or -1 if none of the runes in chars are present in s.

## replace


```lua
function strings.replace(s: string, old: string, new: string, n: number)
  -> new: string
```

 Replaces the first n instances of old with new in s.

@*param* `s` — The string to replace in.

@*param* `old` — The string to replace.

@*param* `new` — The string to replace with.

@*param* `n` — The number of instances to replace.

@*return* `new` — The new string.

## replace_all


```lua
function strings.replace_all(s: string, old: string, new: string)
  -> new: string
```

 Replaces all instances of old with new in s.

@*param* `s` — The string to replace in.

@*param* `old` — The string to replace.

@*param* `new` — The string to replace with.

@*return* `new` — The new string.

## split


```lua
function strings.split(s: string, sep: string)
  -> table: table
```

 Splits s around each instance of sep and returns a table of the substrings between those instances.

@*param* `s` — The string to split.

@*param* `sep` — The string to split around.

@*return* `table` — A table of the substrings between each instance of sep.

## title


```lua
function strings.title(s: string)
  -> s: string
```

 Returns a copy of the string s with all Unicode letters that begin words mapped to their title case.

@*param* `s` — The string to convert.

@*return* `s` — A copy of the string s with all Unicode letters that begin words mapped to their title case.

## to_lower


```lua
function strings.to_lower(s: string)
  -> new: string
```

 Returns a copy of the string s with all Unicode letters mapped to their lower case.

@*param* `s` — The string to convert.

@*return* `new` — The new string.

## to_upper


```lua
function strings.to_upper(s: string)
  -> new: string
```

 Returns a copy of the string s with all Unicode letters mapped to their upper case.

@*param* `s` — The string to convert.

@*return* `new` — The new string.

## trim


```lua
function strings.trim(s: string, cutset: string)
  -> new: string
```

 Returns a copy of the string s with all leading and trailing Unicode code points contained in cutset removed.

@*param* `s` — The string to trim.

@*param* `cutset` — The set of Unicode code points to remove.

@*return* `new` — The new string.

## trim_left


```lua
function strings.trim_left(s: string, cutset: string)
  -> new: string
```

 Returns a copy of the string s with all leading Unicode code points contained in cutset removed.

@*param* `s` — The string to trim.

@*param* `cutset` — The set of Unicode code points to remove.

@*return* `new` — The new string.

## trim_prefix


```lua
function strings.trim_prefix(s: string, prefix: string)
  -> new: string
```

 Returns a copy of the string s without the provided leading prefix string. If s doesn't start with prefix, s is returned unchanged.

@*param* `s` — The string to trim.

@*param* `prefix` — The prefix to remove.

@*return* `new` — The new string.

## trim_right


```lua
function strings.trim_right(s: string, cutset: string)
  -> new: string
```

 Returns a copy of the string s with all trailing Unicode code points contained in cutset removed.

@*param* `s` — The string to trim.

@*param* `cutset` — The set of Unicode code points to remove.

@*return* `new` — The new string.

## trim_space


```lua
function strings.trim_space(s: string)
  -> new: string
```

 Returns a copy of the string s with all leading and trailing white space removed, as defined by Unicode.

@*param* `s` — The string to trim.

@*return* `new` — The new string.

## trim_suffix


```lua
function strings.trim_suffix(s: string, suffix: string)
  -> new: string
```

 Returns a copy of the string s without the provided trailing suffix string. If s doesn't end with suffix, s is returned unchanged.

@*param* `s` — The string to trim.

@*param* `suffix` — The suffix to remove.

@*return* `new` — The new string.


---

# table




[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-table)



```lua
tablelib
```


---

# table.concat


Given a list where all elements are strings or numbers, returns the string `list[i]..sep..list[i+1] ··· sep..list[j]`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-table.concat)


```lua
function table.concat(list: table, sep?: string, i?: integer, j?: integer)
  -> string
```


---

# table.foreach


Executes the given f over all elements of table. For each element, f is called with the index and respective value as arguments. If f returns a non-nil value, then the loop is broken, and this value is returned as the final value of foreach.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-table.foreach)


```lua
function table.foreach(list: any, callback: fun(key: string, value: any):<T>|nil)
  -> <T>|nil
```


---

# table.foreachi


Executes the given f over the numerical indices of table. For each index, f is called with the index and respective value as arguments. Indices are visited in sequential order, from 1 to n, where n is the size of the table. If f returns a non-nil value, then the loop is broken and this value is returned as the result of foreachi.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-table.foreachi)


```lua
function table.foreachi(list: any, callback: fun(key: string, value: any):<T>|nil)
  -> <T>|nil
```


---

# table.getn


Returns the number of elements in the table. This function is equivalent to `#list`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-table.getn)


```lua
function table.getn(list: <T>[])
  -> integer
```


---

# table.insert


Inserts element `value` at position `pos` in `list`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-table.insert)


```lua
function table.insert(list: table, pos: integer, value: any)
```


---

# table.maxn


Returns the largest positive numerical index of the given table, or zero if the table has no positive numerical indices.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-table.maxn)


```lua
function table.maxn(table: table)
  -> integer
```


---

# table.move


Moves elements from table `a1` to table `a2`.
```lua
a2[t],··· =
a1[f],···,a1[e]
return a2
```


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-table.move)


```lua
function table.move(a1: table, f: integer, e: integer, t: integer, a2?: table)
  -> a2: table
```


---

# table.pack


Returns a new table with all arguments stored into keys `1`, `2`, etc. and with a field `"n"` with the total number of arguments.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-table.pack)


```lua
function table.pack(...any)
  -> table
```


---

# table.remove


Removes from `list` the element at position `pos`, returning the value of the removed element.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-table.remove)


```lua
function table.remove(list: table, pos?: integer)
  -> any
```


---

# table.sort


Sorts list elements in a given order, *in-place*, from `list[1]` to `list[#list]`.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-table.sort)


```lua
function table.sort(list: <T>[], comp?: fun(a: <T>, b: <T>):boolean)
```


---

# table.unpack


Returns the elements from the given list. This function is equivalent to
```lua
    return list[i], list[i+1], ···, list[j]
```
By default, `i` is `1` and `j` is `#list`.


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-table.unpack)


```lua
function table.unpack(list: <T>[], i?: integer, j?: integer)
  -> ...<T>
```


---

# time

 Time library

## hour


```lua
number
```

Duration constant. 60 * minute

## microsecond


```lua
number
```

Duration constant. 1000 * nanosecond

## millisecond


```lua
number
```

Duration constant. 1000 * microsecond

## minute


```lua
number
```

Duration constant. 60 * second

## nanosecond


```lua
number
```

Duration constant

## second


```lua
number
```

Duration constant. 1000 * millisecond

## sleep


```lua
function time.sleep(duration: number)
```

 Sleep for the given duration


---

# tonumber


When called with no `base`, `tonumber` tries to convert its argument to a number. If the argument is already a number or a string convertible to a number, then `tonumber` returns this number; otherwise, it returns `fail`.

The conversion of strings can result in integers or floats, according to the lexical conventions of Lua (see [§3.1](http://www.lua.org/manual/5.2/manual.html#3.1)). The string may have leading and trailing spaces and a sign.


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-tonumber)


```lua
function tonumber(e: any)
  -> number?
```


---

# tostring


Receives a value of any type and converts it to a string in a human-readable format.

If the metatable of `v` has a `__tostring` field, then `tostring` calls the corresponding value with `v` as argument, and uses the result of the call as its result. Otherwise, if the metatable of `v` has a `__name` field with a string value, `tostring` may use that string in its final result.

For complete control of how numbers are converted, use [string.format](http://www.lua.org/manual/5.2/manual.html#pdf-string.format).


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-tostring)


```lua
function tostring(v: any)
  -> string
```


---

# type


Returns the type of its only argument, coded as a string. The possible results of this function are `"nil"` (a string, not the value `nil`), `"number"`, `"string"`, `"boolean"`, `"table"`, `"function"`, `"thread"`, and `"userdata"`.


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-type)


```lua
type:
    | "nil"
    | "number"
    | "string"
    | "boolean"
    | "table"
    | "function"
    | "thread"
    | "userdata"
```


```lua
function type(v: any)
  -> type: "boolean"|"function"|"nil"|"number"|"string"...(+3)
```


---

# unpack


Returns the elements from the given `list`. This function is equivalent to
```lua
    return list[i], list[i+1], ···, list[j]
```


[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-unpack)


```lua
function unpack(list: <T>[], i?: integer, j?: integer)
  -> ...<T>
```


---

# url_url

 Structured URL

## hostname


```lua
(method) url_url:hostname()
  -> hostname: string
```

 Return hostname without port numbers.

@*return* `hostname` — URLs hostname

## join_path


```lua
(method) url_url:join_path(...string)
  -> url: url_url
```

 Returns a new URL with the provided path elements joined to any existing path and the resulting path cleaned of any ./ or ../ elements. Any sequences of multiple / characters will be reduced to a single /.

## parse


```lua
(method) url_url:parse(ref: string)
  -> url: url_url
```

 Parses a URL in the context of the receiver. The provided URL may be relative or absolute.

## query


```lua
(method) url_url:query(query?: url_values)
  -> query: url_values?
```

## string


```lua
(method) url_url:string()
  -> url: string
```

 Reassembles the URL into a valid URL string.


---

# url_values

 Values maps a string key to a list of values. It is typically used for query parameters and form values. Unlike in the `http.header` map, the keys in a `values` map are case-sensitive.

## add


```lua
(method) url_values:add(key: string, value: string)
```

 Adds the key and value to the values. It appends to any existing values associated with key.

@*param* `key` — The key to add. It must not be empty.

@*param* `value` — The value to add. It must not be empty.

## del


```lua
(method) url_values:del(key: string)
```

 Deletes the values associated with key.

@*param* `key` — The key to delete.

## get


```lua
(method) url_values:get(key: string)
  -> value: string
```

 Gets the first value associated with the given key. If there are no values associated with the key, Get returns "".

@*param* `key` — The key to get.

@*return* `value` — The first value associated with the given key.

## has


```lua
(method) url_values:has(key: string)
  -> has: boolean
```

 Returns true if the values contains the specified key, false otherwise.

@*param* `key` — The key to check.

@*return* `has` — True if the values contains the specified key, false otherwise.

## parse


```lua
(method) url_values:parse(encoded: string)
  -> values: url_values
```

 Creates a values from the URL encoded form. It is the inverse operation of string.

@*param* `encoded` — The URL encoded form of the values.

@*return* `values` — The values created from the URL encoded form.

## set


```lua
(method) url_values:set(key: string, value: string)
```

 Sets the key to value. It replaces any existing values associated with key.

@*param* `key` — The key to add. It must not be empty.

@*param* `value` — The value to add. It must not be empty.

## string


```lua
(method) url_values:string()
  -> encoded: string
```

 Encodes the values into "URL encoded" form sorted by key.

@*return* `encoded` — The URL encoded form of the values.


---

# urls

 URLs is a library for working with URLs.

## parse


```lua
function urls.parse(raw_url: string)
  -> url: url_url
```

 Parses URL

@*param* `raw_url` — URL string to parse

@*return* `url` — Parsed URL

## path_escape


```lua
function urls.path_escape(path: string)
  -> escaped: string
```

 Escapes the string so it can be safely placed inside a URL path segment, replacing special characters (including /) with %XX sequences as needed.

@*param* `path` — The path to escape.

@*return* `escaped` — The escaped path.

## path_unescape


```lua
function urls.path_unescape(escaped: string)
  -> path: string
```

 Unescapes a string; the inverse operation of path_escape. It converts each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB.

@*param* `escaped` — The escaped path.

@*return* `path` — The unescaped path.

## query_escape


```lua
function urls.query_escape(query: string)
  -> escaped: string
```

 Escapes the string so it can be safely placed inside a URL query parameter, replacing special characters (including /) with %XX sequences as needed.

@*param* `query` — The query to escape.

@*return* `escaped` — The escaped query.

## query_unescape


```lua
function urls.query_unescape(escaped: string)
  -> query: string
```

 Unescapes a string; the inverse operation of query_escape. It converts each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB.

@*param* `escaped` — The escaped query.

@*return* `query` — The unescaped query.

## values


```lua
function urls.values()
  -> values: url_values
```

 Creates a new values.

@*return* `values` — The new values.


---

# util

 Functional helpers

## chunk


```lua
function util.chunk(list: <T>[], size?: number)
  -> chunks: <T>[][]
```

 Split list into chunks

@*param* `list` — List to split

@*param* `size` — Chunk size

@*return* `chunks` — List of chunks

## contains_by


```lua
function util.contains_by(list: <T>[], predicate: fun(value: <T>):boolean)
  -> contains: boolean
```

 Checks if a list contains a value using predicate

@*param* `list` — List to check

@*param* `predicate` — Predicate function

@*return* `contains` — Whether the list contains the value

## drop


```lua
function util.drop(list: <T>[], n: number)
  -> dropped: <T>[]
```

 Drop n elements from the beginning list

@*param* `list` — List to drop from

@*param* `n` — Number of elements to drop

@*return* `dropped` — List of dropped elements

## drop_right


```lua
function util.drop_right(list: <T>[], n: number)
  -> dropped: <T>[]
```

 Drop n elements from the end of list

@*param* `list` — List to drop from

@*param* `n` — Number of elements to drop

@*return* `dropped` — List of dropped elements

## drop_right_while


```lua
function util.drop_right_while(list: <T>[], predicate: fun(value: <T>):boolean)
```

 Drop elements from the end of list while predicate is true

@*param* `list` — List to drop from

@*param* `predicate` — Predicate function

## drop_while


```lua
function util.drop_while(list: <T>[], predicate: fun(value: <T>):boolean)
```

 Drop elements from the beginning of list while predicate is true

@*param* `list` — List to drop from

@*param* `predicate` — Predicate function

## filter


```lua
function util.filter(list: <T>[], predicate: fun(value: <T>, index: number):boolean)
  -> filtered: <T>[]
```

 Filter list by predicate

@*param* `list` — List to filter

@*param* `predicate` — Predicate function

@*return* `filtered` — List of filtered elements

## find


```lua
function util.find(list: <T>[], predicate: fun(value: <T>):boolean)
  -> element: <T>
  2. ok: boolean
```

 Find first element in list that satisfies predicate.

@*param* `list` — List to search

@*param* `predicate` — Predicate function

@*return* `element` — First element that satisfies predicate

@*return* `ok` — True if element was found

## find_index


```lua
function util.find_index(list: <T>[], predicate: fun(value: <T>):boolean)
  -> index: number
  2. ok: boolean
```

 Find index of first element in list that satisfies predicate.

@*param* `list` — List to search

@*param* `predicate` — Predicate function

@*return* `index` — Index of first element that satisfies predicate

@*return* `ok` — True if element was found

## find_last_index


```lua
function util.find_last_index(list: <T>[], predicate: fun(value: <T>):boolean)
  -> index: number
  2. ok: boolean
```

 Find index of last element in list that satisfies predicate.

@*param* `list` — List to search

@*param* `predicate` — Predicate function

@*return* `index` — Index of last element that satisfies predicate

@*return* `ok` — True if element was found

## for_each


```lua
function util.for_each(list: <T>[], func: fun(value: <T>, index: number))
```

 Calls a function for each element in a list

@*param* `list` — List to iterate over

@*param* `func` — Function to call

## head


```lua
function util.head(list: <T>[])
  -> element: <T>?
```

 Get first element of list

@*param* `list` — List to get first element from

@*return* `element` — First element of list

## id


```lua
function util.id(value: <T>)
  -> result: <T>
```

 Returns the first argument

@*param* `value` — Value to return

@*return* `result` — The first argument

## init


```lua
function util.init(list: <T>[])
  -> initial: <T>[]?
```

 Get all elements of list except last

@*param* `list` — List to get initial from

@*return* `initial` — List of all elements except last

## last


```lua
function util.last(list: <T>[])
  -> element: <T>?
```

 Get last element of list

@*param* `list` — List to get last element from

@*return* `element` — Last element of list

## map


```lua
function util.map(table: table<<T>, <G>>, func: fun(value: <G>, key: <T>):<Q>)
  -> mapped: table<<T>, <Q>>
```

 Maps a function over a table values

@*param* `table` — Table to map over

@*param* `func` — Mapping function

@*return* `mapped` — Mapped list

## map_async


```lua
function util.map_async(table: table<<T>, <G>>, func: fun(value: <G>, key: <T>):<Q>)
  -> mapped: table<<T>, <Q>>
```

 Maps a function over a table values asynchronously

@*param* `table` — Table to map values over

@*param* `func` — Mapping function. Takes a value of each key

@*return* `mapped` — Mapped table

## max_by


```lua
function util.max_by(list: <T>[], func: fun(a: <T>, b: <T>):boolean)
  -> result: <T>
```

 Returns the maximum value of a list, based on a function

@*param* `list` — List to find the maximum value of

@*param* `func` — Function to use to compare values. Returns true if the first argument is less than the second

@*return* `result` — The maximum value

## min_by


```lua
function util.min_by(list: <T>[], func: fun(a: <T>, b: <T>):boolean)
  -> result: <T>
```

 Returns the minimum value of a list, based on a function

@*param* `list` — List to find the minimum value of

@*param* `func` — Function to use to compare values. Returns true if the first argument is less than the second

@*return* `result` — The minimum value

## reduce


```lua
function util.reduce(list: <T>[], func: fun(accumulator: <G>, value: <T>, index: number):<G>, inital: <G>)
  -> result: <G>
```

 Reduces a list to a single value

@*param* `list` — List to reduce

@*param* `func` — Function to reduce with

@*param* `inital` — An initial value

@*return* `result` — Result of the reduction

## reduce_right


```lua
function util.reduce_right(list: <T>[], func: fun(accumulator: <G>, value: <T>, index: number):<G>, initial: <G>)
  -> result: <G>
```

 Reduces a list to a single value

@*param* `list` — List to reduce

@*param* `func` — Function to reduce with

@*param* `initial` — An initial value

@*return* `result` — Result of the reduction

## reject


```lua
function util.reject(list: <T>[], predicate: fun(value: <T>, index: number):boolean)
  -> result: <T>[]
```

 Rejects all elements that match a predicate

@*param* `list` — List to reject from

@*param* `predicate` — Predicate function

@*return* `result` — List of rejected elements

## reverse


```lua
function util.reverse(list: <T>[])
  -> reversed: <T>[]
```

 Reverse list

@*param* `list` — List to reverse

@*return* `reversed` — Reversed list

## slice


```lua
function util.slice(list: <T>[], start: number, finish?: number)
  -> slice: <T>[]
```

 Get slice of list

@*param* `list` — List to get slice from

@*param* `start` — Start index of slice

@*param* `finish` — End index of slice. Defaults to length of list

@*return* `slice` — Slice of list

## sort


```lua
function util.sort(list: <T>[], cmp: fun(a: <T>, b: <T>):boolean)
  -> sorted: <T>[]
```

 Returns a sorted list

@*param* `list` — List to sort

@*param* `cmp` — Comparison function. Defaults to <

@*return* `sorted` — Sorted list

## tail


```lua
function util.tail(list: <T>[])
  -> tail: <T>[]?
```

 Get all elements of list except first

@*param* `list` — List to get tail from

@*return* `tail` — List of all elements except first

## take


```lua
function util.take(list: <T>[], n: number)
  -> elements: <T>[]?
```

 Get first n elements of list

@*param* `list` — List to get elements from

@*param* `n` — Number of elements to get

@*return* `elements` — First n elements of list

## take_while


```lua
function util.take_while(list: <T>[], predicate: fun(value: <T>):boolean)
  -> elements: <T>[]?
```

 Take elements from list while predicate is true

@*param* `list` — List to get elements from

@*param* `predicate` — Predicate function

@*return* `elements` — Elements of list before predicate is false


---

# warn


Emits a warning with a message composed by the concatenation of all its arguments (which should be strings).

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-warn)


```lua
function warn(message: string, ...any)
```


---

# xpcall


Calls function `f` with the given arguments in protected mode with a new message handler.

[View documents](http://www.lua.org/manual/5.4/manual.html#pdf-xpcall)


```lua
function xpcall(f: fun(...any):...unknown, msgh: function, arg1?: any, ...any)
  -> success: boolean
  2. result: any
  3. ...any
```