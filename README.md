# Network Statistic Tracker

This is a small Go learning experiment which makes use of the httptrace features to track timing of different parts of an HTTP request and store them in a data structure

The end state is the following:

1. The user can provide an URL to hit, and the service will track all data points and store it in a database
2. Those statistic will be retrievable, and then can be used for whatever.

## Running the App

You can build by running 
```
docker build -t network_statistic_tracker .
```

Then run the app by running
```
docker run -p 8080:8080 network_statistic_tracker
```