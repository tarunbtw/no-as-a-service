# ❌ no-as-a-service

A fast, minimal API that returns random rejection reasons.
Built in Go — zero runtime dependencies, single binary, microsecond latency.

---

## 🚀 API

**Base URL**
```
https://no-as-a-service-5fv3.onrender.com
```

**Method:** `GET`
**Rate Limit:** `150 requests per burst per IP`

### Example Request
```bash
GET /no
```

### Example Response
```json
{
  "reason": "I once said yes. It didn't end well. There were llamas involved."
}
```

---

## 🛠️ Self Hosting

### 1. Clone the repository
```bash
git clone https://github.com/tarunbtw/no-as-a-service.git
cd no-as-a-service
```

### 2. Build and run
```bash
go build -o naas .
./naas
```

The API will be live at:
```
http://localhost:8080/no
```

Custom port:
```bash
PORT=5000 ./naas
```

### Or run with Docker
```bash
docker build -t no-as-a-service .
docker run -p 8080:8080 no-as-a-service
```

---

## 📁 Project Structure
```
no-as-a-service/
├── main.go          
├── handlers.go      
├── middleware.go    
├── ratelimit.go     
├── data.go          
├── reasons.json     
├── Dockerfile       
└── README.md
```

---
Do whatever, just don't say yes when you should say no.
