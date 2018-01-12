# motominderapi

# Background
I have a long history using variants of the C programming language, and the similarities between C and GO spiked my interest in the latter.  I am a staunch advocate for test driven development and Uncle Bob's Screaming Clean Architecture.  During my sebatical between contracted software assignments, I decided to create a web service in GO that will record the maintenance events and suggest required services for a fleet of motorcycles.  Since the vast majority of the operations will manipulate textual JSON objects, a RESTful web service will be created.  

I have used Google RPC ("grpc") on a  few web service projects where the it was better to send data in a binary format, rather than as megabytes of JSON.  A good example for the use of grpc is when I needed to draw a line on a satellite image to indicate the path of a machine.  The image was sent in binary format (png), so it was easy to draw the line and return the updated image to the caller.  I will probably integrate a few grpc calls in this API, too.

