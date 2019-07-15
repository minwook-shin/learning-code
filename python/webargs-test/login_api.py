from webargs import fields
from flask import request, Flask, jsonify
from webargs.flaskparser import parser

app = Flask(__name__)

login_field = {
    "username": fields.Str(required=True),
    "password": fields.Str(required=True, validate=lambda p: len(p) >= 5),
}


@app.route("/login", methods=["POST"])
def register():
    args = parser.parse(login_field, request)
    return jsonify({"id": args["username"], "password": len(args["password"]) * "*"})


if __name__ == "__main__":
    app.run()
