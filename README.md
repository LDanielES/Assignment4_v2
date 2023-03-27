## News Website - Assignment 4
Our final project is News Website writen in Golang, using three different design patterns (Dependency Injection Pattern, Builder Pattern and Observer Pattern). We use an API key from newsapi.org.


## Prerequisites to run & use the app

* A free PORT to serve the website (In this case, I specified the PORT :8088 in the build.env file. If that port isn't available in your computer, change the port number in the build.env file)

* An API key (Also in the build.env file)

## Instructions to run the program

* Open the terminal, go to the root directory of the project and write the following command: *./main.exe*
* Search *localhost:8088* in your browser

## Instructions  to use the website

* Write any topic in the search bar
* Navigate through all the result with the buttons at the bottom of the page
* To clean the page and start a new search, click the logo on the top left corner
* To go to the Github Repository, click the button on the top right corner


## Problems encountered during coding

We had a lot of problems working with local packages. Go was searching the packages in GOROOT instead of in my project directory.

We also had troubles working with syscall/js package, to the point where we took the decision of changing it.

Dependecy cycles were also a small problem for a short time, but we quickly found the solution to that.