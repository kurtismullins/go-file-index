# Description

Crawls a (currently hard-coded) directory of files, grabs as much data as possible about each file, and dumps the data to a (currently hard-coded location) `db.json` file.

# Requirements

Requires `go-exif` package.

# To Do

* Clean up code -- this is my first Go App, and I'm not following any standards.
* Fix memory leak or whatever caused the app to use up all available memory on a test run.
* Find an alternative to using Linux/Unix's `file` command to grab the MimeType since
  that doesn't exist on Windows.
* Specify the *root* directory as a command-line or configuration file option.
* Other outputs including MongoDB, Elastic, STDOut, or allow the user to specify an
  output file.
* Use concurrency / async operations
* Optimize heavy operations (e.g. MD5 Sum generation)
* Better exception handling
