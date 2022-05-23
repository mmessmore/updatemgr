#!/bin/bash

prog=${0##*/}

SW=updatemgr
ARCHES=()
REV=1
ALL_ARCHES=( x86_64 armv7l aarch64 )

# Error codes from errno.h for consistency
EINVAL=22
ENOENT=2

error() {
	local msg="$*"
	echo "${prog} ERROR: ${msg}" >&2
}

usage() {
	cat <<EOM

${prog}
  Make deb packages for updatemgr.

  If run from \`main' branch, it will generate a version using the last tag.
  If run from \`dev' branch, it will do the same appending a patch version of
    YYYYMMDDHHmm.

  A package revision number can be added when only packaging changes have been
  made.

  This must be run from the root directory of the project.

USAGE
  ${prog} [ -a ARCH ] [ -r REV ]
  ${prog} -h

OPTIONS
  -a ARCH   Architecture to build dpkg for.  \`all' specifies all
            This may be specified more than once.
            Default: all: ${ALL_ARCHES[*]}
  -h        This friendly usage message
  -r REV    Package revision
            Default: ${REV}

EXAMPLES
  Build deb packages for all architectures.
  \$ samples/${prog}

  Build just the two ARM architectures with a package revision of 2
  \$ samples/${prog} -a arm -a arm64 -r 2

EOM
}

while getopts a:hr: OPT; do
	case "$OPT" in
		a)
			ARCHES+=("$OPTARG")
			;;
		h)
			usage
			exit 0
			;;
		r)
			REV="$OPTARG"
			;;
		*)
			error "Unrecognized option ${OPT}"
			usage
			exit $EINVAL
			;;
	esac
done

shift $(( OPTIND - 1))

if ! [ -d release ] && ! [ -d sample ]; then
	error "Must be run from project root directory"
	exit 1
fi

if ! command -v dpkg-deb >/dev/null 2>&1; then
	error "\`dpkg-deb' command not present"
	exit $ENOENT
fi

array_contains() {
  local e match="$1"
  shift
  for e; do
	  [[ "$e" == "$match" ]] && return 0
  done
  return 1
}

make_layout() {
	if ! cd "$tdir"; then
		error "Could not \`cd' to temp directory"
		exit $ENOENT
	fi
	mkdir -p usr/lib/systemd/system
	mkdir -p usr/bin
	mkdir -p etc/updatemgr
	mkdir -p var/lib/updatemgr
	mkdir DEBIAN
}

copy_bin() {
	BIN_FILE="release/${SW}.linux.${GO_ARCH}"
	if ! [ -f "$BIN_FILE" ]; then
		error "No ${BIN_FILE} exists, exiting"
		exit $ENOENT
	fi
	install -m 0755 "$BIN_FILE" "${tdir}/usr/bin/updatemgr"
}

copy_systemd() {
	install -m 0644 sample/*.service "${tdir}/usr/lib/systemd/system/"
}

copy_config() {
	for c in sample/*.yaml; do
		f=${c##*/}
		install -m 0644  "$c" "${tdir}/etc/updatemgr/${f}.example"
	done
}

copy_control() {
	m4 -D VERSION="$VERSION" -D DEB_ARCH="$DEB_ARCH" sample/dpkg/control.in > \
		"${tdir}/DEBIAN/control"
	cp sample/dpkg/* "${tdir}/DEBIAN/"
}

build_dpkg() {
	dpkg-deb --build --root-owner-group "$tdir" \
		"release/updatemgr_${VERSION}_${DEB_ARCH}.deb"
}

get_version() {
	local version
	version="$(git tag | tr -d v)"
	if [ -z "$version" ]; then
		error "No git tag to determine version"
		exit $ENOENT
	fi
	case $(git branch --show-current) in
		dev)
			version="$version.$(date +%Y%m%d%H%M)"
			;;
		main)
			true
			;;
		*)
			error "Need to be in \`dev' or \`main' branch to build"
			exit 1
			;;
	esac
	echo "${version}-${REV}"
}

VERSION=$(get_version)

# empty or "all" gets all the arches
if [ "${#ARCHES[@]}" = 0 ]; then
	ARCHES=("${ALL_ARCHES[@]}")
elif array_contains all "${ARCHES[@]}"; then
	ARCHES=("${ALL_ARCHES[@]}")
fi

echo "${prog}: Building packages for ${ARCHES[*]}"

# Now go build packages for every architecture
for ARCH in "${ARCHES[@]}"; do
	echo "********************************"
	echo "* Building package for ${ARCH} *"
	echo "********************************"
	tdir=$(mktemp -d "${prog}.XXXXXX")
	export tdir
	trap 'rm -fr "$tdir"' EXIT

	# translate architecture values to debian names
	case "$ARCH" in
		x86_64|amd64)
			export DEB_ARCH=amd64
			export GO_ARCH=amd64
			;;
		armv7l|armhf|arm)
			export DEB_ARCH=armhf
			export GO_ARCH=arm
			;;
		aarch64|arm64)
			export DEB_ARCH=arm64
			export GO_ARCH=arm64
			;;
		*)
			error "Unsupported architecture: ${ARCH}, skipping"
			exit 1
			;;
	esac

	# do this in a subshell
	(
		make_layout
	)
	copy_bin
	copy_systemd
	copy_config
	copy_control
	build_dpkg
done
