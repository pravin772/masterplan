# Masterplan

Masterplan of activities used in construction. This is coding task.

## Env setup

You need to install [Golang](https://golang.org) and [MySQL](https://www.mysql.com) to use Masterplan.

Once you install both requirements.

Clone GitHub repo

```bash
git clone https://github.com/pravin772/mp-api.git
cd mp-api
```

Create .env file with DB_URL variable that will hold database URL/URI or just put URI in main.go

```go
model.DB, model.Err = gorm.Open("mysql", os.Getenv("DB_URL"))
```

## Usage

Navigate to mp-api folder (if you not)

```bash
cd mp-api
go run main.go
```

Server will be running on http://localhost:8000

## API endpoints

To create Activity in Masterplan do post request

POST http://localhost:8000/addData

Body

```json
{
  "SrNo": "1.5",
  "Activity": "Boundary wall",
  "StartDate": "2018-08-01T00:00:01+10:00",
  "EndDate": "2018-09-02T00:00:01+10:00"
}
```

To get all the Activities in JSON array

GET http://localhost:8000/getAll

To download Activities in csv file by WBS no. do get request

GET http://localhost:8000/dcsv

To download Activities in csv file by StartDate (MM-DD-YYY) do get request

GET http://localhost:8000/dcsvbystartdate

## Help

Please contact to pravinbendre772@gmail.com if you wish to suggest updates or having issue while running.
