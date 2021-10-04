# Introduction

Pet project heavily inspired from the excellent book [Your Code as a Crime Scene](https://pragprog.com/titles/atcrime/your-code-as-a-crime-scene/) written by Adam Tornhill.

This cli tool will allow you to build some of the charts described in that book in order to start conversations regarding the design or organisation of some application.

It has very similar interface to [code-maat](https://github.com/adamtornhill/code-maat), the tool created by Adam Tornhill. Here are the main differences:

- written in golang
- only support git
- use a database to store the information
- include the visualisations

gocan was an opportunity for me to learn the Go language so don't be too harsh with the source code :-)

_Note_: to understand the visualisations generated by the tool, it is better to read the book first.

# Installation

For MacOS & Linux users, you can use [homebrew](https://brew.sh) to install the application:

```
brew tap fouadh/tap
brew install gocan
```

For other platforms, you will have to build the binary from the source code (see the section below).

# Using the app

## Start the embedded database

Gocan comes with an embedded database that can be used to store the metrics data. To start it, execute the following
command:

```
gocan start-db
```

## Run the UI

To visualize the different charts, execute the following command that bootstraps an embedded web server:

```
gocan ui
```

## Create a forensics scene

A scene is a collection of applications. Before adding any application, create a scene first. It will allow (not yet
implemented) to compare applications gathered in a same scene. It can be useful to analyze distributed systems for example.

```
gocan create-scene my-scene
```

## Add an application to a scene

Create an application in a scene.

```
gocan create-app my-app -s my-scene
```

## Import an application history

In order to get metrics, it is needed to run this command to import information from a git repository that has been
cloned locally.

```
gocan import-history my-app -s my-scene --after 2021-01-01 --before 2021-06-30
```

It is recommended to limit the data that you want to analyze to a small period. Like mentionned in the book, having
too many data can skew the results and obscure most recent trends.

## Stop the embedded database

When you're done using the CLI, you can eventually stop the database.

```
gocan stop-db
```

## Using an external database

It is possible to use an external database instead of the embedded one. To configure it, use the following command
with the appropriate flags values:

```
gocan setup-db --external-db --host dbhost --user dbuser --password dbpassword --port 5433 --database dbname
```

After configuring the database, execute the next command to create the appropriate structure:

```
gocan migrate-db
```

# A Few Minutes Tutorial

Let's use one of the examples in the book: analyzing Hibernate ORM.

```
gocan create-scene hibernate
gocan create-app orm -s hibernate
git clone https://github.com/hibernate/hibernate-orm.git
gocan import-history orm -s hibernate --after 2011-12-31 --before 2013-09-05 --path ./hibernate-orm
```

Run the UI to visualize the hotspots:

```
gocan ui --port 1234
```

* Open your browser at the appropriate location (here, it will be http://localhost:1234).
* Select the `hibernate` scene
* A short summary of the apps will be displayed
* Select the `orm` application
* The `Revisions` tab will be displayed: it might take a few seconds to display the chart, be patient (until we improve the performance) :-)

![Revisions](doc/images/revisions.png)

* Select the `Hotspots` tab to visualize the hotspots like in the book. You can zoom-in in the hierarchy by clicking on
the different zones.

![Hotspots](doc/images/hotspots.png)

* Mine the data with the `revisions` command:

```
gocan revisions orm -s hibernate
```

![Revisions](doc/images/mining-revisions.png)

We are going to focus on the `Configuration` class that has a big number of lines for a file related to configuration and that also has been revised quite a lot.

Let's analyze the complexity of this file by running the following command:

```
gocan create-complexity-analysis configuration-analysis --app orm --scene hibernate --directory /code/hibernate-orm/ --filename hibernate-core/src/main/java/org/hibernate/cfg/Configuration.java --spaces 4 
```

The `directory` argument specifies the local folder where the git repo can be found and the `filename` argument specified the location of the file to analyze relative to that directory.

The command will return the complexity calculated of that file for the specified time period.

![Complexity](doc/images/complexity2.png)

We can see here that there is a maximum number of 14 indentations in some line(s) which might indicated some complicated code there.

Now, if you go to the `Complexity` tab in the UI, you should be able to select the analysis we have just created and named `configuration-analysis`.

![Complexity](doc/images/complexity.png)

We can see how the complexity increased with time.

Now, let's take a look at change coupling: the files that have a tendency to be commited together, thanks to the following command.

```
gocan coupling orm -s hibernate -r 20
```

The `-r` flag allows to limit to the query to files that has been modified at least 20 times in average.

![coupling](doc/images/coupling.png)

# Building the app

## Requirements

* golang 1.16
* nodejs
* yarn

## Build

```
make build
```
# Troubleshooting

## Stopping the database

You might have some issue with the database when trying to start it after having being stopped.

You would get a message similar to that one:

```
Starting the embedded database...

FAILED
Cannot start the database: process already listening on port 5432
```

For some reason, stopping the database didn't work well, you will have to manually kill it
first and start it after.

The MacOS command to retrieve the process to kill is:

```
lsof -i tcp:5432
```


## Fail to import the history

If for some reason, the history import failed and when you run the command again, it complains
about some database key issue, the simplest action to do is to delete the application and
reimport the history.