# drumMachine
Coding exercise for Splice interview

To run:

 - First, get the project using:

    ```go get https://github.com/kdrombosky/drumMachine```

 - From the drumMachine directory, run using: 
    
    ```./run.sh```

To test:
 - From the drumMachine directory, run using: 
    
    ```./test.sh```
 

Details:

I decided to implement this project using go, even though my it is not necessarily my strongest language. Thus, the amount of time taken on this project was a bit longer than it would have been using a language that I code in more regularly, like Java or Groovy.

The main program runs in main.go. First, it prints a text based menu for the user to choose options from. I incorporated 4 different playable patterns (one of which is from an old and tired joke :)).

I represented the patterns in a struct containing the necessary information to output the pattern metadata and play the easily via looping. More detailed information on this struct is documented in struct.go

I attempted to handle all unexpected input, throwing errors when necessary, but allowing the user another chance to enter appropriate input whenever possible.

I included tests for the validating the content of each preset pattern. These tests are trivial, but I wanted to include a decent amount of code coverage while practicing writing tests in go. I also think it is useful to maintain tests for expected content because a change to that content could have the potential to derail an entire system.

I also included tests for the validateTempo method, which has some of the greatest potential for error, considering the amount of parameter manipulation invloved.

This project was super fun! I am a classically trained musician, and thinking about how my music-brain would interpret sheet music had a large influence on how my coding-brain structured the beat patterns.