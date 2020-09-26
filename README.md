# Short URL Service
Service for shortening long urls

## Installation
```
git clone https://github.com/maratishimbaev/short-url-service.git
```

## Usage
Run server:
```
docker-compose up
```

Run tests:
```
make test-cover
```

## API
Method | Path | Parameters | Description
------ | ---- | ---------- | -----------
POST | / | body | Create new short url
GET | /{short_url} | - |Go to original url

Create url body example:
```json
{
  "old_url": "https://google.com",
  "new_url":  "goo"
}
```