# Checkout API test exercise
golang version: `v1.16.5`

## Run tests
### By local golang
```bash
go test ./tests/...
```

### By docker
```bash
docker build -t {img_name}:{img_tag} .
docker run --rm {img_name}:{img_tag}
```

## Improvements
- Better to mocked provider api response.