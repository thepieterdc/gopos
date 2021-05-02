# gopos

Address and place information service in Go.

## Features

### Given a coordinate pair, find the timezone.

**Request:**

```http request
GET /timezone?latitude=37.97153995920827&longitude=23.726713776643596
```

**Example response:**

```json
{
  "latitude": 37.97153995920827,
  "longitude": 23.726713776643596,
  "timezone": "Europe/Athens"
}
```