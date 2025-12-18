# Use a slim Python image to keep it lightweight
FROM python:3.11-slim

# Set the working directory
WORKDIR /app

# Copy requirements and install them
COPY . /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy the rest of your application code
COPY . .

# Run the script
CMD ["python","main.py"]