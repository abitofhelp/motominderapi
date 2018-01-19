# motominder API

# Background
Motominderapi is a web service that I created in Go as a good template for other projects.  It implements a RESTful API for manipulating a repository of motorcycles.  The RESTful endpoints will use JSON.  

I am going to implement an endpoint using Google RPC ("grpc"), which is a fantastic technology for working with data in binary format rather than dealing with megabytes of textual JSON.  A good example for the use of grpc was when I needed to draw a line on a satellite image to indicate the path of a machine.  In the grpc call, the image was sent in binary format (png), so it was easy to draw the line and return the updated image to the caller. 

