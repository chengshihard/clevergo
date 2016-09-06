// RESTful API Controller Example.

package main

import (
	"github.com/headwindfly/clevergo"
	"html/template"
)

var (
	html = `<html>
	<head></head>
	<body>
		<h3>RESTful API Controller Example.</h3>

		<h4>Requests</h4>
		<ul>
			<li><a target="_blank" href="javascript:get();">GET</a></li>
			<li><a target="_blank" href="javascript:post('POST');">POST</a></li>
			<li><a target="_blank" href="javascript:post('DELETE');">DELETE</a></li>
			<li><a target="_blank" href="javascript:post('PUT');">PUT</a></li>
			<li><a target="_blank" href="javascript:post('HEAD');">HEAD</a></li>
			<li><a target="_blank" href="javascript:post('OPTIONS');">OPTIONS</a></li>
			<li><a target="_blank" href="javascript:post('PATCH');">PATCH</a></li>
		</ul>

		<h4>Result:</h4>
		<textarea rows="5" cols="100" id="result"></textarea>
		<br>

		<script>
			var resultEle = document.getElementById("result");

			var get = function(){
				resultEle.value = 'Pending';
				xmlHttp = new XMLHttpRequest();
    				xmlHttp.open("GET", '/users');
    				xmlHttp.send(null);
    				xmlHttp.onreadystatechange = function () {
        				resultEle.value = "GET: " + xmlHttp.responseText;
    				}
			}

			var post = function(type){
				resultEle.value = 'Pending';
				var url = '/users';
				switch(type){
					case 'POST':
						break;
					case 'DELETE':
						url += '?_method=DELETE';
						break;
					case 'PUT':
						url += '?_method=PUT';
						break;
					case 'HEAD':
						url += '?_method=HEAD';
						break;
					case 'OPTIONS':
						url += '?_method=OPTIONS';
						break;
					case 'PATCH':
						url += '?_method=PATCH';
						break;
				}

				xmlHttp = new XMLHttpRequest();
    				xmlHttp.open("POST", url);
				xmlHttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");
				xmlHttp.send();
    				xmlHttp.onreadystatechange = function () {
        				resultEle.value = type + ": " + xmlHttp.responseText;
    				}
			}
		</script>
	</body>
	</html>`
	tpl = template.Must(template.New("").Parse(html))
)

type userController struct {
	clevergo.Controller
	allowOrigin  string
	allowMethods string
}

func NewUserController() userController {
	return userController{
		allowOrigin:  "*",
		allowMethods: "GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD",
	}
}

func (c userController) Handle(next clevergo.Handler) clevergo.Handler {
	return clevergo.HandlerFunc(func(ctx *clevergo.Context) {
		// Do anything what you want.
		ctx.Text("Prepare.\n")

		// Set Access-Control-Allow-Origin and Access-Control-Allow-Methods for ajax request.
		ctx.Response.Header.Set("Access-Control-Allow-Origin", c.allowOrigin)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", c.allowMethods)

		// Using param named '_method' to simulate the other request, such as PUT, DELETE etc.
		if !ctx.IsGet() {
			switch string(ctx.FormValue("_method")) {
			case "PUT":
				c.PUT(ctx)
				return
			case "DELETE":
				c.DELETE(ctx)
				return
			case "HEAD":
				c.HEAD(ctx)
				return
			case "OPTIONS":
				c.OPTIONS(ctx)
				return
			case "PATCH":
				c.PATCH(ctx)
				return
			}
		}

		// Invoke the request handler.
		next.Handle(ctx)
	})
}

func (c userController) GET(ctx *clevergo.Context) {
	ctx.Text("GET REQUEST.\n")
}

func (c userController) POST(ctx *clevergo.Context) {
	ctx.Text("POST REQUEST.\n")
}

func (c userController) DELETE(ctx *clevergo.Context) {
	ctx.Text("DELETE REQUEST.\n")
}

func (c userController) PUT(ctx *clevergo.Context) {
	ctx.Text("PUT REQUEST.\n")
}

func (c userController) OPTIONS(ctx *clevergo.Context) {
	ctx.Text("OPTIONS REQUEST.\n")
}

func (c userController) PATCH(ctx *clevergo.Context) {
	ctx.Text("PATCH REQUEST.\n")
}

func (c userController) HEAD(ctx *clevergo.Context) {
	ctx.Text("HEAD REQUEST.\n")
}

func index(ctx *clevergo.Context) {
	ctx.SetContentTypeToHTML()
	tpl.Execute(ctx, nil)
}

func main() {
	app := clevergo.NewApplication()

	// Create a router instance.
	router := clevergo.NewRouter()
	app.AddRouter("", router)

	// Register route handler.
	router.GET("/", clevergo.HandlerFunc(index))
	router.RegisterController("/users", NewUserController())

	// Start server.
	app.Run()
}
