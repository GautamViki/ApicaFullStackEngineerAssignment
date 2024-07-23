Introduction:
    These API provides basic Create, Read, Update, and Delete (CRUD) operations for managing a collection of resources. It is built using Golang.

Features:
    List the main features of your API. For example:
        1. Set/update the value in to cache
        2. Get the value from cache
        3. Delete resources

Getting Started:
    Prerequisites:
        Golang (version 1.22.x)

Installation:
    1. Clone the repository:
        gh repo clone GautamViki/ApicaFullStackEngineerAssignment
    2. Install dependencies:
        go mod tidy

Running the Application:
    go run main.go

Usage:
    POST: http://localhost:3009/lru?key={key}&value={value}
        - Set the cache key
        curl --location --request POST 'http://localhost:3009/lru?key=20&value=30'
        Response: 
            {
                "code": "0",
                "message": "Set Value in cache key successfully.",
                "messages": []
            }

    GET: http://localhost:3009/lru
        - Get all cache data
        curl --location 'http://localhost:3009/lru'
        Response: 
        {
            "code": "0",
            "message": "Cache Fetched Successfully.",
            "messages": [],
            "Lrus": [
                {
                    "Key": 20,
                    "Value": 30
                }
            ]
        }

    GET: http://localhost:3009/lru/{key}
        - Get cache data by key
        curl --location 'http://localhost:3009/lru/20'
        Response: 
        {
            "code": "0",
            "message": "Key Fetched Successfully.",
            "messages": [],
            "LRU": {
                "Key": 20,
                "Value": 30
            }
        }

    DELETE: http://localhost:3009/lru/{key}
        - Delete cache data by key
        curl --location --request DELETE 'http://localhost:3009/lru/20'
        Response:
        {
            "code": "0",
            "message": "Cache key deleted successfully.",
            "messages": []
        }
