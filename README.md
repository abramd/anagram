# anagram
-

Implementation of server for anagram search.

####Algorithm complexity: 
Search: Sort + Map Lookup = O(n\*log(n)) + O(1) ~~ O(n\*log(n)), where "n" - count of chars in word

####Handlers:
1. `[POST] /add` body: `{"data": ["",""]} ` - words adding
2. `[GET] /search?word=anyword` - search (C.O.)
3. `[GET] /list` - list of saved words