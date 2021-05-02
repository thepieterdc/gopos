# gopos

Address and place information service in Go.

## Features

### Format a given address.

_This route depends on [libpostal](https://github.com/openvenues/libpostal)._

**Example request:**

```http request
GET /format?input=Google%20California
```

**Response:**

```json
{}
```

### Given a coordinate pair, find the timezone.

**Example request:**

```http request
GET /timezone?latitude=37.97153995920827&longitude=23.726713776643596
```

**Response:**

```json
{
  "latitude": 37.97153995920827,
  "longitude": 23.726713776643596,
  "timezone": "Europe/Athens"
}
```