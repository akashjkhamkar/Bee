from flask import Flask, request, jsonify
from function import entry

app = Flask(__name__)

@app.route('/', methods=['POST', 'GET'])
def hello():
    response = entry(request)
    return jsonify(
        status="ok",
        result=response
    )

if __name__ == "__main__":
    app.run(debug=False, host="0.0.0.0", port=8000)