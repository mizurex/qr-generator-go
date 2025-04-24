# QR Code Generator API

This is a simple API built using Go and the Gin framework that generates QR codes from URLs. The API allows users to submit a URL and receive a generated QR code image in return. The image is stored and can be accessed via a generated URL.

## Features

- **Generate QR Code**: Submit a URL and get a QR code as an image.
- **Error Handling**: Handles errors like missing URL or server issues gracefully.
- **Store Images**: The generated QR codes are saved locally and can be accessed through a URL.

## Prerequisites

- Go (1.16 or higher)
- Gin Framework
- Go QR Code Library ([go-qrcode](https://github.com/skip2/go-qrcode))

## Installation

1. Clone the repository to your local machine:

    ```bash
    git clone https://github.com/yourusername/qr-code-generator.git
    ```

2. Navigate to the project folder:

    ```bash
    cd qr-code-generator
    ```

3. Install the dependencies:

    ```bash
    go mod tidy
    ```

4. Run the server:

    ```bash
    go run main.go
    ```

   The server will start running at `http://localhost:9090`.

## API Usage

### Generate QR Code

- **Endpoint**: `/generate`
- **Method**: `GET`
- **Query Parameter**: `url` (The URL for which the QR code will be generated)

#### Example Request

```bash
http://localhost:9090/generate?url=https://www.example.com


{
  "message": "QR code generated successfully",
  "qr_image_url": "http://localhost:9090/images/qr_1625497287_123.png"
}

{
  "message": "enter valid url"
}

{
  "error": "cannot generate qr"
}


### Notes:
1. Replace `yourusername` with your GitHub username or the repository name if you are sharing this on GitHub.
2. Make sure the paths in your `README` reflect your project structure accurately.

Feel free to adjust or add any specific details about your project!
