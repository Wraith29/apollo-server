# TODO

- [x] Auth
  - [x] Login
  - [x] Register
- [ ] Artist
  - [x] Add
  - [x] Update (Refreshes their data from MusicBrainz i.e. check for new albums / tags)
  - [ ] Remove
- [x] Recommendation
  - [x] Get
- [x] List
  *These will include the users rating*
  - [x] Artists
  - [x] Albums
  - [ ] Genres
  - [x] Recommendations

- [ ] General unit tests
- [ ] Api testing

- [ ] Improve logging / error handling

> Rather than directly returning errors to the user
> The app should handle them internally, logging
> Them out to a file, and then returning
> A relevant error to the user (i.e. information on the failure for bad req)
> Or what has conflicted, or generic error message for internal server
