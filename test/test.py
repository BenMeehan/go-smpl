from flask import Flask

app = Flask(__name__)

@app.route("/")
def get_key_value_pairs():
    data = "key1:value1\nkey2:value2\nkey3:value3"
    return data

if __name__ == "__main__":
    app.run()
