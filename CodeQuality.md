# Code Quality Review: App / Module Name
## Introduction
### Code Style.
### Code Standards.

### Areas to review:
1. satisfies functional requirements
2. satisfies microservices philosophy
3. code is written according to Go language style
4. code is written consistent
5. code is easy to understand
6. code is easy to maintain, install and operate
7. parts of code can be reused
8. code is well documented
9. code can be easily tested
10. unit test coverage is above accepted threshold

## Code Metrics
### LOC
$ find . -name "*.go" -exec cat {} + | wc -l

### Number of functions
### Number of methods
### Number of defined types
### Number of interfaces
### Number of uncommented functions / methods
### Number of uncommented interfaces
### Number of packages
#### Highest number of LOC packages
#### Lowest number of LOC packages
#### Average number of LOC per package

## Area 1:  Satisfies Functional Requirements
### 1.1 Code is easy to configure during the application lifecycle
### 1.2 Code is easy to adjust its capacity
### 1.3 Current technical debt 
### 1.4 Dataflow
#### 1.4.1 Validation of input data
#### 1.4.2 Internal data movement
Law of DemeterL Obj1().Obj2().Obj3(). Each object should speak with the one adjacent.

Resources
```html
https://en.wikipedia.org/wiki/Law_of_Demeter
```

#### 1.4.3 Data formatting
#### 1.4.4 Data persistence

## Area 2:  Satisfies Microservices Philosophy
Microservices Architecture Pattern attributes:
1. loosely coupled
2. independently deployable
3. discoverable
4. easily manageable
5. horizontally scalable

## Area 3:  Code is written according to Language Style
### Uber style guide
### Thanos
```
https://thanos.io/contributing/coding-style-guide.md/
```

## Area 4:  Code is written consistent

## Area 5:  Code is easy to understand
```
https://www.sonarsource.com/docs/CognitiveComplexity.pdf
```
### Variable naming convention
#### Meaningful variable / type names, describing the data it represents
#### Avoiding generic names for variables
Example: <br/>
generic: quantity <br/>
explicit: total_purchases
#### Variable names not longer than 20 characters
#### Units of measure included in variable names
### Error messages
Error messages stick to the error messages policy.

#### Meaningful file names to import / export

### Number of variables defined by function / method

### Boolean handling
#### Extract complicated boolean tests to variables
#### Compare positive boolean expressions 

### Fail fast returning as soon as possible in a function
### Easy to understand Data Structures
Data Structures do not have many levels.

### Functions / methods have a limited number of parameters
Methods with 0 and 1 arguments are fine
2 parameters still good
3 parameters can be considered OK
4 and more parameters are usually too much

Resources
```html
http://www.principles-wiki.net/anti-patterns:long_parameter_list
```

## Area 6:  Code is easy to maintain, install and operate
### Lower complexity increases maintainability
### Lower number of layers makes the code easier to debug
### Lower number of lines of code makes the code faster to read


## Area 7:  Parts of code can be reused

## Area 8:  Code is well documented
Generally comments explain what the code is about and what cannot be expressed in code.
### Code contains documentation for package
### Code contains documentation at top of file
### Code contains comments for functions / methods
### Code contains comments for interfaces
### Code contains summary comments
Summary comments summarize few lines of code or an alghorithm.

### Detail module architecture.
a. Detail the inside of the module to easily understand how it is designed and built.<br/>
b. How the module has been decomposed into building blocks (ex. layers, packages, models).<br/>
Detail the organisation of the code.<br/>
c. Detail the path from input to output for each input through the application blocks.

## Area 9:  Code can be easily tested
How use cases are served? Is there a set of controllers / handlers?

### End to end testing
### Integration testing
### Assurance

## Area 10: Unit test coverage is above agreed threshold
