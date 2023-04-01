#!/bin/sh
set -e
echo "Searching latest version..."
RELEASES_URL="https://github.com/ThankRain/shook/releases"
BASENAME="shook"
LATEST="$(curl -sI https://github.com/ThankRain/shook/releases/latest | grep location: | awk '{printf $2}' | cut -d '/' -f8 | tr -d '\r\n')"

test -z "$VERSION" && VERSION="$LATEST"

test -z "$VERSION" && {
	echo "Unable to get shook version." >&2
	exit 1
}

test -z "$TMPDIR" && TMPDIR="$(mktemp -d)"
export TAR_FILE="$TMPDIR/${BASENAME}_$(uname -s)_$(uname -m).tar.gz"

(
	cd "$TMPDIR"
	echo "Downloading Shook $VERSION"
	echo "$RELEASES_URL/download/$VERSION/${BASENAME}_$(uname -s)_$(uname -m).tar.gz"
	curl -fLo "$TAR_FILE" \
		"$RELEASES_URL/download/$VERSION/${BASENAME}_$(uname -s)_$(uname -m).tar.gz"
	curl -fLo "checksums.txt" "$RELEASES_URL/download/$VERSION/checksums.txt"
	echo "Verifying checksums..."
	sha256sum --ignore-missing --quiet --check checksums.txt
)
tar -xf "$TAR_FILE" -C "$TMPDIR"
echo "Installing to /usr/local/bin/shook"
sudo mv "${TMPDIR}/shook" "${TMPDIR}/shook-server" "/usr/local/bin/"
echo "Congratulations!"
echo "shook & shook-server installed successfully!"