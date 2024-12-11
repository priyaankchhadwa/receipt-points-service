# Receipt Points Service

## Prerequisites
- Go 1.23
- Docker (optional)

## Setup and Running

### Local Setup
1. Clone the repository
2. Navigate to the project directory
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

### Docker Setup
1. Build the Docker image:
   ```bash
   docker build -t receipt-points-service .
   ```
2. Run the Docker container:
   ```bash
   docker run -p 8080:8080 receipt-points-service
   ```

## Endpoints
- `POST /receipts/process`: Submit a receipt
- `GET /receipts/{id}/points`: Retrieve points for a processed receipt

## Example Curl Commands
### Process a Receipt
```bash
curl -X POST http://localhost:8080/receipts/process \
     -H "Content-Type: application/json" \
     -d '{
            "retailer": "M&M Corner Market",
            "purchaseDate": "2022-03-20",
            "purchaseTime": "14:33",
            "items": [
                {
                "shortDescription": "Gatorade",
                "price": "2.25"
                },{
                "shortDescription": "Gatorade",
                "price": "2.25"
                },{
                "shortDescription": "Gatorade",
                "price": "2.25"
                },{
                "shortDescription": "Gatorade",
                "price": "2.25"
                }
            ],
            "total": "9.00"
     }'
```

### Get Points
```bash
curl http://localhost:8080/receipts/{id}/points
```

## Notes
- Receipts and points are stored in memory and do not persist between application restarts