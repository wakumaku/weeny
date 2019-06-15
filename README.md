# weeny

Start application

* $ `make start`

Shortern a URL
* $ `curl -X POST -H 'application/json' -d '{"url":"http://google.com"}' http://localhost:8080/shortern`
    
    ```
    {"message":"Success","data":"OeULF4sJHN"}
    ```
* browser: http://localhost:8080/lookup/OeULF4sJHN
    ```
    {"message":"Success","data":"http://google.com"}
    ```
* browser: http://localhost:8080/OeULF4sJHN
    ```
    < HTTP/1.1 307 Temporary Redirect
    < Content-Type: text/html; charset=utf-8
    < Location: http://google.com
    ```

Stop application

* $ `make stop`

View docker logs

* $ `make logs`
