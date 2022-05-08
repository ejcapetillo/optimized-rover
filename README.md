# optimized-rover
The main goal of this project is to show the different ways that we can retrieve and find the most imaged days on Mars using NASA's Mars Rover API.

This will be accomplished in a few ways.

The first will be a synchronous retrieval of image metadata through the NASA API.

The second will be through the use of goroutines.

The third will be with the aid of a database that will shorten the data retrieval time after the first run.

Each of these steps will be built on top of the previous step and benchmarked to see how execution time improves.
