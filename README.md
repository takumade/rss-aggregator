# RSS Aggregator

![RSS Aggregator Image](./RSS%20Aggregator.png)

RSS Aggregator is a simple RSS feed reader that fetches RSS feeds from multiple sources and stores them in a database. It is built using Go and Postgres. It exposes a REST API to fetch users, feeds, feed follows, posts.

## Tech Stack
- Go
- Postgres


## Installation

1. Clone this repo:

```sh
git clone https://github.com/takumade/rss-aggregator.git
```

2. Change directory to the project folder:

```sh
cd rss-aggregator
```

3. Create a `.env` file in the root of the project and add the following environment variables:

```sh
PORT=POST_HERE
DB_URL=DB_URL_HERE
```

4. Install the dependencies:

```sh
go mod tidy
```

5. Run the project:

```sh
go build && ./rss-aggregator
```

## API Endpoints
Check the `RSS_Aggregator_API(Thunder Client).json` file for the API endpoints. Its in the root folder

## Other useful commands
### Create migration 

To add schema migration go to `./sql/schema` and create a file and add your changes. 

Then run the following command in root folder:

```sh
./migrate up
```

### Add query

To add a query go to `./sql/queries` and create a file and add your query. 

Then run the following command in root folder to create a query file:

```sh
./generate_query
```

## Special Thanks To
[WagsLane](https://github.com/wagslane)