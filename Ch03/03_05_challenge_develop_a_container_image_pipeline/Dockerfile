# Use an official Python runtime as a parent image
FROM python:3.11-slim-buster

# Set the working directory to /app
WORKDIR /app

# Copy the requirements file into the container at /app
COPY requirements.txt .

# Install the required packages
RUN pip install --no-cache-dir -r requirements.txt

# Copy the current directory contents into the container at /app
COPY . /app

# Expose port 3000 for the Flask web API
ENV FLASK_RUN_PORT=3000
EXPOSE 3000

# Define environment variable for Flask app
ENV FLASK_APP app.py

# Run the command to start the Flask app
CMD ["flask", "run", "--host=0.0.0.0"]

