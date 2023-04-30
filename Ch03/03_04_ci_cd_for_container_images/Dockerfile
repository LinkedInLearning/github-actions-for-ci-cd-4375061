# Use an official Python runtime as a parent image
FROM python:alpine3.16

# Install tini to manage processes
RUN apk add --no-cache tini

# Set the environment variables
ENV TINI_SUBREAPER=1

# Set the working directory to /workspace
WORKDIR /workspace

# Copy the project files into the container at /workspace
COPY requirements.txt \
     main.py \
     main_test.py \
     data.json .

# Install the requirements
RUN pip install --no-cache-dir --requirement requirements.txt

# Expose port 3000
EXPOSE 3000

# Start the server using tini, then uvicorn
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "3000"]

