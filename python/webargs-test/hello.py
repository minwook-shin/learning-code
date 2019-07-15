from flask import Flask
from webargs import fields
from webargs.flaskparser import use_args

app = Flask(__name__)

hello_field = {"hello": fields.Str(required=True)}


@app.route("/")
@use_args(hello_field)
def index(args):
    return "Hello " + args["hello"]


if __name__ == "__main__":
    app.run()
