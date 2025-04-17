import std/[httpclient, json]

const baseUrl = "http://localhost:5000/"

type AuthBody = ref object
  authToken: string

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
  let authResponse = getAuthToken(client)

  for artist in @[
      "Muse", "You me at 6", "Royal Blood", "The Amazons",
      "Polaris", "The 1975", "Architects", "Spiritbox",
      "Bad Omens", "Boston Manor", "Bring me the Horizon",
      "The Clause", "The Sheratons", "Dead Pony", "Lewis Capaldi",
      "My Chemical Romance", "Paramore", "Pierce the Veil",
      "Sam Fender", "Sea Girls"
  ]:
    sendAddArtistRequest(client, artist, authResponse)

when isMainModule:
  main()

