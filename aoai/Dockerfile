# Use an official Python runtime as a parent image
FROM python

 

# Set the working directory in the container
WORKDIR /usr/src/app

 

# Copy the requirements file into the container at /usr/src/app
COPY requirements.txt ./

 

# Install any needed packages specified in requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

 

# Copy the current directory contents into the container at /usr/src/app
COPY . .

 

# Expose the port the app runs on
EXPOSE 8000

 

# Define the command to run your Django application
CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]