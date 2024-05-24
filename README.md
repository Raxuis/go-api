````markdown
# How to Make a POST Request

This document explains how to make a POST request to add a new book entry to a server. We will use the following JSON data as an example:

```json
{
  "id": "5",
  "title": "Romeo and Juliet",
  "author": "William Shakespeare",
  "quantity": 4
}
```
````

## Possibilities:

- Use Postman
- Use cURL

### Using Postman

Postman is a popular tool for testing APIs. Follow these steps to make a POST request using Postman:

1. **Open Postman**: Launch the Postman application on your computer.
2. **Create a New Request**: Click on the "New" button and select "Request".
3. **Set the Request Type**: Change the request type to `POST` from the dropdown menu next to the URL field.
4. **Enter the URL**: Type `http://localhost:8080/books` in the URL field.
5. **Set the Headers**: Go to the "Headers" tab and add a new header with the key `Content-Type` and the value `application/json`.
6. **Add the JSON Body**: Go to the "Body" tab, select "raw" and ensure "JSON" is selected from the dropdown. Copy and paste the JSON data provided above into the body section.
7. **Send the Request**: Click the "Send" button to make the request. If the server is running and set up correctly, you should receive a response indicating that the book has been added successfully.

### Using cURL

cURL is a command-line tool for making network requests. You can use it to make a POST request with the following command:

1. **Create the JSON File**: Save the JSON data in a file named `body.json`.

   ```json
   {
     "id": "5",
     "title": "Romeo and Juliet",
     "author": "William Shakespeare",
     "quantity": 4
   }
   ```

2. **Open the Terminal**: Open your terminal or command prompt.
3. **Run the cURL Command**: Use the following cURL command to make the POST request:

   ```sh
   curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
   ```

   - `localhost:8080/books`: This is the URL where the POST request is sent.
   - `--include`: This option includes the HTTP response headers in the output.
   - `--header "Content-Type: application/json"`: This sets the `Content-Type` header to `application/json`.
   - `-d @body.json`: This specifies the data to be sent with the request, which is read from the `body.json` file.
   - `--request "POST"`: This specifies that the request method is POST.

After running the command, you should see a response in the terminal indicating the result of the request.

## Conclusion

By following the steps above, you can make a POST request to add a new book to the server using either Postman or cURL. Make sure your server is running and accessible at `localhost:8080` before making the request.

```
This revised version ensures clarity and correctness in explaining the two methods for making a POST request.
```

# How to Make a PATCH Request

This document explains how to make a PATCH request to update an existing book entry on a server. We will use the following JSON data as an example:

```json
{
  "quantity": 3
}
```

## Possibilities:

- Use Postman
- Use cURL

### Using Postman

Postman is a popular tool for testing APIs. Follow these steps to make a PATCH request using Postman:

1. **Open Postman**: Launch the Postman application on your computer.
2. **Create a New Request**: Click on the "New" button and select "Request".
3. **Set the Request Type**: Change the request type to `PATCH` from the dropdown menu next to the URL field.
4. **Enter the URL**: Type `http://localhost:8080/checkout?id=2` in the URL field.
5. **Set the Headers**: Go to the "Headers" tab and add a new header with the key `Content-Type` and the value `application/json`.
6. **Add the JSON Body**: Go to the "Body" tab, select "raw" and ensure "JSON" is selected from the dropdown. Copy and paste the JSON data provided above into the body section.
7. **Send the Request**: Click the "Send" button to make the request. If the server is running and set up correctly, you should receive a response indicating that the book entry has been updated successfully.

### Using cURL

cURL is a command-line tool for making network requests. You can use it to make a PATCH request with the following command:

1. **Open the Terminal**: Open your terminal or command prompt.
2. **Run the cURL Command**: Use the following cURL command to make the PATCH request:

   ```sh
   curl localhost:8080/checkout?id=2 --request "PATCH"
   ```

   - `localhost:8080/checkout?id=2`: This is the URL where the PATCH request is sent, including the query parameter `id` to specify the book entry to update.
   - `--request "PATCH"`: This specifies that the request method is PATCH.

After running the command, you should see a response in the terminal indicating the result of the request.

## Conclusion

By following the steps above, you can make a PATCH request to update an existing book entry on the server using either Postman or cURL. Make sure your server is running and accessible at `localhost:8080` before making the request.
