# Movies
webserver with golang ( simple movies api )  

# install
    go run ./main.go
    or
    go build ./main.go -> run movieServer.exe
    
    webserver ready on localhost:8000

# routes
    /movies             method = GET        List of movies
    /movies             method = POST       Create movie
    /movies/{movieId}   method = GET        Get movie 
    /movies/{movieId}   method = PUT        Update movie
    /movies/{movieId}   method = DELETE     Delete movie
