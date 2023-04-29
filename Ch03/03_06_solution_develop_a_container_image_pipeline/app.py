from flask import Flask, jsonify
import random
import re
import locations

app = Flask(__name__)


@app.route("/<ip_address>")
def get_location(ip_address):
    # validate the IP address using a regular expression
    if not re.match(r"^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$", ip_address):
        return jsonify({"error": "Invalid IP address."}), 400

    random_city = random.choice(locations.cities)

    return jsonify({"ip_address": ip_address, "location": random_city})


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=3000)
