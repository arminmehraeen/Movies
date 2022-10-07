# Movies
webserver with golang ( simple movies api )  

# install
    go run ./main.go
    or
    go build ./main.go -> run movieServer.exe
   
# routes
    localhost:8000/movies             method = GET        # List of movies
    localhost:8000/movies             method = POST       # Create movie
    localhost:8000/movies/{movieId}   method = GET        # Get movie 
    localhost:8000/movies/{movieId}   method = PUT        # Update movie
    localhost:8000/movies/{movieId}   method = DELETE     # Delete movie
