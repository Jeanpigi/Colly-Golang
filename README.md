## Colly-Golang

With Colly you can easily extract structured data from websites, which can be used for a wide range of applications, like data mining, data processing or archiving. [Colly](http://go-colly.org/).

To create a new project similar to this one, just follow steps in web site [Install](http://go-colly.org/docs/introduction/install/).

## Features

Clean API
Fast (>1k request/sec on a single core)
Manages request delays and maximum concurrency per domain
Automatic cookie and session handling
Sync/async/parallel scraping
Distributed scraping
Caching
Automatic encoding of non-unicode responses
Robots.txt support
Google App Engine support


## Get started with Go

Install the Go Programming language [Install](https://golang.org/)

# Setting GOPATH in Unix system 

The `GOPATH` environment variable specifies the location of your workspace. Open a terminal or console:

```
go env -w GOPATH=$HOME/[filename]
```

Example:

```
go env -w GOPATH=$HOME/go
```

Next export the GOPATH. 

__NOTE:__  It is important to know that depends on shell has computer (bash, zsh, fish) for do this step. Open terminal and follow the instructions for any case:

# Bash

Edit your `~/.bash_profile` to add the following line:

```bash
export GOPATH=$HOME/go
```

Save and exit your editor. Then, source your `~/.bash_profile`.
```bash
source ~/.bash_profile
```

# Zsh

Edit your `~/.zshrc` file to add the following line:

```bash
export GOPATH=$HOME/go
```
Save and exit your editor. Then, source your `~/.zshrc`.
```bash
source ~/.zshrc
```

# fish

```bash
set -x -U GOPATH $HOME/go
```

# Create a folder

Once it is setting GOPATH. A package's import path corresponds to its location inside a workspace or in a remote repository, for that reason you have to create that folder which identify GOPATH.

```
mkdir -p $GOPATH/src/github.com/user/hello
```

Then get inside the folder a create your first project

```
cd $HOME/go/src/github.com/user/hello
```

__NOTE:__ that you don't need to publish your code to a remote repository before you can build it. For this example We'll use github.com/user/hello as our base path. 

More [Information](https://golang.org/doc/gopath_code.html)


# Setting GOPATH in Window system

Your workspace can be located wherever you want,
but we'll use `C:\go-work` in this example.

__NOTE:__ `GOPATH` must not be the same path as your Go installation.

First, confirm your Go binaries

* Create folder at `C:\go-work`.
* Right click on "Start" and click on "Control Panel". Select "System and Security", then click on "System".
* From the menu on the left, select the "Advanced system settings".
* Click the "Environment Variables" button at the bottom.
* Click "New" from the "User variables" section.
* Type `GOPATH` into the "Variable name" field.
* Type `C:\go-work` into the "Variable value" field.
* Click OK.

# Windows 10 (command line)
* Open a command prompt (`Win` + `r` then type `cmd`) or a powershell window (`Win` + `i`).
* Type `setx GOPATH %USERPROFILE%\go`. (This will set the `GOPATH` to your `[home folder]\go`, such as `C:\Users\yourusername\go`.)
* Close the command or powershell window. (The environment variable is only available for new command or powershell windows, not for the current window.)

## Get started with Colly

First, you need to import Colly to your codebase:

```
import "github.com/gocolly/colly"
```

Colly does not require any special setup outside of the normal installation for go packages. The easiest way to install is of course:

```
go get github.com/gocolly/colly
```
 
The core of a Colly scraper is the colly.Collector. Collector manages the network communication and responsible for the execution of the attached callbacks while a collector job is running. This initiates the calls to websites and tracks page visits, beside controls access rules, controls rate limits, sets proxies, etc.. Setting up a simple Collector starts looks something like:

```
c := colly.NewCollector()
```

Colly is a highly customizable scraping framework. It has sane defaults and provides plenty of options to change them.

Full list of collector attributes can be found [here](https://godoc.org/github.com/gocolly/colly#Collector). The recommended way to initialize a collector is using colly.NewCollector(options...).

example:

```
c := colly.NewCollector(
	colly.UserAgent("xy"),
	colly.AllowURLRevisit(),
)
```


Now you need to setup your callbacks for get specific HTML elements. You can attach different type of callback functions to a Collector to control a collecting job or retrieve information. This is done by setting up onHTML or onXML. First one is used if you are looking for elements using CSS queries, but you can also use second one which implements XPath if it is more your style.


Examples:

```
c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    e.Request.Visit(e.Attr("href"))
})
```

```
c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
    fmt.Println("First column of a table row:", e.Text)
})
```

```
c.OnXML("//h1", func(e *colly.XMLElement) {
    fmt.Println(e.Text)
})
```

Start scraping on website which want to visit

```	
c.Visit("https://hackerspaces.org/")
```

## Building

If you just want to run, you can use a command which is:

```	
$ go run main.go
```	

Then build it with the go tool: 

```	
$ go build main.go
```	

The command above will build an executable named main in the current directory alongside your source code. Execute it to see extraction of data:

```	
$  ./main.go
```	

__NOTE:__  Just is executable for UNIX system. For execute into Windows system click on the main.exe

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.


## Authors
* **Jean Pierre Giovanni Arenas Ortiz** - *Initial work* -

## License
[MIT](https://choosealicense.com/licenses/mit/)