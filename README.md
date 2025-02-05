# Apollo

Music Recommendation App.

Add all of your favourite artists, and ones you want to get to know better - and simply run `apollo recommend` to get a random Album to listen to.

> [!TODO]
> Investigate the potential migration from [gorm](https://gorm.io) to [sqlc](https://github.com/sqlc-dev/sqlc) instead

## Features

- [x] [Add](#add-new-artist)
- [x] [Recommend](#recommend)
- [x] [Rate](#rate)
- [ ] [List](#list)
- [ ] [Update](#update)

### Add new artist

Usage:

```sh
apollo add <artist_name>
```

Add a new artist into your library.

### Recommend

Usage:

```sh
apollo recommend [genres...] [-l / --listened]
```

Aliases:

```sh
apollo rec
```

Recommend an album from your library.

By default only recommends album's you haven't marked as listened yet, and includes all genres in your library.

Pass the `-l` flag to include albums you've already heard, and pass in up to 3 genre filters to narrow down your recommendations

### Rate

Usage:

```sh
apollo rate (1 | 2 | 3)`
```

Rate the most recent recommendation given to you by apollo.

#### Ratings

| Rating | Meaning |
| ------ | ------- |
| 1      | I didn't like this album |
| 2      | This album was ok |
| 3      | I liked this album |

### List

The following commands will all list 10 of the specified item, unless the `[-a / --all]` flag is provided.

#### Artists

Usage:

```sh
apollo list artist
```

#### Genres

Usage:

```sh
apollo list genres
```

#### Recommendations

Usage:

```sh
apollo list recommendations
```

Or

```sh
apollo list recs
```

### Update

This command will go through each of your artists, and check to see if they have released any new albums.

> [!Dev Note]
> Need to consider the best way to do this while keeping MusicBrainz rate limit (1 req/s) in mind.
> May need to try and run this as a background process with some sort of debouncer to prevent hitting the rate limit,
> without blocking the user for (artist count) seconds

