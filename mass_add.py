from pathlib import Path
from time import sleep
from subprocess import run

def main() -> int:
    artists = Path("./artists.txt").read_text().splitlines()

    failed_artists = []

    for idx, artist in enumerate(artists):
        print(f"Adding artist {idx}: {artist}")
        result = run(["apollo", "add", artist], capture_output=True)

        if result.returncode != 0:
            failed_artists.append(artist)

            print(str(result.stderr))

        sleep(1)

    print(f"The following artists failed: {failed_artists}")

    return len(failed_artists)


if __name__ == "__main__":
    raise SystemExit(main())
