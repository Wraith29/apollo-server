import std/[httpclient, json]

const baseUrl = "http://localhost:5000/"

type
  AuthBody = ref object
    authToken: string

  RatingBody = ref object
    albumId: string
    rating: int

proc getAuthToken(client: HttpClient): string =
  const url = baseUrl & "auth/login"

  let response = client.post(url, $(%*{"username": "wraith", "password": "ActionDog2002!"}))

  let body = to(response.body().parseJson(), AuthBody)

  return body.authToken

proc sendAddArtistRequest(client: HttpClient; artistName, authToken: string): void =
  const url = baseUrl & "artist"

  client.headers.add("Authorization", authToken)
  let response = client.post(url, $(%*{"artistName": artistName}))

  echo response.status

proc main(): void =
  let client = newHttpClient()
  let authToken = getAuthToken(client)

  client.headers.add("Authorization", authToken)

  let albumRecommendation = client.get(baseUrl & "album/recommendation")

  echo albumRecommendation.body()
  
  let response = json.parseJson(albumRecommendation.body())

  let albumId = response["AlbumId"]

  let res = client.putContent(baseUrl & "album/rating", $(%*{"albumId": albumId, "rating": 5}))

  echo res
  

when isMainModule:
  main()

