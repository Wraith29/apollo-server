# Apollo

Music Recommendation App.

Add all of your favourite artists, and ones you want to get to know better - and simply run `apollo recommend` to get a random Album to listen to.

## Features

- [x] [Add](#-add-new-artist)
- [x] [Recommend](#-recommend)
- [x] [Rate](#-rate)

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
