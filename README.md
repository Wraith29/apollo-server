# Apollo

A personal music library management tool, in the console.

---

## Overview

Apollo is a CLI tool which allows you to manage your favourite artists.

### Add an Artist

Apollo uses [MusicBrainz](https://musicbrainz.org/) behind the scenes.

```sh
apollo add "An Artist"
```

The `add` command allows you to add the given artist + all of their albums to your library.

> [!NOTE]
> If the artist is coming up as someone other than you expect,
> Try running the command with the `[-i / --interactive]` flag.
> This will prompt you with the top 3 results from MusicBrainz and you can pick which one matches

### Get an Album Recommendation

```sh
apollo recommend [genres...]
```

The `recommend` command will give you a random album (It will prioritise albums you haven't listened to first).

Optionally pass in up to 3 genres to get a filtered recommendation

### Rate an album

Rate your most recent recommendation with this command.

```sh
apollo rate (1-3)
```

The `rate` command allows you to rate the most recent recommendation from apollo.

The ratings are as follows:

| Value | Meaning |
| ----- | ------- |
| 1     | I didn't like this album |
| 2     | This album was OK |
| 3     | I liked this album |

Rating an album will affect the overall rating score for:

- The artist who created the album
- Any genres associated with the album

## Installation

> [!CAUTION]
> Apollo is untested on Windows.

The only way to install Apollo currently is through:

```sh
go install github.com/Wraith29/apollo@latest
```
