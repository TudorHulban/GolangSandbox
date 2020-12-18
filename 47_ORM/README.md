### Performance results
 Response time seconds histogram for 500 requests:
 ```
  0.000 [1]		|
  0.001 [214]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [159]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [89]     |■■■■■■■■■■■■■■■■■
  0.005 [18]     |■■■
  0.007 [11]    	|■■
  0.008 [2]		|
  0.009 [3]		|■
  0.011 [1]		|
  0.012 [0]		|
  0.013 [2]
```
 Tested with: 
 ```bash
 ./hey_linux_amd64 -n 500 http://localhost:8080/authors/4 
 ```