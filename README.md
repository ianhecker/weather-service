# Weather Service

This server fetches weather from the National Weather Service
https://www.weather.gov/documentation/services-web-api

## Weather APIs

There are 2 API calls made to the National Weather Service. They fetch:
1. The grid zone for latittude & longitude coordinates
2. 12 hour forecast of the grided zone (*based on previous latitude & longitude*)

The APIs called, in order:
+ https://api.weather.gov/points/{latitude},{longitude}
+ https://api.weather.gov/gridpoints/{office}/{gridX},{gridY}/forecast

### Grid Forecast for a Point Location
https://api.weather.gov/points/{latitude},{longitude}

> "To obtain the grid forecast for a point location, use the `/points` endpoint
to retrieve the current grid forecast endpoint by coordinates"

### 12 Hour Forecast for a 2.5km x 2.5km Grid
https://api.weather.gov/gridpoints/{office}/{gridX},{gridY}/forecast

> "Forecasts are created at each NWS Weather Forecast Office (WFO) on their own
grid definition, at a resolution of about 2.5km x 2.5km. The API endpoint for
the 12h forecast periods at a specific grid location is formatted as:"

## Server

### Run

```bash
go run main.go
```

Press Ctrl+C to stop

### Requests

The server requests a weather forecast of a latitude & longitude, and the
National Weather Service responds with a

The endpoint is: **"/weather?latitude=&longitude="**

An example:
```bash
curl "localhost:8080/weather?latitude=39.7456&longitude=-97.0892"
```

Or with httpie!
*See https://httpie.io/docs/cli to install the tool*:
```bash
	 http :8080/weather latitude==39.7456 longitude==-97.0892
```

### Response

The server returns JSON of the weather forecast, and looks like:
```json
{
  "name": "This Afternoon",
  "startTime": "2025-08-13T16:00:00-05:00",
  "endTime": "2025-08-13T18:00:00-05:00",
  "isDaytime": true,
  "temperature": 89,
  "temperatureUnit": "F",
  "windSpeed": "10 mph",
  "windDirection": "SE",
  "shortForecast": "Sunny",
  "detailedForecast": "Sunny, with a high near 89. Southeast wind around 10 mph.",
  "characterization": "hot"
}
```
A characterization of the temperature is also returned, and can be:
+ "hot"
+ "cold"
+ "moderate"

and follows the logic:
```golang
switch {
case temperature > 75:
	return Hot
case temperature < 30:
	return Cold
default:
	return Moderate
}
```
