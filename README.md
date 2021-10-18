# gopos

[![Release](https://github.com/thepieterdc/gopos/actions/workflows/release.yml/badge.svg)](https://github.com/thepieterdc/gopos/actions/workflows/release.yml)

_From Ancient Greek τόπος (tópos, “place”)._

Address and place information service in Go.

---

## Running via docker
To use this in your own project:
```shell
docker run --rm ghcr.io/thepieterdc/gopos:latest
```

## Running via docker compose
Using the `docker-compose.yml` file in this repository:
```shell
docker-compose up --build
```

## Dependencies
This module requires [libpostal](https://github.com/openvenues/libpostal) for address resolving.

## Features

### Parse an input string into a formatted address.
If the country of the address is known, it is recommended to pass this along in the request to improve the formatting process (see the second example).

This route depends on `libpostal`.

**Example request:**

```http request
GET /address/parse?query=Apple%2010955%20N%20Tantau%20Ave,%20Cupertino,%20CA%2095014,United%20States
```

**Example request (explicitly providing the country):**

```http request
GET /address/parse?country=fr&query=Champ%20de%20Mars%2C%205%20Av.%20Anatole%20France%2C%2075007%20Paris
```

**Response:**

```json
{
  "city": "cupertino",
  "country": "united states",
  "house": "apple",
  "house_number": "10955",
  "postcode": "95014",
  "road": "n tantau ave",
  "state": "ca"
}
```

### Lookup a Google Place ID.

The result of each call can be cached to a database to avoid costs when sending repeated requests. 

_This route requires [environment variables](CONFIGURATION.md)._

**Example request:**

```http request
GET /google/place/ChIJ37HL3ry3t4kRv3YLbdhpWXE
```

**Response:**

```json
{
  "address_components": [
    {
      "long_name": "1600",
      "short_name": "1600",
      "types": [
        "street_number"
      ]
    },
    {
      "long_name": "Pennsylvania Avenue Northwest",
      "short_name": "Pennsylvania Avenue NW",
      "types": [
        "route"
      ]
    },
    {
      "long_name": "Northwest Washington",
      "short_name": "Northwest Washington",
      "types": [
        "neighborhood",
        "political"
      ]
    },
    {
      "long_name": "Washington",
      "short_name": "Washington",
      "types": [
        "locality",
        "political"
      ]
    },
    {
      "long_name": "District of Columbia",
      "short_name": "DC",
      "types": [
        "administrative_area_level_1",
        "political"
      ]
    },
    {
      "long_name": "United States",
      "short_name": "US",
      "types": [
        "country",
        "political"
      ]
    },
    {
      "long_name": "20500",
      "short_name": "20500",
      "types": [
        "postal_code"
      ]
    }
  ],
  "business_status": "OPERATIONAL",
  "formatted_address": "1600 Pennsylvania Avenue NW, Washington, DC 20500, USA",
  "geometry": {
    "location": {
      "lat": 38.8976763,
      "lng": -77.0365298
    }
  },
  "icon": "https://maps.gstatic.com/mapfiles/place_api/icons/v1/png_71/civic_building-71.png",
  "name": "The White House",
  "place_id": "ChIJ37HL3ry3t4kRv3YLbdhpWXE",
  "plus_code": {
    "compound_code": "VXX7+39 Washington, DC, USA",
    "global_code": "87C4VXX7+39"
  },
  "types": [
    "tourist_attraction",
    "point_of_interest",
    "establishment"
  ],
  "url": "https://maps.google.com/?cid=8167675777476425407",
  "utc_offset": -240,
  "vicinity": "1600 Pennsylvania Avenue Northwest, Washington"
}
```

### Health check.

**Example request:**

```http request
GET /health
```

**Response:**

```json
{
  "status": true
}
```

### Prometheus metrics.

**Example request:**

```http request
GET /metrics
```

**Response (truncated):**

```text
gopos_request_duration_seconds_count{code="200",method="GET",url="/timezone"} 5
gopos_request_duration_seconds_count{code="422",method="GET",url="/timezone"} 1
...
```

### Find the timezone of the given location.

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

## Releasing
This process is automated via GitHub Actions. In order to make a new release, trigger the `Release` workflow.
