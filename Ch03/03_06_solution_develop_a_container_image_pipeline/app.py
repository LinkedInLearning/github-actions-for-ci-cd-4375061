from fastapi import FastAPI, HTTPException
from fastapi.responses import HTMLResponse
from random import choice
from re import match
import locations

PORT = 3000

app = FastAPI()


@app.get('/favicon.ico', include_in_schema=False, status_code=204)
async def favicon():
    return True


@app.get("/", response_class=HTMLResponse)
async def index():
    return """
<html>
    <head>
        <title>Welcome to the Location Service API!</title>
    </head>
    <body>
        <font size="7">
            <b>Welcome to the Location Service API!</b>
        </font>
        <font size="5">
            <p>This API provides a backend service for a mobile application.</p></font>
            <p>Eventually, it will return a location based on a given IP address.</p>
            <p>Right now all it does is return a random location! :D</p>
        </font>
        <font size="4">
            <p>This API is implemented using FastAPI, a modern, high-performance web framework for building APIs with Python.</p>
            <p>FastAPI is designed to provide:</p>
            <ul>
                <li>ease of use</li>
                <li>high performance</li>
                <li>automatic validation of request and response data</li>
                <li>eautomatic generation of OpenAPI and JSON Schema documentation</li>
                <li>many other features.</li>
            </ul>
            </p>
        </font>
        <font size="7">
            <b>Using the Location Service API</b>
        </font>
        <font size="4">
            <p>Submit a request to the API using the following format: 'http://hostname/IP_ADDR'.</p>
            <p>The IP_ADDR value must be a correctly formatted IP address.</p>
            <p>For example: <a href="/127.0.0.1" target="_blank">http://HOSTNAME_AND_PORT/127.0.0.1</a></p>
            <p>You can access the automatic API documentation and interactive API explorer provided by FastAPI at the following links:</p>
            <ul>
                <li>Developer Panel: <a href="/docs" target="_blank">/docs</a></li>
                <li>Documentation: <a href="/redoc" target="_blank">/redoc</a></li>
            </ul>
        </font>
    </body>
</html>
"""


@app.get("/{ip_address}")
async def get_location(ip_address: str):
    # validate the IP address using a regular expression
    if not match(r"^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$", ip_address):
        raise HTTPException(status_code=400, detail="Invalid IP address.")

    random_city = choice(locations.cities)

    return {"ip_address": ip_address, "location": random_city}


if __name__ == "__main__":
    import uvicorn

    uvicorn.run(app, host="0.0.0.0", port=PORT)
