# Short URL Service
Service for shortening long urls

## Commands
Start app:
```
docker-compose up
```

Run tests:
```
make test-cover
```

## API
| Method | Path | Description |
| ------ | ---- | ----------- |
| POST | / | Create new short url |
| GET | /{short_url} | Go to original url |